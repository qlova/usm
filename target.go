//Package usm provides the specification of the usm assembly language.
package usm

import (
	"math/big"
	"reflect"
)

//Value is any usm value.
type Value interface{}

type (
	//Bit is either 'true' or 'false'.
	Bit Value

	//Number is a numeric value.
	Number Value

	//String is a string of bytes.
	String Value

	//Array is a sequence of values.
	Array Value

	//Table is an association between strings and values.
	Table Value

	//Pointer is a shared reference to a value.
	Pointer Value

	//Stream is a external io connection that can be read from, written to or stat.
	Stream Value

	//Function is a binded label.
	Function Value
)

//Register is a reference to a value.
type Register int

//Block is a block of code.
type Block func()

//Label is like a function.
type Label int

//Type is a usm type.
type Type reflect.Type

//ElseIfChain is an ElseIfChain chain.
type ElseIfChain struct {
	Bit
	Block
}

//Target is a usm target.
type Target interface {

	//Var creates a new variable set to the provided value.
	//Returns the register for future reference to the variable.
	Var(Value) Register

	//Set sets the variable in the given register to be the given value.
	Set(Register, Value)

	//Discard allows a value to be used as a statement.
	Discard(Value)

	//Main is the entrypoint of the program.
	Main(body Block)

	//If branches to the body Block if the condition is not zero.
	//If the condition is zero, this process follows the chain, treating them as elseif's.
	//The last block is branched to if none of the previous branches were followed.
	If(condition Bit, body Block, chain ElseIfChain, last Block)

	//Loop loops the body while an optional condition is true.
	//If condition is nil, then the loop is infinite.
	Loop(condition Number, body Block)

	//Each loops over an array, placing the index into 'i' and the value into 'v'.
	Each(array Array, body func(i Number, v Value))

	//Break breaks the inenr-most loop.
	Break()

	//Define defines a function, returning the label to the function.
	//arguments is the number of the arguments the function expects.
	Define(arguments int, body Block) Label

	//Return returns the result to the caller.
	//Pass nil to return without passing a value.
	Return(result Value)

	//JumpTo jumps to the label passing the provided arguments.
	//JumpTo ignores any return values.
	//If the label is 0, then the first argument is treated as a label bind and subsequent arguments are passed.
	JumpTo(label Label, arguments ...Value)

	//Throw throws an Value onto the thread-local Errors stack.
	Throw(Value)

	//Seek attempts to advance the stream by discarding a specified number of bytes from the stream.
	Seek(Stream, Number)

	//Delete frees the memory of the given Value.
	//Has no effect in garbage collected targets.
	Delete(Type, Value)

	//Change changes the pointer value to the provided Value.
	Change(Pointer, Value)

	//Mutate mutates the array at the given index to be set to the given value.
	Mutate(Array, Number, Value)

	//Insert sets the table value at the given string key to be set to the given value.
	Insert(Table, String, Value)

	//Remove removes the given key from the table.
	Remove(table, key Value)

	//Modify mutates a string and sets the index to be set to the given number.
	//If the number's byte representaion is greater than 1.
	Modify(String, Number, Number)

	//Number returns the Number given by the go.big.Int
	Number(*big.Int) Number

	//String returns the String given by the go.string
	String(string) Value

	//Bit returns the Bit given by the go.bool
	Bit(bool) Value

	//Get returns the value inside of the given register.
	Get(Register) Value

	//Bind returns the label as a value that can be passed to a Call, JumpTo or Fork by passing an empty function argument
	Bind(Label) Value

	//Catch removes and returns the latest error on the thread-local error stack.
	Catch() Value

	//Call calls the provided label, passing the provided argument values and returns the result.
	//If the label is 0, then the first argument is treated as a label bind and subsequent arguments are passed.
	Call(label Label, args ...Value) Value

	//Fork jumps to the label in an independant parallel runtime, the arguments are passed.
	//A connected stream is returned, this connects to the Stdin and Stdout of the new runtime.
	Fork(label Label, args ...Value) Stream

	//Pointer retuns a pointer to the provided value.
	Pointer(Value) Pointer

	//Alloc creates a new array of the given size.
	Alloc(Number) Array

	//Array creates a new array with the given elements.
	Array(elements ...Value) Array

	//Table creates a new table with the given elements.
	Table(elements map[Value]Value) Table

	//Count returns the number of elements in the array.
	Count(Array) Number

	//Index returns the value at the given index in the array.
	Index(Array, Number) Value

	//Append adds an element to the end of the array.
	Append(Array, Value) Array

	//Amount returns the number of items in the Table.
	Amount(Table) Value

	//Lookup returns the value at the given key in the Table.
	Lookup(Table, String) Value

	//Create creates a new String of the given size.
	Create(Number) String

	//Equals returns 1 is the two Strings are equal. Returns 0 otherwise.
	Equals(a, b String) Bit

	//Length returns the length of the String in bytes.
	Length(String) Number

	//Symbol returns the byte at the given index in the String.
	Symbol(String, Number) Number

	//Concat creates a new String that is the concatenation of the given strings.
	Concat(a, b String) String

	//Follow returns the value that the pointer is pointing at.
	Follow(Pointer) Value

	//Open returns a stream from the given platform-dependent URI.
	//This may throw an error.
	Open(String) Stream

	//Stat performs a platform-dependent stat on the stream and returns the result.
	Stat(Stream) String

	//Read reads stream data into the given string, returns the number of bytes read.
	//This may throw an error.
	Read(Stream, String) Value

	//Send writes the string data into the stream, returns the number of bytes written.
	//This may throw an error.
	Send(Stream, String) Value

	//Add returns the sum of a and b.
	Add(a, b Number) Number

	//Mul returns the product of a and b.
	Mul(a, b Number) Number

	//Sub returns the difference between a and b.
	Sub(a, b Number) Number

	//Div returns the quotient of a and b.
	Div(a, b Number) Number

	//Mod returns the modulos of a and b. Must mimic Go % operator.
	Mod(a, b Number) Number

	//Pow returns a to the power of b.
	Pow(a, b Number) Number

	//Less returns 1 if a is smaller than b, otherwise 0.
	Less(a, b Number) Bit

	//More returns 1 if a is larger than b, otherwise 0.
	More(a, b Number) Bit

	//Same returns 1 if a is equal to b, otherwise 0.
	Same(a, b Number) Bit

	//And returns a && b
	And(a, b Bit) Bit

	//Or returns a || b
	Or(a, b Bit) Bit

	//Not returns !Bit
	Not(Bit) Bit
}
