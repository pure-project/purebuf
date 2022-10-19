package main

import "C"
import (
	"fmt"
	"github.com/pure-project/purebuf"
	"github.com/pure-project/purebuf/cbuf"
	"github.com/pure-project/purebuf/beta"
	"strings"
)

func simple() {
	var buf purebuf.Buffer
	buf.Write([]byte("hello "))
	buf.WriteString("pure-project")
	buf.WriteByte('!')

	fmt.Printf("buf.Data: %v\n", buf.Data)
	fmt.Printf("buf.Bytes(): %v\n", buf.Bytes())

	//temp string. unsafe but fast, for just-read-only
	fmt.Printf("buf.TempString(): %s\n", buf.TempString())

	//deep-copy string. safe buf slow.
	fmt.Printf("buf.String(): %s\n", buf.String())

	//clear
	buf.Reset()

	//read from reader
	buf.ReadFrom(strings.NewReader("happy to use"))
}

func pool() {
	var pool purebuf.Pool

	//get from pool and do not forget put back
	buf := pool.Get()
	defer pool.Put(buf)

	//same as simple use...
}

func cgo() {
	var buf cbuf.Buffer

	//do not forget free the c memory
	defer buf.Free()

	//pre-alloc memory
	buf.Grow(1024)

	buf.Write([]byte("hello "))
	buf.WriteString("pure-project!")

	//for c string
	buf.WriteByte(0)

	//cgo call
	C.puts((*C.char)(buf.Pointer()))

	//non-copy bytes, attention for use
	buf.TempBytes()

	//non-copy string, attention for use
	buf.TempString()

	//deep-copy bytes
	buf.Bytes()

	//deep-copy string
	buf.String()
}

func useBeta() {
	var pool beta.Pool

	buf := pool.Get()
	defer pool.Put(buf)

	//same as purebuf.Buffer
	//but it can be automatically recycle memory to the pool when re-alloc space
	//it maybe better for performance
}