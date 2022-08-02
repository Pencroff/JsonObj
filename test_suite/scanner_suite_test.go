package test_suite

import (
	"bytes"
	"fmt"
	djs "github.com/Pencroff/JsonStruct"
	tl "github.com/Pencroff/JsonStruct/tool"
	"github.com/stretchr/testify/suite"
	"io"
	"os"
	"testing"
)

func TestJStruct_Scanner(t *testing.T) {
	s := new(ScannerTestSuite)
	suite.Run(t, s)
}

type ScannerTestSuite struct {
	suite.Suite
	buf *bytes.Buffer
	rd  djs.JStructScanner
}

func (s *ScannerTestSuite) SetupTest() {
	size := 16
	th := size >> 2
	s.buf = bytes.NewBuffer([]byte(`ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz`))
	s.rd = djs.NewJStructScannerWithParam(s.buf, size, th)
}

func (s *ScannerTestSuite) TestScanner_BasicBehavior() {
	data := []byte(`ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz`)
	for idx, ch := range data {
		err := s.rd.Next()
		s.NoError(err)
		s.Equal(ch, s.rd.Current(), "idx:%d %s != %s", idx, string(ch), string(s.rd.Current()))
		s.Equal(idx, s.rd.Index())
	}
	s.Equal(data, s.rd.Bytes())
	s.Equal(len(data)-1, s.rd.Index())
}

func (s *ScannerTestSuite) TestScanner_Offset() {
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
		e := s.rd.Scan(tc.offset)
		//fmt.Printf("%d: %s\n", tc.offset, string(s.rd.Buffer()))
		//fmt.Println(s.rd.Window())
		s.Equal(tc.data, s.rd.Bytes())
		s.Equal(tc.idx, s.rd.Index())
		s.Equal(tc.total, s.rd.Index()+1)
		s.Equal(tc.e, e)
	}
}

func (s *ScannerTestSuite) TestScanner_EOF() {
	e := s.rd.Scan(61)
	s.NoError(e)
	st, ptr := s.rd.Window()
	//fmt.Println(st, ptr, string(*s.rd.Buffer()))
	s.Equal(0, st)
	s.Equal(60, ptr)
	s.Equal(60, s.rd.Index())
	s.Equal([]byte(`ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz`), s.rd.Buffer())
	s.Equal([]byte(`ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxy`), s.rd.Bytes())
	e = s.rd.Next()
	st, ptr = s.rd.Window()
	s.NoError(e)
	s.Equal(61, st)
	s.Equal(61, ptr)
	s.Equal(byte('z'), s.rd.Current())
	//fmt.Printf("-------\n")
	e = s.rd.Next()
	s.ErrorIs(e, io.EOF)
	s.Equal(byte('z'), s.rd.Current())
}

func (s *ScannerTestSuite) TestScanner_Three_Byte_Issue() {
	data := []byte(`nnn`)
	b := bytes.NewBuffer(data)
	sc := djs.NewJStructScanner(b)
	e := sc.Next()
	s.NoError(e)
	s.Equal(byte('n'), sc.Current())
	e = sc.Scan(3)
	//fmt.Printf("%s\n", string(s.rd.Buffer()))
	//fmt.Println(s.rd.Window())
	s.ErrorIs(e, io.EOF)
	st, ptr := sc.Window()
	s.Equal(0, st)
	s.Equal(2, ptr)
	s.Equal(data, sc.Bytes())
}

func (s *ScannerTestSuite) TestScanner_should_keep_correct_idx() {
	data := []byte(`abc`)
	b := bytes.NewBuffer(data)
	sc := djs.NewJStructScanner(b)
	e := sc.Next()
	s.NoError(e)
	s.Equal(byte('a'), sc.Current())
	s.Equal(0, sc.Index())
	e = sc.Scan(3)
	s.ErrorIs(e, io.EOF)
	st, ptr := sc.Window()
	s.Equal(0, st)
	s.Equal(2, ptr)
	s.Equal(2, sc.Index())
	s.Equal(data, sc.Bytes())
	s.Equal(2, sc.Index())
	e = sc.Next()
	s.ErrorIs(e, io.EOF)
	s.Equal(2, sc.Index())
	s.Equal(byte(0), sc.Current())
	s.Equal([]byte{}, sc.Bytes())
	e = sc.Scan(5)
	s.ErrorIs(e, io.EOF)
	s.Equal(2, sc.Index())
	s.Equal(byte(0), sc.Current())
	s.Equal([]byte{}, sc.Bytes())
}

func (s *ScannerTestSuite) TestScanner_DifferentBuffers() {
	bufSizes := []struct {
		size      int
		threshold int
	}{
		{8, 2},
		{16, 4},
		{32, 8},
		{512, 128},
	}
	testCases := []struct {
		in      []byte
		offset  int
		out     []byte
		current byte
		idx     int
		e       error
	}{
		{[]byte(`ABC`), 1, []byte(`A`), 0, 0, nil},
		{[]byte(`ABC`), 2, []byte(`AB`), 0, 1, nil},
		{[]byte(`ABC`), 3, []byte(`ABC`), 0, 2, nil},
		{[]byte(`ABC`), 4, []byte(`ABC`), 0, 2, io.EOF},
		{[]byte(`ABC`), 16, []byte(`ABC`), 0, 2, io.EOF},
	}
	for _, bufSize := range bufSizes {
		for _, tc := range testCases {
			sc := djs.NewJStructScannerWithParam(bytes.NewBuffer(tc.in), bufSize.size, bufSize.threshold)
			e := sc.Scan(tc.offset)
			s.ErrorIs(e, tc.e)
			s.Equal(tc.idx, sc.Index())
			s.Equal(tc.out, sc.Bytes())
			s.Equal(tc.current, sc.Current())
		}
	}

}

func Benchmark_Scanner(b *testing.B) {
	data, _ := tl.ReadGzip("../benchmark/data/canada.json.gz")
	data2, _ := tl.ReadGzip("../benchmark/data/large-file.json.gz")
	offsetSize := 20
	fmt.Printf("Data size: %.2f Mb\n", float64(len(data))/1024/1024)
	var e error
	cnt := 0
	total := 0
	b.Run("Scanner___canada", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buf := bytes.NewBuffer(data)
			rd := djs.NewJStructScanner(buf)
			for {
				e = rd.Scan(offsetSize)
				d := rd.Bytes()
				cnt += len(d)
				if e == io.EOF {
					total = rd.Index() + 1
					break
				}
				if e != nil {
					b.Fatal(e)
				}
			}
		}
		b.StopTimer()
	})
	fmt.Println("From memory")
	fmt.Printf("Readed size: %.2f Mb\n", float64(total)/1024/1024)
	if len(data) != total {
		b.Fatal("Data size mismatch", total)
	}
	fmt.Println("Times:", cnt/total)

	b.Run("Scanner___canada_no_threshold", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buf := bytes.NewBuffer(data)
			rd := djs.NewJStructScannerWithParam(buf, djs.JStructScannerBufferSize, 0)
			for {
				e = rd.Scan(offsetSize)
				d := rd.Bytes()
				cnt += len(d)
				if e == io.EOF {
					total = rd.Index() + 1
					break
				}
				if e != nil {
					b.Fatal(e)
				}
			}
		}
		b.StopTimer()
	})
	b.Run("Scanner___canada_large_threshold", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buf := bytes.NewBuffer(data)
			rd := djs.NewJStructScannerWithParam(buf, djs.JStructScannerBufferSize, djs.JStructScannerBufferSize/3)
			for {
				e = rd.Scan(offsetSize)
				d := rd.Bytes()
				cnt += len(d)
				if e == io.EOF {
					total = rd.Index() + 1
					break
				}
				if e != nil {
					b.Fatal(e)
				}
			}
		}
		b.StopTimer()
	})

	b.Run("Scanner___large-file-mem", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data2)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buf := bytes.NewBuffer(data2)
			rd := djs.NewJStructScanner(buf)
			for {
				e = rd.Scan(offsetSize)
				d := rd.Bytes()
				cnt += len(d)
				if e == io.EOF {
					total = rd.Index() + 1
					break
				}
				if e != nil {
					b.Fatal(e)
				}
			}
		}
		b.StopTimer()
	})
	fmt.Println("From memory")
	fmt.Printf("Readed size: %.2f Mb\n", float64(total)/1024/1024)
	if len(data2) != total {
		b.Fatal("Data size mismatch", total)
	}
	fmt.Println("Times:", cnt/total)

	b.Run("Scanner___large-file-drive", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			frd, _ := os.Open("../benchmark/data/large-file.json")
			rd := djs.NewJStructScanner(frd)
			for {
				e = rd.Scan(offsetSize)
				d := rd.Bytes()
				cnt += len(d)
				if e == io.EOF {
					total = rd.Index() + 1
					break
				}
				if e != nil {
					b.Fatal(e)
				}
			}
		}
		b.SetBytes(int64(total))
	})
	fmt.Println("From drive")
	fmt.Printf("Readed size: %.2f Mb\n", float64(total)/1024/1024)
}
