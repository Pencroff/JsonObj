package helper

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type CheckerTestSuite struct {
	suite.Suite
}

func Test_CheckerTestSuite(t *testing.T) {
	s := new(CheckerTestSuite)
	suite.Run(t, s)
}

func (s *CheckerTestSuite) TestFloatToInt() {
	tbl := []struct {
		in  string
		out bool
	}{
		{"2016-01-19T15:21:32.59+02:00", true},
		{"2015-05-14T12:34:56+02:00", true},
		{"2015-05-14T12:34:56Z", true},
		{"1970-01-01T00:00:00Z", true},
		{"1970-01-01T00:00:00+00:00", true},
		{"0001-01-01T00:00:00Z", true},
		{"hello", false},
		{"", false},
		{"2015+05-14T12:34:56.789+02:00", false},
	}
	for _, el := range tbl {
		s.Equal(el.out, IsTimeRe.MatchString(el.in))
	}
}
