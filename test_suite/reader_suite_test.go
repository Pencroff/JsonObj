package test_suite

import (
	"bytes"
	"fmt"
	djs "github.com/Pencroff/JsonStruct"
	h "github.com/Pencroff/JsonStruct/helper"
	"github.com/stretchr/testify/suite"
	"io"
	"testing"
)

func TestJStruct_Reader(t *testing.T) {
	s := new(ReaderTestSuite)
	suite.Run(t, s)
}

type ReaderTestSuite struct {
	suite.Suite
	buf *bytes.Buffer
	rd  djs.JStructReader
}

func (s *ReaderTestSuite) SetupTest() {
	djs.JStructReaderBufferSize = 16
	djs.JSStructReaderBufferThreshold = djs.JStructReaderBufferSize >> 2
	s.buf = bytes.NewBuffer([]byte(`ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz`))
	s.rd = djs.NewJStructReader(s.buf)
}

func (s *ReaderTestSuite) TestScanner_BasicBehavior() {
	data := []byte(`ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz`)
	for idx, ch := range data {
		err := s.rd.Next()
		s.NoError(err)
		s.Equal(ch, s.rd.Current(), "idx:%d %s != %s", idx, string(ch), string(s.rd.Current()))
		s.Equal(idx, s.rd.Index())
	}
	s.Equal(0, s.rd.Total())
	s.Equal(data, s.rd.Release())
	s.Equal(len(data), s.rd.Total())
}

func (s *ReaderTestSuite) TestScanner_Offset() {
	testCases := []struct {
		offset int
		data   []byte
		idx    int
		total  int
		e      error
	}{
		{2, []byte(`AB`), 1, 2, nil},
		{3, []byte(`CDE`), 4, 5, nil},
		{4, []byte(`FGHI`), 8, 9, nil},
		{5, []byte(`JKLMN`), 13, 14, nil},
		{6, []byte(`OPQRST`), 19, 20, nil},
		{7, []byte(`UVWXYZ0`), 26, 27, nil},
		{8, []byte(`12345678`), 34, 35, nil},
		{9, []byte(`9abcdefgh`), 43, 44, nil},
		{10, []byte(`ijklmnopqr`), 53, 54, nil},
		{11, []byte(`stuvwxyz`), 61, 62, io.EOF},
	}
	for _, tc := range testCases {
		e := s.rd.Offset(tc.offset)
		//fmt.Printf("%d: %s\n", tc.offset, string(s.rd.Buffer()))
		s.Equal(tc.data, s.rd.Release())
		s.Equal(tc.idx, s.rd.Index())
		s.Equal(tc.total, s.rd.Total())
		s.Equal(tc.e, e)
	}
}

func (s *ReaderTestSuite) TestScanner_EOF() {
	e := s.rd.Offset(61)
	s.NoError(e)
	s.Equal(60, s.rd.Index())
	s.Equal([]byte(`ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxy`), s.rd.Release())
	e = s.rd.Next()
	s.NoError(e)
	s.Equal(byte('z'), s.rd.Current())
	fmt.Printf("-------\n")
	e = s.rd.Next()
	s.ErrorIs(e, io.EOF)
	s.Equal(byte('z'), s.rd.Current())
}

func Benchmark_Reader_canada(b *testing.B) {
	data, _ := h.ReadData("../benchmark/data/canada.json.gz")
	fmt.Printf("Data size: %.2f Mb\n", float64(len(data))/1024/1024)
	var e error
	cnt := 0
	total := 0
	b.Run("Reader___canada", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buf := bytes.NewBuffer(data)
			rd := djs.NewJStructReader(buf)
			for {
				e = rd.Offset(8)
				d := rd.Release()
				cnt += len(d)
				if e != nil {
					total = rd.Total()
					break
				}
			}
		}
		b.StopTimer()
	})

	fmt.Printf("Readed size: %.2f Mb\n", float64(total)/1024/1024)
	if len(data) != total {
		b.Fatal("Data size mismatch", total)
	}
	b.Log("Times", cnt/total)
}
