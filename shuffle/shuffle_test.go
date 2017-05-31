package shuffle

import (
	"testing"
)

func TestShuffle(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(Shuffle(intRange(10)))
	}
}

func BenchmarkShuffle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Shuffle(intRange(100))
	}
}

func intRange(n int) []int {
	r := make([]int, n)
	for i := range r {
		r[i] = i
	}
	return r
}
