package tree

import (
	"math/rand"
	"testing"
)

func TestInsert(t *testing.T) {
	b := randomBinary(10)
	t.Log(b)
}

func TestHeight(t *testing.T) {
	b := randomBinary(10)
	t.Log(b.Height())
}

func TestTrickle(t *testing.T) {
	for i := 0; i < 50; i++ {
		b := randomBinary(1000)
		bl, br := b.Height()
		// t.Log(b)
		b = b.Trickle()
		al, ar := b.Height()
		t.Log(max(bl, br)-max(al, ar), max(bl, br))
		// t.Log(b)
	}
}

func TestToSortedList(t *testing.T) {
	n := 10
	b := randomBinary(n)
	l := b.ToSortedList()
	s := make([]int, 10)
	for i := range s {
		s[i] = i
	}
	if !sliceEquals(s, l) {
		t.Errorf("%v should be %v", l, s)
	}
}

func TestBalance(t *testing.T) {
	n := 100
	for i := 0; i < 50; i++ {
		b := randomBinary(n)
		// bl, br := b.Height()
		// t.Log(b)
		b = b.Balance()
		al, ar := b.Height()
		// t.Log(max(bl, br)-max(al, ar), max(bl, br), max(al, ar))
		// t.Log(b)
		if max(al, ar) > 6 {
			t.Errorf("Binary tree with %d elements should be at most %d height when balanced", n, 6)
		}
	}
}

func randomBinary(n int) *Binary {
	b := &Binary{nil, nil, nil, n - 1}
	ints := rand.Perm(n - 1)
	for _, v := range ints {
		b.Insert(v)
	}
	return b
}
