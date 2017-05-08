package search

import (
	"testing"
)

func TestBinary(t *testing.T) {
	for i := 0; i < 1000; i += 11 {
		testRange(t, i)
	}
}

func testRange(t testing.TB, n int) {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	for key := 0; key < n; key++ {
		idx := Binary(arr, key)
		if idx != key {
			t.Errorf("Index should be %d but was %d on search with []int of length %d", key, idx, len(arr))
		}
	}
}

func BenchmarkBinary2(b *testing.B) {
	var size uint = 2
	benchmarkBinaryN(b, size)
}

func BenchmarkBinary4(b *testing.B) {
	var size uint = 4
	benchmarkBinaryN(b, size)
}

func BenchmarkBinary8(b *testing.B) {
	var size uint = 8
	benchmarkBinaryN(b, size)
}

func BenchmarkBinary16(b *testing.B) {
	var size uint = 16
	benchmarkBinaryN(b, size)
}

func benchmarkBinaryN(b *testing.B, size uint) {
	n := 1 << size
	key := 0
	for i := uint(0); i < size-1; i++ {
		key += 1 << i
	}
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Binary(arr, key)
	}
}
