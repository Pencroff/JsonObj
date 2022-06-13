package test_suite

import (
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/suite"
	"time"
)

type PrimitiveOpsTestSuite struct {
	suite.Suite
	factory func() djs.JsonStructOps
	testJS  djs.PrimitiveOps
}

func (s *PrimitiveOpsTestSuite) SetFactory(fn func() djs.JsonStructOps) {
	s.factory = fn
}

func (s *PrimitiveOpsTestSuite) SetupTest() {
	if s.factory == nil {
		panic("factory not provided")
	}
	s.testJS = s.factory()
}

func (s *PrimitiveOpsTestSuite) TestNullOps() {
	s.Equal(true, s.testJS.IsNull())
	s.testJS.SetInt(1)
	s.Equal(false, s.testJS.IsNull())
	s.testJS.SetNull()
	s.Equal(true, s.testJS.IsNull())
	// default value
	s.Equal(false, s.testJS.Bool())
	s.Equal(int64(0), s.testJS.Int())
	s.Equal(uint64(0), s.testJS.Uint())
	s.Equal(0.0, s.testJS.Float())
	s.Equal(time.Time{}, s.testJS.Time())
	s.Equal("null", s.testJS.String())
}

func (s *PrimitiveOpsTestSuite) TestBoolOps() {
	s.testJS.SetBool(true)
	s.Equal(true, s.testJS.IsBool())
	s.Equal(true, s.testJS.Bool())
	s.Equal(int64(1), s.testJS.Int())
	s.Equal(1.0, s.testJS.Float())
	t, _ := time.Parse(time.RFC3339, "true")
	s.Equal(t, s.testJS.Time())
	s.Equal("true", s.testJS.String())
	s.testJS.SetBool(false)
	s.Equal(false, s.testJS.Bool())
	s.Equal(int64(0), s.testJS.Int())
	s.Equal(uint64(0), s.testJS.Uint())
	s.Equal(0.0, s.testJS.Float())
	t, _ = time.Parse(time.RFC3339, "false")
	s.Equal(t, s.testJS.Time())
	s.Equal("false", s.testJS.String())
}

func (s *PrimitiveOpsTestSuite) TestIntOps() {
	s.testJS.SetInt(1)
	s.Equal(true, s.testJS.IsNumber())
	s.Equal(true, s.testJS.IsInt())
	s.Equal(false, s.testJS.IsFloat())

	s.Equal(true, s.testJS.Bool())
	s.Equal(int64(1), s.testJS.Int())
	s.Equal(uint64(1), s.testJS.Uint())
	s.Equal(1.0, s.testJS.Float())
	t, _ := time.Parse(time.RFC3339, "1")
	s.Equal(t, s.testJS.Time())
	s.Equal("1", s.testJS.String())

	s.testJS.SetInt(0)
	s.Equal(false, s.testJS.Bool())
}

func (s *PrimitiveOpsTestSuite) TestFloatOps() {
	s.testJS.SetFloat(3.1415926535897932385)
	s.Equal(true, s.testJS.IsNumber())
	s.Equal(false, s.testJS.IsInt())
	s.Equal(true, s.testJS.IsFloat())

	s.Equal(true, s.testJS.Bool())
	s.Equal(int64(3), s.testJS.Int())
	s.Equal(uint64(3), s.testJS.Uint())
	s.Equal(3.141592653589793, s.testJS.Float())
	t, _ := time.Parse(time.RFC3339, "3.141592653589793")
	s.Equal(t, s.testJS.Time())
	s.Equal("3.141592653589793", s.testJS.String())
	s.testJS.SetFloat(0)
	s.Equal(false, s.testJS.Bool())
}

func (s *PrimitiveOpsTestSuite) TestStringOps() {
	s.testJS.SetString("hello")
	s.Equal("hello", s.testJS.String())
	s.Equal(true, s.testJS.IsString())

	s.Equal(true, s.testJS.Bool())
	s.Equal(int64(0), s.testJS.Int())
	s.Equal(uint64(0), s.testJS.Uint())
	s.Equal(0.0, s.testJS.Float())
	t, _ := time.Parse(time.RFC3339, "hello")
	s.Equal(t, s.testJS.Time())
	s.Equal("hello", s.testJS.String())

	s.testJS.SetString("")
	s.Equal(false, s.testJS.Bool())
	s.Equal(int64(0), s.testJS.Int())
	s.Equal(uint64(0), s.testJS.Uint())
	s.Equal(0.0, s.testJS.Float())
	t, _ = time.Parse(time.RFC3339, "")
	s.Equal(t, s.testJS.Time())
	s.Equal("", s.testJS.String())

	s.testJS.SetString("3.141592653589793")
	s.Equal(true, s.testJS.Bool())
	s.Equal(int64(3), s.testJS.Int())
	s.Equal(uint64(3), s.testJS.Uint())
	s.Equal(3.141592653589793, s.testJS.Float())
	t, _ = time.Parse(time.RFC3339, "3.141592653589793")
	s.Equal(t, s.testJS.Time())
	s.Equal("3.141592653589793", s.testJS.String())

}

func (s *PrimitiveOpsTestSuite) TestTimeOps() {
	timeStr := "2015-01-01T12:34:56Z"
	tm, err := time.Parse(time.RFC3339, timeStr)
	s.NoError(err)
	s.testJS.SetTime(tm)
	s.Equal(tm, s.testJS.Time())
	s.Equal(true, s.testJS.IsTime())

	s.testJS.SetString(timeStr)
	s.Equal(tm, s.testJS.Time())
}
