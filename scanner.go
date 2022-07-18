package JsonStruct

import (
	"io"
)

var (
	// Buffers
	JStructScannerBufferSize       = 4 * 1024
	JSStructScannerBufferThreshold = JStructScannerBufferSize >> 3
)

type JStructScanner interface {
	Current() byte    // show from buffer
	Bytes() []byte    // copy to data and release
	Next() error      // scan 1 byte
	Scan(n int) error // move pointer, if required read data

	Index() int                        // position in file
	Buffer() []byte                    // current buffer
	Window() (int, int)                // window of current buffer (start, end)
	FillBuffFrom(idx int) (int, error) // fill buffer from idx, should be in interval [start, size)
}

func NewJStructScanner(rd io.Reader) JStructScanner {
	return NewJStructScannerWithParam(rd, JStructScannerBufferSize, JSStructScannerBufferThreshold)
}

func NewJStructScannerWithParam(rd io.Reader, size, threshold int) JStructScanner {
	return &JStructScannerImpl{
		rd:        rd,
		buf:       make([]byte, size),
		ptr:       -1,
		idx:       -1,
		size:      size,
		threshold: threshold,
		finished:  false,
	}
}

type JStructScannerImpl struct {
	rd  io.Reader
	buf []byte // data buffer

	start int // start window position in buffer
	ptr   int // current position in buffer
	//=======
	size      int // buffer size
	threshold int // buffer threshold
	//=======
	idx      int // index of last read byte
	finished bool
	released bool
}

// Buffer returns current buffer
func (j *JStructScannerImpl) Buffer() []byte {
	return j.buf
}

// Window returns window of current buffer (start, ptr)
func (j *JStructScannerImpl) Window() (int, int) {
	return j.start, j.ptr
}

// Current returns pointed byte
func (j *JStructScannerImpl) Current() byte {
	if j.size == 0 || j.ptr < 0 || j.ptr < j.start {
		return 0
	}
	return j.buf[j.ptr]
}

// Index returns current position in file
func (j *JStructScannerImpl) Index() int {
	return j.idx
}

// Bytes returns data in window (start, ptr)
func (j *JStructScannerImpl) Bytes() []byte {
	idx := j.ptr + 1
	res := j.buf[j.start:idx]
	j.start = idx
	return res
}

func (j *JStructScannerImpl) Next() error {
	return j.Scan(1)
}

func (j *JStructScannerImpl) Scan(n int) error {

	var e error
	var rn int

	if j.finished {
		j.ptr = j.size - 1
		return io.EOF
	}

	if j.ptr == -1 {
		rn, e = j.FillBuffFrom(0)
		if n > rn {
			n -= rn
			j.ptr = rn - 1
			j.idx = j.ptr
		}
	}

	idx := j.ptr + 1
	tailSpace := j.size - idx // tail

	// Count rest of available buffer if required
	if n > tailSpace {
		n -= tailSpace
		j.idx += tailSpace
		j.ptr += tailSpace
		tailSpace = 0
		idx = j.size
	}

	goto firstLoop

loop:
	idx = j.ptr + 1
	tailSpace = j.size - idx // tail

firstLoop:
	// Inside current buffer
	if n <= tailSpace {
		j.ptr += n
		j.idx += n
		return e
	}

	// Outside of tail
	// but still might in buffer

	// move to front data if not there
	if j.start != 0 {
		j.moveDataToFront()
		// update idx by moved values
		idx = j.ptr + 1
	}

	// expand by default,
	// skip if free space larger than
	// threshold or offset value
	availableSpace := tailSpace + j.start // tail + head
	if availableSpace < max(n, j.threshold) {
		j.expand()
	}

	rn, e = j.FillBuffFrom(idx)

	if n > rn {
		n -= rn
		j.ptr = j.size - 1
		j.idx += rn
		if j.finished {
			return io.EOF
		}
	}

	goto loop
}

func (j *JStructScannerImpl) FillBuffFrom(idx int) (int, error) {
	b := j.buf[idx:]
	l := j.size - idx
	rn, err := j.rd.Read(b)
	if rn < l {
		j.size = idx + rn
		j.buf = j.buf[:j.size]
	}
	if err == io.EOF {
		j.finished = true
	}
	return rn, err
}

func (j *JStructScannerImpl) moveDataToFront() {
	copy(j.buf, j.buf[j.start:])
	delta := j.size - j.ptr - 1
	j.idx += delta
	j.ptr = j.size - 1 - j.start
	j.start = 0
}

func (j *JStructScannerImpl) expand() {
	j.buf = append(j.buf, make([]byte, j.size)...)
	j.size *= 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
