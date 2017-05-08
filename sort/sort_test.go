package sort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSortToIndexNoop(t *testing.T) {
	i := []int{0, 1, 2, 3, 4}
	order := SortToIndex(sort.IntSlice(i))
	if !sliceEquals(order, i) {
		t.Errorf("%s should be %s", order, i)
	}
}

func TestSortToIndexReverse(t *testing.T) {
	i := []int{0, 1, 2, 3, 4}
	j := make([]int, len(i))
	copy(j, i)
	reverse(i)
	SortToIndex(sort.IntSlice(i))
	if !sliceEquals(i, j) {
		t.Errorf("%s should be %s", i, j)
	}
}

func TestSortToIndexJumble(t *testing.T) {
	i := []int{0, 1, 2, 3, 4}
	j := []int{2, 1, 4, 3, 0}
	order := SortToIndex(sort.IntSlice(j))
	t.Log(order)
	if !sliceEquals(i, j) {
		t.Errorf("%s should be %s", order, i)
	}
}

func BenchmarkSortToIndexRandom1(b *testing.B) {
	benchmarkSortToIndexRandom(b, 1)
}

func BenchmarkSortToIndexRandom2(b *testing.B) {
	benchmarkSortToIndexRandom(b, 2)
}

func BenchmarkSortToIndexRandom4(b *testing.B) {
	benchmarkSortToIndexRandom(b, 4)
}

func BenchmarkSortToIndexRandom8(b *testing.B) {
	benchmarkSortToIndexRandom(b, 8)
}

func BenchmarkSortToIndexRandom16(b *testing.B) {
	benchmarkSortToIndexRandom(b, 16)
}

func benchmarkSortToIndexRandom(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffled := sort.IntSlice(rand.Perm(n))
		b.StartTimer()
		SortToIndex(shuffled)
	}
}

func sliceEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// reverses int slice in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
