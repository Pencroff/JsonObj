package JsonStruct

import "io"

type JStructScanner interface {
	Peek() []byte
	Current() byte
	Index() int
	Release()
	Offset(n int) error
	Next() error
	Total() int
}

func NewJStructScanner(rd io.Reader) JStructScanner {
	return NewJStructScannerWithSize(rd, JStructScannerBufferSize)
}

func NewJStructScannerWithSize(rd io.Reader, size int) JStructScanner {
	return &JStructScannerImpl{
		rd:  rd,
		buf: make([]byte, size),
	}
}

type JStructScannerImpl struct {
	rd    io.Reader
	buf   []byte
	ptr   uint64
	total uint64
}

func (j *JStructScannerImpl) Peek() []byte {
	return j.buf[:j.ptr]
}

func (j *JStructScannerImpl) Current() byte {
	return j.buf[j.ptr]
}

func (j *JStructScannerImpl) Index() int {
	return int(j.ptr + j.total)
}

func (j *JStructScannerImpl) Release() {
	j.buf = j.buf[j.ptr:]
}

func (j *JStructScannerImpl) Next() error {
	return j.Offset(1)
}

func (j *JStructScannerImpl) Offset(n int) error {
	l := len(j.buf)
	if n >= l {
		tmp := j.buf[:]
		j.buf = make([]byte, cap(j.buf))
		rn, err := j.rd.Read(j.buf)
		if rn > 0 {
			j.buf = append(tmp, j.buf[:rn]...)
		}
		if err != nil {
			return err
		}
	} else {
		j.ptr += uint64(n)
	}
	return nil
}

func (j *JStructScannerImpl) Total() int {
	return int(j.total)
}
