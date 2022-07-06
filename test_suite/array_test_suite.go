package test_suite

import (
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/suite"
	"time"
)

type ArrayOpsTestSuite struct {
	suite.Suite
	factory func() djs.JStructOps
	js      djs.JStructOps
}

func (s *ArrayOpsTestSuite) SetFactory(fn func() djs.JStructOps) {
	s.factory = fn
}

func (s *ArrayOpsTestSuite) SetupTest() {
	if s.factory == nil {
		panic("factory not provided")
	}
	s.js = s.factory()
}

func (s *ArrayOpsTestSuite) TestPushPop() {
	timeStr := "2015-01-01T12:34:56Z"
	tm, _ := time.Parse(time.RFC3339, timeStr)
	tbl := []struct {
		value    interface{}
		resValue interface{}
	}{
		{nil, nil},
		{false, false},
		{true, true},
		{int64(1), int64(1)},
		{uint64(1), uint64(1)},
		{1.0, 1.0},
		{"1", "1"},
		{tm, tm},
	}
	e := s.js.Push(nil)
	s.ErrorIs(e, djs.NotArrayError)
	s.js.AsArray()
	for _, el := range tbl {
		s.js.Push(el.value)
		s.Equal(1, s.js.Size())
		v := s.js.Pop()
		s.Equal(el.resValue, v.Value())
	}
}
func (s *ArrayOpsTestSuite) TestPushPopObjectArray() {
	obj := s.factory()
	obj.AsObject()
	obj.SetKey("a", 1)
	obj.SetKey("b", "2")
	arr := s.factory()
	arr.AsArray()
	arr.Push(1)
	arr.Push("2")
	tbl := []struct {
		value interface{}
		err   error
	}{
		{obj, nil},
		{arr, nil},
	}
	s.js.AsArray()
	for _, el := range tbl {
		s.js.Push(el.value)
		s.Equal(1, s.js.Size())
		v := s.js.Pop()
		s.Equal(el.value, v)
	}
	v := s.js.Pop()
	s.Equal(nil, v)
	v = obj.Pop()
	s.Equal(nil, v)
}

func (s *ArrayOpsTestSuite) TestPushPopUnsupportedType() {
	someObj := map[string]interface{}{"boolKey": true, "intKey": -10, "uintKey": 10}
	s.js.AsArray()
	e := s.js.Push(someObj)
	s.ErrorIs(e, djs.UnsupportedTypeError)
	v := s.js.Pop()
	s.Equal(nil, v)
}

func (s *ArrayOpsTestSuite) TestShiftPop() {
	s.js.AsArray()
	s.js.Push(1)
	s.js.Push(2)
	s.js.Push(3)
	v := s.js.Pop()
	s.Equal(int64(3), v.Value())
	v = s.js.Shift()
	s.Equal(int64(1), v.Value())
	v = s.js.Shift()
	s.Equal(int64(2), v.Value())
	v = s.js.Shift()
	s.Equal(nil, v)
	js := s.factory()
	v = js.Shift()
	s.Equal(nil, v)
}

func (s *ArrayOpsTestSuite) TestGetIndex() {
	s.js.AsArray()
	s.js.Push(1)
	s.js.Push(2)
	s.js.Push(3)
	v := s.js.GetIndex(0)
	s.Equal(int64(1), v.Value())
	v = s.js.GetIndex(2)
	s.Equal(int64(3), v.Value())
	v = s.js.GetIndex(3)
	s.Equal(nil, v)
	v = s.js.GetIndex(100)
	s.Equal(nil, v)
	js := s.factory()
	v = js.GetIndex(0)
	s.Equal(nil, v)
}
func (s *ArrayOpsTestSuite) TestSetIndex() {
	s.js.AsArray()
	s.js.Push(1)
	s.js.Push(2)
	e := s.js.SetIndex(5, 3)
	s.Equal(nil, e)
	e = s.js.SetIndex(10, 4)
	s.Equal(nil, e)
	e = s.js.SetIndex(-1, 3)
	s.ErrorIs(e, djs.IndexOutOfRangeError)
	s.Equal(11, s.js.Size())
	v := s.js.GetIndex(0)
	s.Equal(int64(1), v.Value())
	v = s.js.GetIndex(1)
	s.Equal(int64(2), v.Value())
	v = s.js.GetIndex(3)
	s.Equal(nil, v)
	v = s.js.GetIndex(5)
	s.Equal(int64(3), v.Value())
	v = s.js.GetIndex(10)
	s.Equal(int64(4), v.Value())
	v = s.js.GetIndex(9)
	s.Equal(nil, v)
	js := s.factory()
	e = js.SetIndex(0, 1)
	s.ErrorIs(e, djs.NotArrayError)
}

func (s *ObjectOpsTestSuite) TestSetIndexUnsupportedType() {
	t := map[string]interface{}{"boolKey": true, "intKey": -10, "uintKey": 10}
	s.js.AsArray()
	err := s.js.SetIndex(0, t)
	s.ErrorIs(err, djs.UnsupportedTypeError)
	s.Equal(0, s.js.Size())
}
