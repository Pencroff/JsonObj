package test_suite

import (
	"bytes"
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/suite"
	"io"
	"testing"
)

func TestJStruct_Scanner(t *testing.T) {
	s := new(ScannerTestSuite)
	suite.Run(t, s)
}

type ScannerTestSuite struct {
	suite.Suite
	rd *bytes.Buffer
	sc djs.JStructScanner
}

func (s *ScannerTestSuite) SetupTest() {
	s.rd = bytes.NewBuffer([]byte(`ABCDEFJHIJKLMNOPQRSTUVWXYZ`))
	s.sc = djs.NewJStructScannerWithSize(s.rd, 4)
}

func (s *ScannerTestSuite) TestScanner_BasicBehavior() {
	data := []byte(`ABCDEFJHIJKLMNOPQRSTUVWXYZ`)
	for idx, ch := range data {
		err := s.sc.Next()
		s.NoError(err)
		s.Equal(idx, s.sc.Index())
		s.Equal(ch, s.sc.Current())

	}
	s.Equal(data, s.sc.Peek())
	s.Equal(0, s.sc.Total())
	s.sc.Release()
	s.Equal(len(data), s.sc.Total())
}

func (s *ScannerTestSuite) TestScanner_Offset() {
	s.sc.Offset(3)
	s.Equal([]byte(`ABC`), s.sc.Peek())
	s.Equal(2, s.sc.Index())
	s.sc.Release()
	e := s.sc.Offset(26)
	s.ErrorIs(e, io.EOF)
	s.Equal(25, s.sc.Index())
	s.Equal([]byte(`DEFJHIJKLMNOPQRSTUVWXYZ`), s.sc.Peek())
	s.Equal(3, s.sc.Total())
	s.sc.Release()
	s.Equal(26, s.sc.Total())
}
