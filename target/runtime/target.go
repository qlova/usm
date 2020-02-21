package runtime

import (
	"errors"
	"io"
	"math/big"
	"os"

	"github.com/qlova/usm"
	"github.com/qlova/usm/template"
)

//Target is a Go target for u
type Target struct {
	template.Target

	Runtime
}

//Block returns a Block from a usm.Block
func (t *Target) Block(body usm.Block) Block {
	var old = t.Current
	t.Current = new(Block)
	body()
	var block = *t.Current
	t.Current = old
	return block
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
		return []byte(s)
	})
}

//Create creates a new String of the given size.
func (t *Target) Create(n usm.Number) usm.String {
	var size = n.(Value)
	return Value(func() interface{} {
		return make([]byte, size().(*big.Int).Int64())
	})
}

//Bit returns the Bit given by the go.bool
func (t *Target) Bit(b bool) usm.Value {
	return Value(func() interface{} {
		return b
	})
}

//Read reads stream data into the given string, returns the number of bytes read.
//This may throw an error.
func (t *Target) Read(stream usm.Stream, data usm.String) usm.Value {
	if stream == nil {
		var data = data.(Value)
		return Value(func() interface{} {
			n, _ := os.Stdin.Read(data().([]byte))
			return n
		})
	}
	panic("not implemented")
}

//Send writes the string data into the stream, returns the number of bytes written.
//This may throw an error.
func (t *Target) Send(stream usm.Stream, s usm.String) usm.Value {
	if stream == nil {
		var s = s.(Value)
		return Value(func() interface{} {
			n, _ := os.Stdout.Write(s().([]byte))
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

//Define defines a function, returning the label to the function.
//arguments is the number of the arguments the function expects.
func (t *Target) Define(arguments int, body usm.Block) usm.Label {
	t.Labels++
	var old = t.Current
	t.Current = new(Block)
	body()
	t.Blocks = append(t.Blocks, *t.Current)
	t.Current = old
	return usm.Label(len(t.Blocks))
}

//Var creates a new variable set to the provided value.
//Returns the register for future reference to the variable.
func (t *Target) Var(value usm.Value) usm.Register {
	t.Registers++

	var val = value.(Value)

	t.Write(func() {
		t.Variables[t.Registers] = val()
	})

	return t.Registers
}

//JumpTo jumps to the label passing the provided arguments.
//JumpTo ignores any return values.
//If the label is 0, then the first argument is treated as a label bind and subsequent arguments are passed.
func (t *Target) JumpTo(label usm.Label, arguments ...usm.Value) {
	t.Write(func() {
		t.Blocks[label-1].RunWith(&t.Runtime)
	})
}

//Number returns the Number given by the go.big.Int
func (t *Target) Number(b *big.Int) usm.Number {
	return Value(func() interface{} {
		return b
	})
}

//Call calls the provided label, passing the provided argument values and returns the result.
//If the label is 0, then the first argument is treated as a label bind and subsequent arguments are passed.
func (t *Target) Call(label usm.Label, args ...usm.Value) usm.Value {
	return Value(func() interface{} {

		var converted []interface{}
		for i := range args {
			converted[i] = args[i].(Value)()
		}

		t.Blocks[label-1].RunWith(&t.Runtime, converted...)
		return t.ReturnValue
	})
}

//Get returns the value inside of the given register.
func (t *Target) Get(r usm.Register) usm.Value {
	if r < 0 {
		return Value(func() interface{} {
			return t.Args[-r]
		})
	}
	return Value(func() interface{} {
		return t.Variables[r]
	})
}

//Return returns the result to the caller.
//Pass nil to return without passing a value.
func (t *Target) Return(result usm.Value) {
	if result == nil {
		t.Write(func() {
			t.ReturnValue = nil
			t.Returning = true
		})
		return
	}
	var r = result.(Value)
	t.Write(func() {
		t.ReturnValue = r()
		t.Returning = true
	})

}

//If branches to the body Block if the condition is not zero.
//If the condition is zero, this process follows the chain, treating them as elseif's.
//The last block is branched to if none of the previous branches were followed.
func (t *Target) If(condition usm.Bit, body usm.Block, chain []usm.ElseIf, last usm.Block) {

	var c = condition.(Value)
	var first = t.Block(body)

	type ElseIf struct {
		Value
		Block
	}

	var converted []ElseIf
	for i := range chain {
		converted[i] = ElseIf{
			Value: chain[i].Bit.(Value),
			Block: t.Block(chain[i].Block),
		}
	}

	if last == nil {
		t.Write(func() {
			if c().(bool) {
				first.RunWith(&t.Runtime)
			}

			for i := range converted {
				if (converted[i].Value()).(bool) {
					converted[i].Block.RunWith(&t.Runtime)
					return
				}
			}
		})
		return
	}

	var l = t.Block(last)

	t.Write(func() {
		if c().(bool) {
			first.RunWith(&t.Runtime)
		}

		for i := range converted {
			if (converted[i].Value()).(bool) {
				converted[i].Block.RunWith(&t.Runtime)
				return
			}
		}

		l.RunWith(&t.Runtime)
	})
}

//Loop loops the body while an optional condition is true.
//If condition is nil, then the loop is infinite.
func (t *Target) Loop(condition usm.Number, body usm.Block) {
	var block = t.Block(body)

	if condition == nil {
		t.Write(func() {
			for {
				block.RunWith(&t.Runtime)
			}
		})
	} else {
		panic("not implimented")
	}
}

//Symbol returns the byte at the given index in the String.
func (t *Target) Symbol(data usm.String, index usm.Number) usm.Number {
	var d = data.(Value)
	var i = index.(Value)
	return Value(func() interface{} {
		return (d().([]byte))[i().(*big.Int).Int64()]
	})
}

//Same returns 1 if a is equal to b, otherwise 0.
func (t *Target) Same(a usm.Number, b usm.Number) usm.Bit {
	var A = a.(Value)
	var B = b.(Value)
	return Value(func() interface{} {
		return (A().(*big.Int)).Cmp(B().(*big.Int)) == 0
	})
}
