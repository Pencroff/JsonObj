package test_suite

import (
	"fmt"
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/suite"
	"reflect"
	"time"
)

type ObjectOpsTestSuite struct {
	suite.Suite
	factory func() djs.JsonStructOps
	testJS  djs.JsonStructOps
}

func (s *ObjectOpsTestSuite) SetFactory(fn func() djs.JsonStructOps) {
	s.factory = fn
}

func (s *ObjectOpsTestSuite) SetupTest() {
	if s.factory == nil {
		panic("factory not provided")
	}
	s.testJS = s.factory()
}

func (s *ObjectOpsTestSuite) TestIsAsObject() {
	s.Equal(true, s.testJS.IsNull())
	s.Equal(false, s.testJS.IsObject())
	s.testJS.AsObject()
	s.Equal(false, s.testJS.IsNull())
	s.Equal(true, s.testJS.IsObject())

	// default value
	s.Equal(false, s.testJS.Bool())
	s.Equal(int64(0), s.testJS.Int())
	s.Equal(uint64(0), s.testJS.Uint())
	s.Equal(0.0, s.testJS.Float())
	s.Equal(time.Time{}, s.testJS.Time())
	s.Equal("[object]", s.testJS.String())
}

func (s *ObjectOpsTestSuite) TestNullObject() {
	key := "someKey"
	s.Equal(true, s.testJS.IsNull())
	s.Equal(false, s.testJS.IsObject())
	v := s.testJS.Get(key)
	s.Nil(v)
	s.Equal(false, s.testJS.Has(key))
	s.testJS.AsObject()
	v = s.testJS.Get(key)
	s.Nil(v)
	s.Equal(false, s.testJS.Has(key))
}

func (s *ObjectOpsTestSuite) TestSetJustObject() {
	key := "someKey"
	s.Equal(true, s.testJS.IsNull())
	s.Equal(false, s.testJS.IsObject())
	e := s.testJS.Set(key, "value")
	s.ErrorIs(e, djs.NotObjectError)
	s.testJS.AsObject()
	s.Equal(false, s.testJS.IsNull())
	s.Equal(true, s.testJS.IsObject())
	e = s.testJS.Set(key, "value")
	s.Nil(e)
}

func (s *ObjectOpsTestSuite) TestSetObject() {
	timeStr := "2015-01-01T12:34:56Z"
	tm, err := time.Parse(time.RFC3339, timeStr)
	s.NoError(err)
	s.Equal(true, s.testJS.IsNull())
	obj := s.factory()
	obj.AsObject()
	tbl := []struct {
		key         string
		val         interface{}
		keyType     djs.Type
		e           error
		isMethod    string
		valueMethod string
	}{
		{"a", true, djs.True, nil, "IsBool", "Bool"},
		{"b", int64(-1), djs.Int, nil, "IsInt", "Int"},
		{"c", uint64(1), djs.Uint, nil, "IsUint", "Uint"},
		{"d", 1.0, djs.Float, nil, "IsFloat", "Float"},
		{"e", "hello", djs.String, nil, "IsString", "String"},
		{"f", tm, djs.Time, nil, "IsTime", "Time"},
		{"g", nil, djs.Null, nil, "IsNull", ""},
		//{"h",
		//	map[string]interface{}{"boolKey": true, "intKey": -10, "uintKey": 10},
		//	djs.Null, djs.UnsupportedTypeError, "", ""},
		{"h", obj, djs.Object, nil, "IsObject", ""},
	}
	s.testJS.AsObject()
	for _, t := range tbl {
		fmt.Printf("%+v\n", t)
		e := s.testJS.Set(t.key, t.val)
		v := s.testJS.Get(t.key)
		s.Equal(t.e, e)

		s.Equal(t.keyType.String(), v.Type().String(), "key: %s", t.key)
		s.Equal(true, s.testJS.Has(t.key))
		if t.isMethod != "" {
			s.Equal(true, CallMethod(v, t.isMethod))
		}
		if t.valueMethod != "" {
			s.Equal(t.val, CallMethod(v, t.valueMethod))
		}

	}
	// default value
	s.Equal(false, s.testJS.Bool())
	s.Equal(int64(0), s.testJS.Int())
	s.Equal(uint64(0), s.testJS.Uint())
	s.Equal(0.0, s.testJS.Float())
	s.Equal(time.Time{}, s.testJS.Time())
	s.Equal("[object]", s.testJS.String())
}

func CallMethod(obj interface{}, method string) interface{} {
	ptr := reflect.ValueOf(obj)
	methodValue := ptr.MethodByName(method)
	res := methodValue.Call([]reflect.Value{})
	return res[0].Interface()
}
