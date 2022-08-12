package JsonStruct

import (
	"bytes"
	"io"
)

// Injection inspired by github.com/Rhymond/go-money

// Injection points for testing purpose and custom implementation.
// If you would like to override JSON marshal/unmarshal implementation, please follow bellow approach.
//   js.UnmarshalJSON = func (bytes []byte, v js.JStructOps) error { ... }
//   js.MarshalJSON = func (v js.JStructOps) ([]byte, error) { ... }
var (
	// UnmarshalJSON Func is injection point of json.Unmarshaller for JsonStruct
	UnmarshalJSON = UnmarshalJSONFn
	// MarshalJSON Func is injection point of json.Marshaller for JsonStruct
	MarshalJSON = MarshalJSONFn
	// JStructParse Func provides io.Reader based parsing of JSON data
	JStructParse = JStructParseFn
	// JStructSerialize Func provides io.Writer based serialization of JSON data
	JStructSerialize = JStructSerializeFn
)

//type ParseState byte
//
//const (
//	None ParseState = iota
//	PrimitiveValue
//	Obj
//	Key
//	Value
//	Arr
//	ArrElement
//)

// Initial implementation of the JSON parser supported Standard ECMA-404 JSON format.
// https://www.ecma-international.org/publications-and-standards/standards/ecma-404/

func UnmarshalJSONFn(b []byte, v JStructOps) error {
	var rd = bytes.NewReader(b)
	return JStructParse(rd, v)
}

func JStructParseFn(rd io.Reader, v JStructOps) (e error) {
	sc := NewJStructScanner(rd)
	tc := NewJStructTokenizer(sc)
	e = tc.Next()
	if e != nil {
		return
	}
	switch tc.Kind() {
	case KindNull:
		v.SetNull()
	case KindTrue:
		v.SetBool(true)
	case KindFalse:
		v.SetBool(false)
	default:
		e = InvalidJsonError{}
	}
	return
}

func MarshalJSONFn(v JStructOps) ([]byte, error) {
	b := bytes.NewBuffer([]byte{})
	err := JStructSerialize(v, b)
	return b.Bytes(), err
}

func JStructSerializeFn(v JStructOps, wr io.Writer) (e error) {
	return
}

//func (s *JsonStruct) ToJson() string {
//	switch s.valType {
//	case Integer:
//		return strconv.Itoa(s.intNum)
//	case Float:
//		return strconv.FormatFloat(s.floatNum, 'f', -1, 64)
//	case Bool:
//		if s.intNum == 1 {
//			return "true"
//		} else {
//			return "false"
//		}
//	case String:
//		return `"` + s.str + `"`
//	case Time:
//		return `"` + s.dt.Format(time.RFC3339) + `"`
//	}
//	return "null"
//}
