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

func (s *JsonStructPtr) Value() interface{} {
	switch s.valType {
	default:
		return nil
	case djs.False:
		return false
	case djs.True:
		return true
	case djs.Int:
		return *(*int64)(s.ptr)
	case djs.Uint:
		return *(*uint64)(s.ptr)
	case djs.Float:
		return *(*float64)(s.ptr)
	case djs.String:
		return *(*string)(s.ptr)
	case djs.Time:
		return *(*time.Time)(s.ptr)
	case djs.Object:
		return *(*map[string]djs.JsonStructOps)(s.ptr)
	case djs.Array:
		return *(*[]djs.JsonStructOps)(s.ptr)
	}
}

func (s *JsonStructPtr) Size() int {
	switch s.valType {
	default:
		return -1
	case djs.String:
		v := *(*string)(s.ptr)
		return len(v)
	case djs.Object:
		v := *(*map[string]djs.JsonStructOps)(s.ptr)
		return len(v)
	case djs.Array:
		v := *(*[]djs.JsonStructOps)(s.ptr)
		return len(v)
	}
}

//region Primitive operations

func (s *JsonStructPtr) IsBool() bool {
	return s.valType == djs.False || s.valType == djs.True
}

func (s *JsonStructPtr) SetBool(v bool) {
	s.valType = djs.False
	if v == true {
		s.valType = djs.True
	}
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
	case djs.Array:
		return "[array]"
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
		return time.Time{}
	case djs.String:
		v := *(*string)(s.ptr)
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

func (s *JsonStructPtr) SetKey(key string, v interface{}) error {
	if s.valType != djs.Object {
		return djs.NotObjectError
	}
	m := *(*map[string]djs.JsonStructOps)(s.ptr)
	pjs, err := s.populatePjs(v, m[key])
	if err != nil {
		return err
	}
	m[key] = pjs
	return nil
}

func (s *JsonStructPtr) GetKey(key string) djs.JsonStructOps {
	if s.valType != djs.Object {
		return nil
	}
	m := *(*map[string]djs.JsonStructOps)(s.ptr)
	return m[key]
}

func (s *JsonStructPtr) RemoveKey(key string) djs.JsonStructOps {
	m := *(*map[string]djs.JsonStructOps)(s.ptr)
	v, _ := m[key]
	delete(m, key)
	return v
}

func (s *JsonStructPtr) HasKey(key string) bool {
	if s.valType != djs.Object {
		return false
	}
	m := *(*map[string]djs.JsonStructOps)(s.ptr)
	_, ok := m[key]
	return ok
}

func (s *JsonStructPtr) Keys() []string {
	if s.valType != djs.Object {
		return []string{}
	}
	m := *(*map[string]djs.JsonStructOps)(s.ptr)
	keys := make([]string, len(m))
	var idx uint64
	for k := range m {
		keys[idx] = k
		idx++
	}
	return keys
}

func (s *JsonStructPtr) IsObject() bool {
	return s.valType == djs.Object
}

func (s *JsonStructPtr) AsObject() {
	if s.valType == djs.Object {
		return
	}
	s.valType = djs.Object
	s.ptr = unsafe.Pointer(&map[string]djs.JsonStructOps{})
}

//endregion Object operations

//region Array operations

// https://github.com/golang/go/wiki/SliceTricks

func (s *JsonStructPtr) Push(v interface{}) error {
	if s.valType != djs.Array {
		return djs.NotArrayError
	}
	m := *(*[]djs.JsonStructOps)(s.ptr)
	el, err := s.populatePjs(v, nil)
	if err != nil {
		return err
	}
	m = append(m, el)
	s.ptr = unsafe.Pointer(&m)
	return nil
}

func (s *JsonStructPtr) Pop() djs.JsonStructOps {
	if s.valType != djs.Array {
		return nil
	}
	m := *(*[]djs.JsonStructOps)(s.ptr)
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

func (s *JsonStructPtr) Shift() djs.JsonStructOps {
	if s.valType != djs.Array {
		return nil
	}
	m := *(*[]djs.JsonStructOps)(s.ptr)
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

func (s *JsonStructPtr) SetIndex(i int, v interface{}) error {
	if s.valType != djs.Array {
		return djs.NotArrayError
	}
	if i < 0 {
		return djs.IndexOutOfRangeError
	}
	el, err := s.populatePjs(v, nil)
	if err != nil {
		return err
	}
	m := *(*[]djs.JsonStructOps)(s.ptr)
	l := len(m)
	if i >= l {
		m = append(m, make([]djs.JsonStructOps, i-l+1)...)
		s.ptr = unsafe.Pointer(&m)
	}
	m[i] = el
	return nil
}

func (s *JsonStructPtr) GetIndex(i int) djs.JsonStructOps {
	if s.valType != djs.Array {
		return nil
	}
	m := *(*[]djs.JsonStructOps)(s.ptr)
	l := len(m)
	if i >= l {
		return nil
	}
	return m[i]
}

func (s *JsonStructPtr) IsArray() bool {
	return s.valType == djs.Array
}

func (s *JsonStructPtr) AsArray() {
	if s.valType == djs.Array {
		return
	}
	s.valType = djs.Array
	s.ptr = unsafe.Pointer(&[]JsonStructPtr{})
}

//endregion Array operations

// region Helper functions

func (s *JsonStructPtr) populatePjs(v interface{}, pjs djs.JsonStructOps) (djs.JsonStructOps, error) {
	switch data := v.(type) {
	case djs.JsonStructOps:
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
		return nil, djs.UnsupportedTypeError
	}
	return pjs, nil
}

func resolvePointer(v djs.JsonStructOps) djs.JsonStructOps {
	if v == nil {
		return &JsonStructPtr{}
	}
	return v
}

// endregion Helper functions

//endregion JsonStructPtr
