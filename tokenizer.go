package JsonStruct

import (
	"fmt"
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

const (
	minTimeLen = 22
	maxTimeLen = 35
)

var spaceCh = [256]bool{
	0x09: true, // tab
	0x0A: true, // line feed
	0x0D: true, // carriage return
	0x20: true, // space
}

var hexCh = [256]bool{
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
	'a': true,
	'b': true,
	'c': true,
	'd': true,
	'e': true,
	'f': true,
	'A': true,
	'B': true,
	'C': true,
	'D': true,
	'E': true,
	'F': true,
}

var numCh = [256]bool{
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

var exponentCh = [256]bool{
	'-': true,
	'+': true,
	'e': true,
	'E': true,
}

var nullTokenData = []byte(`null`)
var falseTokenData = []byte(`false`)
var trueTokenData = []byte(`true`)

const (
	backSlashCh = 0x5c // '\'
	quoteCh     = 0x22 // '"'
)

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
		if spaceCh[b] {
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
		if spaceCh[b] {
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
	t.scType = TokenIntNumber
	first := t.sc.Index()
	e := t.sc.Next()
	ch := t.sc.Current()
	hasIntPart := !hasMinus || numCh[ch]
	if e != nil || !numCh[ch] {
		goto afterLoop
	}
	// integer part
	for {
		e = t.sc.Next()
		ch = t.sc.Current()
		if !numCh[ch] || e != nil {
			break
		}
	}
afterLoop:
	if e == io.EOF && numCh[ch] {
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

	if e == nil && spaceCh[ch] {
		idx := t.sc.Index()
		e = t.nextKeepWhiteSpace()
		if spaceCh[t.sc.Current()] {
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
	t.scType = TokenFloatNumber
	e := t.sc.Next()
	ch := t.sc.Current()
	hasFractionPart := numCh[ch]
	if e != nil || !numCh[ch] {
		goto afterLoop
	}
	for {
		e = t.sc.Next()
		ch = t.sc.Current()
		if !numCh[ch] || e != nil {
			break
		}
	}
afterLoop:
	if e == io.EOF && numCh[ch] {
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

	if e == nil && spaceCh[ch] {
		idx := t.sc.Index()
		e = t.nextKeepWhiteSpace()
		ch = t.sc.Current()
		if spaceCh[ch] {
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
	t.scType = TokenFloatNumber
	hasExpNumPart := false
	e := t.sc.Next()
	ch := t.sc.Current()
	if ch == '+' || ch == '-' {
		e = t.sc.Next()
		ch = t.sc.Current()
	}
	hasExpNumPart = numCh[ch]
	if e != nil || !numCh[ch] {
		goto afterLoop
	}

	for {
		e = t.sc.Next()
		ch = t.sc.Current()
		if !numCh[ch] || e != nil {
			break
		}
	}
afterLoop:
	if e == io.EOF && numCh[ch] {
		t.v = t.sc.Bytes()
		return nil
	}
	if !hasExpNumPart {
		goto errLbl
	}
	if e == nil && spaceCh[ch] {
		idx := t.sc.Index()
		e = t.nextKeepWhiteSpace()
		ch = t.sc.Current()
		if spaceCh[ch] {
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
	fmt.Println(isTimeStrFn(t.v), string(t.v))
	if isTimeStrFn(t.v) {
		t.scType = TokenTime
	}
	return nil
}

func (t *JStructTokenizerImpl) ReadString() error {
	t.scType = TokenString
	first := t.sc.Index()
	var e error
	var ch, prev byte
	for {
		e = t.sc.Next()
		ch = t.sc.Current()
		if ch == quoteCh && prev != backSlashCh || e != nil {
			break
		}
		if ch == 'u' && prev == backSlashCh {
			e = t.readHex(4)
			ch = t.sc.Current()
			if e != nil {
				break
			}
		}
		prev = ch
	}
	idx := t.sc.Index()
	if e != nil && e != io.EOF {
		goto errLbl
	}
	e = t.sc.Next()
	ch = t.sc.Current()
	if (e == nil && !spaceCh[ch]) ||
		(e == io.EOF && ch != quoteCh) ||
		(ch == quoteCh && t.sc.Index() != idx) {
		goto errLbl
	}
	if e == nil && spaceCh[ch] {
		idx := t.sc.Index()
		e = t.nextKeepWhiteSpace()
		if spaceCh[t.sc.Current()] {
			l := idx - first
			t.v = t.sc.Bytes()[:l]
			return nil
		}
	}
	t.v = t.sc.Bytes()
	return nil
errLbl:
	t.scType = TokenUnknown
	t.v = t.sc.Bytes()
	return InvalidJsonPtrError{Pos: t.sc.Index(), Err: e}
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
	if idx == t.sc.Index() || spaceCh[t.sc.Current()] {
		t.v = t.sc.Bytes()[:l]
		return nil
	}
	t.scType = TokenUnknown
	t.v = t.sc.Bytes()
	return InvalidJsonPtrError{Pos: t.sc.Index(), Err: e}
}

func (t *JStructTokenizerImpl) readHex(n int) error {
	var ch byte
	for i := 0; i < n; i++ {
		e := t.sc.Next()
		ch = t.sc.Current()
		if e != nil {
			return e
		}
		if !hexCh[ch] {
			return InvalidHexNumberError
		}
	}
	return nil
}

func isTimeStrFn(v []byte) bool {
	l := len(v)

	if l < minTimeLen || l > maxTimeLen {
		return false
	}
	if v[0] != '"' || v[l-1] != '"' {
		return false
	}
	if !numCh[v[1]] || !numCh[v[2]] || !numCh[v[3]] || !numCh[v[4]] ||
		v[5] != '-' ||
		!numCh[v[6]] || !numCh[v[7]] ||
		v[8] != '-' ||
		!numCh[v[9]] || !numCh[v[10]] ||
		v[11] != 'T' ||
		!numCh[v[12]] || !numCh[v[13]] ||
		v[14] != ':' ||
		!numCh[v[15]] || !numCh[v[16]] ||
		v[17] != ':' ||
		!numCh[v[18]] || !numCh[v[19]] {
		return false
	}
	timeZonePos := 20
	for idx := 20; idx < l-1; idx++ {
		ch := v[idx]
		if idx == timeZonePos && ch == 'Z' {
			idx += 1
			ch = v[idx]
			if ch != '"' {
				return false
			}
			break
		} else if idx == 20 && ch == '.' {
			idx += 1
			ch = v[idx]
			if !numCh[ch] {
				return false
			}
			for {
				idx += 1
				ch = v[idx]
				if ch == 'Z' || ch == '+' || ch == '-' {
					timeZonePos = idx
					idx -= 1
					break
				}
				if !numCh[ch] {
					return false
				}
			}
			timeZonePos = idx + 1
		} else if idx == timeZonePos && (ch == '+' || ch == '-') {
			if !numCh[v[idx+1]] || !numCh[v[idx+2]] ||
				v[idx+3] != ':' ||
				!numCh[v[idx+4]] || !numCh[v[idx+5]] ||
				v[idx+6] != '"' {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
