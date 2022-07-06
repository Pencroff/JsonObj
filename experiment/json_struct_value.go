package experiment

import (
	djs "github.com/Pencroff/JsonStruct"
	h "github.com/Pencroff/JsonStruct/helper"
	"strconv"
	"time"
)

type JsonStructValue struct {
	valType djs.Type

	// data
	iNum int64
	uNum uint64
	fNum float64
	str  string
	tm   time.Time
	// ref
	props map[string]djs.JStructOps
	elms  []djs.JStructOps
}

func (s *JsonStructValue) Type() djs.Type {
	return s.valType
}

func (s *JsonStructValue) Value() interface{} {
	switch s.valType {
	default:
		return nil
	case djs.Null:
		return nil
	case djs.True, djs.False:
		return s.Bool()
	case djs.Int:
		return s.Int()
	case djs.Uint:
		return s.Uint()
	case djs.Float:
		return s.Float()
	case djs.String:
		return s.String()
	case djs.Time:
		return s.Time()
	case djs.Object:
		return s.props
	case djs.Array:
		return s.elms
	}
}

func (s *JsonStructValue) IsNull() bool {
	return s.valType == djs.Null
}

func (s *JsonStructValue) SetNull() {
	s.valType = djs.Null
}

func (s *JsonStructValue) IsObject() bool {
	return s.valType == djs.Object
}

func (s *JsonStructValue) AsObject() {
	if s.valType == djs.Object {
		return
	}
	s.valType = djs.Object
	s.props = make(map[string]djs.JStructOps)
}

func (s *JsonStructValue) IsArray() bool {
	return s.valType == djs.Array
}

func (s *JsonStructValue) AsArray() {
	if s.valType == djs.Array {
		return
	}
	s.valType = djs.Array
	s.elms = make([]djs.JStructOps, 0)
}

func (s *JsonStructValue) Size() int {
	switch s.valType {
	default:
		return -1
	case djs.String:
		return len(s.str)
	case djs.Object:
		return len(s.props)
	case djs.Array:
		return len(s.elms)
	}
}

func (s *JsonStructValue) IsBool() bool {
	return s.valType == djs.True || s.valType == djs.False
}

func (s *JsonStructValue) SetBool(v bool) {
	s.valType = djs.False
	if v {
		s.valType = djs.True
	}
}

func (s *JsonStructValue) Bool() bool {
	switch s.valType {
	default:
		return false
	case djs.True:
		return true
	case djs.Int:
		return s.iNum != 0
	case djs.Uint:
		return s.uNum != 0
	case djs.Float:
		return s.fNum != 0
	case djs.String:
		return s.str != ""
	case djs.Time:
		return s.tm.UnixMilli() != 0
	}
}

func (s *JsonStructValue) IsNumber() bool {
	return s.valType == djs.Int || s.valType == djs.Uint || s.valType == djs.Float
}

func (s *JsonStructValue) IsInt() bool {
	return s.valType == djs.Int
}

func (s *JsonStructValue) SetInt(v int64) {
	s.valType = djs.Int
	s.iNum = v
}

func (s *JsonStructValue) Int() int64 {
	switch s.valType {
	default:
		return 0
	case djs.True:
		return 1
	case djs.Int:
		return s.iNum
	case djs.Uint:
		return int64(s.uNum)
	case djs.Float:
		return int64(s.fNum)
	case djs.String:
		n, _ := h.StringToInt(s.str)
		return n
	case djs.Time:
		return s.tm.UnixMilli()
	}
}

func (s *JsonStructValue) IsUint() bool {
	return s.valType == djs.Uint
}

func (s *JsonStructValue) SetUint(v uint64) {
	s.valType = djs.Uint
	s.uNum = v
}

func (s *JsonStructValue) Uint() uint64 {
	switch s.valType {
	default:
		return 0
	case djs.True:
		return 1
	case djs.Int:
		return uint64(s.iNum)
	case djs.Uint:
		return s.uNum
	case djs.Float:
		return uint64(s.fNum)
	case djs.String:
		n, _ := h.StringToUint(s.str)
		return n
	case djs.Time:
		return uint64(s.tm.UnixMilli())
	}
}

func (s *JsonStructValue) IsFloat() bool {
	return s.valType == djs.Float
}

func (s *JsonStructValue) SetFloat(v float64) {
	s.valType = djs.Float
	s.fNum = v
}

func (s *JsonStructValue) Float() float64 {
	switch s.valType {
	default:
		return 0
	case djs.True:
		return 1
	case djs.Int:
		return float64(s.iNum)
	case djs.Uint:
		return float64(s.uNum)
	case djs.Float:
		return s.fNum
	case djs.String:
		n, _ := strconv.ParseFloat(s.str, 64)
		return n
	}
}

func (s *JsonStructValue) IsString() bool {
	return s.valType == djs.String
}

func (s *JsonStructValue) SetString(v string) {
	s.valType = djs.String
	s.str = v
}

func (s *JsonStructValue) String() string {
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
		return strconv.FormatInt(s.iNum, 10)
	case djs.Uint:
		return strconv.FormatUint(s.uNum, 10)
	case djs.Float:
		return strconv.FormatFloat(s.fNum, 'f', -1, 64)
	case djs.String:
		return s.str
	case djs.Time:
		return s.tm.Format(time.RFC3339)
	case djs.Object:
		return "[object]"
	case djs.Array:
		return "[array]"
	}
}

func (s *JsonStructValue) IsTime() bool {
	return s.valType == djs.Time
}

func (s *JsonStructValue) SetTime(v time.Time) {
	s.valType = djs.Time
	s.tm = v
}

func (s *JsonStructValue) Time() time.Time {
	switch s.valType {
	default:
		return time.Time{}
	case djs.String:
		t, _ := time.Parse(time.RFC3339, s.str)
		return t
	case djs.Time:
		return s.tm
	}
}

func (s *JsonStructValue) SetKey(key string, v interface{}) error {
	if s.valType != djs.Object {
		return djs.NotObjectError
	}
	m := s.props
	pjs, err := s.populateVjs(v, m[key])
	if err != nil {
		return err
	}
	m[key] = pjs
	return nil
}

func (s *JsonStructValue) GetKey(key string) djs.JStructOps {
	if s.valType != djs.Object {
		return nil
	}
	return s.props[key]
}

func (s *JsonStructValue) RemoveKey(key string) djs.JStructOps {
	m := s.props
	v, _ := m[key]
	delete(m, key)
	return v
}

func (s *JsonStructValue) HasKey(key string) bool {
	if s.valType != djs.Object {
		return false
	}
	m := s.props
	_, ok := m[key]
	return ok
}

func (s *JsonStructValue) Keys() []string {
	if s.valType != djs.Object {
		return []string{}
	}
	m := s.props
	keys := make([]string, len(m))
	var idx uint64
	for k := range m {
		keys[idx] = k
		idx++
	}
	return keys
}

func (s *JsonStructValue) Push(v interface{}) error {
	if s.valType != djs.Array {
		return djs.NotArrayError
	}
	el, err := s.populateVjs(v, nil)
	if err != nil {
		return err
	}
	s.elms = append(s.elms, el)
	return nil
}

func (s *JsonStructValue) Pop() djs.JStructOps {
	if s.valType != djs.Array {
		return nil
	}
	m := s.elms
	lIdx := len(m) - 1
	if lIdx == -1 {
		return nil
	}
	v := m[lIdx]
	m[lIdx] = nil
	s.elms = m[:lIdx]
	return v
}

func (s *JsonStructValue) Shift() djs.JStructOps {
	if s.valType != djs.Array {
		return nil
	}
	m := s.elms
	l := len(m)
	if l == 0 {
		return nil
	}
	v := m[0]
	m[0] = nil
	s.elms = m[1:]
	return v
}

func (s *JsonStructValue) SetIndex(i int, v interface{}) error {
	if s.valType != djs.Array {
		return djs.NotArrayError
	}
	if i < 0 {
		return djs.IndexOutOfRangeError
	}
	el, err := s.populateVjs(v, nil)
	if err != nil {
		return err
	}
	m := s.elms
	l := len(m)
	if i >= l {
		m = append(m, make([]djs.JStructOps, i-l+1)...)
		s.elms = m
	}
	m[i] = el
	return nil
}

func (s *JsonStructValue) GetIndex(i int) djs.JStructOps {
	if s.valType != djs.Array {
		return nil
	}
	m := s.elms
	l := len(m)
	if i >= l {
		return nil
	}
	return m[i]
}

func (s *JsonStructValue) populateVjs(v interface{}, vjs djs.JStructOps) (djs.JStructOps, error) {
	switch data := v.(type) {
	case djs.JStructOps:
		vjs = data
	case nil:
		vjs = resolveValue(vjs)
		vjs.SetNull()
	case bool:
		vjs = resolveValue(vjs)
		vjs.SetBool(data)
	case int8:
		vjs = resolveValue(vjs)
		vjs.SetInt(int64(data))
	case int16:
		vjs = resolveValue(vjs)
		vjs.SetInt(int64(data))
	case int32:
		vjs = resolveValue(vjs)
		vjs.SetInt(int64(data))
	case int64:
		vjs = resolveValue(vjs)
		vjs.SetInt(data)
	case int:
		vjs = resolveValue(vjs)
		vjs.SetInt(int64(data))
	case uint8:
		vjs = resolveValue(vjs)
		vjs.SetUint(uint64(data))
	case uint16:
		vjs = resolveValue(vjs)
		vjs.SetUint(uint64(data))
	case uint32:
		vjs = resolveValue(vjs)
		vjs.SetUint(uint64(data))
	case uint64:
		vjs = resolveValue(vjs)
		vjs.SetUint(data)
	case uint:
		vjs = resolveValue(vjs)
		vjs.SetUint(uint64(data))
	case float64:
		vjs = resolveValue(vjs)
		vjs.SetFloat(data)
	case string:
		vjs = resolveValue(vjs)
		vjs.SetString(data)
	case time.Time:
		vjs = resolveValue(vjs)
		vjs.SetTime(data)
	default:
		return nil, djs.UnsupportedTypeError
	}
	return vjs, nil
}

func resolveValue(v djs.JStructOps) djs.JStructOps {
	if v == nil {
		return &JsonStructPtr{}
	}
	return v
}
