package comparison

import "testing"

var res2 []int

func BenchmarkSliceAppendSet(b *testing.B) {
	b.Run("append", func(b *testing.B) {
		a := make([]int, 0, b.N)
		for i := 0; i < b.N; i++ {
			a = append(a, i)
		}
		res2 = a
	})
	b.Run("set", func(b *testing.B) {
		a := make([]int, b.N)
		for i := 0; i < b.N; i++ {
			a[i] = i
		}
		res2 = a
	})
}
