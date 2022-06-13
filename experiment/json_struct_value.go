package experiment

import (
	djs "github.com/Pencroff/JsonStruct"
	"time"
)

//region JsonStructValue
type JsonStructValue struct {
	props    map[string]*JsonStructValue
	elements []*JsonStructValue

	valType djs.Type

	// data
	intNum   int
	floatNum float64
	str      string
	dt       time.Time
}

func (s *JsonStructValue) IsNumber() bool {
	return s.valType == djs.Int || s.valType == djs.Uint || s.valType == djs.Float
}

func (s *JsonStructValue) IsInt() bool {
	return s.valType == djs.Int
}

func (s *JsonStructValue) SetInt(v int) {
	s.SetNull()
	s.valType = djs.Int
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
	return s.valType == djs.False || s.valType == djs.True
}

func (s *JsonStructValue) SetBool(v bool) {
	s.SetNull()
	s.valType = djs.False
	if v {
		s.valType = djs.True
	}
}

func (s *JsonStructValue) Bool() bool {
	return s.valType == djs.True
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
