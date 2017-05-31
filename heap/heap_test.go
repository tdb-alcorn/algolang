package heap

import (
	"testing"
)

func TestInsert(t *testing.T) {
	h := &Heap{nil, make([]*Heap, 0), 1}
	h.Insert(3)
	h.Insert(5)
	h.Insert(6)
	for i := 0; i < 10; i++ {
		h.Insert(i)
	}
	t.Log(h)
}
