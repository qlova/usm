package golang

//Runtime is the Go `u` runtime.
const Runtime = `package main

import (
	"os"
	"reflect"
	"unsafe"
)

type Value struct {
	Pointer unsafe.Pointer
	int64
}

func (v Value) Bytes() []byte {
	var slice []byte
	var header = (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	header.Data = uintptr(v.Pointer)
	header.Len = int(v.int64)
	header.Cap = int(v.int64)
	return slice
}

type Runtime struct {
	Errors []Value
}

func (r *Runtime) Stdin(s Value) Value {
	n, err := os.Stdin.Read(s.Bytes())
	if err != nil {
		r.Errors = append(r.Errors, r.String(err.Error()))
	}
	return Value{
		int64: int64(n),
	}
}

func (r *Runtime) Stdout(s Value) Value {
	n, err := os.Stdout.Write(s.Bytes())
	if err != nil {
		r.Errors = append(r.Errors, r.String(err.Error()))
	}
	return Value{
		int64: int64(n),
	}
}

func (Runtime) String(s string) (v Value) {
	var slice = []byte(s)
	var header = (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	v.Pointer = unsafe.Pointer(header.Data)
	v.int64 = int64(header.Len)
	return
}
`
