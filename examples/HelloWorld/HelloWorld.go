package main

import (
	"bytes"
	"fmt"

	"github.com/qlova/usm/target/bytecode"
	"github.com/qlova/usm/target/golang"
	"github.com/qlova/usm/target/runtime"
)

func main() {
	var c bytecode.Target
	c.Main(func() {
		hello := c.Define(0, func() {
			c.Discard(c.Send(nil, c.String("Hello World\n")))
		})
		c.JumpTo(hello)
	})
	var buffer bytes.Buffer
	c.WriteTo(&buffer)

	fmt.Println("Bytecode: ", buffer.Bytes())

	var r runtime.Target
	bytecode.NewReader(bytes.NewReader(buffer.Bytes())).Target(&r)

	fmt.Println("\nRuntime: ")
	if err := r.Run(); err != nil {
		fmt.Println(err)
	}

	var t golang.Target
	bytecode.NewReader(bytes.NewReader(buffer.Bytes())).Target(&t)

	//t.WriteTo(os.Stdout)

	fmt.Println("\nGo: ")
	if err := t.Run(); err != nil {
		fmt.Println(err)
	}
}
