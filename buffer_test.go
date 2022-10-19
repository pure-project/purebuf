package purebuf

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
)

func TestBuffer_ReadFrom(t *testing.T) {
	data := make([]byte, 4096)
	var buf Buffer
	n, err := buf.ReadFrom(bytes.NewReader(data))
	fmt.Println(n, err)
}

func BenchmarkBuffer_Pool(b *testing.B) {
	const Bytes = 32 * 1024   //32KB
	data := make([]byte, Bytes)
	var pool Pool
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := pool.Get()
			for i := 0; i < 10; i++ {
				buf.Write(data[:rand.Intn(Bytes)])
			}
			pool.Put(buf)
		}
	})
}