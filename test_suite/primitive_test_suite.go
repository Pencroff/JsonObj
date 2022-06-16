package test_suite

import (
	djs "github.com/Pencroff/JsonStruct"
	"github.com/Pencroff/JsonStruct/helper"
	"github.com/stretchr/testify/suite"
	"time"
)

type PrimitiveOpsTestSuite struct {
	suite.Suite
	factory func() djs.JsonStructOps
	js      djs.JsonStructOps
}

func (s *PrimitiveOpsTestSuite) SetFactory(fn func() djs.JsonStructOps) {
	s.factory = fn
}

func (s *PrimitiveOpsTestSuite) SetupTest() {
	if s.factory == nil {
		panic("factory not provided")
	}
	s.js = s.factory()
}

func (s *PrimitiveOpsTestSuite) TestIsMethods() {
	timeStr := "2015-01-01T12:34:56Z"
	tm, _ := time.Parse(time.RFC3339, timeStr)
	tbl := []struct {
		val       interface{}
		setMethod string
		isBool    bool
		isNumber  bool
		isInt     bool
		isUint    bool
		isFloat   bool
		isTime    bool
		isString  bool
	}{
		{false, "SetBool", true, false, false, false, false, false, false},
		{true, "SetBool", true, false, false, false, false, false, false},
		{int64(1), "SetInt", false, true, true, false, false, false, false},
		{uint64(1), "SetUint", false, true, false, true, false, false, false},
		{3.1415, "SetFloat", false, true, false, false, true, false, false},
		{tm, "SetTime", false, false, false, false, false, true, false},
		{"hello", "SetString", false, false, false, false, false, false, true},
	}
	for _, el := range tbl {
		CallMethod(s.js, el.setMethod, el.val)
		s.Equal(el.isBool, s.js.IsBool())
		s.Equal(el.isNumber, s.js.IsNumber())
		s.Equal(el.isInt, s.js.IsInt())
		s.Equal(el.isUint, s.js.IsUint())
		s.Equal(el.isFloat, s.js.IsFloat())
		s.Equal(el.isTime, s.js.IsTime())
	}
}

func (s *PrimitiveOpsTestSuite) TestGetSetValueTable() {
	timeStr := "2015-01-01T12:34:56Z"
	tm, _ := time.Parse(time.RFC3339, timeStr)
	tbl := []struct {
		val       interface{}
		res       interface{}
		setMethod string
		getMethod string
	}{
		{false, false, "SetBool", "Bool"},
		{true, true, "SetBool", "Bool"},
		{int64(1), int64(1), "SetInt", "Int"},
		{uint64(1), uint64(1), "SetUint", "Uint"},
		{3.1415, 3.1415, "SetFloat", "Float"},
		{tm, tm, "SetTime", "Time"},
		{"hello", "hello", "SetString", "String"},
	}
	for _, el := range tbl {
		CallMethod(s.js, el.setMethod, el.val)
		s.Equal(el.res, CallMethod(s.js, el.getMethod))
	}
}

func (s *PrimitiveOpsTestSuite) TestGetMethods() {
	//timeStr := "2015-01-01T12:34:56Z"
	//tm, _ := time.Parse(time.RFC3339, timeStr)
	emptyTime := time.Time{}
	tbl := []struct {
		idx       string
		val       interface{}
		setMethod string
		boolVal   bool
		intVal    int64
		uintVal   uint64
		floatVal  float64
		timeVal   time.Time
		stringVal string
	}{
		// bool
		{"bool:0", false, "SetBool", false, 0, 0, 0, emptyTime, "false"},
		{"bool:1", true, "SetBool", true, 1, 1, 1, emptyTime, "true"},
		// int
		{"int:0", int64(0), "SetInt", false, 0, 0, 0, emptyTime, "0"},
		{"int:1", int64(1), "SetInt", true, 1, 1, 1, emptyTime, "1"},
		{"int:2", int64(-1), "SetInt", true, -1, helper.MaxUint, -1, emptyTime, "-1"},
	}
	for _, el := range tbl {
		CallMethod(s.js, el.setMethod, el.val)
		s.Equal(el.boolVal, s.js.Bool(), "#%s %s(%v) => Bool() = %v != %v", el.idx, el.setMethod, el.val, el.boolVal, s.js.Bool(), el.boolVal)
		s.Equal(el.intVal, s.js.Int(), "#%s %s(%v) => Int() = %v != %v", el.idx, el.setMethod, el.val, s.js.Int(), el.intVal)
		s.Equal(el.uintVal, s.js.Uint(), "#%s %s(%v) => Uint() = %v != %v", el.idx, el.setMethod, el.val, s.js.Uint(), el.uintVal)
		s.Equal(el.floatVal, s.js.Float(), "#%s %s(%v) => Float() = %v != %v", el.idx, el.setMethod, el.val, s.js.Float(), el.floatVal)
		s.Equal(el.timeVal, s.js.Time(), "#%s %s(%v) => Time() = %v != %v", el.idx, el.setMethod, el.val, s.js.Time(), el.timeVal)
		s.Equal(el.stringVal, s.js.String(), "#%s %s(%v) => String() = %v != %v", el.idx, el.setMethod, el.val, s.js.String(), el.stringVal)
	}
}

func (s *PrimitiveOpsTestSuite) TestNullOps() {
	s.Equal(true, s.js.IsNull())
	s.js.SetInt(1)
	s.Equal(false, s.js.IsNull())
	s.js.SetNull()
	s.Equal(true, s.js.IsNull())
	// default value
	s.Equal(false, s.js.Bool())
	s.Equal(int64(0), s.js.Int())
	s.Equal(uint64(0), s.js.Uint())
	s.Equal(0.0, s.js.Float())
	s.Equal(time.Time{}, s.js.Time())
	s.Equal("null", s.js.String())
}

//func (s *PrimitiveOpsTestSuite) TestBoolOps() {
//	s.js.SetBool(true)
//	s.Equal(true, s.js.IsBool())
//	s.Equal(true, s.js.Bool())
//	s.Equal(int64(1), s.js.Int())
//	s.Equal(uint64(1), s.js.Uint())
//	s.Equal(1.0, s.js.Float())
//	t, _ := time.Parse(time.RFC3339, "true")
//	s.Equal(t, s.js.Time())
//	s.Equal("true", s.js.String())
//	s.js.SetBool(false)
//	s.Equal(false, s.js.Bool())
//	s.Equal(int64(0), s.js.Int())
//	s.Equal(uint64(0), s.js.Uint())
//	s.Equal(0.0, s.js.Float())
//	t, _ = time.Parse(time.RFC3339, "false")
//	s.Equal(t, s.js.Time())
//	s.Equal("false", s.js.String())
//}
//
//func (s *PrimitiveOpsTestSuite) TestIntOps() {
//	s.js.SetInt(1)
//	s.Equal(true, s.js.IsNumber())
//	s.Equal(true, s.js.IsInt())
//	s.Equal(false, s.js.IsFloat())
//
//	s.Equal(true, s.js.Bool())
//	s.Equal(int64(1), s.js.Int())
//	s.Equal(uint64(1), s.js.Uint())
//	s.Equal(1.0, s.js.Float())
//	t, _ := time.Parse(time.RFC3339, "1")
//	s.Equal(t, s.js.Time())
//	s.Equal("1", s.js.String())
//
//	s.js.SetInt(0)
//	s.Equal(false, s.js.Bool())
//}
//
//func (s *PrimitiveOpsTestSuite) TestUintOps() {
//	s.js.SetUint(1)
//	s.Equal(true, s.js.IsNumber())
//	s.Equal(true, s.js.IsUint())
//	s.Equal(false, s.js.IsFloat())
//
//	s.Equal(true, s.js.Bool())
//	s.Equal(int64(1), s.js.Int())
//	s.Equal(uint64(1), s.js.Uint())
//	s.Equal(1.0, s.js.Float())
//	t, _ := time.Parse(time.RFC3339, "1")
//	s.Equal(t, s.js.Time())
//	s.Equal("1", s.js.String())
//
//	s.js.SetInt(0)
//	s.Equal(false, s.js.Bool())
//}
//
//func (s *PrimitiveOpsTestSuite) TestFloatOps() {
//	s.js.SetFloat(3.1415926535897932385)
//	s.Equal(true, s.js.IsNumber())
//	s.Equal(false, s.js.IsInt())
//	s.Equal(true, s.js.IsFloat())
//
//	s.Equal(true, s.js.Bool())
//	s.Equal(int64(3), s.js.Int())
//	s.Equal(uint64(3), s.js.Uint())
//	s.Equal(3.141592653589793, s.js.Float())
//	t, _ := time.Parse(time.RFC3339, "3.141592653589793")
//	s.Equal(t, s.js.Time())
//	s.Equal("3.141592653589793", s.js.String())
//	s.js.SetFloat(0)
//	s.Equal(false, s.js.Bool())
//}
//
//func (s *PrimitiveOpsTestSuite) TestStringOps() {
//	s.js.SetString("hello")
//	s.Equal("hello", s.js.String())
//	s.Equal(true, s.js.IsString())
//
//	s.Equal(true, s.js.Bool())
//	s.Equal(int64(0), s.js.Int())
//	s.Equal(uint64(0), s.js.Uint())
//	s.Equal(0.0, s.js.Float())
//	t, _ := time.Parse(time.RFC3339, "hello")
//	s.Equal(t, s.js.Time())
//	s.Equal("hello", s.js.String())
//
//	s.js.SetString("")
//	s.Equal(false, s.js.Bool())
//	s.Equal(int64(0), s.js.Int())
//	s.Equal(uint64(0), s.js.Uint())
//	s.Equal(0.0, s.js.Float())
//	t, _ = time.Parse(time.RFC3339, "")
//	s.Equal(t, s.js.Time())
//	s.Equal("", s.js.String())
//
//	s.js.SetString("3.141592653589793")
//	s.Equal(true, s.js.Bool())
//	s.Equal(int64(3), s.js.Int())
//	s.Equal(uint64(3), s.js.Uint())
//	s.Equal(3.141592653589793, s.js.Float())
//	t, _ = time.Parse(time.RFC3339, "3.141592653589793")
//	s.Equal(t, s.js.Time())
//	s.Equal("3.141592653589793", s.js.String())
//
//}
//
//func (s *PrimitiveOpsTestSuite) TestTimeOps() {
//	timeStr := "2015-01-01T12:34:56Z"
//	tm, err := time.Parse(time.RFC3339, timeStr)
//	s.NoError(err)
//	s.js.SetTime(tm)
//	s.Equal(tm, s.js.Time())
//	s.Equal(true, s.js.IsTime())
//
//	s.js.SetString(timeStr)
//	s.Equal(tm, s.js.Time())
//}
