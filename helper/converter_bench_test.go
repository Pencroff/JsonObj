package helper

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

var result int64
var uresult uint64

func BenchmarkFloatToIntConverter(b *testing.B) {
	b.Run("math round", func(b *testing.B) {
		n := int64(0)
		for i := 0; i < b.N; i++ {
			n = int64(math.Round(1.99))
			n += int64(math.Round(-1.99))
		}
		result = n
	})
	b.Run("math floor", func(b *testing.B) {
		n := int64(0)
		for i := 0; i < b.N; i++ {
			n = int64(math.Floor(1.99))
			n += int64(math.Floor(-1.99))
		}
		result = n
	})
	b.Run("custom round", func(b *testing.B) {
		n := int64(0)
		for i := 0; i < b.N; i++ {
			n = FloatToInt(1.99)
			n += FloatToInt(-1.99)
		}
		result = n
	})
	b.Run("type cast", func(b *testing.B) {
		n := int64(0)
		for i := 0; i < b.N; i++ {
			a := 1.99
			b := -1.99
			n = int64(a)
			n += int64(b)
		}
		result = n
	})
}

/*
λ go test -bench=BenchmarkStringToIntConverter
goos: windows
goarch: amd64
pkg: github.com/Pencroff/JsonStruct/helper
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkStringToIntConverter/StringToUint-12           63075232                18.42 ns/op
BenchmarkStringToIntConverter/ParseUint-12              41217283                33.59 ns/op
BenchmarkStringToIntConverter/ParseUint(0)-12           40003066                29.64 ns/op

BenchmarkStringToIntConverter/StringToInt-12            79572430                17.95 ns/op
BenchmarkStringToIntConverter/strconv.Atoi-12           34287967                38.27 ns/op
BenchmarkStringToIntConverter/ParseInt-12               34227234                33.97 ns/op
BenchmarkStringToIntConverter/ParseInt(0)-12            34930328                31.58 ns/op
BenchmarkStringToIntConverter/fmt.Sscan-12               1381633               861.6 ns/op
PASS
ok      github.com/Pencroff/JsonStruct/helper   11.256s
*/

func BenchmarkStringToIntConverter(b *testing.B) {
	str := "-9223372036854775808"
	ustr := "18446744073709551615"
	b.Run("StringToUint", func(b *testing.B) {
		n := uint64(0)
		for i := 0; i < b.N; i++ {
			n, _ = StringToUint(ustr)
		}
		uresult = n
	})
	b.Run("ParseUint", func(b *testing.B) {
		n := uint64(0)
		for i := 0; i < b.N; i++ {
			n, _ = strconv.ParseUint(ustr, 10, 64)
		}
		uresult = n
	})
	b.Run("ParseUint(0)", func(b *testing.B) {
		n := uint64(0)
		for i := 0; i < b.N; i++ {
			n, _ = strconv.ParseUint(ustr, 10, 0)
		}
		uresult = n
	})
	fmt.Println("")
	b.Run("StringToInt", func(b *testing.B) {
		n := int64(0)
		for i := 0; i < b.N; i++ {
			n, _ = StringToInt(str)
		}
		result = n
	})

	b.Run("strconv.Atoi", func(b *testing.B) {
		n := 0
		for i := 0; i < b.N; i++ {
			n, _ = strconv.Atoi(str)
		}
		result = int64(n)
	})
	b.Run("ParseInt", func(b *testing.B) {
		n := int64(0)
		for i := 0; i < b.N; i++ {
			n, _ = strconv.ParseInt(str, 10, 64)
		}
		result = n
	})
	b.Run("ParseInt(0)", func(b *testing.B) {
		n := int64(0)
		for i := 0; i < b.N; i++ {
			n, _ = strconv.ParseInt(str, 10, 0)
		}
		result = n
	})
	// Takes  860 - 890 ns/op, compare around 30 - 40 ns/op for rest of the benchmarks
	b.Run("fmt.Sscan", func(b *testing.B) {
		n := 0
		for i := 0; i < b.N; i++ {
			fmt.Sscan(str, &n)
		}
		result = int64(n)
	})

}
