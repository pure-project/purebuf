package beta

import (
	"github.com/pure-project/purebuf/internal"
	"io"
)

//write only byte buffer
//
//automatic recycle unused memories to binding Pool
type Buffer struct {
	Data []byte
	pool *Pool
}

const minReadSize = 512

//length
func (buf *Buffer) Len() int {
	return len(buf.Data)
}

//capacity
func (buf *Buffer) Cap() int {
	return cap(buf.Data)
}

//grow to target capacity
func (buf *Buffer) Grow(size int) {
	if c := cap(buf.Data); c < size {
		grown := c * 2
		if grown < size {
			grown = size
		}

		data := buf.Data
		buf.Data = make([]byte, len(data), grown)
		if len(data) != 0 {
			copy(buf.Data, data)
		}

		//recycle memory
		if c != 0 && buf.pool != nil {
			buf.pool.Put(&Buffer{ Data: data })
		}
	}
}

//reset buffer size
func (buf *Buffer) Reset() {
	buf.Data = buf.Data[:0]
}

//resize to target length
func (buf *Buffer) Resize(size int) {
	buf.Grow(size)
	buf.Data = buf.Data[:size]
}

//read from reader until EOF
func (buf *Buffer) ReadFrom(reader io.Reader) (int64, error) {
	var size int64
	for {
		l := len(buf.Data)
		buf.Grow(l + minReadSize)
		n, err := reader.Read(buf.Data[l:cap(buf.Data)])
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
		size += int64(n)
		buf.Data = buf.Data[:l + n]
	}
	return size, nil
}

func (buf *Buffer) Write(b []byte) (int, error) {
	data := append(buf.Data, b...)
	if internal.SliceHeader(data).Data != internal.SliceHeader(buf.Data).Data && cap(buf.Data) != 0 && buf.pool != nil {
		buf.pool.Put(&Buffer{ Data: buf.Data })
	}
	buf.Data = data
	return len(b), nil
}

func (buf *Buffer) WriteString(s string) (int, error) {
	return buf.Write(internal.S2B(s))
}

func (buf *Buffer) WriteByte(b byte) error {
	data := append(buf.Data, b)
	if internal.SliceHeader(data).Data != internal.SliceHeader(buf.Data).Data && cap(buf.Data) != 0 && buf.pool != nil {
		buf.pool.Put(&Buffer{ Data: buf.Data })
	}
	buf.Data = data
	return nil
}

func (buf *Buffer) Bytes() []byte {
	return buf.Data
}

//copy to string
func (buf *Buffer) String() string {
	return string(buf.Data)
}

//temp string (unsafe)
func (buf *Buffer) TempString() string {
	return internal.B2S(buf.Data)
}

