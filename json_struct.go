package JsonStruct

import (
	"strconv"
	"time"
)

type JsonStructType byte

const (
	Null    JsonStructType = 0
	Integer                = 'i'
	Float                  = 'f'
	Bool                   = 'b'
	String                 = 's'
	Time                   = 't'
	//Object                 = 'o'
	//Array                  = 'a'
)

type JsonStructMap map[string]JsonStruct

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
	JsonStructMap

	valType JsonStructType

	// data
	intNum   int
	floatNum float64
	str      string
	dt       time.Time
}

func (s *JsonStruct) ToJson() string {
	switch s.valType {
	case Integer:
		return strconv.Itoa(s.intNum)
	case Float:
		return strconv.FormatFloat(s.floatNum, 'f', -1, 64)
	case Bool:
		if s.intNum == 1 {
			return "true"
		} else {
			return "false"
		}
	case String:
		return `"` + s.str + `"`
	case Time:
		return `"` + s.dt.Format(time.RFC3339) + `"`
	}
	return "null"
}

func (s *JsonStruct) IsNull() bool {
	return s.valType == Null
}

func (s *JsonStruct) SetNull() {
	s.reset()
}

func (s *JsonStruct) IsNumber() bool {
	return s.valType == Integer || s.valType == Float
}

func (s *JsonStruct) SetInt(i int) {
	s.reset()
	s.valType = Integer
	s.intNum = i
}
func (s *JsonStruct) GetInt() int {
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
func (s *JsonStruct) GetFloat() float64 {
	return s.floatNum
}

func (s *JsonStruct) IsFloat() bool {
	return s.valType == Float
}

func (s *JsonStruct) SetBool(v bool) {
	s.reset()
	s.valType = Bool
	if v {
		s.intNum = 1
	}
}
func (s *JsonStruct) GetBool() bool {
	return s.intNum == 1
}

func (s *JsonStruct) IsBool() bool {
	return s.valType == Bool
}

func (s *JsonStruct) SetString(v string) {
	s.reset()
	s.valType = String
	s.str = v
}
func (s *JsonStruct) GetString() string {
	return s.str
}

func (s *JsonStruct) IsString() bool {
	return s.valType == String
}

func (s *JsonStruct) SetTime(v time.Time) {
	s.reset()
	s.valType = Time
	s.dt = v
}
func (s *JsonStruct) GetTime() time.Time {
	return s.dt
}

func (s *JsonStruct) IsTime() bool {
	return s.valType == Time
}

func (s *JsonStruct) reset() {
	s.valType = Null
	s.intNum = 0
	s.floatNum = 0
	s.str = ""
	s.dt = time.Time{}
}
