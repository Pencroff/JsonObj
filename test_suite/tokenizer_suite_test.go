package test_suite

import (
	"bufio"
	"bytes"
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TokenizerTestElement struct {
	idx  string
	in   []byte
	kind djs.TokenizerKind
	out  []byte
	err  error
}

func TestJStruct_Tokenizer(t *testing.T) {
	s := new(TokenizerTestSuite)
	suite.Run(t, s)
}

type TokenizerTestSuite struct {
	suite.Suite
}

func (s *TokenizerTestSuite) SetupTest() {
}

func (s *TokenizerTestSuite) TestNewTokenizerInstance() {
	rd := &bufio.Reader{}
	sc := djs.NewJStructScanner(rd)
	tk := djs.NewJSStructTokenizer(sc)
	s.Equal(cap(tk.Value()), djs.JStructScannerBufferSize)
	sc = djs.NewJStructScannerWithSize(rd, 100)
	s.Equal(cap(tk.Value()), 100)
}

func (s *TokenizerTestSuite) TestTokenizer_Next_null() {
	tbl := []TokenizerTestElement{
		{"null:0", []byte("null"), djs.TokenNull, []byte("null"), nil},
		{"null:1", []byte("null "), djs.TokenNull, []byte("null"), nil},
		{"null:2", []byte("null\n"), djs.TokenNull, []byte("null"), nil},
		{"null:3", []byte("null\r"), djs.TokenNull, []byte("null"), nil},
		{"null:4", []byte("null\t"), djs.TokenNull, []byte("null"), nil},
		{"null:5", []byte("null\r\n"), djs.TokenNull, []byte("null"), nil},
		{"null:6", []byte("\nnull\t\n"), djs.TokenNull, []byte("null"), nil},
		{"null:7", []byte("\nnull\t\r"), djs.TokenNull, []byte("null"), nil},
		{"null:8", []byte(" null "), djs.TokenNull, []byte("null"), nil},
		{"null:9", []byte(" null\n"), djs.TokenNull, []byte("null"), nil},
		// Invalid cases
		{"null:50", []byte(""), djs.TokenNull, []byte{}, djs.InvalidJsonError},
		{"null:50", []byte("n"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"null:51", []byte("nill"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"null:52", []byte("nnn"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
	}
	for _, el := range tbl {
		RunTokeniserTest(el, s)
	}
}

func (s *TokenizerTestSuite) TestTokenizer_Next_bool() {
	tbl := []TokenizerTestElement{
		// False cases
		{"bool:f00", []byte("false"), djs.TokenFalse, []byte("false"), nil},
		{"bool:f01", []byte(" false "), djs.TokenFalse, []byte("false"), nil},
		// Invalid cases
		{"bool:f50", []byte(" folse "), djs.TokenFalse, []byte{}, djs.InvalidJsonError},
		{"bool:f51", []byte("falze"), djs.TokenFalse, []byte{}, djs.InvalidJsonError},
		{"bool:f52", []byte("fals"), djs.TokenFalse, []byte{}, djs.InvalidJsonError},
		{"bool:f53", []byte("f "), djs.TokenFalse, []byte{}, djs.InvalidJsonError},
		// True cases
		{"bool:t00", []byte("true"), djs.TokenTrue, []byte("true"), nil},
		{"bool:t01", []byte("\n\rtrue\n\r"), djs.TokenTrue, []byte("true"), nil},
		// Invalid cases
		{"bool:t50", []byte("truae "), djs.TokenTrue, []byte{}, djs.InvalidJsonError},
		{"bool:t51", []byte("trues"), djs.TokenTrue, []byte{}, djs.InvalidJsonError},
		{"bool:t52", []byte(" t "), djs.TokenTrue, []byte{}, djs.InvalidJsonError},
	}
	for _, el := range tbl {
		RunTokeniserTest(el, s)
	}
}

func (s *TokenizerTestSuite) TestTokenizer_Next_number() {
	tbl := []TokenizerTestElement{
		{"num:00", []byte("123"), djs.TokenIntNumber, []byte("123"), nil},
		{"num:01", []byte("0"), djs.TokenIntNumber, []byte("0"), nil},
		{"num:02", []byte("-0"), djs.TokenIntNumber, []byte("-0"), nil},
		{"num:03", []byte("1"), djs.TokenIntNumber, []byte("1"), nil},
		{"num:04", []byte("-1"), djs.TokenIntNumber, []byte("-1"), nil},
		{"num:05", []byte(" -1 "), djs.TokenIntNumber, []byte("-1"), nil},
		{"num:06", []byte("123456789"), djs.TokenIntNumber, []byte("123456789"), nil},
		{"num:07", []byte("-123456789"), djs.TokenIntNumber, []byte("-123456789"), nil},
		{"num:08", []byte("9223372036854775807"), djs.TokenIntNumber, []byte("9223372036854775807"), nil},
		{"num:09", []byte("-9223372036854775808"), djs.TokenIntNumber, []byte("-9223372036854775808"), nil},
		{"num:10", []byte("9223372036854775808"), djs.TokenIntNumber, []byte("9223372036854775808"), nil},
		{"num:11", []byte("-9223372036854775809"), djs.TokenIntNumber, []byte("-9223372036854775809"), nil},   // -9.223372036854776e+18
		{"num:12", []byte("18446744073709551615"), djs.TokenIntNumber, []byte("18446744073709551615"), nil},   // 1.8446744073709552e+19
		{"num:13", []byte("-18446744073709551615"), djs.TokenIntNumber, []byte("-18446744073709551615"), nil}, // -1.8446744073709552e+19
		{"num:14", []byte("\n9064\n\r"), djs.TokenIntNumber, []byte("9064"), nil},
		{"num:15", []byte("340282366920938463463374607431768211455"), djs.TokenIntNumber, []byte("340282366920938463463374607431768211455"), nil},
		// Num errors
		{"num:50", []byte("9 0 6 4"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"num:51", []byte("-e"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"num:52", []byte("$E1"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"num:53", []byte("1i1l1"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"num:54", []byte("1e"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"num:55", []byte("11$!"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
	}
	for _, el := range tbl {
		RunTokeniserTest(el, s)
	}
}

func (s *TokenizerTestSuite) TestTokenizer_Next_float() {
	tbl := []TokenizerTestElement{
		{"float:00", []byte("123.45"), djs.TokenFloatNumber, []byte("123.45"), nil},
		{"float:01", []byte("0.0"), djs.TokenFloatNumber, []byte("0.0"), nil},
		{"float:02", []byte("-0.0"), djs.TokenFloatNumber, []byte("-0.0"), nil},
		{"float:03", []byte("1.0"), djs.TokenFloatNumber, []byte("1.0"), nil},
		{"float:04", []byte("-1.0"), djs.TokenFloatNumber, []byte("-1.0"), nil},
		{"float:03", []byte("3.1415"), djs.TokenFloatNumber, []byte("3.1415"), nil},
		{"float:04", []byte("-3.1415"), djs.TokenFloatNumber, []byte("-3.1415"), nil},
		{"float:05", []byte("3.141592653589793238462643383279502884197169"), djs.TokenFloatNumber, []byte("3.141592653589793238462643383279502884197169"), nil},   // 3.141592653589793
		{"float:06", []byte("-3.141592653589793238462643383279502884197169"), djs.TokenFloatNumber, []byte("-3.141592653589793238462643383279502884197169"), nil}, // -3.141592653589793
		{"float:05", []byte("3.141592653589793238462643383279502884197169e15"), djs.TokenFloatNumber, []byte("3.141592653589793238e15"), nil},
		{"float:06", []byte("-141592653589793238462643383279502884197169e+10"), djs.TokenFloatNumber, []byte("-141592653589793238462643383279502884197169e+10"), nil},
		{"float:07", []byte("3.141592653589793238462643383279502884197169e-10"), djs.TokenFloatNumber, []byte("3.141592653589793238462643383279502884197169e-10"), nil},
		{"float:08", []byte("-3.141592653589793238462643383279502884197169e-10"), djs.TokenFloatNumber, []byte("-3.141592653589793238462643383279502884197169e-10"), nil},
		{"float:09", []byte("92653589793238462643383279502884197169e-10"), djs.TokenFloatNumber, []byte("92653589793238462643383279502884197169e-10"), nil},
		{"float:10", []byte("-926535897932384626433.83279502884197169e-10"), djs.TokenFloatNumber, []byte("-926535897932384626433.83279502884197169e-10"), nil},
		{"float:10", []byte(" 3.1415E5 "), djs.TokenFloatNumber, []byte("3.1415E5"), nil},
		{"float:11", []byte("\n-3.1415E+5\n"), djs.TokenFloatNumber, []byte("-3.1415E+5"), nil},
		{"float:12", []byte("-3.1415E-5"), djs.TokenFloatNumber, []byte("-3.1415E-5"), nil},
		{"float:13", []byte("3.1415E-5"), djs.TokenFloatNumber, []byte("3.1415E-5"), nil},
		{"float:14", []byte("1.6180339887498948482045868343656381e999"), djs.TokenFloatNumber, []byte("1.6180339887498948482045868343656381e999"), nil},
		{"float:15", []byte("-1.6180339887498948482045868343656381e-999"), djs.TokenFloatNumber, []byte("-1.6180339887498948482045868343656381e-999"), nil},
		{"float:16", []byte("0.01"), djs.TokenFloatNumber, []byte("0.01"), nil},
		{"float:17", []byte("-0.01"), djs.TokenFloatNumber, []byte("-0.01"), nil},
		{"float:18", []byte(" 0.01 "), djs.TokenFloatNumber, []byte("0.01"), nil},
		{"float:19", []byte("\n\r-0.01\n\r"), djs.TokenFloatNumber, []byte("-0.01"), nil},
		// Float errors
		{"float:50", []byte("-"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:51", []byte("-e"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:52", []byte("0."), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:53", []byte("0.e"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:54", []byte("0.e1"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:55", []byte("0.1e"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:56", []byte("0.1e-1"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:57", []byte(".01"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:58", []byte("1i1.4l1"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:59", []byte("-3."), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:60", []byte("-3.e"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:61", []byte("3.1e"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:62", []byte("3.1415926535.89793"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:63", []byte("3.14159265Ee589793"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:64", []byte("3.14159265E+"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:65", []byte("3.14159265E-"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:66", []byte("161803398.874989zzz8204e28"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"float:67", []byte("16180.3398.874989e8204e+28"), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
	}
	for _, el := range tbl {
		RunTokeniserTest(el, s)
	}
}

func (s *TokenizerTestSuite) TestTokenizer_Next_string() {
	testCases := []TokenizerTestElement{
		{"str:00", []byte(`""`), djs.TokenString, []byte{}, nil},
		{"str:01", []byte(`"abc"`), djs.TokenString, []byte("abc"), nil},
		{"str:02", []byte(` "abc" `), djs.TokenString, []byte("abc"), nil},
		{"str:03", []byte(`
									"abc"
								`), djs.TokenString, []byte("abc"), nil},
		{"str:04", []byte(`"abc xyz"`), djs.TokenString, []byte("abc xyz"), nil},
		{"str:05", []byte(`"hello world!"`), djs.TokenString, []byte("hello world!"), nil},
		{"str:06", []byte(`"The quick brown fox jumps over the lazy dog"`), djs.TokenString, []byte("The quick brown fox jumps over the lazy dog"), nil},
		{"str:06", []byte(`"a\"z"`), djs.TokenString, []byte{0x61, 0x5c, 0x22, 0x7a}, nil},
		{"str:07", []byte(`"a\\z"`), djs.TokenString, []byte(`a\z`), nil},
		{"str:08", []byte(`"a\/z"`), djs.TokenString, []byte(`a\/z`), nil},
		{"str:09", []byte(`"a/z"`), djs.TokenString, []byte(`a/z`), nil},
		{"str:10", []byte(`"a\bz"`), djs.TokenString, []byte(`a\bz`), nil},
		{"str:11", []byte(`"a\fz"`), djs.TokenString, []byte(`a\fz`), nil},
		{"str:12", []byte(`"a\nz"`), djs.TokenString, []byte(`a\nz`), nil},
		{"str:13", []byte(`"a\rz"`), djs.TokenString, []byte(`a\rz`), nil},
		{"str:14", []byte(`"a\tz"`), djs.TokenString, []byte(`a\tz`), nil},
		{"str:15", []byte(`"abc\u00A0xyz"`), djs.TokenString, []byte(`abc\u00A0xyz`), nil},
		{"str:16", []byte(`"abc\u002Fxyz"`), djs.TokenString, []byte(`abc\u002Fxyz`), nil},
		{"str:17", []byte(`"abc\u002fxyz"`), djs.TokenString, []byte(`abc\u002fxyz`), nil},
		{"str:18", []byte(`"\u2070"`), djs.TokenString, []byte(`\u2070`), nil},
		{"str:19", []byte(`"\u0008"`), djs.TokenString, []byte(`\u0008`), nil},
		{"str:20", []byte(`"\u000C"`), djs.TokenString, []byte(`\u000C`), nil},
		{"str:21", []byte(`"\uD834\uDD1E"`), djs.TokenString, []byte(`\uD834\uDD1E`), nil},
		{"str:22", []byte(`"D'fhuascail Íosa, Úrmhac na hÓighe Beannaithe, pór Éava agus Ádhaimh"`), djs.TokenString, []byte(`D'fhuascail Íosa, Úrmhac na hÓighe Beannaithe, pór Éava agus Ádhaimh`), nil},
		{"str:23", []byte(`"いろはにほへとちりぬるを"`), djs.TokenString, []byte(`いろはにほへとちりぬるを`), nil},
		{"str:24", []byte(`"\uD834\n"`), djs.TokenString, []byte(`\uD834\n`), nil},

		// String errors
		{"str:50", []byte(`"abc`), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"str:51", []byte(`"abc"xyz`), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"str:52", []byte(`abc"`), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"str:53", []byte(`"""`), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"str:54", []byte(`""\"`), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"str:55", []byte(`"\u2O70"`), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"str:56", []byte(`"\uD8Y4\uDU1E"`), djs.TokenUnknown, []byte{}, djs.InvalidJsonError},
		{"str:57", []byte(`"\uD834\q"`), djs.TokenString, []byte{}, nil},
	}
	for _, el := range testCases {
		RunTokeniserTest(el, s)
	}
}

func (s *TokenizerTestSuite) TestTokenizer_Next_time() {
	tbl := []TokenizerTestElement{
		{"time:00", []byte(`"2015-05-14T12:34:56.379+02:00"`), djs.TokenTime, []byte("2015-05-14T12:34:56.379+02:00"), nil},
		{"time:01", []byte(`"1970-01-01T00:00:00Z"`), djs.TokenTime, []byte("1970-01-01T00:00:00Z"), nil},
		{"time:02", []byte(`"0001-01-01T00:00:00Z"`), djs.TokenTime, []byte("0001-01-01T00:00:00Z"), nil},
		{"time:03", []byte(`"1985-04-12T23:20:50.52Z"`), djs.TokenTime, []byte("1985-04-12T23:20:50.52Z"), nil},
		{"time:04", []byte(`"1996-12-19T16:39:57-08:00"`), djs.TokenTime, []byte("1996-12-19T16:39:57-08:00"), nil},
		{"time:05", []byte(`"1990-12-31T23:59:60Z"`), djs.TokenTime, []byte("1990-12-31T23:59:60Z"), nil},
		{"time:06", []byte(`"1990-12-31T15:59:60-08:00"`), djs.TokenTime, []byte("1990-12-31T15:59:60-08:00"), nil},
		{"time:07", []byte(`"1937-01-01T12:00:27.87+00:20"`), djs.TokenTime, []byte("1937-01-01T12:00:27.87+00:20"), nil},
		{"time:08", []byte(`"2022-02-24T04:00:00+02:00"`), djs.TokenTime, []byte("2022-02-24T04:00:00+02:00"), nil},
		{"time:09", []byte(`"2022-07-12T21:55:16+01:00"`), djs.TokenTime, []byte("2022-02-24T04:00:00+02:00"), nil},
		// Invalid cases fall back to string
		{"time:50", []byte(`"2015-05-14E12:34:56.379+02:00"`), djs.TokenString, []byte(`2015-05-14E12:34:56.379+02:00`), nil},
		{"time:51", []byte(`"2O15-O5-14T12:34:56.379+02:00"`), djs.TokenString, []byte(`2O15-O5-14T12:34:56.379+02:00`), nil},
		{"time:52", []byte(`"1985-04-12T23:20:50.52ZZZZ"`), djs.TokenString, []byte(`1985-04-12T23:20:50.52ZZZZ`), nil},
		{"time:53", []byte(`"2022-07-12 21:55:16"`), djs.TokenString, []byte(`2022-07-12 21:55:16`), nil},
		{"time:54", []byte(`"20220712T215516Z"`), djs.TokenString, []byte(`20220712T215516Z`), nil},
	}
	for _, el := range tbl {
		RunTokeniserTest(el, s)
	}
}

func RunTokeniserTest(el TokenizerTestElement, s *TokenizerTestSuite) {
	s.T().Run(el.idx, func(t *testing.T) {
		b := bytes.NewBuffer(el.in)
		bf := bufio.NewReader(b)
		sc := djs.NewJStructScanner(bf)
		tk := djs.NewJSStructTokenizer(sc)
		e := tk.Next()
		v := tk.Value()
		k := tk.Kind()
		assert.ErrorIs(t, e, el.err, "%s Next err: %v", el.idx, e)
		assert.Equal(t, v, el.out, "%s Value %v != %v", el.idx, v, el.out)
		assert.Equal(t, k, el.kind, "%s Kind %v != %v", el.idx, k, el.kind)
	})
}
