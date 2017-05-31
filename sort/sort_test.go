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

func TestSortToIndexRandom10(t *testing.T) {
	testSortToIndexRandom(t, 100, 10)
}

func TestSortToIndexRandom100(t *testing.T) {
	testSortToIndexRandom(t, 10, 100)
}

func TestSortToIndexRandom1000(t *testing.T) {
	testSortToIndexRandom(t, 1, 1000)
}

func testSortToIndexRandom(t *testing.T, numTests int, size int) {
	sorted := make([]int, size)
	for i := range sorted {
		sorted[i] = i
	}
	for i := 0; i < numTests; i++ {
		shuffled := sort.IntSlice(rand.Perm(size))
		SortToIndex(shuffled)
		if !sliceEquals(sorted, shuffled) {
			t.Errorf("%s should be %s", shuffled, sorted)
		}
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
