package runtime

import (
	"errors"
	"io"
	"os"

	"github.com/qlova/usm"
	"github.com/qlova/usm/template"
)

//Target is a Go target for u
type Target struct {
	template.Target

	Runtime
}

func (t *Target) Write(f func()) {
	t.Current.Statements = append(t.Current.Statements, f)
}

//WriteTo writes the target.
func (t *Target) WriteTo(writer io.Writer) (int64, error) {
	return 0, errors.New("runtime.WriteTo: impossible to write runtime")
}

//Main is the entrypoint of the program.
func (t *Target) Main(body usm.Block) {
	t.Entrypoint = new(Block)
	t.Current = t.Entrypoint
	body()
	t.Current = nil
}

//String returns the String given by the go.string
func (t *Target) String(s string) usm.Value {
	return Value(func() interface{} {
		return s
	})
}

//Send writes the string data into the stream, returns the number of bytes written.
//This may throw an error.
func (t *Target) Send(stream usm.Stream, s usm.String) usm.Value {
	if stream == nil {
		var s = s.(Value)
		return Value(func() interface{} {
			n, _ := os.Stdout.Write([]byte(s().(string)))
			return n
		})
	}
	panic("not implemented")
}

//Discard allows a value to be used as a statement.
func (t *Target) Discard(value usm.Value) {
	var f = value.(Value)
	t.Write(func() {
		_ = f()
	})
}
