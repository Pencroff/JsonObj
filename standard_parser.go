package JsonStruct

import (
	"bufio"
	"bytes"
	"github.com/Pencroff/JsonStruct/parsing"
	"io"
)

// Injection inspired by github.com/Rhymond/go-money

// Injection points for testing purpose and custom implementation.
// If you would like to override JSON marshal/unmarshal implementation, please follow bellow approach.
//   js.UnmarshalJSON = func (bytes []byte, v js.JStructOps) error { ... }
//   js.MarshalJSON = func (v js.JStructOps) ([]byte, error) { ... }
var (
	// UnmarshalJSON Func is injection point of json.Unmarshaller for JsonStruct
	UnmarshalJSON = JStructParseByte
	// MarshalJSON Func is injection point of json.Marshaller for JsonStruct
	MarshalJSON = JStructSerializeByte
	// JStructParse Func provides io.Reader based parsing of JSON data
	JStructParse = JStructParseReader
	// JStructSerialize Func provides io.Writer based serialization of JSON data
	JStructSerialize        = JStructSerializeWriter
	JStructReaderBufferSize = 1024
)

// Initial implementation of the JSON parser supported Standard ECMA-404 JSON format.
// https://www.ecma-international.org/publications-and-standards/standards/ecma-404/

func JStructParseByte(b []byte, v JStructOps) error {
	rd := bytes.NewReader(b)
	return JStructParse(rd, v)
}

type ParseState byte

const (
	None ParseState = iota
	PrimitiveValue
	Obj
	Key
	Value
	Arr
	ArrElement
)

func JStructParseReader(r io.Reader, v JStructOps) (e error) {
	rd := bufio.NewReaderSize(r, JStructReaderBufferSize)
	finishLoop := false
	buf := make([]byte, 1)
	var bt byte
	for {
		n, err := rd.Read(buf)
		if err != nil {
			if err == io.EOF {
				finishLoop = true
			} else {
				return err
			}
		}
		if n == 0 {
			continue
		}
		bt = buf[0]
		switch bt {
		//case '{':
		//	state = Obj
		//case '[':
		//	state = Arr
		default:
			parsing.ParsePrimitiveValue(rd, v)
		}
		if finishLoop {
			break
		}
	}
	return
}

func JStructSerializeByte(v JStructOps) ([]byte, error) {
	var b bytes.Buffer
	err := JStructSerialize(v, &b)
	return b.Bytes(), err
}

func JStructSerializeWriter(v JStructOps, wr io.Writer) error {
	return nil
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
