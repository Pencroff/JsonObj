package JsonStruct

import (
	"io"
)

type TokenizerKind byte

const (
	TokenUnknown TokenizerKind = iota
	TokenNull
	TokenFalse
	TokenTrue
	TokenIntNumber
	TokenFloatNumber
	TokenTime
	TokenString
	TokenObject
	TokenArray
	TokenKey
	TokenValue
	TokenValueLast
)

var whiteSpaces = [256]bool{
	0x09: true, // tab
	0x0A: true, // line feed
	0x0D: true, // carriage return
	0x20: true, // space
}

var numberInt = [256]bool{
	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
}

var numberExponent = [256]bool{
	'-': true,
	'+': true,
	'e': true,
	'E': true,
}

var nullTokenData = []byte(`null`)
var falseTokenData = []byte(`false`)
var trueTokenData = []byte(`true`)

func (k *TokenizerKind) String() string {
	switch *k {
	case TokenNull:
		return "TokenNull"
	case TokenFalse:
		return "TokenFalse"
	case TokenTrue:
		return "TokenTrue"
	case TokenIntNumber:
		return "TokenIntNumber"
	case TokenFloatNumber:
		return "TokenFloatNumber"
	case TokenTime:
		return "TokenTime"
	case TokenString:
		return "TokenString"
	case TokenObject:
		return "TokenObject"
	case TokenArray:
		return "TokenArray"
	case TokenKey:
		return "TokenKey"
	case TokenValue:
		return "TokenValue"
	case TokenValueLast:
		return "TokenValueLast"
	default:
		return "TokenUnknown"
	}
}

type JStructTokenizer interface {
	Next() error
	Value() []byte
	Kind() TokenizerKind
}

func NewJSStructTokenizer(sc JStructScanner) JStructTokenizer {
	return &JStructTokenizerImpl{sc: sc}
}

type JStructTokenizerImpl struct {
	sc     JStructScanner
	scType TokenizerKind
	v      []byte
}

func (t *JStructTokenizerImpl) nextSkipWhiteSpace() error {
	for {
		err := t.sc.Next()
		if err != nil {
			return err
		}
		b := t.sc.Current()
		if whiteSpaces[b] {
			t.sc.Bytes()
			continue
		}
		return nil
	}
}

func (t *JStructTokenizerImpl) nextKeepWhiteSpace() error {
	for {
		err := t.sc.Next()
		if err != nil {
			return err
		}
		b := t.sc.Current()
		if whiteSpaces[b] {
			continue
		}
		return nil
	}
}

func (t *JStructTokenizerImpl) Next() error {
	t.scType = TokenUnknown
	for {
		err := t.nextSkipWhiteSpace()
		if err != nil {
			return InvalidJsonError{Err: err}
		}
		switch t.sc.Current() {
		//case '{':
		//	t.scType = TokenObject
		//	return nil
		//case '[':
		//	t.scType = TokenArray
		//	return nil
		case 'n':
			return t.ReadNull()
		case 'f':
			return t.ReadFalse()
		case 't':
			return t.ReadTrue()
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return t.ReadNumber(false)
		case '-':
			return t.ReadNumber(true)
		case '"':
			return t.ReadStringTime()
		//case ':':
		//	t.scType = TokenKey
		//	return nil
		//case ',':
		//	t.scType = TokenValue
		//	return nil
		//case ']':
		//	t.scType = TokenValueLast
		//	return nil
		//case '}':
		//	t.scType = TokenValueLast
		//	return nil
		default:
			return InvalidJsonError{Err: err}
		}
	}
	return nil
}

func (t *JStructTokenizerImpl) Value() []byte {
	return t.v
}

func (t *JStructTokenizerImpl) Kind() TokenizerKind {
	return t.scType
}

func (t *JStructTokenizerImpl) ReadNull() error {
	return t.hardcodedToken(TokenNull, nullTokenData)
}

func (t *JStructTokenizerImpl) ReadFalse() error {
	return t.hardcodedToken(TokenFalse, falseTokenData)
}

func (t *JStructTokenizerImpl) ReadTrue() error {
	return t.hardcodedToken(TokenTrue, trueTokenData)
}

func (t *JStructTokenizerImpl) ReadNumber(hasMinus bool) error {
	var e error
	t.scType = TokenIntNumber
	first := t.sc.Index()
	e = t.sc.Next()
	ch := t.sc.Current()
	hasIntPart := !hasMinus || numberInt[ch]
	if e != nil || !numberInt[ch] {
		goto afterLoop
	}
	// integer part
	for {
		e = t.sc.Next()
		ch = t.sc.Current()
		if !numberInt[ch] || e != nil {
			break
		}
	}
afterLoop:
	if e == io.EOF && numberInt[ch] {
		t.v = t.sc.Bytes()
		return nil
	}
	if hasIntPart {
		if ch == '.' {
			return t.ReadFractionPart(first)
		}
		if ch == 'e' || ch == 'E' {
			return t.ReadExponentPart(first)
		}
	} else {
		goto errLbl
	}

	if e == nil && whiteSpaces[ch] {
		idx := t.sc.Index()
		e = t.nextKeepWhiteSpace()
		if whiteSpaces[t.sc.Current()] {
			l := idx - first
			t.v = t.sc.Bytes()[:l]
			return nil
		}
	}

errLbl:
	t.scType = TokenUnknown
	t.v = t.sc.Bytes()
	return InvalidJsonPtrError{Pos: t.sc.Index(), Err: e}
}

func (t *JStructTokenizerImpl) ReadFractionPart(firstIdx int) error {
	var e error
	t.scType = TokenFloatNumber

	e = t.sc.Next()
	ch := t.sc.Current()
	hasFractionPart := numberInt[ch]
	if e != nil || !numberInt[ch] {
		goto afterLoop
	}
	for {
		e = t.sc.Next()
		ch = t.sc.Current()
		if !numberInt[ch] || e != nil {
			break
		}
	}
afterLoop:
	if e == io.EOF && numberInt[ch] {
		t.v = t.sc.Bytes()
		return nil
	}
	if hasFractionPart {
		if ch == 'e' || ch == 'E' {
			return t.ReadExponentPart(firstIdx)
		}
	} else {
		goto errLbl
	}

	if e == nil && whiteSpaces[ch] {
		idx := t.sc.Index()
		e = t.nextKeepWhiteSpace()
		ch = t.sc.Current()
		if whiteSpaces[ch] {
			l := idx - firstIdx
			t.v = t.sc.Bytes()[:l]
			return nil
		}
	}

errLbl:
	t.scType = TokenUnknown
	t.v = t.sc.Bytes()

	return InvalidJsonPtrError{Pos: t.sc.Index(), Err: e}
}

func (t *JStructTokenizerImpl) ReadExponentPart(first int) error {
	var e error
	t.scType = TokenFloatNumber
	hasExpNumPart := false
	e = t.sc.Next()
	ch := t.sc.Current()
	if ch == '+' || ch == '-' {
		e = t.sc.Next()
		ch = t.sc.Current()
	}
	hasExpNumPart = numberInt[ch]
	if e != nil || !numberInt[ch] {
		goto afterLoop
	}

	for {
		e = t.sc.Next()
		ch = t.sc.Current()
		if !numberInt[ch] || e != nil {
			break
		}
	}
afterLoop:
	if e == io.EOF && numberInt[ch] {
		t.v = t.sc.Bytes()
		return nil
	}
	if !hasExpNumPart {
		goto errLbl
	}
	if e == nil && whiteSpaces[ch] {
		idx := t.sc.Index()
		e = t.nextKeepWhiteSpace()
		ch = t.sc.Current()
		if whiteSpaces[ch] {
			l := idx - first
			t.v = t.sc.Bytes()[:l]
			return nil
		}
	}
errLbl:
	t.scType = TokenUnknown
	t.v = t.sc.Bytes()
	return InvalidJsonPtrError{Pos: t.sc.Index(), Err: e}
}

// Read json string or time in RFC3339 format
func (t *JStructTokenizerImpl) ReadStringTime() error {
	err := t.ReadString()
	if err != nil {
		return err
	}
	//l := len(t.buf)
	//if l > 20 { // RFC3339 format
	//	if t.buf[4] == '-' && t.buf[7] == '-' && t.buf[10] == 'T' &&
	//		t.buf[13] == ':' && t.buf[16] == ':' {
	//		b19 := t.buf[19]
	//		if b19 == 'Z' || b19 == '+' || b19 == '-' || b19 == '.' {
	//			t.scType = TokenTime
	//		}
	//	}
	//}
	return nil
}

func (t *JStructTokenizerImpl) ReadString() error {
	t.scType = TokenString
	//for {
	//	b, err := t.rd.ReadByte()
	//	if err != nil {
	//		return err
	//	}
	//	if b == 0x22 {
	//		last := len(t.buf) - 1
	//		if t.buf[last] != 0x5c { // "\"
	//			return nil
	//		}
	//	}
	//	t.buf = append(t.buf, b)
	//}
	return nil
}

func (t *JStructTokenizerImpl) hardcodedToken(kind TokenizerKind, origin []byte) error {
	t.scType = kind
	l := len(origin)
	idx := t.sc.Index()
	for i, b := range origin[1:] {
		e := t.sc.Next()

		if b != t.sc.Current() || e != nil {
			t.v = t.sc.Bytes()
			t.scType = TokenUnknown
			return InvalidJsonPtrError{Pos: idx + i + 1, Err: e}
		}
	}
	idx = t.sc.Index()
	// Check if it is a valid token ended by whitespaces
	e := t.nextKeepWhiteSpace()
	if idx == t.sc.Index() || whiteSpaces[t.sc.Current()] {
		t.v = t.sc.Bytes()[:l]
		return nil
	}
	t.scType = TokenUnknown
	t.v = t.sc.Bytes()
	return InvalidJsonPtrError{Pos: t.sc.Index(), Err: e}
}
