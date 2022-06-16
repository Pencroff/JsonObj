package test_suite

import (
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/suite"
	"time"
)

type GeneralOpsTestSuite struct {
	suite.Suite
	factory func() djs.JsonStructOps
	js      djs.JsonStructOps
}

func (s *GeneralOpsTestSuite) SetFactory(fn func() djs.JsonStructOps) {
	s.factory = fn
}

func (s *GeneralOpsTestSuite) SetupTest() {
	if s.factory == nil {
		panic("factory not provided")
	}
	s.js = s.factory()
}

func (s *GeneralOpsTestSuite) TestTypeValue() {
	timeStr := "2015-01-01T12:34:56Z"
	tm, _ := time.Parse(time.RFC3339, timeStr)

	tbl := []struct {
		valueType djs.Type
		value     interface{}
		method    string
	}{
		{djs.False, false, "SetBool"},
		{djs.True, true, "SetBool"},
		{djs.Int, int64(1), "SetInt"},
		{djs.Uint, uint64(1), "SetUint"},
		{djs.Float, float64(1.0), "SetFloat"},
		{djs.String, "1", "SetString"},
		{djs.Time, tm, "SetTime"},
	}

	s.Equal(true, s.js.IsNull())
	for _, el := range tbl {
		CallMethod(s.js, el.method, el.value)

		s.Equal(el.valueType, s.js.Type(), "%s(%v) => %v type", el.method, el.value, el.valueType)
		s.Equal(el.value, s.js.Value(), "%s(%v) => %v type", el.method, el.value, el.valueType)
	}
}

func (s *GeneralOpsTestSuite) TestAsTypeValue() {
	obj := make(map[string]djs.JsonStructOps)
	arr := make([]djs.JsonStructOps, 0)
	tbl := []struct {
		valueType djs.Type
		value     interface{}
		method    string
		isMethod  string
	}{
		{djs.Null, nil, "SetNull", "IsNull"},
		{djs.Object, obj, "AsObject", "IsObject"},
		{djs.Array, arr, "AsArray", "IsArray"},
	}

	s.Equal(true, s.js.IsNull())
	for _, el := range tbl {
		CallMethod(s.js, el.method)

		s.Equal(el.valueType, s.js.Type())
		s.Equal(el.value, s.js.Value())
		s.Equal(true, CallMethod(s.js, el.isMethod))
	}
}

func (s *GeneralOpsTestSuite) TestSetKeyNullObjArr() {
	null := s.factory()
	obj := s.factory()
	obj.AsObject()
	arr := s.factory()
	arr.AsArray()
	tbl := []struct {
		key      string
		value    djs.JsonStructOps
		resValue djs.JsonStructOps
		isMethod string
		keyType  djs.Type
	}{
		{"a", nil, null, "IsNull", djs.Null},
		{"b", obj, obj, "IsObject", djs.Object},
		{"c", arr, arr, "IsArray", djs.Array},
	}

	s.Equal(true, s.js.IsNull())
	s.js.AsObject()
	for _, el := range tbl {
		err := s.js.SetKey(el.key, el.value)
		v := s.js.GetKey(el.key)
		s.NoError(err)
		s.Equal(el.keyType.String(), v.Type().String(), "[%s]: %s != %s", el.key, el.keyType, v.Type())
		s.Equal(true, s.js.HasKey(el.key))
		s.Equal(el.resValue.Value(), v.Value())
		s.Equal(true, CallMethod(v, el.isMethod))
	}
}

func (s *GeneralOpsTestSuite) TestSizePrimitive() {
	obj := s.factory()
	obj.AsObject()

	arr := s.factory()
	arr.AsArray()
	tbl := []struct {
		value  interface{}
		size   int
		method string
	}{
		{false, -1, "SetBool"},
		{true, -1, "SetBool"},
		{int64(1), -1, "SetInt"},
		{uint64(1), -1, "SetUint"},
		{1.0, -1, "SetFloat"},
		{"hello", 5, "SetString"},
		{time.Now(), -1, "SetTime"},
	}
	for _, el := range tbl {
		CallMethod(s.js, el.method, el.value)

		s.Equal(el.size, s.js.Size(), "len(%v) => %v", el.value, el.size)
	}
}

func (s *GeneralOpsTestSuite) TestSizeArrayObj() {
	emptyObj := s.factory()
	emptyObj.AsObject()
	obj := s.factory()
	obj.AsObject()
	obj.SetKey("a", true)
	obj.SetKey("b", 1)
	obj.SetKey("c", 1.0)

	emptyArr := s.factory()
	emptyArr.AsArray()
	arr := s.factory()
	arr.AsArray()
	arr.Push(true)
	arr.Push(1)
	arr.Push(1.0)
	arr.Push(obj)
	tbl := []struct {
		value interface{}
		size  int
	}{
		{emptyObj, 0},
		{obj, 3},
		{emptyArr, 0},
		{arr, 4},
	}
	for _, el := range tbl {
		s.Equal(el.size, el.value.(djs.JsonStructOps).Size(), "len(%v) => %v", el.value, el.size)
	}
}
