//+build cgo

package cbuf

//#include <stdio.h>
import "C"
import "unsafe"

//C.puts for test
func cPuts(p unsafe.Pointer) {
	C.puts((*C.char)(p))
}