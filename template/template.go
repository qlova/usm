//Package template provides a usm.Target template.
package template

import (
	"bytes"
	"fmt"
	"io"
	"math/big"

	"github.com/qlova/usm"
)

//Target is a Target Target
type Target struct {
	bytes.Buffer
	Head bytes.Buffer

	Tabs      int
	Labels    usm.Label
	Registers usm.Register
}

//Indent processes and indents the given block.
func (t *Target) Indent(block usm.Block) {
	t.Tabs++
	block()
	t.Tabs--
}

//WriteStatement writes and indents a statement.
func (t *Target) WriteStatement(format string, args ...interface{}) {
	t.Write(bytes.Repeat([]byte{'\t'}, t.Tabs))
	fmt.Fprintf(t, format, args...)
}

//WriteTo writes the target's output to the given writer.
func (t *Target) WriteTo(writer io.Writer) (int64, error) {
	i, err := writer.Write(t.Bytes())
	return int64(i), err
}

//Var creates a new variable set to the provided value.
//Returns the register for future reference to the variable.
func (t *Target) Var(_ usm.Value) usm.Register {
	panic("not implemented")
}

//Set sets the variable in the given register to be the given value.
func (t *Target) Set(_ usm.Register, _ usm.Value) {
	panic("not implemented")
}

//Discard allows a value to be used as a statement.
func (t *Target) Discard(_ usm.Value) {
	panic("not implemented")
}

//Main is the entrypoint of the program.
func (t *Target) Main(body usm.Block) {
	panic("not implemented")
}

//If branches to the body Block if the condition is not zero.
//If the condition is zero, this process follows the chain, treating them as elseif's.
//The last block is branched to if none of the previous branches were followed.
func (t *Target) If(condition usm.Bit, body usm.Block, chain []usm.ElseIf, last usm.Block) {
	panic("not implemented")
}

//Loop loops the body while an optional condition is true.
//If condition is nil, then the loop is infinite.
func (t *Target) Loop(condition usm.Number, body usm.Block) {
	panic("not implemented")
}

//Each loops over an array, placing the index into 'i' and the value into 'v'.
func (t *Target) Each(array usm.Array, body func(i usm.Number, v usm.Value)) {
	panic("not implemented")
}

//Break breaks the inenr-most loop.
func (t *Target) Break() {
	panic("not implemented")
}

//Define defines a function, returning the label to the function.
//arguments is the number of the arguments the function expects.
func (t *Target) Define(arguments int, body usm.Block) usm.Label {
	panic("not implemented")
}

//Return returns the result to the caller.
//Pass nil to return without passing a value.
func (t *Target) Return(result usm.Value) {
	panic("not implemented")
}

//JumpTo jumps to the label passing the provided arguments.
//JumpTo ignores any return values.
//If the label is 0, then the first argument is treated as a label bind and subsequent arguments are passed.
func (t *Target) JumpTo(label usm.Label, arguments ...usm.Value) {
	panic("not implemented")
}

//Throw throws an Value onto the thread-local Errors stack.
func (t *Target) Throw(_ usm.Value) {
	panic("not implemented")
}

//Seek attempts to advance the stream by discarding a specified number of bytes from the stream.
func (t *Target) Seek(_ usm.Stream, _ usm.Number) {
	panic("not implemented")
}

//Delete frees the memory of the given Value.
//Has no effect in garbage collected targets.
func (t *Target) Delete(_ usm.Type, _ usm.Value) {
	panic("not implemented")
}

//Change changes the pointer value to the provided Value.
func (t *Target) Change(_ usm.Pointer, _ usm.Value) {
	panic("not implemented")
}

//Mutate mutates the array at the given index to be set to the given value.
func (t *Target) Mutate(_ usm.Array, _ usm.Number, _ usm.Value) {
	panic("not implemented")
}

//Insert sets the table value at the given string key to be set to the given value.
func (t *Target) Insert(_ usm.Table, _ usm.String, _ usm.Value) {
	panic("not implemented")
}

//Remove removes the given key from the table.
func (t *Target) Remove(table usm.Value, key usm.Value) {
	panic("not implemented")
}

//Modify mutates a string and sets the index to be set to the given number.
//If the number's byte representaion is greater than 1.
func (t *Target) Modify(_ usm.String, _ usm.Number, _ usm.Number) {
	panic("not implemented")
}

//Number returns the Number given by the go.big.Int
func (t *Target) Number(_ *big.Int) usm.Number {
	panic("not implemented")
}

//String returns the String given by the go.string
func (t *Target) String(_ string) usm.Value {
	panic("not implemented")
}

//Bit returns the Bit given by the go.bool
func (t *Target) Bit(_ bool) usm.Value {
	panic("not implemented")
}

//Get returns the value inside of the given register.
func (t *Target) Get(_ usm.Register) usm.Value {
	panic("not implemented")
}

//Bind returns the label as a value that can be passed to a Call, JumpTo or Fork by passing an empty function argument
func (t *Target) Bind(_ usm.Label) usm.Value {
	panic("not implemented")
}

//Catch removes and returns the latest error on the thread-local error stack.
func (t *Target) Catch() usm.Value {
	panic("not implemented")
}

//Call calls the provided label, passing the provided argument values and returns the result.
//If the label is 0, then the first argument is treated as a label bind and subsequent arguments are passed.
func (t *Target) Call(label usm.Label, args ...usm.Value) usm.Value {
	panic("not implemented")
}

//Fork jumps to the label in an independant parallel runtime, the arguments are passed.
//A connected stream is returned, this connects to the Stdin and Stdout of the new runtime.
func (t *Target) Fork(label usm.Label, args ...usm.Value) usm.Stream {
	panic("not implemented")
}

//Pointer retuns a pointer to the provided value.
func (t *Target) Pointer(_ usm.Value) usm.Pointer {
	panic("not implemented")
}

//Alloc creates a new array of the given size.
func (t *Target) Alloc(_ usm.Number) usm.Array {
	panic("not implemented")
}

//Array creates a new array with the given elements.
func (t *Target) Array(elements ...usm.Value) usm.Array {
	panic("not implemented")
}

//Table creates a new table with the given elements.
func (t *Target) Table(elements map[usm.Value]usm.Value) usm.Table {
	panic("not implemented")
}

//Count returns the number of elements in the array.
func (t *Target) Count(_ usm.Array) usm.Number {
	panic("not implemented")
}

//Index returns the value at the given index in the array.
func (t *Target) Index(_ usm.Array, _ usm.Number) usm.Value {
	panic("not implemented")
}

//Append adds an element to the end of the array.
func (t *Target) Append(_ usm.Array, _ usm.Value) usm.Array {
	panic("not implemented")
}

//Amount returns the number of items in the Table.
func (t *Target) Amount(_ usm.Table) usm.Value {
	panic("not implemented")
}

//Lookup returns the value at the given key in the Table.
func (t *Target) Lookup(_ usm.Table, _ usm.String) usm.Value {
	panic("not implemented")
}

//Create creates a new String of the given size.
func (t *Target) Create(_ usm.Number) usm.String {
	panic("not implemented")
}

//Equals returns 1 is the two Strings are equal. Returns 0 otherwise.
func (t *Target) Equals(a usm.String, b usm.String) usm.Bit {
	panic("not implemented")
}

//Length returns the length of the String in bytes.
func (t *Target) Length(_ usm.String) usm.Number {
	panic("not implemented")
}

//Symbol returns the byte at the given index in the String.
func (t *Target) Symbol(_ usm.String, _ usm.Number) usm.Number {
	panic("not implemented")
}

//Concat creates a new String that is the concatenation of the given strings.
func (t *Target) Concat(a usm.String, b usm.String) usm.String {
	panic("not implemented")
}

//Follow returns the value that the pointer is pointing at.
func (t *Target) Follow(_ usm.Pointer) usm.Value {
	panic("not implemented")
}

//Open returns a stream from the given platform-dependent URI.
//This may throw an error.
func (t *Target) Open(_ usm.String) usm.Stream {
	panic("not implemented")
}

//Stat performs a platform-dependent stat on the stream and returns the result.
func (t *Target) Stat(_ usm.Stream) usm.String {
	panic("not implemented")
}

//Read reads stream data into the given string, returns the number of bytes read.
//This may throw an error.
func (t *Target) Read(_ usm.Stream, _ usm.String) usm.Value {
	panic("not implemented")
}

//Send writes the string data into the stream, returns the number of bytes written.
//This may throw an error.
func (t *Target) Send(_ usm.Stream, _ usm.String) usm.Value {
	panic("not implemented")
}

//Add returns the sum of a and b.
func (t *Target) Add(a usm.Number, b usm.Number) usm.Number {
	panic("not implemented")
}

//Mul returns the product of a and b.
func (t *Target) Mul(a usm.Number, b usm.Number) usm.Number {
	panic("not implemented")
}

//Sub returns the difference between a and b.
func (t *Target) Sub(a usm.Number, b usm.Number) usm.Number {
	panic("not implemented")
}

//Div returns the quotient of a and b.
func (t *Target) Div(a usm.Number, b usm.Number) usm.Number {
	panic("not implemented")
}

//Mod returns the modulos of a and b. Must mimic Go % operator.
func (t *Target) Mod(a usm.Number, b usm.Number) usm.Number {
	panic("not implemented")
}

//Pow returns a to the power of b.
func (t *Target) Pow(a usm.Number, b usm.Number) usm.Number {
	panic("not implemented")
}

//Less returns 1 if a is smaller than b, otherwise 0.
func (t *Target) Less(a usm.Number, b usm.Number) usm.Bit {
	panic("not implemented")
}

//More returns 1 if a is larger than b, otherwise 0.
func (t *Target) More(a usm.Number, b usm.Number) usm.Bit {
	panic("not implemented")
}

//Same returns 1 if a is equal to b, otherwise 0.
func (t *Target) Same(a usm.Number, b usm.Number) usm.Bit {
	panic("not implemented")
}

//And returns a && b
func (t *Target) And(a usm.Bit, b usm.Bit) usm.Bit {
	panic("not implemented")
}

//Or returns a || b
func (t *Target) Or(a usm.Bit, b usm.Bit) usm.Bit {
	panic("not implemented")
}

//Not returns !Bit
func (t *Target) Not(_ usm.Bit) usm.Bit {
	panic("not implemented")
}

//Range creates a loop that runs the iterator from 'from' to 'to'
//under the relationship constraint with a given step.
//Relationship -2: <, -1:<=, 0: =, 1: >=, 2: >
func (t *Target) Range(from usm.Number, relationship int, to usm.Number, step usm.Number,
	body func(i usm.Number)) {
	panic("not implemented")
}

//Errors returns the number of errors on the thread-local error stack.
func (t *Target) Errors() usm.Number {
	panic("not implemented")
}

//Native returns b as a native.
func (t *Target) Native(b []byte) usm.Native {
	return b
}

//Writer returns a writer to the target.
func (t *Target) Writer() *bytes.Buffer {
	return &t.Buffer
}
