package JsonStruct

import (
	"errors"
	"time"
)

type JsonStructPrimitiveOps interface {
	IsBool() bool
	SetBool(bool)
	Bool() bool

	IsNumber() bool
	IsInt() bool
	SetInt(int)
	Int() int

	IsUint() bool
	SetUint(uint)
	Uint() uint

	IsFloat() bool
	SetFloat(float64)
	Float() float64

	IsString() bool
	SetString(string)
	String() string

	IsTime() bool
	SetTime(time.Time)
	Time() time.Time

	IsNull() bool
	SetNull()
}

type JsonStructObjectOps interface {
	Set(string, interface{}) error
	Get(string) *JsonStruct
	Delete(string) bool
	Has(string) bool
	Keys() []string
	IsObject() bool
	AsObject()
}

type JsonStructArrayOps interface {
	Len() int
	Push(interface{}) error
	Pop() *JsonStruct
	GetIndex(int) *JsonStruct
	SetIndex(int, interface{}) error
	IsArray() bool
	AsArray()
}

type JsonStructOps interface {
	JsonStructPrimitiveOps
	JsonStructObjectOps
	JsonStructArrayOps
}

type JsonStructType byte

const (
	Null JsonStructType = iota
	False
	True
	Integer
	Float
	Time
	String
	Object
	Array
)

func (t JsonStructType) String() string {
	switch t {
	default:
		return "Null"
	case False:
		return "False"
	case True:
		return "True"
	case Integer:
		return "Integer"
	case Float:
		return "Float"
	case Time:
		return "Time"
	case String:
		return "String"
	case Object:
		return "Object"
	case Array:
		return "Array"
	}
}

var UnsupportedTypeError = errors.New("unsupported value type, resolved as null")

// JsonStruct is a struct that can be converted to JSON.
//// It implements the json.Marshaler and json.Unmarshaler interfaces.
//// It also implements the sql.Scanner and sql.Valuer interfaces.
// It supports JSON types like:
// 	- string
// 	- int / int64
// 	- float64
// 	- bool
// 	- [extra] DateTime (ISO 8601, rfc3339)
// 	- Object
// 	- Array
type JsonStruct struct {
	m map[string]*JsonStruct
	a []*JsonStruct

	valType JsonStructType

	// data
	intNum   int
	floatNum float64
	str      string
	dt       time.Time
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

func (s *JsonStruct) Set(key string, value interface{}) (err error) {
	if s.valType != Object {
		s.AsObject()
	}
	v, ok := s.m[key]
	if !ok {
		v = &JsonStruct{}
	}
	switch data := value.(type) {
	case nil:
	case JsonStruct:
		v = &data
	case *JsonStruct:
		v = data
	case int:
		v.SetInt(data)
	case float64:
		v.SetFloat(data)
	case bool:
		v.SetBool(data)
	case string:
		v.SetString(data)
	case time.Time:
		v.SetTime(data)
	default:
		return UnsupportedTypeError
	}
	s.m[key] = v
	return nil
}

func (s *JsonStruct) Get(key string) *JsonStruct {
	if s.valType == Object {
		return s.m[key]
	}
	return nil
}

func (s *JsonStruct) Delete(key string) bool {
	_, ok := s.m[key]
	delete(s.m, key)
	return ok
}

func (s *JsonStruct) Has(key string) bool {
	_, ok := s.m[key]
	return ok
}

func (s JsonStruct) Keys() []string {
	if s.valType == Object {
		keys := make([]string, len(s.m))
		var idx uint64
		for k := range s.m {
			keys[idx] = k
			idx++
		}
		return keys
	}
	return []string{}
}

func (s *JsonStruct) IsObject() bool {
	return s.valType == Object
}

func (s *JsonStruct) AsObject() {
	if s.valType != Object {
		s.reset()
		s.m = make(map[string]*JsonStruct)
		s.valType = Object
	}
}

func (s *JsonStruct) IsArray() bool {
	return s.valType == Array
}

func (s *JsonStruct) AsArray() {
	if s.valType != Array {
		s.reset()
		s.a = make([]*JsonStruct, 0)
		s.valType = Array
	}
}

//region Null ops

func (s *JsonStruct) IsNull() bool {
	return s.valType == Null
}

func (s *JsonStruct) SetNull() {
	s.reset()
}

//endregion

//region Number ops
func (s *JsonStruct) IsNumber() bool {
	return s.valType == Integer || s.valType == Float
}

func (s *JsonStruct) SetInt(i int) {
	s.reset()
	s.valType = Integer
	s.intNum = i
}
func (s *JsonStruct) Int() int {
	return s.intNum
}

func (s *JsonStruct) IsInt() bool {
	return s.valType == Integer
}

func (s *JsonStruct) SetFloat(i float64) {
	s.reset()
	s.valType = Float
	s.floatNum = i
}
func (s *JsonStruct) Float() float64 {
	return s.floatNum
}

func (s *JsonStruct) IsFloat() bool {
	return s.valType == Float
}

//endregion Number ops

//region Boolean ops
func (s *JsonStruct) SetBool(v bool) {
	s.reset()
	s.valType = False
	if v {
		s.valType = True
	}
}
func (s *JsonStruct) Bool() bool {
	return s.intNum == 1
}

func (s *JsonStruct) IsBool() bool {
	return s.valType == False || s.valType == True
}

//endregion Boolean ops

//region String ops
func (s *JsonStruct) SetString(v string) {
	s.reset()
	s.valType = String
	s.str = v
}
func (s *JsonStruct) String() string {
	return s.str
}

func (s *JsonStruct) IsString() bool {
	return s.valType == String
}

//endregion String ops

//region Time ops

func (s *JsonStruct) SetTime(v time.Time) {
	s.reset()
	s.valType = Time
	s.dt = v
}
func (s *JsonStruct) Time() time.Time {
	return s.dt
}

func (s *JsonStruct) IsTime() bool {
	return s.valType == Time
}

//endregion

//region Private methods

func (s *JsonStruct) reset() {
	s.valType = Null
	s.intNum = 0
	s.floatNum = 0
	s.str = ""
	s.dt = time.Time{}
	s.m = nil
	s.a = nil
}

//endregion Private methods
