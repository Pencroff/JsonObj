package JsonStruct

import (
	h "github.com/Pencroff/JsonStruct/helper"
	"strconv"
	"time"
	"unsafe"
)

// JsonStruct is a struct that can be converted to JSON.
// It implements the json.Marshaler and json.Unmarshaler interfaces.
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
	valType Type
	// data
	ptr unsafe.Pointer
}

//region Json Unmarshal / Marshal

func (s *JsonStruct) UnmarshalJSON(bytes []byte) error {
	return UnmarshalJSON(bytes, s)
}

func (s *JsonStruct) MarshalJSON() ([]byte, error) {
	return MarshalJSON(s)
}

//endregion

func (s *JsonStruct) Type() Type {
	return s.valType
}

func (s *JsonStruct) Value() interface{} {
	switch s.valType {
	default:
		return nil
	case False:
		return false
	case True:
		return true
	case Int:
		return *(*int64)(s.ptr)
	case Uint:
		return *(*uint64)(s.ptr)
	case Float:
		return *(*float64)(s.ptr)
	case String:
		return *(*string)(s.ptr)
	case Time:
		return *(*time.Time)(s.ptr)
	case Object:
		return *(*map[string]JStructOps)(s.ptr)
	case Array:
		return *(*[]JStructOps)(s.ptr)
	}
}

func (s *JsonStruct) Size() int {
	switch s.valType {
	default:
		return -1
	case String:
		v := *(*string)(s.ptr)
		return len(v)
	case Object:
		v := *(*map[string]JStructOps)(s.ptr)
		return len(v)
	case Array:
		v := *(*[]JStructOps)(s.ptr)
		return len(v)
	}
}

//region Primitive operations

func (s *JsonStruct) IsBool() bool {
	return s.valType == False || s.valType == True
}

func (s *JsonStruct) SetBool(v bool) {
	s.valType = False
	if v == true {
		s.valType = True
	}
}

func (s *JsonStruct) Bool() bool {
	switch s.valType {
	default:
		return false
	case True:
		return true
	case Int:
		v := *(*int)(s.ptr)
		return v != 0
	case Uint:
		v := *(*uint)(s.ptr)
		return v != 0
	case Float:
		v := *(*float64)(s.ptr)
		return v != 0
	case String:
		v := *(*string)(s.ptr)
		return v != ""
	case Time:
		v := *(*time.Time)(s.ptr)
		return v.UnixMilli() != 0
	}
}

func (s *JsonStruct) IsNumber() bool {
	return s.valType == Int || s.valType == Uint || s.valType == Float
}

func (s *JsonStruct) IsInt() bool {
	return s.valType == Int
}

func (s *JsonStruct) SetInt(v int64) {
	s.valType = Int
	s.ptr = unsafe.Pointer(&v)
}

func (s *JsonStruct) Int() int64 {
	switch s.valType {
	default:
		return 0
	case True:
		return 1
	case Int:
		return *(*int64)(s.ptr)
	case Uint:
		return *(*int64)(s.ptr)
	case Float:
		v := *(*float64)(s.ptr)
		return int64(v)
	case String:
		v := *(*string)(s.ptr)
		n, _ := h.StringToInt(v)
		return n
	case Time:
		v := *(*time.Time)(s.ptr)
		return v.UnixMilli()
	}
}

func (s *JsonStruct) IsUint() bool {
	return s.valType == Uint
}

func (s *JsonStruct) SetUint(v uint64) {
	s.valType = Uint
	s.ptr = unsafe.Pointer(&v)
}

func (s *JsonStruct) Uint() uint64 {
	switch s.valType {
	default:
		return 0
	case True:
		return 1
	case Int:
		return *(*uint64)(s.ptr)
	case Uint:
		return *(*uint64)(s.ptr)
	case Float:
		v := *(*float64)(s.ptr)
		return uint64(v)
	case String:
		v := *(*string)(s.ptr)
		n, _ := h.StringToUint(v)
		return n
	case Time:
		v := *(*time.Time)(s.ptr)
		return uint64(v.UnixMilli())
	}
}

func (s *JsonStruct) IsFloat() bool {
	return s.valType == Float
}

func (s *JsonStruct) SetFloat(v float64) {
	s.valType = Float
	s.ptr = unsafe.Pointer(&v)
}

func (s *JsonStruct) Float() float64 {
	switch s.valType {
	default:
		return 0
	case True:
		return 1
	case Int:
		v := *(*int)(s.ptr)
		return float64(v)
	case Uint:
		v := *(*uint)(s.ptr)
		return float64(v)
	case Float:
		return *(*float64)(s.ptr)
	case String:
		v := *(*string)(s.ptr)
		n, _ := strconv.ParseFloat(v, 64)
		return n
	}
}

func (s *JsonStruct) IsString() bool {
	return s.valType == String
}

func (s *JsonStruct) SetString(v string) {
	s.valType = String
	s.ptr = unsafe.Pointer(&v)
}

func (s *JsonStruct) String() string {
	switch s.valType {
	default:
		return ""
	case Null:
		return "null"
	case False:
		return "false"
	case True:
		return "true"
	case Int:
		v := *(*int64)(s.ptr)
		return strconv.FormatInt(v, 10)
	case Uint:
		v := *(*uint64)(s.ptr)
		return strconv.FormatUint(v, 10)
	case Float:
		v := *(*float64)(s.ptr)
		return strconv.FormatFloat(v, 'f', -1, 64)
	case String:
		return *(*string)(s.ptr)
	case Time:
		v := *(*time.Time)(s.ptr)
		return v.Format(time.RFC3339)
	case Object:
		return "[object]"
	case Array:
		return "[array]"
	}
}

func (s *JsonStruct) IsTime() bool {
	return s.valType == Time
}

func (s *JsonStruct) SetTime(v time.Time) {
	s.valType = Time
	s.ptr = unsafe.Pointer(&v)
}

func (s *JsonStruct) Time() time.Time {
	switch s.valType {
	default:
		return time.Time{}
	case String:
		v := *(*string)(s.ptr)
		t, _ := time.Parse(time.RFC3339, v)
		return t
	case Time:
		return *(*time.Time)(s.ptr)
	}
}

func (s *JsonStruct) IsNull() bool {
	return s.valType == Null
}

func (s *JsonStruct) SetNull() {
	s.valType = Null
	s.ptr = nil
}

//endregion Primitive operations

//region Object operations

func (s *JsonStruct) SetKey(key string, v interface{}) error {
	if s.valType != Object {
		return NotObjectError
	}
	m := *(*map[string]JStructOps)(s.ptr)
	pjs, err := s.populatePjs(v, m[key])
	if err != nil {
		return err
	}
	m[key] = pjs
	return nil
}

func (s *JsonStruct) GetKey(key string) JStructOps {
	if s.valType != Object {
		return nil
	}
	m := *(*map[string]JStructOps)(s.ptr)
	return m[key]
}

func (s *JsonStruct) RemoveKey(key string) JStructOps {
	m := *(*map[string]JStructOps)(s.ptr)
	v, _ := m[key]
	delete(m, key)
	return v
}

func (s *JsonStruct) HasKey(key string) bool {
	if s.valType != Object {
		return false
	}
	m := *(*map[string]JStructOps)(s.ptr)
	_, ok := m[key]
	return ok
}

func (s *JsonStruct) Keys() []string {
	if s.valType != Object {
		return []string{}
	}
	m := *(*map[string]JStructOps)(s.ptr)
	keys := make([]string, len(m))
	var idx uint64
	for k := range m {
		keys[idx] = k
		idx++
	}
	return keys
}

func (s *JsonStruct) IsObject() bool {
	return s.valType == Object
}

func (s *JsonStruct) AsObject() {
	if s.valType == Object {
		return
	}
	s.valType = Object
	s.ptr = unsafe.Pointer(&map[string]JStructOps{})
}

//endregion Object operations

//region Array operations

// https://github.com/golang/go/wiki/SliceTricks

func (s *JsonStruct) Push(v interface{}) error {
	if s.valType != Array {
		return NotArrayError
	}
	m := *(*[]JStructOps)(s.ptr)
	el, err := s.populatePjs(v, nil)
	if err != nil {
		return err
	}
	m = append(m, el)
	s.ptr = unsafe.Pointer(&m)
	return nil
}

func (s *JsonStruct) Pop() JStructOps {
	if s.valType != Array {
		return nil
	}
	m := *(*[]JStructOps)(s.ptr)
	lIdx := len(m) - 1
	if lIdx == -1 {
		return nil
	}
	v := m[lIdx]
	m[lIdx] = nil
	m = m[:lIdx]
	s.ptr = unsafe.Pointer(&m)
	return v
}

func (s *JsonStruct) Shift() JStructOps {
	if s.valType != Array {
		return nil
	}
	m := *(*[]JStructOps)(s.ptr)
	l := len(m)
	if l == 0 {
		return nil
	}
	v := m[0]
	m[0] = nil
	m = m[1:]
	s.ptr = unsafe.Pointer(&m)
	return v
}

func (s *JsonStruct) SetIndex(i int, v interface{}) error {
	if s.valType != Array {
		return NotArrayError
	}
	if i < 0 {
		return IndexOutOfRangeError
	}
	el, err := s.populatePjs(v, nil)
	if err != nil {
		return err
	}
	m := *(*[]JStructOps)(s.ptr)
	l := len(m)
	if i >= l {
		m = append(m, make([]JStructOps, i-l+1)...)
		s.ptr = unsafe.Pointer(&m)
	}
	m[i] = el
	return nil
}

func (s *JsonStruct) GetIndex(i int) JStructOps {
	if s.valType != Array {
		return nil
	}
	m := *(*[]JStructOps)(s.ptr)
	l := len(m)
	if i >= l {
		return nil
	}
	return m[i]
}

func (s *JsonStruct) IsArray() bool {
	return s.valType == Array
}

func (s *JsonStruct) AsArray() {
	if s.valType == Array {
		return
	}
	s.valType = Array
	s.ptr = unsafe.Pointer(&[]JsonStruct{})
}

//endregion Array operations

// region Helper functions

func (s *JsonStruct) populatePjs(v interface{}, pjs JStructOps) (JStructOps, error) {
	switch data := v.(type) {
	case JStructOps:
		pjs = data
	case nil:
		pjs = resolvePointer(pjs)
		pjs.SetNull()
	case bool:
		pjs = resolvePointer(pjs)
		pjs.SetBool(data)
	case int8:
		pjs = resolvePointer(pjs)
		pjs.SetInt(int64(data))
	case int16:
		pjs = resolvePointer(pjs)
		pjs.SetInt(int64(data))
	case int32:
		pjs = resolvePointer(pjs)
		pjs.SetInt(int64(data))
	case int64:
		pjs = resolvePointer(pjs)
		pjs.SetInt(data)
	case int:
		pjs = resolvePointer(pjs)
		pjs.SetInt(int64(data))
	case uint8:
		pjs = resolvePointer(pjs)
		pjs.SetUint(uint64(data))
	case uint16:
		pjs = resolvePointer(pjs)
		pjs.SetUint(uint64(data))
	case uint32:
		pjs = resolvePointer(pjs)
		pjs.SetUint(uint64(data))
	case uint64:
		pjs = resolvePointer(pjs)
		pjs.SetUint(data)
	case uint:
		pjs = resolvePointer(pjs)
		pjs.SetUint(uint64(data))
	case float64:
		pjs = resolvePointer(pjs)
		pjs.SetFloat(data)
	case string:
		pjs = resolvePointer(pjs)
		pjs.SetString(data)
	case time.Time:
		pjs = resolvePointer(pjs)
		pjs.SetTime(data)
	default:
		return nil, UnsupportedTypeError
	}
	return pjs, nil
}

func resolvePointer(v JStructOps) JStructOps {
	if v == nil {
		return &JsonStruct{}
	}
	return v
}

// endregion Helper functions

//endregion JsonStructPtr
