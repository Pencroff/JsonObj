package test_suite

import (
	"fmt"
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
		s.Equal(el.isString, s.js.IsString())
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
	tm := time.Date(2015, 5, 14, 12, 34, 56, 379000000, time.FixedZone("CEST", 2*60*60))
	fmt.Printf("%v\n", tm.Format(time.RFC3339))
	unixStart := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("%v\n", unixStart.Format(time.RFC3339))
	emptyTime := time.Time{}
	fmt.Printf("%v\n", emptyTime.Format(time.RFC3339))
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
		{"int:3", helper.MaxInt, "SetInt", true, helper.MaxInt, helper.MaxIntUint, float64(helper.MaxInt), emptyTime, fmt.Sprintf("%d", helper.MaxInt)},
		{"int:4", helper.MinInt, "SetInt", true, helper.MinInt, helper.MinIntUint, float64(helper.MinInt), emptyTime, fmt.Sprintf("%d", helper.MinInt)},
		// uint
		{"uint:0", uint64(0), "SetUint", false, 0, 0, 0, emptyTime, "0"},
		{"uint:1", uint64(1), "SetUint", true, 1, 1, 1, emptyTime, "1"},
		{"uint:2", helper.MaxUint, "SetUint", true, -1, helper.MaxUint, float64(helper.MaxUint), emptyTime, fmt.Sprintf("%d", helper.MaxUint)},
		// float
		{"float:0", 0.0, "SetFloat", false, 0, 0, 0, emptyTime, "0"},
		{"float:1", 1.0, "SetFloat", true, 1, 1, 1, emptyTime, "1"},
		{"float:2", -1.0, "SetFloat", true, -1, helper.MaxUint, -1, emptyTime, "-1"},
		{"float:3", 3.1415926535897932385, "SetFloat", true, 3, 3, 3.141592653589793, emptyTime, "3.141592653589793"},
		{"float:4", -3.1415926535897932385, "SetFloat", true, -3, helper.MaxUint - 2, -3.141592653589793, emptyTime, "-3.141592653589793"},
		{"float:5", 3.1415, "SetFloat", true, 3, 3, 3.1415, emptyTime, "3.1415"},
		// time
		{"time:0", emptyTime, "SetTime", true, emptyTime.UnixMilli(), uint64(emptyTime.UnixMilli()), 0, emptyTime, "0001-01-01T00:00:00Z"},
		{"time:1", tm, "SetTime", true, tm.UnixMilli(), uint64(tm.UnixMilli()), 0, tm, "2015-05-14T12:34:56+02:00"},
		{"time:2", unixStart, "SetTime", false, 0, 0, 0, unixStart, "1970-01-01T00:00:00Z"},
		// string
		{"string:0", "", "SetString", false, 0, 0, 0, emptyTime, ""},
		{"string:1", "hello", "SetString", true, 0, 0, 0, emptyTime, "hello"},
		{"string:2", "3.1415926535897932385", "SetString", true, 3, 3, 3.141592653589793, emptyTime, "3.1415926535897932385"},
		{"string:3", "3.1415", "SetString", true, 3, 3, 3.1415, emptyTime, "3.1415"},
		{"string:4", "-3.1415", "SetString", true, -3, 0, -3.1415, emptyTime, "-3.1415"},
	}
	for _, el := range tbl {
		CallMethod(s.js, el.setMethod, el.val)
		s.Equal(el.boolVal, s.js.Bool(), "#%s %s(%v) => Bool() = %v != %v", el.idx, el.setMethod, el.val, s.js.Bool(), el.boolVal)
		s.Equal(el.intVal, s.js.Int(), "#%s %s(%v) => Int() = %v != %v", el.idx, el.setMethod, el.val, s.js.Int(), el.intVal)
		s.Equal(el.uintVal, s.js.Uint(), "#%s %s(%v) => Uint() = %v != %v", el.idx, el.setMethod, el.val, s.js.Uint(), el.uintVal)
		s.Equal(el.floatVal, s.js.Float(), "#%s %s(%v) => Float() = %v != %v", el.idx, el.setMethod, el.val, s.js.Float(), el.floatVal)
		s.Equal(el.timeVal, s.js.Time(), "#%s %s(%v) => Time() = %v != %v", el.idx, el.setMethod, el.val, s.js.Time(), el.timeVal)
		s.Equal(el.stringVal, s.js.String(), "#%s %s(%v) => String() = %v != %v", el.idx, el.setMethod, el.val, s.js.String(), el.stringVal)
	}
}
