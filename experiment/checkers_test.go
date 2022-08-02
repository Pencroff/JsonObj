package experiment

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

var tiCheckerTestCases = []struct {
	in  []byte
	out bool
}{
	{[]byte(`"2016-01-19T15:21:32Z"`), true},
	{[]byte(`"2015-05-14T12:34:56.789+05:00"`), true},
	{[]byte(`"2015-05-14T12:34:56.789Z"`), true},
	{[]byte(`"1970-01-01T00:00:00Z"`), true},
	{[]byte(`"1970-01-01T00:00:00+00:00"`), true},
	{[]byte(`"0001-01-01T00:00:00Z"`), true},
	// examples from rfc3339
	{[]byte(`"1985-04-12T23:20:50.52Z"`), true},
	{[]byte(`"1996-12-19T16:39:57-08:00"`), true},
	{[]byte(`"1990-12-31T23:59:59Z"`), true},      // not passed by time.Parse with leap second, keep 59 instead of 60
	{[]byte(`"1990-12-31T15:59:59-08:00"`), true}, // not passed by time.Parse with leap second, keep 59 instead of 60
	{[]byte(`"1937-01-01T12:00:27.87+00:20"`), true},
	// end examples from rfc3339
	{[]byte(`"2022-02-24T04:00:00+02:00"`), true},
	{[]byte(`"2022-02-24T02:00:00Z"`), true},
	{[]byte(`"2022-07-12T21:55:16+01:00"`), true},
	{[]byte(`"2022-07-12T21:55:16.1-02:00"`), true},
	{[]byte(`"2022-07-12T21:55:16.12+03:00"`), true},
	{[]byte(`"2022-07-12T21:55:16.123-04:00"`), true},
	{[]byte(`"2022-07-12T21:55:16.1234+05:00"`), true},
	{[]byte(`"2022-07-12T21:55:16.12345-06:00"`), true},
	{[]byte(`"2022-07-12T21:55:16.123456+07:00"`), true},
	{[]byte(`"2022-07-12T21:55:16Z"`), true},
	{[]byte(`"2022-07-12T21:55:16.1Z"`), true},
	{[]byte(`"2022-07-12T21:55:16.12Z"`), true},
	{[]byte(`"2022-07-12T21:55:16.123Z"`), true},
	{[]byte(`"2022-07-12T21:55:16.1234Z"`), true},
	{[]byte(`"2022-07-12T21:55:16.12345Z"`), true},
	{[]byte(`"2022-07-12T21:55:16.123456Z"`), true},

	// invalid
	{[]byte(`"2015-05-14E12:34:56.379+02:00"`), false},
	{[]byte(`"2015-05-14E12:34:56.379+02:00"`), false},
	{[]byte(`"2015-05-14 12:34:56.379+02:00"`), false},
	{[]byte(`"2015+05-14T12:34:56.789-02:00"`), false},
	{[]byte(`"2015-05-14T12:34:56.789_02:00"`), false},
	{[]byte(`2022-07-12T21:55:16.12+03:00"`), false},
	{[]byte(`"2022-07-12T21:55:16.12+03:00`), false},
	{[]byte(`"2022-07-12T21:55:16.12+03:00""`), false},
	{[]byte(`2022-07-12T21:55:16.12Z"`), false},
	{[]byte(`"2022-07-12T21:55:16.12Z`), false},
	{[]byte(`"2022-07-12T21:55:16.12Z""`), false},
	{[]byte(`"2022-07-12T21:55:16.12ab56Z"`), false},
	{[]byte(`"2022-07-12T21:55:16.12xy56+07:00"`), false},
	{[]byte(`"2015-05-14T12:34:56.789-5:0"`), false},
	{[]byte(`"2O15-O5-14T12:34:56.379+02:00"`), false},
	{[]byte(`"1985-04-12T23:20:50.52ZZZZ"`), false},
	{[]byte(`"2022-07-12 21:55:16"`), false},
	{[]byte(`"20220712T215516Z"`), false},
	{[]byte(`"20220712T215516+01:00"`), false},
	{[]byte(`"1985-04-12T23:20:50.Z"`), false},
	{[]byte(`"not a Timestamps"`), false},
	{[]byte(`"Hello World"`), false},
	{[]byte(`"One morning, when Gregor Samsa woke from troubled dream"`), false},
	{[]byte(`"The quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymphs grab quick-jived waltz. Brick quiz whangs jumpy veldt fox. Bright vixens jump; dozy fowl quack. Quick wafting zephyrs vex bold Jim. Quick zephyrs blow, vexing daft Jim. Sex-charged fop blew my junk TV quiz. How quickly daft jumping zebras vex. Two driven jocks help fax my big quiz. Quick, Baz, get my woven flax jodhpurs! \"Now fax quiz Jack!\" my b"`), false},
	{[]byte(`"31415926535.897932385"`), false},
}

type CheckerTestSuite struct {
	suite.Suite
	tbl []struct {
		in  []byte
		out bool
	}
	size int
}

func Test_CheckerTestSuite(t *testing.T) {
	s := new(CheckerTestSuite)
	suite.Run(t, s)
}

func (s *CheckerTestSuite) SetupTest() {
	s.tbl = tiCheckerTestCases
	s.size = 0
	for _, el := range s.tbl {
		s.size += len(el.in)
	}
	fmt.Println()
	fmt.Println(s.size)
}

func (s *CheckerTestSuite) TestReFn() {
	for idx, el := range s.tbl {
		l := len(el.in)
		s.T().Run("RE__"+string(el.in[1:l-1]), func(t *testing.T) {
			assert.Equal(t, el.out, IsTimeStrReFn(el.in), "%d - %s", idx, string(el.in))
		})
	}
}
func (s *CheckerTestSuite) TestRe6Fn() {
	for idx, el := range s.tbl {
		l := len(el.in)
		s.T().Run("RE7__"+string(el.in[1:l-1]), func(t *testing.T) {
			assert.Equal(t, el.out, IsTimeStrRe6Fn(el.in), "%d - %s", idx, string(el.in))
		})
	}
}

func (s *CheckerTestSuite) TestManual6Fn() {
	for idx, el := range s.tbl {
		l := len(el.in)
		s.T().Run("Fn7__"+string(el.in[1:l-1]), func(t *testing.T) {
			assert.Equal(t, el.out, IsTimeStr6Fn(el.in), "%d - %s", idx, string(el.in))
		})
	}
}

func (s *CheckerTestSuite) TestHeadTailFn() {
	for idx, el := range s.tbl {
		l := len(el.in)
		s.T().Run("Fn__"+string(el.in[1:l-1]), func(t *testing.T) {
			assert.Equal(t, el.out, IsTimeStrHeadTailFn(el.in), "%d - %s", idx, string(el.in))
		})
	}
}

func (s *CheckerTestSuite) TestIsTimeStrTime() {
	for idx, el := range s.tbl {
		l := len(el.in)
		s.T().Run("Time__"+string(el.in[1:l-1]), func(t *testing.T) {
			assert.Equal(t, el.out, IsTimeStrTime(el.in), "%d - %s", idx, string(el.in))
		})
	}
}
