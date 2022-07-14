package JsonStruct

import (
	"io"
)

var (
	// Buffers
	JStructReaderBufferSize       = 4096
	JSStructReaderBufferThreshold = JStructReaderBufferSize >> 2
)

type JStructReader interface {
	Current() byte                     // show from buffer
	Index() int                        // position in file = index + total
	Release() []byte                   // copy to data and release
	Offset(n int) error                // move pointer, if required copy to data
	Next() error                       // offset with 1
	Total() int                        // size of released bytes
	Buffer() []byte                    // current buffer
	FillBuffFrom(idx int) (int, error) // fill buffer from idx
}

func NewJStructReader(rd io.Reader) JStructReader {
	return NewJStructReaderWithSize(rd, JStructReaderBufferSize)
}

func NewJStructReaderWithSize(rd io.Reader, size int) JStructReader {
	return &JStructReaderImpl{
		rd:       rd,
		buf:      make([]byte, size),
		ptr:      -1,
		idx:      -1,
		size:     size,
		finished: false,
	}
}

type JStructReaderImpl struct {
	rd    io.Reader
	buf   []byte // data buffer
	size  int    // buffer size
	start int    // start window position in buffer
	ptr   int    // current position in buffer
	//=======
	idx      int // index of last read byte
	total    int // total bytes read
	finished bool
}

func (j *JStructReaderImpl) Buffer() []byte {
	return j.buf
}

func (j *JStructReaderImpl) Current() byte {
	return j.buf[j.ptr]
}

func (j *JStructReaderImpl) Index() int {
	return j.idx
}

func (j *JStructReaderImpl) Release() []byte {
	idx := j.ptr + 1
	res := j.buf[j.start:idx]
	j.total += idx - j.start
	j.start = idx
	return res
}

func (j *JStructReaderImpl) Next() error {
	return j.Offset(1)
}

func (j *JStructReaderImpl) Offset(n int) error {
	j.idx += n

	var e error
	var rn int

	if j.ptr == -1 {
		rn, e = j.FillBuffFrom(0)
		if n > rn {
			n -= rn
			j.ptr = rn - 1
		}
	}

	// Inside current buffer
loop:

	if j.finished {
		j.ptr = j.size - 1
		j.idx -= n
		n = 0
		return io.EOF
	}

	availableSpace := j.size - j.ptr - 1 // tail

	if availableSpace >= n {
		j.ptr += n
		return e
	}

	// Outside current buffer
	// but still inside full buffer size
	availableSpace += j.start // tail + head
	// move to front data if not there
	if j.start != 0 {
		j.moveDataToFront()
		n -= j.ptr + 1
	}

	// expand by default,
	// skip if free space larger than
	// threshold or offset value
	if availableSpace < max(n, JSStructReaderBufferThreshold) {
		j.expand()
	}
	idx := j.ptr + 1
	rn, e = j.FillBuffFrom(idx)

	if rn > 0 && rn < n {
		n -= rn
		j.ptr = j.size - 1
	}
	goto loop
}

func (j *JStructReaderImpl) Total() int {
	return j.total
}

func (j *JStructReaderImpl) FillBuffFrom(idx int) (int, error) {
	b := j.buf[idx:]
	l := len(b)
	rn, err := j.rd.Read(b)
	if rn < l {
		j.buf = j.buf[:idx+rn]
		j.size = len(j.buf)
	}
	if err == io.EOF {
		j.finished = true
	}
	return rn, err
}

func (j *JStructReaderImpl) moveDataToFront() {
	j.ptr = j.size - 1
	copy(j.buf, j.buf[j.start:])
	j.ptr = j.ptr - j.start
	j.start = 0
}

func (j *JStructReaderImpl) expand() {
	j.buf = append(j.buf, make([]byte, j.size)...)
	j.size *= 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
