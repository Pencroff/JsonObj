package experiment

import (
	"testing"
)

var bResult bool

/**
λ go test -bench=BenchmarkCheckers -timeout 30m
goos: windows
goarch: amd64
pkg: github.com/Pencroff/JsonStruct/experiment
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkCheckers/Sm_string_Re-12               538956687                2.305 ns/op    5640.82 MB/s
BenchmarkCheckers/Sm_string_Re6-12              598628541                2.126 ns/op    6113.53 MB/s
BenchmarkCheckers/Sm_string_Fn6-12              627275506                1.957 ns/op    6642.56 MB/s
BenchmarkCheckers/Sm_string_Fn-12               724225711                1.946 ns/op    6679.08 MB/s
BenchmarkCheckers/Sm_string_Time-12             608100198                1.844 ns/op    7049.93 MB/s
BenchmarkCheckers/Md_string_Re-12               579713504                2.163 ns/op    26350.57 MB/s
BenchmarkCheckers/Md_string_Re6-12              718574491                1.970 ns/op    28929.75 MB/s
BenchmarkCheckers/Md_string_Fn6-12              657972292                1.691 ns/op    33700.41 MB/s
BenchmarkCheckers/Md_string_Fn-12               622671144                1.789 ns/op    31859.90 MB/s
BenchmarkCheckers/Md_string_Time-12             533247302                2.043 ns/op    27901.98 MB/s
BenchmarkCheckers/Lg_string_Re-12               621202575                2.007 ns/op    276067.47 MB/s
BenchmarkCheckers/Lg_string_Re6-12              627695493                1.881 ns/op    294485.49 MB/s
BenchmarkCheckers/Lg_string_Fn6-12              622457325                1.717 ns/op    322708.58 MB/s
BenchmarkCheckers/Lg_string_Fn-12               608962712                1.755 ns/op    315610.12 MB/s
BenchmarkCheckers/Lg_string_Time-12             576784148                2.053 ns/op    269880.90 MB/s
BenchmarkCheckers/Float_string_Re-12            10701070               111.2 ns/op       251.85 MB/s
BenchmarkCheckers/Float_string_Re6-12           10401118               109.8 ns/op       254.94 MB/s
BenchmarkCheckers/Float_string_Fn6-12           37237706                33.16 ns/op      844.30 MB/s
BenchmarkCheckers/Float_string_Fn-12            326542784                3.611 ns/op    7754.97 MB/s
BenchmarkCheckers/Float_string_Time-12           8996566               133.6 ns/op       209.62 MB/s
BenchmarkCheckers/Time_string_Re-12              2613177               432.6 ns/op        78.59 MB/s
BenchmarkCheckers/Time_string_Re6-12             2775282               420.1 ns/op        80.92 MB/s
BenchmarkCheckers/Time_string_Fn6-12            20677854                60.27 ns/op      564.11 MB/s
BenchmarkCheckers/Time_string_Fn-12             100000000               12.97 ns/op     2621.42 MB/s
BenchmarkCheckers/Time_string_Time-12            2598003               455.6 ns/op        74.63 MB/s
BenchmarkCheckers/Many_variants_Re-12              98056             13387 ns/op         145.36 MB/s
BenchmarkCheckers/Many_variants_Re6-12             89527             12923 ns/op         150.59 MB/s
BenchmarkCheckers/Many_variants_Fn6-12            544622              2290 ns/op         849.88 MB/s
BenchmarkCheckers/Many_variants_Fn-12            3143151               371.7 ns/op      5234.76 MB/s
BenchmarkCheckers/Many_variants_Time-12            95800             12344 ns/op         157.65 MB/s
PASS
ok      github.com/Pencroff/JsonStruct/experiment       42.530s

Many_variants size: 1946 bytes
*/

func BenchmarkCheckers(b *testing.B) {
	tblMethod := []struct {
		name string
		in   []byte
	}{
		{"Sm string", []byte(`"Hello World"`)},
		{"Md string", []byte(`"One morning, when Gregor Samsa woke from troubled dream"`)},
		{"Lg string", []byte(`"The quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymphs grab quick-jived waltz. Brick quiz whangs jumpy veldt fox. Bright vixens jump; dozy fowl quack. Quick wafting zephyrs vex bold Jim. Quick zephyrs blow, vexing daft Jim. Sex-charged fop blew my junk TV quiz. How quickly daft jumping zebras vex. Two driven jocks help fax my big quiz. Quick, Baz, get my woven flax jodhpurs! \"Now fax quiz Jack!\" my b"`)},
		{"Float string", []byte(`"1234567890.0123456789e+123"`)},
		{"Time string", []byte(`"2015-05-14T12:34:56.123456-11:00"`)},
	}
	for _, el := range tblMethod {
		b.Run(el.name+" Re", func(b *testing.B) {
			n := false
			b.SetBytes(int64(len(el.in)))
			for i := 0; i < b.N; i++ {
				n = IsTimeStrReFn(el.in)
			}
			bResult = n
		})
		b.Run(el.name+" Re6", func(b *testing.B) {
			n := false
			b.SetBytes(int64(len(el.in)))
			for i := 0; i < b.N; i++ {
				n = IsTimeStrRe6Fn(el.in)
			}
			bResult = n
		})

		b.Run(el.name+" Fn6", func(b *testing.B) {
			n := false
			b.SetBytes(int64(len(el.in)))
			for i := 0; i < b.N; i++ {
				n = IsTimeStr6Fn(el.in)
			}
			bResult = n
		})
		b.Run(el.name+" Fn", func(b *testing.B) {
			n := false
			b.SetBytes(int64(len(el.in)))
			for i := 0; i < b.N; i++ {
				n = IsTimeStrHeadTailFn(el.in)
			}
			bResult = n
		})
		b.Run(el.name+" Time", func(b *testing.B) {
			n := false
			b.SetBytes(int64(len(el.in)))
			for i := 0; i < b.N; i++ {
				n = IsTimeStrTime(el.in)
			}
			bResult = n
		})
	}
	tbl := tiCheckerTestCases
	cnt := 0
	for _, el := range tbl {
		cnt += len(el.in)
	}
	b.Run("Many variants Re", func(b *testing.B) {
		n := false
		b.SetBytes(int64(cnt))
		for i := 0; i < b.N; i++ {
			for _, el := range tbl {
				n = IsTimeStrReFn(el.in)
			}
		}
		bResult = n
	})
	b.Run("Many variants Re6", func(b *testing.B) {
		n := false
		b.SetBytes(int64(cnt))
		for i := 0; i < b.N; i++ {
			for _, el := range tbl {
				n = IsTimeStrRe6Fn(el.in)
			}
		}
		bResult = n
	})
	b.Run("Many variants Fn6", func(b *testing.B) {
		n := false
		b.SetBytes(int64(cnt))
		for i := 0; i < b.N; i++ {
			for _, el := range tbl {
				n = IsTimeStr6Fn(el.in)
			}
		}
		bResult = n
	})
	b.Run("Many variants Fn", func(b *testing.B) {
		n := false
		b.SetBytes(int64(cnt))
		for i := 0; i < b.N; i++ {
			for _, el := range tbl {
				n = IsTimeStrHeadTailFn(el.in)
			}
		}
		bResult = n
	})
	b.Run("Many variants Time", func(b *testing.B) {
		n := false
		b.SetBytes(int64(cnt))
		for i := 0; i < b.N; i++ {
			for _, el := range tbl {
				n = IsTimeStrTime(el.in)
			}
		}
		bResult = n
	})
}
