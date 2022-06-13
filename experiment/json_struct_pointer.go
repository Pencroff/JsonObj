package experiment

import (
	djs "github.com/Pencroff/JsonStruct"
	h "github.com/Pencroff/JsonStruct/helper"
	"strconv"
	"time"
	"unsafe"
)

//region JsonStructPtr
type JsonStructPtr struct {
	valType djs.Type
	// data
	ptr unsafe.Pointer
}

func (s *JsonStructPtr) Type() djs.Type {
	return s.valType
}

//region Primitive operations

func (s *JsonStructPtr) IsBool() bool {
	return s.valType == djs.False || s.valType == djs.True
}

func (s *JsonStructPtr) SetBool(v bool) {
	s.valType = djs.False
	if v {
		s.valType = djs.True
	}
	s.ptr = nil
}

func (s *JsonStructPtr) Bool() bool {
	switch s.valType {
	default:
		return false
	case djs.True:
		return true
	case djs.Int:
		v := *(*int)(s.ptr)
		return v != 0
	case djs.Uint:
		v := *(*uint)(s.ptr)
		return v != 0
	case djs.Float:
		v := *(*float64)(s.ptr)
		return v != 0
	case djs.String:
		v := *(*string)(s.ptr)
		return v != ""
	case djs.Time:
		v := *(*time.Time)(s.ptr)
		return v.UnixMilli() != 0
	}
}

func (s *JsonStructPtr) IsNumber() bool {
	return s.valType == djs.Int || s.valType == djs.Uint || s.valType == djs.Float
}

func (s *JsonStructPtr) IsInt() bool {
	return s.valType == djs.Int
}

func (s *JsonStructPtr) SetInt(v int64) {
	s.valType = djs.Int
	s.ptr = unsafe.Pointer(&v)
}

func (s *JsonStructPtr) Int() int64 {
	switch s.valType {
	default:
		return 0
	case djs.True:
		return 1
	case djs.Int:
		return *(*int64)(s.ptr)
	case djs.Uint:
		return *(*int64)(s.ptr)
	case djs.Float:
		v := *(*float64)(s.ptr)
		return int64(v)
	case djs.String:
		v := *(*string)(s.ptr)
		n, _ := h.StringToInt(v)
		return n
	case djs.Time:
		v := *(*time.Time)(s.ptr)
		return v.UnixMilli()
	}
}

func (s *JsonStructPtr) IsUint() bool {
	return s.valType == djs.Uint
}

func (s *JsonStructPtr) SetUint(v uint64) {
	s.valType = djs.Uint
	s.ptr = unsafe.Pointer(&v)
}

func (s *JsonStructPtr) Uint() uint64 {
	switch s.valType {
	default:
		return 0
	case djs.True:
		return 1
	case djs.Int:
		return *(*uint64)(s.ptr)
	case djs.Uint:
		return *(*uint64)(s.ptr)
	case djs.Float:
		v := *(*float64)(s.ptr)
		return uint64(v)
	case djs.String:
		v := *(*string)(s.ptr)
		n, _ := h.StringToUint(v)
		return n
	case djs.Time:
		v := *(*time.Time)(s.ptr)
		return uint64(v.UnixMilli())
	}
}

func (s *JsonStructPtr) IsFloat() bool {
	return s.valType == djs.Float
}

func (s *JsonStructPtr) SetFloat(v float64) {
	s.valType = djs.Float
	s.ptr = unsafe.Pointer(&v)
}

func (s *JsonStructPtr) Float() float64 {
	switch s.valType {
	default:
		return 0
	case djs.True:
		return 1
	case djs.Int:
		v := *(*int)(s.ptr)
		return float64(v)
	case djs.Uint:
		v := *(*uint)(s.ptr)
		return float64(v)
	case djs.Float:
		return *(*float64)(s.ptr)
	case djs.String:
		v := *(*string)(s.ptr)
		n, _ := strconv.ParseFloat(v, 64)
		return n
	}
}

func (s *JsonStructPtr) IsString() bool {
	return s.valType == djs.String
}

func (s *JsonStructPtr) SetString(v string) {
	s.valType = djs.String
	s.ptr = unsafe.Pointer(&v)
}

func (s *JsonStructPtr) String() string {
	switch s.valType {
	default:
		return ""
	case djs.Null:
		return "null"
	case djs.False:
		return "false"
	case djs.True:
		return "true"
	case djs.Int:
		v := *(*int64)(s.ptr)
		return strconv.FormatInt(v, 10)
	case djs.Uint:
		v := *(*uint64)(s.ptr)
		return strconv.FormatUint(v, 10)
	case djs.Float:
		v := *(*float64)(s.ptr)
		return strconv.FormatFloat(v, 'f', -1, 64)
	case djs.String:
		return *(*string)(s.ptr)
	case djs.Time:
		v := *(*time.Time)(s.ptr)
		return v.Format(time.RFC3339)
	case djs.Object:
		return "[object]"
	}

}

func (s *JsonStructPtr) IsTime() bool {
	return s.valType == djs.Time
}

func (s *JsonStructPtr) SetTime(v time.Time) {
	s.valType = djs.Time
	s.ptr = unsafe.Pointer(&v)
}

func (s *JsonStructPtr) Time() time.Time {
	switch s.valType {
	default:
		v := s.String()
		t, _ := time.Parse(time.RFC3339, v)
		return t
	case djs.Time:
		return *(*time.Time)(s.ptr)
	}
}

func (s *JsonStructPtr) IsNull() bool {
	return s.valType == djs.Null
}

func (s *JsonStructPtr) SetNull() {
	s.valType = djs.Null
	s.ptr = nil
}

//endregion Primitive operations

//region Object operations

func (s *JsonStructPtr) Set(key string, v interface{}) error {
	if s.valType != djs.Object {
		return djs.NotObjectError
	}
	m := *(*map[string]djs.JsonStructOps)(s.ptr)
	vjs, ok := m[key]
	if !ok {
		vjs = &JsonStructPtr{}
	}
	switch data := v.(type) {
	case JsonStructPtr:
		vjs = &data
	case *JsonStructPtr:
		vjs = data
	case nil:
		vjs.SetNull()
	case bool:
		vjs.SetBool(data)
	case int8:
		vjs.SetInt(int64(data))
	case int16:
		vjs.SetInt(int64(data))
	case int32:
		vjs.SetInt(int64(data))
	case int64:
		vjs.SetInt(data)
	case int:
		vjs.SetInt(int64(data))
	case uint8:
		vjs.SetUint(uint64(data))
	case uint16:
		vjs.SetUint(uint64(data))
	case uint32:
		vjs.SetUint(uint64(data))
	case uint64:
		vjs.SetUint(data)
	case uint:
		vjs.SetUint(uint64(data))
	case float64:
		vjs.SetFloat(data)
	case string:
		vjs.SetString(data)
	case time.Time:
		vjs.SetTime(data)
	default:
		return djs.UnsupportedTypeError
	}
	m[key] = vjs
	return nil
}

func (s *JsonStructPtr) Get(key string) djs.JsonStructOps {
	if s.valType != djs.Object {
		return nil
	}
	m := *(*map[string]djs.JsonStructOps)(s.ptr)
	return m[key]
}

func (s *JsonStructPtr) Remove(key string) bool {
	m := *(*map[string]djs.JsonStructOps)(s.ptr)
	_, ok := m[key]
	delete(m, key)
	return ok
}

func (s *JsonStructPtr) Has(key string) bool {
	if s.valType != djs.Object {
		return false
	}
	m := *(*map[string]djs.JsonStructOps)(s.ptr)
	_, ok := m[key]
	return ok
}

func (s *JsonStructPtr) Keys() []string {
	//TODO implement me
	panic("implement me")
}

func (s *JsonStructPtr) IsObject() bool {
	return s.valType == djs.Object
}

func (s *JsonStructPtr) AsObject() {
	s.valType = djs.Object
	s.ptr = unsafe.Pointer(&map[string]djs.JsonStructOps{})
}

//endregion Object operations

//region Array operations

func (s *JsonStructPtr) Len() int {
	//TODO implement me
	panic("implement me")
}

func (s *JsonStructPtr) Push(v interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (s *JsonStructPtr) Pop() djs.JsonStructOps {
	//TODO implement me
	panic("implement me")
}

func (s *JsonStructPtr) SetIndex(i int, v interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (s *JsonStructPtr) GetIndex(i int) djs.JsonStructOps {
	//TODO implement me
	panic("implement me")
}

func (s *JsonStructPtr) IsArray() bool {
	//TODO implement me
	panic("implement me")
}

func (s *JsonStructPtr) AsArray() {
	//TODO implement me
	panic("implement me")
}

//endregion Array operations

//endregion JsonStructPtr
