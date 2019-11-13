package bytecode

//This is a list of all opcodes for usm bytecode.
const (
	Nil = iota

	//End is a special opcode that ends a block.
	End

	Var
	Set

	Discard

	Main

	If

	Loop
	Each
	Break

	Define
	Return
	JumpTo

	Throw

	Seek

	Delete

	Change
	Mutate
	Insert

	Remove

	Modify

	Number
	String
	Bit

	Get

	Bind

	Catch

	Call
	Fork

	Pointer

	Array
	Alloc
	Count
	Index
	Append

	Table
	Amount
	Lookup

	Create
	Equals
	Length

	Symbol
	Concat
	Follow

	Open
	Stat
	Read
	Send

	Add
	Sub
	Mul
	Div
	Mod
	Pow

	Less
	More
	Same

	And
	Or
	Not
)
