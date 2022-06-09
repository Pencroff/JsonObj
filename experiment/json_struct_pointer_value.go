package experiment

import (
	djs "github.com/Pencroff/JsonStruct"
	"time"
	"unsafe"
)

//region JsonStructValue
type JsonStructValue struct {
	props    map[string]*JsonStructValue
	elements []*JsonStructValue

	valType djs.JsonStructType

	// data
	intNum   int
	floatNum float64
	str      string
	dt       time.Time
}

func (s *JsonStructValue) IsNumber() bool {
	return s.valType == djs.Integer || s.valType == djs.Float
}

func (s *JsonStructValue) IsInt() bool {
	return s.valType == djs.Integer
}

func (s *JsonStructValue) SetInt(v int) {
	s.SetNull()
	s.valType = djs.Integer
	s.intNum = v
}

func (s *JsonStructValue) Int() int {
	return s.intNum
}

func (s *JsonStructValue) IsFloat() bool {
	return s.valType == djs.Float
}

func (s *JsonStructValue) SetFloat(v float64) {
	s.SetNull()
	s.valType = djs.Float
	s.floatNum = v
}

func (s *JsonStructValue) Float() float64 {
	return s.floatNum
}

func (s *JsonStructValue) IsBool() bool {
	return s.valType == djs.Bool
}

func (s *JsonStructValue) SetBool(v bool) {
	s.SetNull()
	s.valType = djs.Bool
	if v {
		s.intNum = 1
	}
}

func (s *JsonStructValue) Bool() bool {
	return s.intNum == 1
}

func (s *JsonStructValue) IsString() bool {
	return s.valType == djs.String
}

func (s *JsonStructValue) SetString(v string) {
	s.SetNull()
	s.valType = djs.String
	s.str = v
}

func (s *JsonStructValue) String() string {
	return s.str
}

func (s *JsonStructValue) IsTime() bool {
	return s.valType == djs.Time
}

func (s *JsonStructValue) SetTime(t time.Time) {
	s.SetNull()
	s.valType = djs.Time
	s.dt = t
}

func (s *JsonStructValue) Time() time.Time {
	return s.dt
}

func (s *JsonStructValue) IsNull() bool {
	return s.valType == djs.Null
}

func (s *JsonStructValue) SetNull() {
	s.valType = djs.Null
	s.intNum = 0
	s.floatNum = 0
	s.str = ""
	s.dt = time.Time{}
}

//endregion JsonStructValue

//region JsonStructPtr
type JsonStructPtr struct {
	valType djs.JsonStructType
	// data
	ptr unsafe.Pointer
}

func (s *JsonStructPtr) IsNumber() bool {
	return s.valType == djs.Integer || s.valType == djs.Float
}

func (s *JsonStructPtr) IsInt() bool {
	return s.valType == djs.Integer
}

func (s *JsonStructPtr) SetInt(v int) {
	s.valType = djs.Integer
	s.ptr = unsafe.Pointer(&v)
}

func (s *JsonStructPtr) Int() int {
	return *(*int)(s.ptr)
}

func (s *JsonStructPtr) IsFloat() bool {
	return s.valType == djs.Float
}

func (s *JsonStructPtr) SetFloat(v float64) {
	s.valType = djs.Float
	s.ptr = unsafe.Pointer(&v)
}

func (s *JsonStructPtr) Float() float64 {
	return *(*float64)(s.ptr)
}

func (s *JsonStructPtr) IsBool() bool {
	return s.valType == djs.Bool
}

func (s *JsonStructPtr) SetBool(v bool) {
	s.valType = djs.Bool
	s.ptr = unsafe.Pointer(&v)
}

func (s *JsonStructPtr) Bool() bool {
	return *(*bool)(s.ptr)
}

func (s *JsonStructPtr) IsString() bool {
	return s.valType == djs.String
}

func (s *JsonStructPtr) SetString(v string) {
	s.valType = djs.String
	s.ptr = unsafe.Pointer(&v)
}

func (s *JsonStructPtr) String() string {
	return *(*string)(s.ptr)
}

func (s *JsonStructPtr) IsTime() bool {
	return s.valType == djs.Time
}

func (s *JsonStructPtr) SetTime(v time.Time) {
	s.valType = djs.Time
	s.ptr = unsafe.Pointer(&v)
}

func (s *JsonStructPtr) Time() time.Time {
	return *(*time.Time)(s.ptr)
}

func (s *JsonStructPtr) IsNull() bool {
	return s.valType == djs.Null
}

func (s *JsonStructPtr) SetNull() {
	s.valType = djs.Null
	s.ptr = nil
}

//endregion JsonStructPtr
