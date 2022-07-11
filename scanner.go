package JsonStruct

import "bufio"

type ScannerKind byte

const (
	ScanUnknown ScannerKind = iota
	ScanNull
	ScanFalse
	ScanTrue
	ScanIntNumber
	ScanFloatNumber
	ScanTime
	ScanString
	ScanObject
	ScanArray
	ScanKey
	ScanValue
	ScanValueLast
)

func (k *ScannerKind) String() string {
	switch *k {
	case ScanNull:
		return "ScanNull"
	case ScanFalse:
		return "ScanFalse"
	case ScanTrue:
		return "ScanTrue"
	case ScanIntNumber:
		return "ScanIntNumber"
	case ScanFloatNumber:
		return "ScanFloatNumber"
	case ScanTime:
		return "ScanTime"
	case ScanString:
		return "ScanString"
	case ScanObject:
		return "ScanObject"
	case ScanArray:
		return "ScanArray"
	case ScanKey:
		return "ScanKey"
	case ScanValue:
		return "ScanValue"
	case ScanValueLast:
		return "ScanValueLast"
	default:
		return "ScanUnknown"
	}
}

type JStructScanner interface {
	Next() error
	Value() []byte
	Kind() ScannerKind
}

func NewJStructScanner(rd *bufio.Reader) JStructScanner {
	return NewJStructScannerWithSize(rd, JStructScannerBufferSize)
}
func NewJStructScannerWithSize(rd *bufio.Reader, initSize int) JStructScanner {
	return &JStructScannerImpl{rd: rd, buf: make([]byte, 0, initSize)}
}

type JStructScannerImpl struct {
	rd     *bufio.Reader
	buf    []byte
	scType ScannerKind
}

func (j *JStructScannerImpl) Next() error {
	j.scType = ScanUnknown
	j.buf = j.buf[:0]
	for {
		b, err := j.rd.ReadByte()
		if err != nil {
			return err
		}
		switch b {
		//case ' ', '\t', '\n', '\r':
		//	continue
		//case '{':
		//	j.scType = ScanObject
		//	return nil
		//case '[':
		//	j.scType = ScanArray
		//	return nil
		case 'n':
			j.scType = ScanNull
			j.buf = append(j.buf, b)
			return j.ReadNBytes(3)
		case 'f':
			j.scType = ScanFalse
			j.buf = append(j.buf, b)
			return j.ReadNBytes(4)
		case 't':
			j.scType = ScanTrue
			j.buf = append(j.buf, b)
			return j.ReadNBytes(3)
			return nil
		case '-':
			j.scType = ScanIntNumber
			return nil
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			j.scType = ScanIntNumber
			return nil
		case '.':
			j.scType = ScanFloatNumber
			return nil
		case '"':
			return j.ReadStringTime()
		//case ':':
		//	j.scType = ScanKey
		//	return nil
		//case ',':
		//	j.scType = ScanValue
		//	return nil
		//case ']':
		//	j.scType = ScanValueLast
		//	return nil
		//case '}':
		//	j.scType = ScanValueLast
		//	return nil
		default:
			return InvalidJsonError
		}
	}
	return nil
}

func (j *JStructScannerImpl) ReadNBytes(n int) error {
	for i := 0; i < n; i++ {
		b, err := j.rd.ReadByte()
		if err != nil {
			return err
		}
		j.buf = append(j.buf, b)
	}
	return nil
}

// Read json string or time in RFC3339 format
func (j *JStructScannerImpl) ReadStringTime() error {
	err := j.ReadString()
	if err != nil {
		return err
	}
	l := len(j.buf)
	if l > 20 { // RFC3339 format
		if j.buf[4] == '-' && j.buf[7] == '-' && j.buf[10] == 'T' &&
			j.buf[13] == ':' && j.buf[16] == ':' {
			b19 := j.buf[19]
			if b19 == 'Z' || b19 == '+' || b19 == '-' || b19 == '.' {
				j.scType = ScanTime
			}
		}
	}
	return nil
}

func (j *JStructScannerImpl) ReadString() error {
	j.scType = ScanString
	for {
		b, err := j.rd.ReadByte()
		if err != nil {
			return err
		}
		if b == 0x22 {
			last := len(j.buf) - 1
			if j.buf[last] != 0x5c { // "\"
				return nil
			}
		}
		j.buf = append(j.buf, b)
	}
	return nil
}

func (j *JStructScannerImpl) Value() []byte {
	return j.buf
}

func (j *JStructScannerImpl) Kind() ScannerKind {
	return j.scType
}
