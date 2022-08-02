package test_suite

import (
	"bytes"
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io"
	"testing"
)

type TokenizerTestExpectation struct {
	idx   string
	kind  djs.TokenizerKind
	level djs.TokenizerLevel
	value []byte
	err   error
}

type TokenizerTestObjArrElement struct {
	idx string
	in  []byte
	out []TokenizerTestExpectation
}

func TestJStruct_TokenizerObjArr(t *testing.T) {
	s := new(TokenizerTestObjArrSuite)
	suite.Run(t, s)
}

type TokenizerTestObjArrSuite struct {
	suite.Suite
}

func (s *TokenizerTestObjArrSuite) SetupTest() {
}

func (s *TokenizerTestObjArrSuite) TestTokenizer_Next_Array_Primitive() {
	tbl := []TokenizerTestObjArrElement{
		{"arr:00", []byte(` [] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindLiteral, djs.LevelArrayEnd, nil, nil},
		}},
		{"arr:01", []byte(` [ ] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindLiteral, djs.LevelArrayEnd, nil, nil},
		}},
		{"arr:02", []byte(` [ null ] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindNull, djs.LevelValueLast, []byte(`null`), nil},
		}},
		{"arr:03", []byte(` [ true , false ] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindTrue, djs.LevelValue, []byte(`true`), nil},
			{"02", djs.KindFalse, djs.LevelValueLast, []byte(`false`), nil},
		}},
		{"arr:04", []byte(` [ 1, \n 2, \t 3 ] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindNumber, djs.LevelValue, []byte(`1`), nil},
			{"02", djs.KindNumber, djs.LevelValue, []byte(`2`), nil},
			{"03", djs.KindNumber, djs.LevelValueLast, []byte(`3`), nil},
		}},
		{"arr:05", []byte(` [
									0.1,
									-1,
									3.14 
								   ] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindFloatNumber, djs.LevelValue, []byte(`0.1`), nil},
			{"02", djs.KindNumber, djs.LevelValue, []byte(`-1`), nil},
			{"03", djs.KindFloatNumber, djs.LevelValueLast, []byte(`-3.14`), nil},
		}},
		{"arr:06", []byte(`["a","b","c"]`), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindString, djs.LevelValue, []byte(`"a"`), nil},
			{"02", djs.KindString, djs.LevelValue, []byte(`"b"`), nil},
			{"03", djs.KindString, djs.LevelValueLast, []byte(`"c"`), nil},
		}},
		{"arr:07", []byte(`["1970-01-01T00:00:00Z","1985-04-12T23:20:50.Z"]`), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindTime, djs.LevelValue, []byte(`"1970-01-01T00:00:00Z"`), nil},
			{"03", djs.KindString, djs.LevelValueLast, []byte(`"1985-04-12T23:20:50.Z"`), nil},
		}},
		// Invalid
		{"arr:100", []byte(` [   `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindUnknown, djs.LevelArray, nil,
				djs.InvalidJsonPtrError{Pos: 4, Err: io.EOF}},
		}},
		{"arr:101", []byte(` [ null, true  `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindNull, djs.LevelArray, []byte(`null`), nil},
			{"02", djs.KindUnknown, djs.LevelArray, []byte(`true  `),
				djs.InvalidJsonPtrError{Pos: 14, Err: io.EOF}},
		}},
		{"arr:102", []byte(` [ 1, 2 } `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindNumber, djs.LevelArray, []byte(`1`), nil},
			{"02", djs.KindUnknown, djs.LevelArray, []byte(`2 }`),
				djs.InvalidJsonPtrError{Pos: 8}},
		}},
		{"arr:103", []byte(` ["extra comma",] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindString, djs.LevelArray, []byte("extra comma"), nil},
			{"02", djs.KindUnknown, djs.LevelArray, []byte(`]`),
				djs.InvalidJsonTokenPtrError{Pos: 16}},
		}},
		{"arr:104", []byte(` ["double extra comma",,] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindString, djs.LevelArray, []byte("double extra comma"), nil},
			{"02", djs.KindUnknown, djs.LevelArray, []byte(`,`),
				djs.InvalidJsonTokenPtrError{Pos: 24}},
		}},
		{"arr:105", []byte(` [, "<-- missing value"] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindUnknown, djs.LevelArray, []byte(`,`),
				djs.InvalidJsonTokenPtrError{Pos: 2}},
		}},
		{"arr:106", []byte(` ["comma after the close"], `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindString, djs.LevelArrayEnd, []byte("comma after the close"), nil},
			{"02", djs.KindUnknown, djs.LevelRoot, []byte(`,`),
				djs.InvalidJsonTokenPtrError{Pos: 26}},
		}},
		{"arr:107", []byte(` ["extra close"]] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindString, djs.LevelArrayEnd, []byte("extra close"), nil},
			{"02", djs.KindUnknown, djs.LevelRoot, []byte(`]`),
				djs.InvalidJsonTokenPtrError{Pos: 16}},
		}},
		{"arr:108", []byte(` ["illegal backslash escape: \x15"] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindUnknown, djs.LevelArray, []byte(`"illegal backslash escape: \x`),
				djs.InvalidJsonTokenPtrError{Pos: 28, Err: djs.InvalidEscapeCharacterError}},
		}},
		{"arr:109", []byte(` ["illegal backslash escape: \017"] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindUnknown, djs.LevelArray, []byte(`"illegal backslash escape: \0`),
				djs.InvalidJsonTokenPtrError{Pos: 28, Err: djs.InvalidEscapeCharacterError}},
		}},
		{"arr:110", []byte(` [\naked] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindUnknown, djs.LevelArray, []byte(`\`),
				djs.InvalidJsonTokenPtrError{Pos: 2}},
		}},
		{"arr:111", []byte(` ["colon instead of comma" : false] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindUnknown, djs.LevelArray, []byte(`"colon instead of comma" :`),
				djs.InvalidJsonTokenPtrError{Pos: 27}},
		}},
		{"arr:112", []byte(` ["bad value", truth] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"01", djs.KindString, djs.LevelValue, []byte(`"bad value"`), nil},
			{"02", djs.KindUnknown, djs.LevelArray, []byte(`"trut`),
				djs.InvalidJsonTokenPtrError{Pos: 27}},
		}},
		{"arr:113", []byte(` ['single quote'] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"02", djs.KindUnknown, djs.LevelArray, []byte(`'`),
				djs.InvalidJsonTokenPtrError{Pos: 2}},
		}},
		{"arr:114", []byte(` ["	tab	character	in	string	"] `), []TokenizerTestExpectation{
			{"00", djs.KindLiteral, djs.LevelArray, nil, nil},
			{"02", djs.KindUnknown, djs.LevelArray, []byte(`"	`),
				djs.InvalidJsonTokenPtrError{Pos: 3, Err: djs.InvalidCharacterError}},
		}},
	}
	for _, el := range tbl {
		RunTokenizerTestCaseAndExpectations(el, s)
	}
}

func RunTokenizerTestCaseAndExpectations(el TokenizerTestObjArrElement, s *TokenizerTestObjArrSuite) {
	s.T().Run(el.idx, func(t *testing.T) {
		b := bytes.NewBuffer(el.in)
		sc := djs.NewJStructScanner(b)
		tk := djs.NewJSStructTokenizer(sc)
		for _, exp := range el.out {
			e := tk.Next()
			v := tk.Value()
			k := tk.Kind()
			l := tk.Level()
			name := el.idx + "__" + exp.idx
			assert.Equal(t, exp.value, v, "%s Value %v != %v (%v)", name, v, exp.value, el.in)
			assert.Equal(t, exp.kind, k, "%s Kind %v != %v", name, k, exp.kind)
			assert.Equal(t, exp.level, l, "%s Level %v != %v", name, l, exp.level)
			assert.ErrorIs(t, e, exp.err, "%s Next err: %v", name, e)
		}

	})
}
