package purebuf

import "sync"

//buffer pool
type Pool struct {
	pool sync.Pool
}

var DefaultPool Pool

//get buffer from pool or create new buffer and bind to pool.
func (p *Pool) Get() *Buffer {
	v := p.pool.Get()
	if v != nil {
		return v.(*Buffer)
	}
	return new(Buffer)
}

//put buffer to pool
func (p *Pool) Put(buf *Buffer) {
	if buf.Cap() != 0 {
		buf.Reset()
		p.pool.Put(buf)
	}
}