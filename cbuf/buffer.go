//+build cgo

package cbuf

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"errors"
	"github.com/pure-project/purebuf/internal"
	"io"
	"reflect"
	"unsafe"
)

//write only byte buffer for cgo
type Buffer struct {
	ptr   unsafe.Pointer
	size  int
	space int
}

const minReadSize = 512

//alloc memory failed
var ErrBadAlloc = errors.New("bad alloc")

//free c memory
func (buf *Buffer) Free() {
	if buf.ptr != nil {
		C.free(buf.ptr)
		buf.ptr = nil
	}
}

//length
func (buf *Buffer) Len() int {
	return buf.size
}

//capacity
func (buf *Buffer) Cap() int {
	return buf.space
}

//grow to target capacity
func (buf *Buffer) Grow(size int) {
	if buf.space < size {
		grown := buf.space * 2
		if grown < size {
			grown = size
		}

		buf.ptr = C.realloc(buf.ptr, C.size_t(grown))
		if buf.ptr == nil {
			panic(ErrBadAlloc)
		}
		buf.space = grown
	}
}

//reset size
func (buf *Buffer) Reset() {
	buf.size = 0
}

//resize to target size
func (buf *Buffer) Resize(size int) {
	buf.Grow(size)
	buf.size = size
}

//read from reader until EOF
func (buf *Buffer) ReadFrom(reader io.Reader) (int64, error) {
	var size int64
	for {
		buf.Grow(buf.size + minReadSize)
		n, err := reader.Read(buf.allBytesOffset(buf.size))
		if n > 0 {
			size += int64(n)
			buf.size += n
		}
		
		if err == io.EOF {
			return size, nil
		}
		
		if err != nil {
			return size, err
		}
	}
	return size, nil
}

func (buf *Buffer) Write(b []byte) (int, error) {
	hdr := internal.SliceHeader(b)
	if hdr.Len != 0 {
		buf.Grow(buf.size + hdr.Len)
		C.memcpy(unsafe.Pointer(uintptr(buf.ptr) + uintptr(buf.size)), unsafe.Pointer(hdr.Data), C.size_t(hdr.Len))
		buf.size += hdr.Len
	}
	return hdr.Len, nil
}

func (buf *Buffer) WriteString(s string) (int, error) {
	return buf.Write(internal.S2B(s))
}

func (buf *Buffer) WriteByte(b byte) error {
	buf.Grow(buf.size + 1)
	C.memset(unsafe.Pointer(uintptr(buf.ptr) + uintptr(buf.size)), C.int(b), 1)
	buf.size++
	return nil
}

//c memory pointer
func (buf *Buffer) Pointer() unsafe.Pointer {
	return buf.ptr
}

//unsafe go slice
func (buf *Buffer) TempBytes() []byte {
	return *(*[]byte)(unsafe.Pointer(buf))
}

//unsafe go string
func (buf *Buffer) TempString() string {
	return *(*string)(unsafe.Pointer(buf))
}

//copy go slice
func (buf *Buffer) Bytes() []byte {
	return []byte(buf.TempString())
}

//copy go string
func (buf *Buffer) String() string {
	return string(buf.TempBytes())
}

func (buf *Buffer) allBytesOffset(offset int) []byte {
	l := buf.space - offset
	hdr := &reflect.SliceHeader{
		Data: uintptr(buf.ptr) + uintptr(offset),
		Len:  l,
		Cap:  l,
	}
	return *(*[]byte)(unsafe.Pointer(hdr))
}
