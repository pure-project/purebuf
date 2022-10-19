package cbuf

import (
	"math/rand"
	"testing"
)

//func TestCBuffer_cgo(t *testing.T) {
//	var buf CBuffer
//	defer buf.Free()
//
//	buf.WriteString("this string can be print by cgo.")
//	buf.WriteByte(0)
//
//	cPuts(buf.Pointer())
//}

func TestCBuffer_Bytes(t *testing.T) {
	var buf CBuffer
	defer buf.Free()

	buf.WriteByte(0)
	buf.WriteByte(1)
	buf.WriteByte(2)

	t.Logf("buf.Bytes(): %v", buf.Bytes())
}

func TestCBuffer_String(t *testing.T) {
	var buf CBuffer
	defer buf.Free()

	buf.WriteByte('a')
	buf.WriteByte('b')
	buf.WriteByte('c')

	t.Logf("buf.String(): %s", buf.String())
}

func TestCBuffer_Free(t *testing.T) {
	const Bytes = 32 * 1024  //32KB

	var data [Bytes]byte
	for {
		var buf CBuffer
		buf.Write(data[:rand.Intn(Bytes)])
		buf.Free()
	}
}