# pure-buf

**一个纯粹的go字节缓冲库。**





### 前言

pure-buf是一个简单的golang字节缓冲库，生产可用。



### 功能

1. 简单好用的写入字节缓冲区(包括用于复用的缓冲池)
2. 可用于cgo的字节缓冲区 (cbuf.Buffer)
3. 自动回收空闲内存的字节缓冲区 (beta.Buffer与对应用于复用的beta.Pool)



### 用法

导入包：

```go
import (
    "github.com/pure-project/purebuf"       //purebuf.Buffer purebuf.Pool
    "github.com/pure-project/purebuf/cbuf"  //cbuf.Buffer
    "github.com/pure-project/purebuf/beta"  //beta.Buffer beta.Pool
)
```



简单用法：

```go
var buf purebuf.Buffer
buf.Write([]byte("hello "))
buf.WriteString("pure-project")
buf.WriteByte('!')

fmt.Printf("buf.Data: %v\n", buf.Data)
fmt.Printf("buf.Bytes(): %v\n", buf.Bytes())

//临时字符串 速度快但不安全 仅用于临时使用
fmt.Printf("buf.TempString(): %s\n", buf.TempString())

//深拷贝字符串 安全可用
fmt.Printf("buf.String(): %s\n", buf.String())

//清空缓冲区
buf.Reset()

//从io.Reader读取
buf.ReadFrom(strings.NewReader("happy to use"))
```



池化复用:

```go
var pool purebuf.Pool

//从池中获取缓冲区，并且在用完后放回去
buf := pool.Get()
defer pool.Put(buf)

//其余与简单用法相同
```



用于cgo:

```go
var buf cbuf.Buffer

//记得释放C内存
defer buf.Free()

//预分配内存
buf.Grow(1024)

buf.Write([]byte("hello "))
buf.WriteString("pure-project!")

//0结尾C字符串
buf.WriteByte(0)

//用于调用C函数
C.puts((*C.char)(buf.Pointer()))

//临时字节切片，速度快但不安全，仅临时使用
buf.TempBytes()

//临时字符串，速度快但不安全，仅临时使用
buf.TempString()

//深拷贝字节切片，安全可用
buf.Bytes()

//深拷贝字符串，安全可用
buf.String()
```



### 许可

MIT许可证

版权所有 (c) 2022 pure-project团队。