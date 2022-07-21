package helper

import (
	"testing"
)

var bResult bool

/**
λ go test -bench=BenchmarkCheckers
goos: windows
goarch: amd64
pkg: github.com/Pencroff/JsonStruct/helper
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkCheckers/Sm_string-12          279017856                4.099 ns/op
BenchmarkCheckers/Md_string-12          25530284                48.53 ns/op
BenchmarkCheckers/Lg_string-12          21995424                50.50 ns/op
BenchmarkCheckers/Float_string-12       13187088                83.40 ns/op
BenchmarkCheckers/Time_String-12         4367382               302.5 ns/op
PASS
ok      github.com/Pencroff/JsonStruct/helper   7.055s
*/

func BenchmarkCheckers(b *testing.B) {
	tblMethod := []struct {
		name string
		in   []byte
	}{
		{"Sm string", []byte(`"Hello World"`)},
		{"Md string", []byte(`"One morning, when Gregor Samsa woke from troubled dream"`)},
		{"Lg string", []byte(`"The quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymphs grab quick-jived waltz. Brick quiz whangs jumpy veldt fox. Bright vixens jump; dozy fowl quack. Quick wafting zephyrs vex bold Jim. Quick zephyrs blow, vexing daft Jim. Sex-charged fop blew my junk TV quiz. How quickly daft jumping zebras vex. Two driven jocks help fax my big quiz. Quick, Baz, get my woven flax jodhpurs! \"Now fax quiz Jack!\" my b"`)},
		{"Float string", []byte(`"31415926535.897932385"`)},
		{"Time String", []byte(`"2015-05-14T12:34:56.123+02:00"`)},
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
		b.Run(el.name+" Re7", func(b *testing.B) {
			n := false
			b.SetBytes(int64(len(el.in)))
			for i := 0; i < b.N; i++ {
				n = IsTimeStrRe7Fn(el.in)
			}
			bResult = n
		})

		b.Run(el.name+" Fn7", func(b *testing.B) {
			n := false
			b.SetBytes(int64(len(el.in)))
			for i := 0; i < b.N; i++ {
				n = IsTimeStr7Fn(el.in)
			}
			bResult = n
		})
		b.Run(el.name+" Fn", func(b *testing.B) {
			n := false
			b.SetBytes(int64(len(el.in)))
			for i := 0; i < b.N; i++ {
				n = IsTimeStrFn(el.in)
			}
			bResult = n
		})
	}
	tbl := []struct {
		in  []byte
		out bool
	}{
		{[]byte(`"2015-05-14T12:34:56.123+02:00"`), true},
		{[]byte(`"2015-05-14T12:34:56.123Z"`), true},
		{[]byte(`"1970-01-01T00:00:00Z"`), true},
		{[]byte(`"0001-01-01T00:00:00Z"`), true},
		{[]byte(`"1985-04-12T23:20:50.52Z"`), true},
		{[]byte(`"1996-12-19T16:39:57-08:00"`), true},
		{[]byte(`"1990-12-31T23:59:60Z"`), true},
		{[]byte(`"1990-12-31T15:59:60-08:00"`), true},
		{[]byte(`"1937-01-01T12:00:27.87+00:20"`), true},
		{[]byte(`"2022-02-24T04:00:00+02:00"`), true},
		{[]byte(`"2022-07-12T21:55:16+01:00"`), true},
		{[]byte(`"2015-05-14T12:34:56+02:00"`), true},
		{[]byte(`"2015-05-14T12:34:56.1-02:00"`), true},
		{[]byte(`"2015-05-14T12:34:56.12+02:00"`), true},
		{[]byte(`"2015-05-14T12:34:56.123-02:00"`), true},
		{[]byte(`"2015-05-14T12:34:56.1234+02:00"`), true},
		{[]byte(`"2015-05-14T12:34:56.12345-02:00"`), true},
		{[]byte(`"2015-05-14T12:34:56.123456+02:00"`), true},
		{[]byte(`"2015-05-14T12:34:56.1234567-02:00"`), true},
		{[]byte(`"2015-05-14T12:34:56.1Z"`), true},
		{[]byte(`"2015-05-14T12:34:56.12Z"`), true},
		{[]byte(`"2015-05-14T12:34:56.123Z"`), true},
		{[]byte(`"2015-05-14T12:34:56.1234Z"`), true},
		{[]byte(`"2015-05-14T12:34:56.12345Z"`), true},
		{[]byte(`"2015-05-14T12:34:56.123456Z"`), true},
		{[]byte(`"2015-05-14T12:34:56.1234567Z"`), true},
		// origin
		{[]byte(`"2016-01-19T15:21:32.59+02:00"`), true},
		{[]byte(`"2015-05-14T12:34:56+02:00"`), true},
		{[]byte(`"2015-05-14T12:34:56Z"`), true},
		{[]byte(`"1970-01-01T00:00:00Z"`), true},
		{[]byte(`"1970-01-01T00:00:00+00:00"`), true},
		{[]byte(`"0001-01-01T00:00:00Z"`), true},
		// invalid
		{[]byte(`"2015-05-14E12:34:56.379+02:00"`), false},
		{[]byte(`"2O15-O5-14T12:34:56.379+02:00"`), false},
		{[]byte(`"1985-04-12T23:20:50.52ZZZZ"`), false},
		{[]byte(`"2022-07-12 21:55:16"`), false},
		{[]byte(`"20220712T215516Z"`), false},
		{[]byte(`"20220712T215516+01:00"`), false},
		{[]byte(`"1985-04-12T23:20:50.Z"`), false},
		// origin
		{[]byte(`"not a Timestamps"`), false},
		{[]byte(`"2015+05-14T12:34:56.789+02:00"`), false},
		// extra
		{[]byte(`"Hello World""`), false},
		{[]byte(`"One morning, when Gregor Samsa woke from troubled dream""`), false},
		{[]byte(`"The quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymphs grab quick-jived waltz. Brick quiz whangs jumpy veldt fox. Bright vixens jump; dozy fowl quack. Quick wafting zephyrs vex bold Jim. Quick zephyrs blow, vexing daft Jim. Sex-charged fop blew my junk TV quiz. How quickly daft jumping zebras vex. Two driven jocks help fax my big quiz. Quick, Baz, get my woven flax jodhpurs! \"Now fax quiz Jack!\" my b"`), false},
		{[]byte(`"31415926535.897932385"`), false},
		{[]byte(`"2015-05-14T12:34:56.123+02:00"`), true},
	}
	b.Run("Many variants Re", func(b *testing.B) {
		n := false
		b.SetBytes(int64(1780))
		for i := 0; i < b.N; i++ {
			for _, el := range tbl {
				n = IsTimeStrReFn(el.in)
			}
		}
		bResult = n
	})
	b.Run("Many variants Re7", func(b *testing.B) {
		n := false
		b.SetBytes(int64(1780))
		for i := 0; i < b.N; i++ {
			for _, el := range tbl {
				n = IsTimeStrRe7Fn(el.in)
			}
		}
		bResult = n
	})
	b.Run("Many variants Fn7", func(b *testing.B) {
		n := false
		b.SetBytes(int64(1780))
		for i := 0; i < b.N; i++ {
			for _, el := range tbl {
				n = IsTimeStr7Fn(el.in)
			}
		}
		bResult = n
	})
	b.Run("Many variants Fn", func(b *testing.B) {
		n := false
		b.SetBytes(int64(1780))
		for i := 0; i < b.N; i++ {
			for _, el := range tbl {
				n = IsTimeStrFn(el.in)
			}
		}
		bResult = n
	})
}
