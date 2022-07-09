package test_suite

import (
	"bufio"
	"bytes"
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ScannerTestSuite struct {
	suite.Suite
}

func (s *ScannerTestSuite) SetupTest() {

}

func (s *ScannerTestSuite) TestNewScannerInstance() {
	rd := &bufio.Reader{}
	sc := djs.NewJStructScanner(rd)
	s.Equal(cap(sc.Value()), djs.JStructScannerBufferSize)
	sc = djs.NewJStructScannerWithSize(rd, 100)
	s.Equal(cap(sc.Value()), 100)
}

func (s *ScannerTestSuite) TestScanner_Next() {
	tbl := []struct {
		idx  string
		in   []byte
		kind djs.ScannerKind
		out  []byte
	}{
		{"null:0", []byte("null"), djs.ScanNull, []byte("null")},
		// ------------------------------------------------------
		{"bool:f", []byte("false"), djs.ScanFalse, []byte("false")},
		{"bool:t", []byte("true"), djs.ScanTrue, []byte("true")},
		// ------------------------------------------------------
		//{"num:00", []byte("123"), djs.ScanIntNumber, []byte("123")},
		//{"num:01", []byte("0"), djs.ScanIntNumber, []byte("0")},
		//{"num:02", []byte("-0"), djs.ScanIntNumber, []byte("-0")},
		//{"num:03", []byte("1"), djs.ScanIntNumber, []byte("1")},
		//{"num:04", []byte("-1"), djs.ScanIntNumber, []byte("-1")},
		//{"num:05", []byte("123456789"), djs.ScanIntNumber, []byte("123456789")},
		//{"num:06", []byte("-123456789"), djs.ScanIntNumber, []byte("-123456789")},
		//{"num:07", []byte("9223372036854775807"), djs.ScanIntNumber, []byte("9223372036854775807")},
		//{"num:08", []byte("-9223372036854775808"), djs.ScanIntNumber, []byte("-9223372036854775808")},
		//{"num:09", []byte("9223372036854775808"), djs.ScanIntNumber, []byte("9223372036854775808")},
		//{"num:10", []byte("-9223372036854775809"), djs.ScanIntNumber, []byte("-9223372036854775809")},     // -9.223372036854776e+18
		//{"num:11", []byte("18446744073709551615"), djs.ScanFloatNumber, []byte("18446744073709551615")},   // 1.8446744073709552e+19
		//{"num:11", []byte("-18446744073709551615"), djs.ScanFloatNumber, []byte("-18446744073709551615")}, // -1.8446744073709552e+19
		//// ------------------------------------------------------
		//{"float:00", []byte("123.45"), djs.ScanFloatNumber, []byte("123.45")},
		//{"float:01", []byte("0.0"), djs.ScanFloatNumber, []byte("0.0")},
		//{"float:02", []byte("-0.0"), djs.ScanFloatNumber, []byte("-0.0")},
		//{"float:03", []byte("1.0"), djs.ScanFloatNumber, []byte("1.0")},
		//{"float:04", []byte("-1.0"), djs.ScanFloatNumber, []byte("-1.0")},
		//{"float:03", []byte("3.1415"), djs.ScanFloatNumber, []byte("3.1415")},
		//{"float:04", []byte("-3.1415"), djs.ScanFloatNumber, []byte("-3.1415")},
		//{"float:05", []byte("3.141592653589793238462643383279502884197169"), djs.ScanFloatNumber, []byte("3.141592653589793238")},   // 3.141592653589793
		//{"float:06", []byte("-3.141592653589793238462643383279502884197169"), djs.ScanFloatNumber, []byte("-3.141592653589793238")}, // -3.141592653589793
		//{"float:05", []byte("3.141592653589793238462643383279502884197169e15"), djs.ScanFloatNumber, []byte("3.141592653589793238e15")},
		//{"float:06", []byte("-141592653589793238462643383279502884197169e+10"), djs.ScanFloatNumber, []byte("-14159265358979323846e+10")},
		//{"float:07", []byte("3.141592653589793238462643383279502884197169e-10"), djs.ScanFloatNumber, []byte("3.14159265358979323846e-10")},
		//{"float:08", []byte("-3.141592653589793238462643383279502884197169e-10"), djs.ScanFloatNumber, []byte("-3.14159265358979323846e-10")},
		//{"float:09", []byte("92653589793238462643383279502884197169e-10"), djs.ScanFloatNumber, []byte("92653589793238462643e-10")},
		//{"float:10", []byte("3.1415E5"), djs.ScanFloatNumber, []byte("3.1415E5")},
		//{"float:11", []byte("-3.1415E+5"), djs.ScanFloatNumber, []byte("-3.1415E+5")},
		//{"float:12", []byte("-3.1415E-5"), djs.ScanFloatNumber, []byte("-3.1415E-5")},
		// ------------------------------------------------------
		{"str:00", []byte(`"abc xyz"`), djs.ScanString, []byte("abc xyz")},
		{"str:01", []byte(`"abc\"xyz"`), djs.ScanString, []byte(`abc\"xyz`)},
		{"str:02", []byte(`"abc\xyz"`), djs.ScanString, []byte(`abc\\xyz`)},
		{"str:03", []byte(`"abc/xyz"`), djs.ScanString, []byte(`abc/xyz`)},
		{"str:04", []byte(`"abc"`), djs.ScanString, []byte(`abc\bxyz`)},
		{"str:05", []byte(`"abc\fxyz"`), djs.ScanString, []byte(`abc\fxyz`)},
		{"str:06", []byte(`"abc\nxyz"`), djs.ScanString, []byte(`abc\nxyz`)},
		{"str:07", []byte(`"abc\rxyz"`), djs.ScanString, []byte(`abc\rxyz`)},
		{"str:08", []byte(`"abc\txyz"`), djs.ScanString, []byte(`abc\txyz`)},
		{"str:09", []byte(`"abc\u00A0xyz"`), djs.ScanString, []byte(`abc\u00A0xyz`)},
		{"str:10", []byte(`"abc\u002Fxyz"`), djs.ScanString, []byte(`abc\u002Fxyz`)},
		{"str:11", []byte(`"abc\u002fxyz"`), djs.ScanString, []byte(`abc\u002fxyz`)},
		// ------------------------------------------------------
		{"time:00", []byte(`"2015-05-14T12:34:56.379+02:00"`), djs.ScanTime, []byte("2015-05-14T12:34:56.379+02:00")},
		{"time:01", []byte(`"1970-01-01T00:00:00Z"`), djs.ScanTime, []byte("1970-01-01T00:00:00Z")},
		{"time:02", []byte(`"0001-01-01T00:00:00Z"`), djs.ScanTime, []byte("0001-01-01T00:00:00Z")},
		{"time:03", []byte(`"1985-04-12T23:20:50.52Z"`), djs.ScanTime, []byte("1985-04-12T23:20:50.52Z")},
		{"time:04", []byte(`"1996-12-19T16:39:57-08:00"`), djs.ScanTime, []byte("1996-12-19T16:39:57-08:00")},
		{"time:05", []byte(`"1990-12-31T23:59:60Z"`), djs.ScanTime, []byte("1990-12-31T23:59:60Z")},
		{"time:06", []byte(`"1990-12-31T15:59:60-08:00"`), djs.ScanTime, []byte("1990-12-31T15:59:60-08:00")},
		{"time:07", []byte(`"1937-01-01T12:00:27.87+00:20"`), djs.ScanTime, []byte("1937-01-01T12:00:27.87+00:20")},
		{"time:08", []byte(`"2022-02-24T04:00:00+02:00"`), djs.ScanTime, []byte("2022-02-24T04:00:00+02:00")},
	}
	for _, el := range tbl {
		s.T().Run(el.idx, func(t *testing.T) {
			b := bytes.NewBuffer(el.in)
			bf := bufio.NewReader(b)
			sc := djs.NewJStructScanner(bf)
			e := sc.Next()
			v := sc.Value()
			k := sc.Kind()
			assert.NoError(t, e, "%s Next err: %v", el.idx, e)
			assert.Equal(t, v, el.out, "%s Value %v != %v", el.idx, v, el.out)
			assert.Equal(t, k, el.kind, "%s Kind %v != %v", el.idx, k, el.kind)
		})
	}
}
