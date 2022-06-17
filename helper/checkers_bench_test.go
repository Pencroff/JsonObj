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
		v    string
	}{
		{"Sm string", "Hello World"},
		{"Md string", "One morning, when Gregor Samsa woke from troubled dream"},
		{"Lg string", "The quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymphs grab quick-jived waltz. Brick quiz whangs jumpy veldt fox. Bright vixens jump; dozy fowl quack. Quick wafting zephyrs vex bold Jim. Quick zephyrs blow, vexing daft Jim. Sex-charged fop blew my junk TV quiz. How quickly daft jumping zebras vex. Two driven jocks help fax my big quiz. Quick, Baz, get my woven flax jodhpurs! \"Now fax quiz Jack!\" my b"},
		{"Float string", "31415926535.897932385"},
		{"Time String", "2015-05-14T12:34:56+02:00"},
	}
	for _, el := range tblMethod {
		b.Run(el.name, func(b *testing.B) {
			n := false
			for i := 0; i < b.N; i++ {
				n = IsTimeRe.MatchString(el.v)
			}
			bResult = n
		})
	}
}
