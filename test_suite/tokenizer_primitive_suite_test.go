package test_suite

import (
	"bytes"
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io"
	"testing"
)

type TokenizerTestPrimitiveElement struct {
	idx   string
	in    []byte
	kind  djs.TokenizerKind
	level djs.TokenizerLevel
	out   []byte
	err   error
}

func TestJStruct_TokenizerPrimitive(t *testing.T) {
	s := new(TokenizerTestPrimitiveSuite)
	suite.Run(t, s)
}

type TokenizerTestPrimitiveSuite struct {
	suite.Suite
}

func (s *TokenizerTestPrimitiveSuite) SetupTest() {
}

func (s *TokenizerTestPrimitiveSuite) TestTokenizer_Next_null() {
	tbl := []TokenizerTestPrimitiveElement{
		{"null:0", []byte("null"), djs.KindNull, djs.LevelRoot, []byte("null"), nil},
		{"null:1", []byte("null           "), djs.KindNull, djs.LevelRoot, []byte("null"), nil},
		{"null:2", []byte("null\n"), djs.KindNull, djs.LevelRoot, []byte("null"), nil},
		{"null:3", []byte("null\r"), djs.KindNull, djs.LevelRoot, []byte("null"), nil},
		{"null:4", []byte("null\t"), djs.KindNull, djs.LevelRoot, []byte("null"), nil},
		{"null:5", []byte("null\r\n"), djs.KindNull, djs.LevelRoot, []byte("null"), nil},
		{"null:6", []byte("\nnull\t\n"), djs.KindNull, djs.LevelRoot, []byte("null"), nil},
		{"null:7", []byte("\nnull\t\r"), djs.KindNull, djs.LevelRoot, []byte("null"), nil},
		{"null:8", []byte(" null "), djs.KindNull, djs.LevelRoot, []byte("null"), nil},
		{"null:9", []byte(" null\n"), djs.KindNull, djs.LevelRoot, []byte("null"), nil},
		// Invalid cases
		{"null:100", []byte(""), djs.KindUnknown, djs.LevelRoot, []byte(nil), djs.InvalidJsonError{Err: io.EOF}},
		{"null:101", []byte("n"), djs.KindUnknown, djs.LevelRoot, []byte("n"), djs.InvalidJsonPtrError{Pos: 1, Err: io.EOF}},
		{"null:102", []byte("   nill"), djs.KindUnknown, djs.LevelRoot, []byte("ni"), djs.InvalidJsonPtrError{Pos: 4}},
		{"null:103", []byte("nnn"), djs.KindUnknown, djs.LevelRoot, []byte("nn"), djs.InvalidJsonPtrError{Pos: 1}},
		{"null:104", []byte("nnnn"), djs.KindUnknown, djs.LevelRoot, []byte("nn"), djs.InvalidJsonPtrError{Pos: 1}},
		{"null:105", []byte("nulle"), djs.KindUnknown, djs.LevelRoot, []byte("nulle"), djs.InvalidJsonPtrError{Pos: 4}},
		{"null:106", []byte("null\t\t\tnull"), djs.KindUnknown, djs.LevelRoot, []byte("null\t\t\tn"), djs.InvalidJsonPtrError{Pos: 7}},
	}
	for _, el := range tbl {
		RunTokenizerTestPrimitiveCase(s, el)
	}
}

func (s *TokenizerTestPrimitiveSuite) TestTokenizer_Next_bool() {
	tbl := []TokenizerTestPrimitiveElement{
		// False cases
		{"bool:f00", []byte("false"), djs.KindFalse, djs.LevelRoot, []byte("false"), nil},
		{"bool:f01", []byte(" false "), djs.KindFalse, djs.LevelRoot, []byte("false"), nil},
		// Invalid cases
		{"bool:f100", []byte(" folse "), djs.KindUnknown, djs.LevelRoot, []byte("fo"), djs.InvalidJsonPtrError{Pos: 2}},
		{"bool:f101", []byte("falze"), djs.KindUnknown, djs.LevelRoot, []byte("falz"), djs.InvalidJsonPtrError{Pos: 3}},
		{"bool:f102", []byte("fals"), djs.KindUnknown, djs.LevelRoot, []byte("fals"), djs.InvalidJsonPtrError{Pos: 4, Err: io.EOF}},
		{"bool:f103", []byte("f "), djs.KindUnknown, djs.LevelRoot, []byte("f "), djs.InvalidJsonPtrError{Pos: 1}},
		{"bool:f104", []byte("falsez"), djs.KindUnknown, djs.LevelRoot, []byte("falsez"), djs.InvalidJsonPtrError{Pos: 5}},
		{"bool:f105", []byte("false\t\t\tfalse"), djs.KindUnknown, djs.LevelRoot, []byte("false\t\t\tf"), djs.InvalidJsonPtrError{Pos: 8}},
		// True cases
		{"bool:t00", []byte("true"), djs.KindTrue, djs.LevelRoot, []byte("true"), nil},
		{"bool:t01", []byte("\n\rtrue\n\r"), djs.KindTrue, djs.LevelRoot, []byte("true"), nil},
		// Invalid cases
		{"bool:t100", []byte("truae "), djs.KindUnknown, djs.LevelRoot, []byte("trua"), djs.InvalidJsonPtrError{Pos: 3}},
		{"bool:t101", []byte("trues"), djs.KindUnknown, djs.LevelRoot, []byte("trues"), djs.InvalidJsonPtrError{Pos: 4}},
		{"bool:t102", []byte(" t "), djs.KindUnknown, djs.LevelRoot, []byte("t "), djs.InvalidJsonPtrError{Pos: 2}},
	}
	for _, el := range tbl {
		RunTokenizerTestPrimitiveCase(s, el)
	}
}

func (s *TokenizerTestPrimitiveSuite) TestTokenizer_Next_number() {
	tbl := []TokenizerTestPrimitiveElement{
		{"num:00", []byte("123"), djs.KindNumber, djs.LevelRoot, []byte("123"), nil},
		{"num:01", []byte("0"), djs.KindNumber, djs.LevelRoot, []byte("0"), nil},
		{"num:02", []byte("-0"), djs.KindNumber, djs.LevelRoot, []byte("-0"), nil},
		{"num:03", []byte("1"), djs.KindNumber, djs.LevelRoot, []byte("1"), nil},
		{"num:04", []byte("-1"), djs.KindNumber, djs.LevelRoot, []byte("-1"), nil},
		{"num:05", []byte(" -1 "), djs.KindNumber, djs.LevelRoot, []byte("-1"), nil},
		{"num:06", []byte("123456789"), djs.KindNumber, djs.LevelRoot, []byte("123456789"), nil},
		{"num:07", []byte("-123456789"), djs.KindNumber, djs.LevelRoot, []byte("-123456789"), nil},
		{"num:08", []byte("9223372036854775807"), djs.KindNumber, djs.LevelRoot, []byte("9223372036854775807"), nil},
		{"num:09", []byte("-9223372036854775808"), djs.KindNumber, djs.LevelRoot, []byte("-9223372036854775808"), nil},
		{"num:10", []byte("9223372036854775808"), djs.KindNumber, djs.LevelRoot, []byte("9223372036854775808"), nil},
		{"num:11", []byte("-9223372036854775809"), djs.KindNumber, djs.LevelRoot, []byte("-9223372036854775809"), nil},   // -9.223372036854776e+18
		{"num:12", []byte("18446744073709551615"), djs.KindNumber, djs.LevelRoot, []byte("18446744073709551615"), nil},   // 1.8446744073709552e+19
		{"num:13", []byte("-18446744073709551615"), djs.KindNumber, djs.LevelRoot, []byte("-18446744073709551615"), nil}, // -1.8446744073709552e+19
		{"num:14", []byte("\n9064\n\r"), djs.KindNumber, djs.LevelRoot, []byte("9064"), nil},
		{"num:15", []byte("340282366920938463463374607431768211455"), djs.KindNumber, djs.LevelRoot, []byte("340282366920938463463374607431768211455"), nil},
		// Num errors
		{"num:100", []byte("9 0 6 4"), djs.KindUnknown, djs.LevelRoot, []byte(`9 0`), djs.InvalidJsonPtrError{Pos: 2}},
		{"num:101", []byte("-e"), djs.KindUnknown, djs.LevelRoot, []byte(`-e`), djs.InvalidJsonPtrError{Pos: 1}},
		{"num:102", []byte("25$E1"), djs.KindUnknown, djs.LevelRoot, []byte(`25$`), djs.InvalidJsonPtrError{Pos: 2}},
		{"num:103", []byte("123l1"), djs.KindUnknown, djs.LevelRoot, []byte(`123l`), djs.InvalidJsonPtrError{Pos: 3}},
		{"num:104", []byte("1e"), djs.KindUnknown, djs.LevelRoot, []byte(`1e`), djs.InvalidJsonPtrError{Pos: 1, Err: io.EOF}},
		{"num:105", []byte("1234e  "), djs.KindUnknown, djs.LevelRoot, []byte(`1234e `), djs.InvalidJsonPtrError{Pos: 5}},
		{"num:106", []byte("11$!"), djs.KindUnknown, djs.LevelRoot, []byte(`11$`), djs.InvalidJsonPtrError{Pos: 2}},
		{"num:107", []byte("- 123"), djs.KindUnknown, djs.LevelRoot, []byte(`- `), djs.InvalidJsonPtrError{Pos: 1}},
	}
	for _, el := range tbl {
		RunTokenizerTestPrimitiveCase(s, el)
	}
}

func (s *TokenizerTestPrimitiveSuite) TestTokenizer_Next_float() {
	tbl := []TokenizerTestPrimitiveElement{
		{"float:00", []byte("123.45"), djs.KindFloatNumber, djs.LevelRoot, []byte("123.45"), nil},
		{"float:01", []byte("0.0"), djs.KindFloatNumber, djs.LevelRoot, []byte("0.0"), nil},
		{"float:02", []byte("-0.0"), djs.KindFloatNumber, djs.LevelRoot, []byte("-0.0"), nil},
		{"float:03", []byte("1.0"), djs.KindFloatNumber, djs.LevelRoot, []byte("1.0"), nil},
		{"float:04", []byte("-1.0"), djs.KindFloatNumber, djs.LevelRoot, []byte("-1.0"), nil},
		{"float:05", []byte("3.1415"), djs.KindFloatNumber, djs.LevelRoot, []byte("3.1415"), nil},
		{"float:06", []byte("-3.1415"), djs.KindFloatNumber, djs.LevelRoot, []byte("-3.1415"), nil},
		{"float:07", []byte("3.141592653589793238462643383279502884197169"), djs.KindFloatNumber, djs.LevelRoot, []byte("3.141592653589793238462643383279502884197169"), nil},   // 3.141592653589793
		{"float:08", []byte("-3.141592653589793238462643383279502884197169"), djs.KindFloatNumber, djs.LevelRoot, []byte("-3.141592653589793238462643383279502884197169"), nil}, // -3.141592653589793
		{"float:09", []byte("3.141592653589793238462643383279502884197169e15"), djs.KindFloatNumber, djs.LevelRoot, []byte("3.141592653589793238462643383279502884197169e15"), nil},
		{"float:10", []byte("-141592653589793238462643383279502884197169e+10"), djs.KindFloatNumber, djs.LevelRoot, []byte("-141592653589793238462643383279502884197169e+10"), nil},
		{"float:11", []byte("3.141592653589793238462643383279502884197169e-10"), djs.KindFloatNumber, djs.LevelRoot, []byte("3.141592653589793238462643383279502884197169e-10"), nil},
		{"float:12", []byte("-3.141592653589793238462643383279502884197169e-10"), djs.KindFloatNumber, djs.LevelRoot, []byte("-3.141592653589793238462643383279502884197169e-10"), nil},
		{"float:13", []byte("92653589793238462643383279502884197169e-10"), djs.KindFloatNumber, djs.LevelRoot, []byte("92653589793238462643383279502884197169e-10"), nil},
		{"float:14", []byte("-926535897932384626433.83279502884197169e-10"), djs.KindFloatNumber, djs.LevelRoot, []byte("-926535897932384626433.83279502884197169e-10"), nil},
		{"float:15", []byte(" 3.1415E5 "), djs.KindFloatNumber, djs.LevelRoot, []byte("3.1415E5"), nil},
		{"float:16", []byte("\n-3.1415E+5\n"), djs.KindFloatNumber, djs.LevelRoot, []byte("-3.1415E+5"), nil},
		{"float:17", []byte("-3.1415E-5"), djs.KindFloatNumber, djs.LevelRoot, []byte("-3.1415E-5"), nil},
		{"float:18", []byte("3.1415E-5"), djs.KindFloatNumber, djs.LevelRoot, []byte("3.1415E-5"), nil},
		{"float:19", []byte("1.6180339887498948482045868343656381e999"), djs.KindFloatNumber, djs.LevelRoot, []byte("1.6180339887498948482045868343656381e999"), nil},
		{"float:20", []byte("-1.6180339887498948482045868343656381e-999"), djs.KindFloatNumber, djs.LevelRoot, []byte("-1.6180339887498948482045868343656381e-999"), nil},
		{"float:21", []byte("0.01"), djs.KindFloatNumber, djs.LevelRoot, []byte("0.01"), nil},
		{"float:22", []byte("-0.01"), djs.KindFloatNumber, djs.LevelRoot, []byte("-0.01"), nil},
		{"float:23", []byte(" 0.01 "), djs.KindFloatNumber, djs.LevelRoot, []byte("0.01"), nil},
		{"float:24", []byte("\n\r-0.01\n\r"), djs.KindFloatNumber, djs.LevelRoot, []byte("-0.01"), nil},
		{"float:25", []byte("0.1e-1"), djs.KindFloatNumber, djs.LevelRoot, []byte("0.1e-1"), nil},
		// Float errors
		{"float:100", []byte("-"), djs.KindUnknown, djs.LevelRoot, []byte("-"), djs.InvalidJsonPtrError{Pos: 0, Err: io.EOF}},
		{"float:101", []byte("-e"), djs.KindUnknown, djs.LevelRoot, []byte(`-e`), djs.InvalidJsonPtrError{Pos: 1}},
		{"float:102", []byte("0."), djs.KindUnknown, djs.LevelRoot, []byte("0."), djs.InvalidJsonPtrError{Pos: 1, Err: io.EOF}},
		{"float:103", []byte("0.e"), djs.KindUnknown, djs.LevelRoot, []byte("0.e"), djs.InvalidJsonPtrError{Pos: 2}},
		{"float:104", []byte("0.e1"), djs.KindUnknown, djs.LevelRoot, []byte("0.e"), djs.InvalidJsonPtrError{Pos: 2}},
		{"float:105", []byte("0.1e"), djs.KindUnknown, djs.LevelRoot, []byte("0.1e"), djs.InvalidJsonPtrError{Pos: 3, Err: io.EOF}},
		{"float:106", []byte(".01"), djs.KindUnknown, djs.LevelRoot, []byte(nil), djs.InvalidJsonError{}},
		{"float:107", []byte("123.4l1"), djs.KindUnknown, djs.LevelRoot, []byte("123.4l"), djs.InvalidJsonPtrError{Pos: 5}},
		{"float:108", []byte("-3."), djs.KindUnknown, djs.LevelRoot, []byte("-3."), djs.InvalidJsonPtrError{Pos: 2, Err: io.EOF}},
		{"float:109", []byte("-3.e"), djs.KindUnknown, djs.LevelRoot, []byte("-3.e"), djs.InvalidJsonPtrError{Pos: 3}},
		{"float:110", []byte("-3.e1"), djs.KindUnknown, djs.LevelRoot, []byte("-3.e"), djs.InvalidJsonPtrError{Pos: 3}},
		{"float:111", []byte("-3.1e"), djs.KindUnknown, djs.LevelRoot, []byte("-3.1e"), djs.InvalidJsonPtrError{Pos: 4, Err: io.EOF}},
		{"float:112", []byte("3.1415926535.89793"), djs.KindUnknown, djs.LevelRoot, []byte("3.1415926535."), djs.InvalidJsonPtrError{Pos: 12}},
		{"float:113", []byte("3.14159265Ee589793"), djs.KindUnknown, djs.LevelRoot, []byte("3.14159265Ee"), djs.InvalidJsonPtrError{Pos: 11}},
		{"float:114", []byte("3.14159265E+"), djs.KindUnknown, djs.LevelRoot, []byte("3.14159265E+"), djs.InvalidJsonPtrError{Pos: 11, Err: io.EOF}},
		{"float:115", []byte("3.14159265E-"), djs.KindUnknown, djs.LevelRoot, []byte("3.14159265E-"), djs.InvalidJsonPtrError{Pos: 11, Err: io.EOF}},
		{"float:116", []byte("161803398.874989opq8204e28"), djs.KindUnknown, djs.LevelRoot, []byte("161803398.874989o"), djs.InvalidJsonPtrError{Pos: 16}},
		{"float:117", []byte("16180.3398.874989e8204e+28"), djs.KindUnknown, djs.LevelRoot, []byte("16180.3398."), djs.InvalidJsonPtrError{Pos: 10}},
	}
	for _, el := range tbl {
		RunTokenizerTestPrimitiveCase(s, el)
	}
}

func (s *TokenizerTestPrimitiveSuite) TestTokenizer_Next_string() {
	testCases := []TokenizerTestPrimitiveElement{
		{"str:00", []byte(`""`), djs.KindString, djs.LevelRoot, []byte(`""`), nil},
		{"str:01", []byte(`"abc"`), djs.KindString, djs.LevelRoot, []byte(`"abc"`), nil},
		{"str:02", []byte(` "abc" `), djs.KindString, djs.LevelRoot, []byte(`"abc"`), nil},
		{"str:03", []byte(`
									"abc"
								`), djs.KindString, djs.LevelRoot, []byte(`"abc"`), nil},
		{"str:04", []byte(`"abc xyz"`), djs.KindString, djs.LevelRoot, []byte(`"abc xyz"`), nil},
		{"str:05", []byte(`"hello world!"`), djs.KindString, djs.LevelRoot, []byte(`"hello world!"`), nil},
		{"str:06", []byte(`"The quick brown fox jumps over the lazy dog"`), djs.KindString, djs.LevelRoot, []byte(`"The quick brown fox jumps over the lazy dog"`), nil},
		{"str:07", []byte(`"a\"z"`), djs.KindString, djs.LevelRoot, []byte{0x22, 0x61, 0x5c, 0x22, 0x7a, 0x22}, nil},
		{"str:08", []byte(`"a\\z"`), djs.KindString, djs.LevelRoot, []byte(`"a\\z"`), nil},
		{"str:09", []byte(`"a\/z"`), djs.KindString, djs.LevelRoot, []byte(`"a\/z"`), nil},
		{"str:10", []byte(`"a/z"`), djs.KindString, djs.LevelRoot, []byte(`"a/z"`), nil},
		{"str:11", []byte(`"a\bz"`), djs.KindString, djs.LevelRoot, []byte(`"a\bz"`), nil},
		{"str:12", []byte(`"a\fz"`), djs.KindString, djs.LevelRoot, []byte(`"a\fz"`), nil},
		{"str:13", []byte(`"a\nz"`), djs.KindString, djs.LevelRoot, []byte(`"a\nz"`), nil},
		{"str:14", []byte(`"a\rz"`), djs.KindString, djs.LevelRoot, []byte(`"a\rz"`), nil},
		{"str:15", []byte(`"a\tz"`), djs.KindString, djs.LevelRoot, []byte(`"a\tz"`), nil},
		{"str:16", []byte(`"abc\u00A0xyz"`), djs.KindString, djs.LevelRoot, []byte(`"abc\u00A0xyz"`), nil},
		{"str:17", []byte(`"abc\u002Fxyz"`), djs.KindString, djs.LevelRoot, []byte(`"abc\u002Fxyz"`), nil},
		{"str:18", []byte(`"abc\u002fxyz"`), djs.KindString, djs.LevelRoot, []byte(`"abc\u002fxyz"`), nil},
		{"str:19", []byte(`"\u2070"`), djs.KindString, djs.LevelRoot, []byte(`"\u2070"`), nil},
		{"str:20", []byte(`"\u0008"`), djs.KindString, djs.LevelRoot, []byte(`"\u0008"`), nil},
		{"str:21", []byte(`"\u000C"`), djs.KindString, djs.LevelRoot, []byte(`"\u000C"`), nil},
		{"str:22", []byte(`"\uD834\uDD1E"`), djs.KindString, djs.LevelRoot, []byte(`"\uD834\uDD1E"`), nil},
		{"str:23", []byte(`"D'fhuascail Íosa, Úrmhac na hÓighe Beannaithe, pór Éava agus Ádhaimh"`), djs.KindString, djs.LevelRoot, []byte(`"D'fhuascail Íosa, Úrmhac na hÓighe Beannaithe, pór Éava agus Ádhaimh"`), nil},
		{"str:24", []byte(`"いろはにほへとちりぬるを"`), djs.KindString, djs.LevelRoot, []byte(`"いろはにほへとちりぬるを"`), nil},

		// String errors
		{"str:100", []byte(`"abc`), djs.KindUnknown, djs.LevelRoot, []byte(`"abc`),
			djs.InvalidJsonPtrError{Pos: 3, Err: io.EOF}},
		{"str:101", []byte(`"abc"xyz`), djs.KindUnknown, djs.LevelRoot, []byte(`"abc"x`),
			djs.InvalidJsonPtrError{Pos: 5}},
		{"str:102", []byte(`abc"`), djs.KindUnknown, djs.LevelRoot, []byte(nil),
			djs.InvalidJsonError{}},
		{"str:103", []byte(`"""`), djs.KindUnknown, djs.LevelRoot, []byte(`"""`),
			djs.InvalidJsonPtrError{Pos: 2}},
		{"str:104", []byte(`""\"`), djs.KindUnknown, djs.LevelRoot, []byte(`""\`),
			djs.InvalidJsonPtrError{Pos: 2}},
		{"str:105", []byte(`"\u2O70"`), djs.KindUnknown, djs.LevelRoot, []byte(`"\u2O`),
			djs.InvalidJsonPtrError{Pos: 4, Err: djs.InvalidHexNumberError}},
		{"str:106", []byte(`"\uD8Y4\uDU1E"`), djs.KindUnknown, djs.LevelRoot, []byte(`"\uD8Y`), djs.InvalidJsonPtrError{Pos: 5, Err: djs.InvalidHexNumberError}},
		{"str:107", []byte(`"\x15"`), djs.KindUnknown, djs.LevelRoot, []byte(`"\x`),
			djs.InvalidJsonPtrError{Pos: 2, Err: djs.InvalidEscapeCharacterError}},
		{"str:110", []byte("\"a\nz\""), djs.KindUnknown, djs.LevelRoot, []byte("\"a\n"),
			djs.InvalidJsonPtrError{Pos: 2, Err: djs.InvalidCharacterError}},
		{"str:111", []byte("\"a\rz\""), djs.KindUnknown, djs.LevelRoot, []byte("\"a\r"),
			djs.InvalidJsonPtrError{Pos: 2, Err: djs.InvalidCharacterError}},
		{"str:112", []byte("\"a\tz\""), djs.KindUnknown, djs.LevelRoot, []byte("\"a\t"),
			djs.InvalidJsonPtrError{Pos: 2, Err: djs.InvalidCharacterError}},
		// Skip invalid UTF-8 characters. Will validate it on next level
		// by strconv.Unquote / strconv.Quote.
		// {"str:57", []byte(`"\uD834\q"`), djs.KindUnknown, djs.LevelRoot, []byte(`"\uD834\q`), djs.InvalidJsonPtrError{Pos: 8}},
		// {"str:25", []byte(`"\uD834\n"`), djs.KindUnknown, djs.LevelRoot, []byte(`"\uD834\n`), djs.InvalidJsonPtrError{Pos: 8}},
	}
	for _, el := range testCases {
		RunTokenizerTestPrimitiveCase(s, el)
	}
}

func (s *TokenizerTestPrimitiveSuite) TestTokenizer_Next_time() {
	tbl := []TokenizerTestPrimitiveElement{
		{"time:00", []byte(`"2015-05-14T12:34:56+02:00"`), djs.KindTime, djs.LevelRoot, []byte(`"2015-05-14T12:34:56+02:00"`), nil},
		{"time:01", []byte(`"2015-05-14T12:34:56.3+02:00"`), djs.KindTime, djs.LevelRoot, []byte(`"2015-05-14T12:34:56.3+02:00"`), nil},
		{"time:02", []byte(`"2015-05-14T12:34:56.37+02:00"`), djs.KindTime, djs.LevelRoot, []byte(`"2015-05-14T12:34:56.37+02:00"`), nil},
		{"time:03", []byte(`"2015-05-14T12:34:56.379+02:00"`), djs.KindTime, djs.LevelRoot, []byte(`"2015-05-14T12:34:56.379+02:00"`), nil},
		{"time:04", []byte(`"1970-01-01T00:00:00Z"`), djs.KindTime, djs.LevelRoot, []byte(`"1970-01-01T00:00:00Z"`), nil},
		{"time:05", []byte(`"0001-01-01T00:00:00Z"`), djs.KindTime, djs.LevelRoot, []byte(`"0001-01-01T00:00:00Z"`), nil},
		{"time:06", []byte(`"1985-04-12T23:20:50.52Z"`), djs.KindTime, djs.LevelRoot, []byte(`"1985-04-12T23:20:50.52Z"`), nil},
		{"time:07", []byte(`"1996-12-19T16:39:57-08:00"`), djs.KindTime, djs.LevelRoot, []byte(`"1996-12-19T16:39:57-08:00"`), nil},
		{"time:08", []byte(`"1990-12-31T23:59:60Z"`), djs.KindTime, djs.LevelRoot, []byte(`"1990-12-31T23:59:60Z"`), nil},
		{"time:09", []byte(`"1990-12-31T15:59:60-08:00"`), djs.KindTime, djs.LevelRoot, []byte(`"1990-12-31T15:59:60-08:00"`), nil},
		{"time:10", []byte(`"1937-01-01T12:00:27.87+00:20"`), djs.KindTime, djs.LevelRoot, []byte(`"1937-01-01T12:00:27.87+00:20"`), nil},
		{"time:11", []byte(`"2022-02-24T04:00:00+02:00"`), djs.KindTime, djs.LevelRoot, []byte(`"2022-02-24T04:00:00+02:00"`), nil},
		{"time:12", []byte(`"2022-07-12T21:55:16+01:00"`), djs.KindTime, djs.LevelRoot, []byte(`"2022-07-12T21:55:16+01:00"`), nil},
		{"time:13", []byte(`"2015-05-14T12:34:56.123Z"`), djs.KindTime, djs.LevelRoot, []byte(`"2015-05-14T12:34:56.123Z"`), nil},
		{"time:14", []byte(`"2015-05-14T12:34:56Z"`), djs.KindTime, djs.LevelRoot, []byte(`"2015-05-14T12:34:56Z"`), nil},
		// Invalid cases fall back to string
		{"time:100", []byte(`"2015-05-14E12:34:56.379+02:00"`), djs.KindString, djs.LevelRoot, []byte(`"2015-05-14E12:34:56.379+02:00"`), nil},
		{"time:101", []byte(`"2O15-O5-14T12:34:56.379+02:00"`), djs.KindString, djs.LevelRoot, []byte(`"2O15-O5-14T12:34:56.379+02:00"`), nil},
		{"time:102", []byte(`"1985-04-12T23:20:50.52ZZZZ"`), djs.KindString, djs.LevelRoot, []byte(`"1985-04-12T23:20:50.52ZZZZ"`), nil},
		{"time:103", []byte(`"2022-07-12 21:55:16"`), djs.KindString, djs.LevelRoot, []byte(`"2022-07-12 21:55:16"`), nil},
		{"time:104", []byte(`"20220712T215516Z"`), djs.KindString, djs.LevelRoot, []byte(`"20220712T215516Z"`), nil},
		{"time:105", []byte(`"20220712T215516+01:00"`), djs.KindString, djs.LevelRoot, []byte(`"20220712T215516+01:00"`), nil},
		{"time:106", []byte(`"1985-04-12T23:20:50.Z"`), djs.KindString, djs.LevelRoot, []byte(`"1985-04-12T23:20:50.Z"`), nil},
		{"time:107", []byte(`"not a Timestamp"`), djs.KindString, djs.LevelRoot, []byte(`"not a Timestamp"`), nil},
	}
	for _, el := range tbl {
		RunTokenizerTestPrimitiveCase(s, el)
	}
}

func RunTokenizerTestPrimitiveCase(s *TokenizerTestPrimitiveSuite, el TokenizerTestPrimitiveElement) {
	s.T().Run(el.idx, func(t *testing.T) {
		b := bytes.NewBuffer(el.in)
		sc := djs.NewJStructScanner(b)
		tk := djs.NewJStructTokenizer(sc)
		e := tk.Next()
		v := tk.Value()
		k := tk.Kind()
		l := tk.Level()

		assert.Equal(t, el.out, v, "%s Value %v != %v (%v)", el.idx, v, el.out, el.in)
		assert.Equal(t, el.kind, k, "%s Kind %v != %v", el.idx, k, el.kind)
		assert.Equal(t, el.level, l, "%s Level %v != %v", el.idx, l, el.level)
		assert.ErrorIs(t, e, el.err, "%s Next err: %v", el.idx, e)
	})
}
