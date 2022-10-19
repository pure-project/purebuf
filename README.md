# pure-buf

**A pure write-only byte buffer library for go.**

[中文](doc/README_cn.md)



### Overview

purebuf is a write-only  byte buffer library for golang.

It is fully production-ready.



### Features

1. simple write-only byte buffer (and with pool)
2. byte buffer for cgo (cbuf.Buffer)
3. automatic recycle unused memory byte buffer (beta.Buffer)



### Usage

imports:

```go
import (
    "github.com/pure-project/purebuf"       //purebuf.Buffer purebuf.Pool
    "github.com/pure-project/purebuf/cbuf"  //cbuf.Buffer
    "github.com/pure-project/purebuf/beta"  //beta.Buffer beta.Pool
)
```



simple:

```go
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
```



pool:

```go
var pool purebuf.Pool

//get from pool and put it back
buf := pool.Get()
defer pool.Put(buf)

//same as simple use...
```



cgo:

```go
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
```



beta:

```go
var pool beta.Pool

buf := pool.Get()
defer pool.Put(buf)

//same as purebuf.Buffer
//but it can be automatically recycle memory to the pool when re-alloc space
//it maybe better for performance :p
```



### Licence

MIT Licence

Copyright (c) 2022 pure-project team.