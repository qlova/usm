package bytecode

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/qlova/usm"
	"github.com/qlova/usm/template"
)

//Target is a Bytecode target for u
type Target struct {
	template.Target
}

//WriteBlock writes a block.
func (t *Target) WriteBlock(block usm.Block) {
	block()
	t.WriteByte(End)
}

//WriteInt64 writes a int64 to the target.
func writeInt64(w io.Writer, i int64) {
	binary.Write(w, binary.LittleEndian, i)
}

//WriteInt64 writes a int64 to the target.
func (t *Target) WriteInt64(i int64) {
	writeInt64(t, i)
}

//WriteTo writes the target.
func (t *Target) WriteTo(writer io.Writer) (int64, error) {
	return t.Target.WriteTo(writer)
}

//Main is the entrypoint of the program.
func (t *Target) Main(body usm.Block) {
	t.WriteByte(Main)
	t.WriteBlock(body)
}

//String returns the String given by the go.string
func (t *Target) String(s string) usm.Value {
	var b bytes.Buffer
	b.WriteByte(String)
	writeInt64(&b, int64(len(s)))
	b.Write([]byte(s))
	return b.Bytes()
}

//Send writes the string data into the stream, returns the number of bytes written.
//This may throw an error.
func (t *Target) Send(stream usm.Stream, s usm.String) usm.Value {
	if stream == nil {
		var b bytes.Buffer
		b.WriteByte(Send)
		b.WriteByte(Nil)
		b.Write(s.([]byte))
		return b.Bytes()
	}
	panic("not implemented")
}

//Discard allows a value to be used as a statement.
func (t *Target) Discard(value usm.Value) {
	t.WriteByte(Discard)
	t.Write(value.([]byte))
}
