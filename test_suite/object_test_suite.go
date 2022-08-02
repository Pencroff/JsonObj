package test_suite

import (
	djs "github.com/Pencroff/JsonStruct"
	tl "github.com/Pencroff/JsonStruct/tool"
	"github.com/stretchr/testify/suite"
	"time"
)

type ObjectOpsTestSuite struct {
	suite.Suite
	factory func() djs.JStructOps
	js      djs.JStructOps
}

func (s *ObjectOpsTestSuite) SetFactory(fn func() djs.JStructOps) {
	s.factory = fn
}

func (s *ObjectOpsTestSuite) SetupTest() {
	if s.factory == nil {
		panic("factory not provided")
	}
	s.js = s.factory()
}

func (s *ObjectOpsTestSuite) TestIsAsObject() {
	s.Equal(true, s.js.IsNull())
	s.Equal(false, s.js.IsObject())
	s.js.AsObject()
	s.Equal(false, s.js.IsNull())
	s.Equal(true, s.js.IsObject())

	// default value
	s.Equal(false, s.js.Bool())
	s.Equal(int64(0), s.js.Int())
	s.Equal(uint64(0), s.js.Uint())
	s.Equal(0.0, s.js.Float())
	s.Equal(time.Time{}, s.js.Time())
	s.Equal("[object]", s.js.String())
}

func (s *ObjectOpsTestSuite) TestNullObject() {
	key := "someKey"
	s.Equal(true, s.js.IsNull())
	s.Equal(false, s.js.IsObject())
	v := s.js.GetKey(key)
	s.Nil(v)
	s.Equal(false, s.js.HasKey(key))
	s.js.AsObject()
	v = s.js.GetKey(key)
	s.Nil(v)
	s.Equal(false, s.js.HasKey(key))
}

func (s *ObjectOpsTestSuite) TestSetJustObject() {
	key := "someKey"
	s.Equal(true, s.js.IsNull())
	s.Equal(false, s.js.IsObject())
	e := s.js.SetKey(key, "value")
	s.ErrorIs(e, djs.NotObjectError)
	s.js.AsObject()
	s.Equal(false, s.js.IsNull())
	s.Equal(true, s.js.IsObject())
	e = s.js.SetKey(key, "value")
	s.Nil(e)
}

func (s *ObjectOpsTestSuite) TestSet() {
	timeStr := "2015-01-01T12:34:56Z"
	tm, _ := time.Parse(time.RFC3339, timeStr)
	obj := s.factory()
	obj.AsObject()
	tbl := []struct {
		key         string
		val         interface{}
		resValue    interface{}
		keyType     djs.Type
		e           error
		isMethod    string
		valueMethod string
	}{
		{"a1", true, true, djs.True, nil, "IsBool", "Bool"},
		{"a2", false, false, djs.False, nil, "IsBool", "Bool"},
		{"b1", int64(-1), int64(-1), djs.Int, nil, "IsInt", "Int"},
		{"b2", int32(-2), int64(-2), djs.Int, nil, "IsInt", "Int"},
		{"b3", int16(-4), int64(-4), djs.Int, nil, "IsInt", "Int"},
		{"b4", int8(-8), int64(-8), djs.Int, nil, "IsInt", "Int"},
		{"b5", 128, int64(128), djs.Int, nil, "IsInt", "Int"},
		{"c1", uint64(128), uint64(128), djs.Uint, nil, "IsUint", "Uint"},
		{"c2", uint32(64), uint64(64), djs.Uint, nil, "IsUint", "Uint"},
		{"c3", uint16(32), uint64(32), djs.Uint, nil, "IsUint", "Uint"},
		{"c4", uint8(16), uint64(16), djs.Uint, nil, "IsUint", "Uint"},
		{"c5", uint(8), uint64(8), djs.Uint, nil, "IsUint", "Uint"},
		{"d", 1.0, 1.0, djs.Float, nil, "IsFloat", "Float"},
		{"e", "hello", "hello", djs.String, nil, "IsString", "String"},
		{"f", tm, tm, djs.Time, nil, "IsTime", "Time"},
	}
	s.js.AsObject()
	for _, el := range tbl {

		e := s.js.SetKey(el.key, el.val)
		v := s.js.GetKey(el.key)
		s.Equal(el.e, e)

		s.Equal(el.keyType.String(), v.Type().String(), "key: %s", el.key)
		s.Equal(true, s.js.HasKey(el.key))
		s.Equal(true, tl.CallMethod(v, el.isMethod))
		s.Equal(el.resValue, tl.CallMethod(v, el.valueMethod))
	}

	// default value
	s.Equal(false, s.js.Bool())
	s.Equal(int64(0), s.js.Int())
	s.Equal(uint64(0), s.js.Uint())
	s.Equal(0.0, s.js.Float())
	s.Equal(time.Time{}, s.js.Time())
	s.Equal("[object]", s.js.String())
}

func (s *ObjectOpsTestSuite) TestSetUnsupportedType() {
	t := map[string]interface{}{"boolKey": true, "intKey": -10, "uintKey": 10}
	key := "someKey"
	s.js.AsObject()

	err := s.js.SetKey(key, t)
	s.ErrorIs(err, djs.UnsupportedTypeError)
	s.Equal(false, s.js.HasKey(key))
}

func (s *ObjectOpsTestSuite) TestRemoveKey() {
	s.js.AsObject()
	s.js.SetKey("a", "aValue")
	s.js.SetKey("b", "bValue")
	s.js.SetKey("c", "cValue")
	keys := s.js.Keys()
	s.Equal(3, len(keys))
	s.Equal(true, s.js.HasKey("b"))
	v := s.js.RemoveKey("b")
	keys = s.js.Keys()
	s.Equal(v.Value(), "bValue")
	s.Equal(2, len(keys))
	s.Equal(false, s.js.HasKey("b"))
	v = s.js.RemoveKey("someKey")
	s.Nil(v)
}

func (s *ObjectOpsTestSuite) TestKeys() {
	timeStr := "2015-01-01T12:34:56Z"
	tm, _ := time.Parse(time.RFC3339, timeStr)
	tbl := []struct {
		key string
		val interface{}
	}{
		{"a", true},
		{"b", int64(-1)},
		{"c", uint64(1)},
		{"d", 1.0},
		{"e", "hello"},
		{"f", tm},
	}
	s.js.AsObject()
	keys := s.js.Keys()
	s.Equal(0, len(keys))
	for _, el := range tbl {
		s.js.SetKey(el.key, el.val)
	}
	keys = s.js.Keys()
	s.Len(keys, len(tbl))
	for _, el := range tbl {
		s.Contains(keys, el.key)
		v := s.js.GetKey(el.key)
		s.Equal(0, len(v.Keys()))
	}
}
