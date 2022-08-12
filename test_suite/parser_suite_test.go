package test_suite

import (
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ParserTestCase struct {
	idx      string
	data     []byte
	err      error
	validate func(t *testing.T, el ParserTestCase, js djs.JStructOps)
}

func TestJsonStructConverter_ParserTestSuite(t *testing.T) {
	s := new(ParserTestSuite)
	s.SetFactory(JsonStructFactory)
	suite.Run(t, s)
}

type ParserTestSuite struct {
	suite.Suite
	factory func() djs.JStructOps
	mock    *MockedParser
}

func (s *ParserTestSuite) SetFactory(fn func() djs.JStructOps) {
	s.factory = fn
}

func (s *ParserTestSuite) SetupTest() {
	if s.factory == nil {
		panic("factory not provided")
	}
}

func (s *ParserTestSuite) TestParsing_PrimitiveValues() {
	testCases := []ParserTestCase{
		{"null", []byte(`null`), nil, func(t *testing.T, el ParserTestCase, js djs.JStructOps) {
			assert.True(t, js.IsNull(), "%s IsNull() != true", el.idx)
		}},
		{"true", []byte(`true`), nil, func(t *testing.T, el ParserTestCase, js djs.JStructOps) {
			assert.True(t, js.IsBool(), "%s IsBool() != true", el.idx)
			assert.True(t, js.Bool(), "%s Bool() != true", el.idx)
		}},
		{"false", []byte(`false`), nil, func(t *testing.T, el ParserTestCase, js djs.JStructOps) {
			assert.True(t, js.IsBool(), "%s IsBool() != true", el.idx)
			assert.False(t, js.Bool(), "%s Bool() != false", el.idx)
		}},
	}
	for _, el := range testCases {
		RunParserTestCase(s, el)
	}
}

func RunParserTestCase(s *ParserTestSuite, el ParserTestCase) {
	s.T().Run(el.idx, func(t *testing.T) {
		js := s.factory()
		js.AsObject()
		err := djs.UnmarshalJSON(el.data, js)
		assert.Equal(t, err, el.err, "%s err != %v", el.idx, el.err)
		el.validate(t, el, js)
	})
}

//func (s *ParserTestSuite) TestParsing_PrimitiveValues() {
//	tm := time.Date(2015, 5, 14, 12, 34, 56, 379000000, time.FixedZone("CEST", 2*60*60))
//	unixStart := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
//	emptyTime := time.Time{}
//	tm1, _ := time.Parse(time.RFC3339, "1985-04-12T23:20:50.52Z")
//	tm2, _ := time.Parse(time.RFC3339, "1996-12-19T16:39:57-08:00")
//	tm3, _ := time.Parse(time.RFC3339, "1990-12-31T23:59:60Z")
//	tm4, _ := time.Parse(time.RFC3339, "1990-12-31T15:59:60-08:00")
//	tm5, _ := time.Parse(time.RFC3339, "1937-01-01T12:00:27.87+00:20")
//	tbl := []struct {
//		idx    string
//		in     []byte
//		value  interface{}
//		method string
//	}{
//		{"null:0", []byte(`null`), nil, "SetNull"},
//		// ----------------------------------------------------------
//		{"bool:0", []byte(`true`), true, "SetBool"},
//		{"bool:1", []byte(`false`), false, "SetBool"},
//		// ----------------------------------------------------------
//		{"int:0", []byte(`0`), int64(0), "SetInt"},
//		{"int:1", []byte(`1`), int64(1), "SetInt"},
//		{"int:2", []byte(`-1`), int64(-1), "SetInt"},
//		{"int:3", []byte(`9223372036854775807`), helper.MaxInt, "SetInt"},
//		{"int:4", []byte(`-9223372036854775808`), helper.MinInt, "SetInt"},
//		// ----------------------------------------------------------
//		{"uint:0", []byte(`9223372036854775808`), uint64(9223372036854775808), "SetUint"},
//		{"uint:1", []byte(`18446744073709551615`), helper.MaxUint, "SetUint"},
//		// ----------------------------------------------------------
//		{"float:0", []byte(`0.0`), float64(0.0), "SetFloat"},
//		{"float:1", []byte(`3.1415`), float64(3.1415), "SetFloat"},
//		{"float:2", []byte(`-3.1415`), float64(-3.1415), "SetFloat"},
//		{"float:3", []byte(`1.0e+308`), float64(1.0e+308), "SetFloat"},
//		{"float:4", []byte(`-1.0e+308`), float64(-1.0e+308), "SetFloat"},
//		{"float:5", []byte(`1.0e-308`), float64(1.0e-308), "SetFloat"},
//		{"float:6", []byte(`-1.0e-308`), float64(-1.0e-308), "SetFloat"},
//		// ----------------------------------------------------------
//		{"string:0", []byte(`"hello"`), "hello", "SetString"},
//		{"string:1", []byte(`"hello world"`), "hello world", "SetString"},
//		{"string:2", []byte(`"hello\nworld"`), "hello\nworld", "SetString"},
//		{"string:3", []byte(`"hello\rworld"`), "hello\rworld", "SetString"},
//		{"string:4", []byte(`"hello\tworld"`), "hello\tworld", "SetString"},
//		{"string:5", []byte(`"hello\bworld"`), "hello\bworld", "SetString"},
//		{"string:6", []byte(`"hello\fworld"`), "hello\fworld", "SetString"},
//		{"string:7", []byte(`"hello\u0020world"`), "hello world", "SetString"},
//		// ----------------------------------------------------------
//		{"time:0", []byte(`"2015-05-14T12:34:56.379+02:00"`), tm, "SetTime"},
//		{"time:1", []byte(`"1970-01-01T00:00:00Z"`), unixStart, "SetTime"},
//		{"time:2", []byte(`"0001-01-01T00:00:00Z"`), emptyTime, "SetTime"},
//		{"time:3", []byte(`"1985-04-12T23:20:50.52Z"`), tm1, "SetTime"},
//		{"time:4", []byte(`"1996-12-19T16:39:57-08:00"`), tm2, "SetTime"},
//		{"time:5", []byte(`"1990-12-31T23:59:60Z"`), tm3, "SetTime"},
//		{"time:6", []byte(`"1990-12-31T15:59:60-08:00"`), tm4, "SetTime"},
//		{"time:7", []byte(`"1937-01-01T12:00:27.87+00:20"`), tm5, "SetTime"},
//	}
//	for _, el := range tbl {
//		rd := bytes.NewReader(el.in)
//		e := djs.JStructParseFn(rd, s.js)
//		v := s.factory()
//		if el.value == nil {
//			tl.CallMethod(v, el.method)
//		} else {
//			tl.CallMethod(v, el.method, el.value)
//		}
//		s.NoError(e)
//		s.Equal(v.Value(), s.js.Value(), "%s %v != %v", el.idx, v.Value(), s.js.Value())
//	}
//}
