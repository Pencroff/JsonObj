package JsonStruct

import (
	h "github.com/Pencroff/JsonStruct/helper"
	"io"
)

type TokenizerKind byte

const (
	KindUnknown TokenizerKind = iota
	KindNull
	KindFalse
	KindTrue
	KindNumber
	KindFloatNumber
	KindTime
	KindString
	KindLiteral
)

func (k TokenizerKind) String() string {
	switch k {
	case KindNull:
		return "KindNull"
	case KindFalse:
		return "KindFalse"
	case KindTrue:
		return "KindTrue"
	case KindNumber:
		return "KindNumber"
	case KindFloatNumber:
		return "KindFloatNumber"
	case KindTime:
		return "KindTime"
	case KindString:
		return "KindString"
	case KindLiteral:
		return "KindLiteral"
	default:
		return "KindUnknown"
	}
}

type TokenizerLevel byte

const (
	LevelUnknown TokenizerLevel = iota
	LevelRoot
	LevelObject
	LevelObjectEnd
	LevelArray
	LevelArrayEnd
	LevelKey
	LevelValue
	LevelValueLast
)

func (l TokenizerLevel) String() string {
	switch l {
	case LevelRoot:
		return "LevelRoot"
	case LevelObject:
		return "LevelObject"
	case LevelObjectEnd:
		return "LevelObjectEnd"
	case LevelArray:
		return "LevelArray"
	case LevelArrayEnd:
		return "LevelArrayEnd"
	case LevelKey:
		return "LevelKey"
	case LevelValue:
		return "LevelValue"
	case LevelValueLast:
		return "LevelValueLast"
	default:
		return "LevelUnknown"
	}
}

type JStructTokenizer interface {
	Next() error
	Value() []byte
	Kind() TokenizerKind
	Level() TokenizerLevel
}

func NewJStructTokenizer(sc JStructScanner) JStructTokenizer {
	return &JStructTokenizerImpl{sc: sc, scLevel: LevelRoot, depth: []TokenizerLevel{LevelRoot}}
}

type JStructTokenizerImpl struct {
	sc      JStructScanner
	scType  TokenizerKind
	scLevel TokenizerLevel
	v       []byte
	depth   []TokenizerLevel
}

func (t *JStructTokenizerImpl) nextSkipWhiteSpace() error {
	for {
		err := t.sc.Next()
		if err != nil {
			return err
		}
		b := t.sc.Current()
		if h.SpaceCh[b] {
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
		ch := t.sc.Current()
		if h.SpaceCh[ch] {
			continue
		}
		return nil
	}
}

func (t *JStructTokenizerImpl) Next() error {
	t.scType = KindUnknown
	//t.scLevel = LevelUnknown
	if t.scLevel == LevelValueLast || t.scLevel == LevelArrayEnd {
		t.PopLevel()
	}
	for {
		err := t.nextSkipWhiteSpace()
		if err != nil {
			idx := t.sc.Index()
			if idx > -1 {
				return InvalidJsonPtrError{Err: err, Pos: t.sc.Index()}
			}
			return InvalidJsonError{Err: err}
		}
		switch t.sc.Current() {
		//case '{':
		//	t.scType = TokenObject
		//	return nil
		case '[':
			t.PushLevel(LevelArray)
			t.scType = KindLiteral
			t.sc.Bytes()
			t.v = nil
			return nil
		case 'n':
			return t.ReadNull()
		case 'f':
			return t.ReadFalse()
		case 't':
			return t.ReadTrue()
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return t.ReadNumber(false)
		case h.MinusCh:
			return t.ReadNumber(true)
		case h.QuoteCh:
			return t.ReadStringTime()
		//case ':':
		//	t.scType = TokenKey
		//	return nil
		//case ',':
		//	t.scType = TokenValue
		//	return nil
		case ']':
			t.PopLevel()
			t.scLevel = LevelArrayEnd
			t.scType = KindLiteral
			return nil
		//case '}':
		//	t.scType = TokenValueLast
		//	return nil
		default:
			return InvalidJsonError{Err: err}
		}
	}
}

func (t *JStructTokenizerImpl) Value() []byte {
	return t.v
}

func (t *JStructTokenizerImpl) Kind() TokenizerKind {
	return t.scType
}

func (t *JStructTokenizerImpl) Level() TokenizerLevel {
	return t.scLevel
}

// PushLevel pushes the current level to the stack and sets the new level to the given value
func (t *JStructTokenizerImpl) PushLevel(l TokenizerLevel) {
	t.depth = append(t.depth, t.scLevel)
	t.scLevel = l
}

// PopLevel pops the current level from the stack and sets the new level to the popped value
func (t *JStructTokenizerImpl) PopLevel() {
	last := len(t.depth) - 1
	if last > 0 {
		t.depth = t.depth[:last]
		last -= 1

	}
	t.scLevel = t.depth[last]
}

func (t *JStructTokenizerImpl) ReadNull() error {
	return t.hardcodedToken(KindNull, h.NullTokenData)
}

func (t *JStructTokenizerImpl) ReadFalse() error {
	return t.hardcodedToken(KindFalse, h.FalseTokenData)
}

func (t *JStructTokenizerImpl) ReadTrue() error {
	return t.hardcodedToken(KindTrue, h.TrueTokenData)
}

func (t *JStructTokenizerImpl) ReadNumber(hasMinus bool) error {
	var idx, l int
	t.scType = KindNumber
	firstIdx := t.sc.Index()
	e := t.sc.Next()
	ch := t.sc.Current()
	hasIntPart := !hasMinus || h.NumCh[ch]
	if e != nil || !h.NumCh[ch] {
		goto afterLoop
	}
	// integer part
	for {
		e = t.sc.Next()
		ch = t.sc.Current()
		if !h.NumCh[ch] || e != nil {
			break
		}
	}
afterLoop:
	idx = t.sc.Index()
	l = idx - firstIdx
	if e == io.EOF && h.NumCh[ch] {
		t.v = t.sc.Bytes()
		return nil
	}
	if !hasIntPart || e != nil {
		goto errLbl
	}
	if ch == h.PointCh {
		return t.ReadFractionPart(firstIdx)
	}
	if ch == h.ExpSmCh || ch == h.ExpCh {
		return t.ReadExponentPart(firstIdx)
	}
	if t.scLevel == LevelRoot && !h.SpaceCh[ch] {
		goto errLbl
	}
	if h.SpaceCh[ch] {
		e = t.nextKeepWhiteSpace()
		ch = t.sc.Current()
		if h.SpaceCh[ch] {
			t.v = t.sc.Bytes()[:l]
			return nil
		}
	}
	if t.scLevel != LevelRoot {
		if ch == h.CommaCh {
			t.v = t.sc.Bytes()[:l]
			t.scLevel = LevelValue
			return nil
		}
		if ch == h.CloseBracketCh {
			t.v = t.sc.Bytes()[:l]
			t.scLevel = LevelValueLast
			return nil
		}
	}
errLbl:
	t.scType = KindUnknown
	t.v = t.sc.Bytes()
	return InvalidJsonPtrError{Pos: t.sc.Index(), Err: e}
}

func (t *JStructTokenizerImpl) ReadFractionPart(firstIdx int) error {
	var idx, l int
	t.scType = KindFloatNumber
	e := t.sc.Next()
	ch := t.sc.Current()
	hasFractionPart := h.NumCh[ch]
	if e != nil || !h.NumCh[ch] {
		goto afterLoop
	}
	for {
		e = t.sc.Next()
		ch = t.sc.Current()
		if !h.NumCh[ch] || e != nil {
			break
		}
	}
afterLoop:
	idx = t.sc.Index()
	l = idx - firstIdx
	if e == io.EOF && h.NumCh[ch] {
		t.v = t.sc.Bytes()
		return nil
	}
	if !hasFractionPart || e != nil {
		goto errLbl
	}
	if ch == h.ExpSmCh || ch == h.ExpCh {
		return t.ReadExponentPart(firstIdx)
	}
	if t.scLevel == LevelRoot && !h.SpaceCh[ch] {
		goto errLbl
	}
	if h.SpaceCh[ch] {
		e = t.nextKeepWhiteSpace()
		ch = t.sc.Current()
		if h.SpaceCh[ch] {
			t.v = t.sc.Bytes()[:l]
			return nil
		}
	}
	if t.scLevel != LevelRoot {
		if ch == h.CommaCh {
			t.v = t.sc.Bytes()[:l]
			t.scLevel = LevelValue
			return nil
		}
		if ch == h.CloseBracketCh {
			t.v = t.sc.Bytes()[:l]
			t.scLevel = LevelValueLast
			return nil
		}
	}
errLbl:
	t.scType = KindUnknown
	t.v = t.sc.Bytes()

	return InvalidJsonPtrError{Pos: t.sc.Index(), Err: e}
}

func (t *JStructTokenizerImpl) ReadExponentPart(firstIdx int) error {
	var idx, l int
	t.scType = KindFloatNumber
	hasExpNumPart := false
	e := t.sc.Next()
	ch := t.sc.Current()
	if ch == '+' || ch == '-' {
		e = t.sc.Next()
		ch = t.sc.Current()
	}
	hasExpNumPart = h.NumCh[ch]
	if e != nil || !h.NumCh[ch] {
		goto afterLoop
	}

	for {
		e = t.sc.Next()
		ch = t.sc.Current()
		if !h.NumCh[ch] || e != nil {
			break
		}
	}
afterLoop:
	idx = t.sc.Index()
	l = idx - firstIdx
	if e == io.EOF && h.NumCh[ch] {
		t.v = t.sc.Bytes()
		return nil
	}
	if !hasExpNumPart || e != nil {
		goto errLbl
	}
	if t.scLevel == LevelRoot && !h.SpaceCh[ch] {
		goto errLbl
	}
	if h.SpaceCh[ch] {
		e = t.nextKeepWhiteSpace()
		ch = t.sc.Current()
		if h.SpaceCh[ch] {
			t.v = t.sc.Bytes()[:l]
			return nil
		}
	}
	if t.scLevel != LevelRoot {
		if ch == h.CommaCh {
			t.v = t.sc.Bytes()[:l]
			t.scLevel = LevelValue
			return nil
		}
		if ch == h.CloseBracketCh {
			t.v = t.sc.Bytes()[:l]
			t.scLevel = LevelValueLast
			return nil
		}
	}
	//if e == nil && h.SpaceCh[ch] {
	//	idx := t.sc.Index()
	//	e = t.nextKeepWhiteSpace()
	//	ch = t.sc.Current()
	//	if h.SpaceCh[ch] {
	//		l := idx - first
	//		t.v = t.sc.Bytes()[:l]
	//		return nil
	//	}
	//}
errLbl:
	t.scType = KindUnknown
	t.v = t.sc.Bytes()
	return InvalidJsonPtrError{Pos: t.sc.Index(), Err: e}
}

// Read json string or time in RFC3339 format
func (t *JStructTokenizerImpl) ReadStringTime() error {
	err := t.ReadString()
	if err != nil {
		return err
	}
	if h.IsTimeFormat(t.v) {
		t.scType = KindTime
	}
	return nil
}

func (t *JStructTokenizerImpl) ReadString() error {
	t.scType = KindString
	first := t.sc.Index()
	var e error
	var ch, prev byte
	var escaped bool
	var l, idx int
	for {
		e = t.sc.Next()
		ch = t.sc.Current()
		if ch == h.QuoteCh && prev != h.BackSlashCh ||
			e != nil {
			break
		}
		if prev == h.BackSlashCh {
			switch ch {
			case 'u':
				e = t.readHex(4)
				if e != nil {
					goto errLbl
				}
			case 0x22, 0x2F, 0x5c, 'b', 'f', 'n', 'r', 't':
				escaped = true
				break
			default:
				if !escaped {
					e = InvalidEscapeCharacterError
					goto errLbl
				}
				escaped = false
			}
		}
		if ch == h.TabCh || ch == h.NewLineCh || ch == h.CarriageReturnCh {
			e = InvalidCharacterError
			goto errLbl
		}
		prev = ch
	}
	idx = t.sc.Index()
	l = idx - first + 1
	e = t.nextKeepWhiteSpace()
	ch = t.sc.Current()
	if idx == t.sc.Index() && ch == h.QuoteCh ||
		h.SpaceCh[ch] {
		t.v = t.sc.Bytes()[:l]
		return nil
	}
	if t.scLevel != LevelRoot {
		if ch == h.CommaCh {
			t.v = t.sc.Bytes()[:l]
			t.scLevel = LevelValue
			return nil
		}
		if ch == h.CloseBracketCh {
			t.v = t.sc.Bytes()[:l]
			t.scLevel = LevelValueLast
			return nil
		}
	}
errLbl:
	t.scType = KindUnknown
	t.v = t.sc.Bytes()
	return InvalidJsonPtrError{Pos: t.sc.Index(), Err: e}
}

func (t *JStructTokenizerImpl) hardcodedToken(
	kind TokenizerKind, origin []byte) error {
	t.scType = kind
	l := len(origin)
	idx := t.sc.Index()
	for i, b := range origin[1:] {
		e := t.sc.Next()
		if b != t.sc.Current() || e != nil {
			t.v = t.sc.Bytes()
			t.scType = KindUnknown
			return InvalidJsonPtrError{Pos: idx + i + 1, Err: e}
		}
	}
	idx = t.sc.Index()
	// Check if it is a valid token ended by whitespaces
	e := t.nextKeepWhiteSpace()
	ch := t.sc.Current()
	if t.scLevel == LevelRoot && (idx == t.sc.Index() || h.SpaceCh[ch]) {
		t.v = t.sc.Bytes()[:l]
		return nil
	}
	if t.scLevel != LevelRoot {
		if ch == h.CommaCh {
			t.v = t.sc.Bytes()[:l]
			t.scLevel = LevelValue
			return nil
		}
		if ch == h.CloseBracketCh {
			t.v = t.sc.Bytes()[:l]
			t.scLevel = LevelValueLast
			return nil
		}
	}

	t.scType = KindUnknown
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
		if !h.HexCh[ch] {
			return InvalidHexNumberError
		}
	}
	return nil
}
