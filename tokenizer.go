package JsonStruct

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
}

func (t *JStructTokenizerImpl) Next() error {
	t.scType = TokenUnknown
	for {
		err := t.sc.Next()
		if err != nil {
			return err
		}
		switch t.sc.Current() {
		//case ' ', '\t', '\n', '\r':
		//	continue
		//case '{':
		//	t.scType = TokenObject
		//	return nil
		//case '[':
		//	t.scType = TokenArray
		//	return nil
		case 'n':
			t.scType = TokenNull
			//return t.ReadNBytes(3)
		case 'f':
			t.scType = TokenFalse
			//return t.ReadNBytes(4)
		case 't':
			t.scType = TokenTrue
			//return t.ReadNBytes(3)
			return nil
		case '-':
			t.scType = TokenIntNumber
			return nil
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			t.scType = TokenIntNumber
			return nil
		case '.':
			t.scType = TokenFloatNumber
			return nil
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
			return InvalidJsonError
		}
	}
	return nil
}

//func (t *JStructTokenizerImpl) ReadNBytes(n int) error {
//	for i := 0; i < n; i++ {
//		b, err := t.rd.ReadByte()
//		if err != nil {
//			return err
//		}
//		t.buf = append(t.buf, b)
//	}
//	return nil
//}

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

func (t *JStructTokenizerImpl) Value() []byte {
	return t.sc.Peek()
}

func (t *JStructTokenizerImpl) Kind() TokenizerKind {
	return t.scType
}
