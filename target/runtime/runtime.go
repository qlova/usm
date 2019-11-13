package runtime

//Value is a runtime usm.Value
type Value func() interface{}

//Block is a runtime usm.Block
type Block struct {
	Statements []func()
}

//Runtime is a runtime object for a runtime `u` target.
type Runtime struct {
	ProgramCounter int

	Current, Entrypoint *Block
}

//Run runs the runtime.
func (r *Runtime) Run() error {
	r.Current = r.Entrypoint

	for {
		if r.ProgramCounter >= len(r.Current.Statements) {
			return nil
		}
		r.Current.Statements[r.ProgramCounter]()
		r.ProgramCounter++
	}
}
