package golang

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/qlova/usm"
	"github.com/qlova/usm/template"
)

//Target is a Go target for u
type Target struct {
	template.Target
}

//WriteTo writes the target.
func (t *Target) WriteTo(writer io.Writer) (int64, error) {
	writer.Write([]byte(Runtime))
	writer.Write(t.Head.Bytes())
	return t.Target.WriteTo(writer)
}

//Main is the entrypoint of the program.
func (t *Target) Main(body usm.Block) {
	t.WriteStatement("func main() {\n")
	t.WriteStatement("\tvar r = new(Runtime)\n")
	t.Indent(body)
	t.WriteStatement("}\n")
}

//Loop loops the body while an optional condition is true.
//If condition is nil, then the loop is infinite.
func (t *Target) Loop(condition usm.Number, body usm.Block) {
	if condition == nil {
		t.WriteStatement("for {\n")
	} else {
		t.WriteStatement("for %v {\n", condition)
	}
	t.Indent(body)
	t.WriteStatement("}\n")
}

//String returns the String given by the go.string
func (t *Target) String(s string) usm.Value {
	return fmt.Sprintf(`r.String(%v)`, strconv.Quote(s))
}

//Bit returns the Bit given by the go.bool
func (t *Target) Bit(b bool) usm.Value {
	if b {
		return "true"
	}
	return "false"
}

//Send writes the string data into the stream, returns the number of bytes written.
//This may throw an error.
func (t *Target) Send(stream usm.Stream, s usm.String) usm.Value {
	if stream == nil {
		return fmt.Sprintf(`r.Stdout(%v)`, s)
	}
	panic("not implemented")
}

//Read reads stream data into the given string, returns the number of bytes read.
//This may throw an error.
func (t *Target) Read(stream usm.Stream, s usm.String) usm.Value {
	if stream == nil {
		return fmt.Sprintf(`r.Stdin(%v)`, s)
	}
	panic("not implemented")
}

//Discard allows a value to be used as a statement.
func (t *Target) Discard(value usm.Value) {
	t.WriteStatement("_ = %v\n", value)
}

//Create creates a new String of the given size.
func (t *Target) Create(n usm.Number) usm.String {
	return fmt.Sprintf(`make([]byte, %v)`, n)
}

//Define defines a function, returning the label to the function.
//arguments is the number of the arguments the function expects.
func (t *Target) Define(arguments int, body usm.Block) usm.Label {
	t.Labels++

	var args = make([]string, arguments)
	for i := 0; i < arguments; i++ {
		args[i] = fmt.Sprintf("a%v", i)
	}

	var backup = t.Buffer
	var old = t.Tabs
	t.Tabs = 0

	t.Buffer = bytes.Buffer{}

	if arguments > 0 {
		t.WriteStatement("func f%v(r *Runtime, %v Value) {\n", t.Labels, strings.Join(args, ","))
	} else {
		t.WriteStatement("func f%v(r *Runtime) {\n", t.Labels)
	}

	t.Indent(body)

	t.WriteStatement("}\n")

	t.Head.Write(t.Buffer.Bytes())

	t.Tabs = old
	t.Buffer = backup

	return t.Labels
}

//Var creates a new variable set to the provided value.
//Returns the register for future reference to the variable.
func (t *Target) Var(value usm.Value) usm.Register {
	t.Registers++

	t.WriteStatement("var v%v = %v\n", t.Registers, value)

	return t.Registers
}

//Get returns the value inside of the given register.
func (t *Target) Get(r usm.Register) usm.Value {
	return fmt.Sprintf(`v%v`, r)
}

//JumpTo jumps to the label passing the provided arguments.
//JumpTo ignores any return values.
//If the label is 0, then the first argument is treated as a label bind and subsequent arguments are passed.
func (t *Target) JumpTo(label usm.Label, arguments ...usm.Value) {
	var converted = make([]string, len(arguments))
	for i := range arguments {
		converted[i] = arguments[i].(string)
	}
	if len(arguments) > 0 {
		t.WriteStatement("f%v(r, %v)\n", label, strings.Join(converted, ","))
	} else {
		t.WriteStatement("f%v(r)\n", label)
	}

}
