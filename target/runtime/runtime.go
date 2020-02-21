package runtime

import "github.com/qlova/usm"

//Value is a runtime usm.Value
type Value func() interface{}

//Block is a runtime usm.Block
type Block struct {
	Function bool

	Statements []func()
}

//RunWith runs a block with the given runtime.
func (block Block) RunWith(r *Runtime, args ...interface{}) error {
	if len(block.Statements) == 0 {
		return nil
	}

	r.Push()
	defer r.Pop()

	r.Args = args
	for {
		if r.ProgramCounter >= len(block.Statements) {
			return nil
		}
		block.Statements[r.ProgramCounter]()
		r.ProgramCounter++

		if r.Returning {
			if block.Function {
				r.Returning = false
			}
			return nil
		}
	}
}

//Runtime is a runtime object for a runtime `u` target.
type Runtime struct {
	Scope
	Scopes []Scope

	Current, Entrypoint *Block
	Blocks              []Block

	ReturnValue interface{}
	Returning   bool
}

//Push pushes a new scope.
func (r *Runtime) Push() {
	r.Scopes = append(r.Scopes, r.Scope)
	r.Scope = NewScope()
	r.Backup = r.ProgramCounter
	r.ProgramCounter = 0
}

//Pop pops the last scope.
func (r *Runtime) Pop() {
	var backup = r.Backup
	r.Scope = r.Scopes[len(r.Scopes)-1]
	r.Scopes = r.Scopes[:len(r.Scopes)-1]
	r.ProgramCounter = backup
}

//Run runs the runtime.
func (r *Runtime) Run() error {
	r.Scope = NewScope()
	r.Scopes = nil

	return r.Entrypoint.RunWith(r)
}

//Scope is the current scope.
type Scope struct {
	Backup         int
	ProgramCounter int

	Args      []interface{}
	Variables map[usm.Register]interface{}
}

//NewScope returns a new scope.
func NewScope() Scope {
	return Scope{
		Variables: make(map[usm.Register]interface{}),
	}
}
