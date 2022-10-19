package internal

import (
	"reflect"
	"unsafe"
)

func SliceHeader(b []byte) reflect.SliceHeader {
	return *(*reflect.SliceHeader)(unsafe.Pointer(&b))
}

func S2B(s string) []byte {
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: hdr.Data,
		Len:  hdr.Len,
		Cap:  hdr.Len,
	}))
}

func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
