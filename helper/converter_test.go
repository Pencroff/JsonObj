package helper

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"strconv"
	"testing"
)

type ConverterTestSuite struct {
	suite.Suite
}

func Test_ConverterTestSuite(t *testing.T) {
	s := new(ConverterTestSuite)
	suite.Run(t, s)
}

func (s *ConverterTestSuite) TestFloatToInt() {
	s.Equal(int64(0), FloatToInt(0.49))
	s.Equal(int64(0), FloatToInt(-0.49))
	s.Equal(int64(1), FloatToInt(0.5))
	s.Equal(int64(-1), FloatToInt(-0.5))
	s.Equal(int64(2), FloatToInt(1.99))
	s.Equal(int64(-2), FloatToInt(-1.99))
}

func (s *ConverterTestSuite) TestStringToInt() {
	str := "-123456789"
	i, err := strconv.Atoi(str)
	s.Equal(-123456789, i)
	s.Nil(err)
	ia, err := strconv.ParseInt(str, 10, 64)
	s.Equal(-123456789, int(ia))
	s.Nil(err)
	var ib int
	_, err = fmt.Sscan(str, &ib)
	s.Nil(err)
	s.Equal(-123456789, ib)
	ic, _ := StringToInt(str)
	s.Equal(-123456789, int(ic))
}

func (s *ConverterTestSuite) TestStringToIntTable() {
	var parseInt64Tests = []struct {
		in  string
		out int64
		ok  bool
	}{
		{"", 0, false},
		{"0", 0, true},
		{"-0", 0, true},
		//{"+0", 0, true}, // not supported in JSON
		{"1", 1, true},
		{"-1", -1, true},
		//{"+1", 1, true},
		{"12345", 12345, true},
		{"-12345", -12345, true},
		{"012345", 12345, true},
		{"-012345", -12345, true},
		{"98765432100", 98765432100, true},
		{"-98765432100", -98765432100, true},
		{"9223372036854775807", 1<<63 - 1, true},
		{"-9223372036854775807", -(1<<63 - 1), true},
		{"9223372036854775808", 1<<63 - 1, false},
		{"-9223372036854775808", -1 << 63, true},
		{"9223372036854775809", 1<<63 - 1, false},
		{"-9223372036854775809", -1 << 63, false},
		{"-1_2_3_4_5", 0, false},
		{"-_12345", 0, false},
		{"_12345", 0, false},
		{"1__2345", 0, false},
		{"12345_", 0, false},
	}
	for idx, el := range parseInt64Tests {
		i, ok := StringToInt(el.in)
		s.Equal(el.out, i, "[%d] StringToInt(%q) = %d, %v", idx, el.in, i, ok)
		s.Equal(el.ok, ok)
	}
}

func (s *ConverterTestSuite) TestStringToUint() {
	str := "18446744073709551615"
	i, ok := StringToUint(str)
	s.Equal(uint(18446744073709551615), uint(i))
	s.Equal(true, ok)
}

func (s *ConverterTestSuite) TestStringToUintTable() {
	var parseUint64Tests = []struct {
		in  string
		out uint64
		ok  bool
	}{
		{"", 0, false},
		{"0", 0, true},
		{"1", 1, true},
		{"12345", 12345, true},
		{"012345", 12345, true},
		{"12345x", 0, false},
		{"98765432100", 98765432100, true},
		{"18446744073709551615", 1<<64 - 1, true},
		{"18446744073709551616", 1<<64 - 1, false},
		{"18446744073709551620", 1<<64 - 1, false},
		{"1_2_3_4_5", 0, false}, // base=10 so no underscores allowed
		{"_12345", 0, false},
		{"1__2345", 0, false},
		{"12345_", 0, false},
		{"-0", 0, false},
		{"-1", 0, false},
		{"+1", 0, false},
	}
	for idx, el := range parseUint64Tests {
		i, ok := StringToUint(el.in)
		s.Equal(el.out, i, "[%d] StringToUint(%q) = %d, %v", idx, el.in, i, ok)
		s.Equal(el.ok, ok)
	}
}

func (s *ConverterTestSuite) TestMinMaxInt() {
	s.Equal(int64(-9223372036854775808), MinInt)
	s.Equal(int64(9223372036854775807), MaxInt)
	s.Equal(uint64(18446744073709551615), MaxUint)
	s.Equal(uint64(9223372036854775807), MaxIntUint)
	s.Equal(uint64(9223372036854775808), MinIntUint)
	s.Equal(int64(9007199254740991), MaxSafeInt)
	s.Equal(int64(-9007199254740991), MinSafeInt)
}

func (s *ConverterTestSuite) TestIntUintToFloat() {
	s.Equal(-9223372036854775808.0, float64(MinInt))
	s.Equal(9223372036854775807.0, float64(MaxInt))
	s.Equal(18446744073709551615.0, float64(MaxUint))
	s.Equal(9223372036854775807.0, float64(MaxIntUint))
	s.Equal(9223372036854775808.0, float64(MinIntUint))
	s.Equal(9007199254740991.0, float64(MaxSafeInt))
	s.Equal(-9007199254740991.0, float64(MinSafeInt))
}
