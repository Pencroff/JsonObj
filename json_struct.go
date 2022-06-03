package JsonStruct

import (
	"errors"
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
	Object                 = 'o'
	Array                  = 'a'
)

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

func (s *JsonStruct) Set(key string, value interface{}) (err error) {
	if s.valType != Object {
		s.AsObject()
	}
	v, ok := value.(JsonStruct)
	if ok {
		s.m[key] = &v
		return nil
	}
	vptr, ok := value.(*JsonStruct)
	if ok {
		s.m[key] = vptr
		return nil
	}
	jsonStruct := JsonStruct{}

	vInt, ok := value.(int)
	if ok {
		jsonStruct.SetInt(vInt)
	} else {
		vFloat, ok := value.(float64)
		if ok {
			jsonStruct.SetFloat(vFloat)
		} else {
			vBool, ok := value.(bool)
			if ok {
				jsonStruct.SetBool(vBool)
			} else {
				vStr, ok := value.(string)
				if ok {
					jsonStruct.SetString(vStr)
				} else {
					vTime, ok := value.(time.Time)
					if ok {
						jsonStruct.SetTime(vTime)
					} else {
						if value != nil {
							return errors.New("unsupported value type, resolved as null")
						}
					}
				}
			}
		}
	}
	s.m[key] = &jsonStruct
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

//endregion Number ops

//region Boolean ops
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

//endregion Boolean ops

//region String ops
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

//endregion String ops

//region Time ops

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
