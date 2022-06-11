package test_suite

import (
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/suite"
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
	s.Equal(0, s.testJS.Int())
	s.Equal(0.0, s.testJS.Float())
	s.Equal(time.Time{}, s.testJS.Time())
	s.Equal("null", s.testJS.String())
}

func (s *ObjectOpsTestSuite) TestSetObject() {
	s.Equal(true, s.testJS.IsNull())
	s.testJS.Set("boolKey", true)
	s.testJS.Set("intKey", -10)
	s.testJS.Set("uintKey", 10)
	s.testJS.Set("floatKey", 1.5)
	s.testJS.Set("stringKey", "string")
	s.testJS.Set("timeKey", time.Now())

	// default value
	s.Equal(false, s.testJS.Bool())
	s.Equal(0, s.testJS.Int())
	s.Equal(0.0, s.testJS.Float())
	s.Equal(time.Time{}, s.testJS.Time())
	s.Equal("null", s.testJS.String())
}
