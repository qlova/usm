package bytecode

import (
	"bytes"
	"encoding/binary"
	"io"
	"math/big"

	"github.com/qlova/usm"
	"github.com/qlova/usm/template"
)

//Target is a Bytecode target for u
type Target struct {
	template.Target
	labels    usm.Label
	registers usm.Register
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

//WriteValue writes a value to the target.
func (t *Target) WriteValue(v usm.Value) {
	t.Write(v.([]byte))
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

//Loop loops the body while an optional condition is true.
//If condition is nil, then the loop is infinite.
func (t *Target) Loop(condition usm.Number, body usm.Block) {
	t.WriteByte(Loop)
	t.WriteValue(condition)
	t.WriteBlock(body)
}

//String returns the String given by the go.string
func (t *Target) String(s string) usm.Value {
	var b bytes.Buffer
	b.WriteByte(String) //Header

	writeInt64(&b, int64(len(s)))
	b.Write([]byte(s))
	return b.Bytes()
}

//Create creates a new String of the given size.
func (t *Target) Create(n usm.Number) usm.String {
	var b bytes.Buffer
	b.WriteByte(Create) //Header
	b.Write(n.([]byte))
	return b.Bytes()
}

//Number returns the Number given by the *go.big.Int
func (t *Target) Number(i *big.Int) usm.Value {
	var b bytes.Buffer
	b.WriteByte(Number) //Header

	var bytes = i.Bytes()
	writeInt64(&b, int64(len(bytes)))
	b.Write([]byte(bytes))
	return b.Bytes()
}

//Bit returns the Bit given by the go.bool
func (t *Target) Bit(bit bool) usm.Value {
	var b bytes.Buffer
	b.WriteByte(Bit) //Header

	if bit {
		b.WriteByte(1)
	} else {
		b.WriteByte(0)
	}

	return b.Bytes()
}

//Define defines a function, returning the label to the function.
//arguments is the number of the arguments the function expects.
func (t *Target) Define(arguments int, body usm.Block) usm.Label {
	t.WriteByte(Define)
	t.WriteInt64(int64(arguments))
	t.WriteBlock(body)
	t.labels++
	return t.labels
}

//Var creates a new variable set to the provided value.
//Returns the register for future reference to the variable.
func (t *Target) Var(value usm.Value) usm.Register {
	t.WriteByte(Var)
	t.WriteValue(value)
	t.registers++
	return t.registers
}

//JumpTo jumps to the label passing the provided arguments.
//JumpTo ignores any return values.
//If the label is 0, then the first argument is treated as a label bind and subsequent arguments are passed.
func (t *Target) JumpTo(label usm.Label, arguments ...usm.Value) {
	t.WriteByte(JumpTo)
	t.WriteInt64(int64(label))
	t.WriteInt64(int64(len(arguments)))
	for _, arg := range arguments {
		t.Write(arg.([]byte))
	}
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

//Read reads stream data into the given string, returns the number of bytes read.
//This may throw an error.
func (t *Target) Read(stream usm.Stream, s usm.String) usm.Value {
	var b bytes.Buffer
	b.WriteByte(Read)
	b.Write(stream.([]byte))
	b.Write(s.([]byte))
	return b.Bytes()
}
