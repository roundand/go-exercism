// package circular provides functions to create and operate circular buffers.
package circular

import "errors"

const testVersion = 3

// Buffer implements a circular buffer.
type Buffer struct {
	b []byte // underlying linear buffer
	c int    // data items counts
	w int    // write pointer
	r int    // read pointer
}

// NewBuffer returns an initialised circular Buffer.
func NewBuffer(size int) *Buffer {
	n := Buffer{}
	n.b = make([]byte, size)
	return &n
}

// ReadByte reads one byte from a circular Buffer, and returns that byte, or an error if empty.
func (buff *Buffer) ReadByte() (byte, error) {
//  buff := *buffp
	if buff.c == 0 {
		return '0', errors.New("can't read from empty buffer")
	}
	buff.c -= 1
	ret := buff.b[buff.r]
	buff.r = (buff.r + 1) % len(buff.b)
	return ret, nil
}

// WriteByte writes one byte to a circular Buffer or returns an error.
func (buff *Buffer) WriteByte(c byte) error {
	if buff.c == len(buff.b) {
		return errors.New("can't write to full buffer")
	}
	buff.c += 1
	buff.b[buff.w] = c
	buff.w = (buff.w + 1) % len(buff.b)
  return nil
}

// Overwrite writes one byte to the buffer, clearing one character out if necessary.
func (buff *Buffer) Overwrite(c byte) {
	// ifbuffer is full, read a byte
	if buff.c == len(buff.b) {
		_, _ = buff.ReadByte()
	}
	// buffer is now guaranteed not to be full (in a single-threaded context)
	_ = buff.WriteByte(c)
}

// Reset puts the buffer back to an empty state.
func (buff *Buffer) Reset() {
	// only fresh content can be read once size and pointers are zeroed
	buff.c, buff.w, buff.r = 0, 0, 0
}
