package bytecode

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"

	"github.com/qlova/usm"
)

//Reader is bytecode reader.
type Reader struct {
	*bufio.Reader
}

//NewReader creates a bytecode.Reader from the reader.
func NewReader(r io.Reader) Reader {
	return Reader{bufio.NewReader(r)}
}

//ReadInt64 reads an int64.
func (r Reader) ReadInt64() (int64, error) {
	var i int64
	err := binary.Read(r, binary.LittleEndian, &i)
	return i, err
}

//ReadValue reads a value.
func (r Reader) ReadValue(t usm.Target) (usm.Value, error) {
	for {
		var opcode, err = r.ReadByte()
		if err == io.EOF {
			return nil, errors.New("bytecode.Reader.ReadBlock: unexpected eof")
		} else if err != nil {
			return nil, err
		}

		switch opcode {
		case Nil:
			return nil, nil
		case String:
			var length, err = r.ReadInt64()
			if err != nil {
				return nil, err
			}
			var buffer = make([]byte, length)
			_, err = r.Read(buffer)
			if err != nil {
				return nil, err
			}
			return t.String(string(buffer)), nil
		case Send:
			var stream, err = r.ReadValue(t)
			if err != nil {
				return nil, err
			}
			value, err := r.ReadValue(t)
			if err != nil {
				return nil, err
			}
			return t.Send(stream, value), nil
		}
	}
}

//ReadStatement reads a usm statement from the reader.
func (r Reader) ReadStatement(opcode byte, t usm.Target) (err error) {
	switch opcode {
	case Main:
		t.Main(func() {
			err = r.ReadBlock(t)
		})
		if err != nil {
			return err
		}
	case Discard:
		value, err := r.ReadValue(t)
		if err != nil {
			return err
		}
		t.Discard(value)
	}
	return
}

//ReadBlock reads a block from the reader.
func (r Reader) ReadBlock(t usm.Target) error {
	for {
		var opcode, err = r.ReadByte()
		if err == io.EOF {
			return errors.New("bytecode.Reader.ReadBlock: unexpected eof")
		} else if err != nil {
			return err
		}

		if err := r.ReadStatement(opcode, t); err != nil {
			return err
		}
	}
}

//Target assembles the bytecode to the specified target.
func (r Reader) Target(t usm.Target) error {
	for {
		var opcode, err = r.ReadByte()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		if err := r.ReadStatement(opcode, t); err != nil {
			return err
		}
	}
}
