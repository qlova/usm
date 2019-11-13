package golang

import (
	"fmt"
	"io"
	"strconv"

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
	return t.Target.WriteTo(writer)
}

//Main is the entrypoint of the program.
func (t *Target) Main(body usm.Block) {
	t.WriteStatement("func main() {\n")
	t.WriteStatement("\tvar r = new(Runtime)\n")
	t.Indent(body)
	t.WriteStatement("}\n")
}

//String returns the String given by the go.string
func (t *Target) String(s string) usm.Value {
	return fmt.Sprintf(`r.String(%v)`, strconv.Quote(s))
}

//Send writes the string data into the stream, returns the number of bytes written.
//This may throw an error.
func (t *Target) Send(stream usm.Stream, s usm.String) usm.Value {
	if stream == nil {
		return fmt.Sprintf(`r.Stdout(%v)`, s)
	}
	panic("not implemented")
}

//Discard allows a value to be used as a statement.
func (t *Target) Discard(value usm.Value) {
	t.WriteStatement("_ = %v\n", value)
}
