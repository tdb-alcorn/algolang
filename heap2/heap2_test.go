package heap2

import (
	"testing"
)

func TestNewHeap(t *testing.T) {
	heap := Heapify([]float64{3, 1, 2, 3, 0, 5})
	t.Log(heap)
	if !isHeap(heap) {
		t.Errorf("Not a heap: %#v", heap)
	}
}

func TestInsert(t *testing.T) {
	heap := Heapify([]float64{5, 4, 3, 2, 1})
	t.Log(heap)
	Insert(&heap, 2)
	t.Log(heap)
	if !isHeap(heap) {
		t.Errorf("Not a heap: %#v", heap)
	}
}

func TestPop(t *testing.T) {
	heap := Heapify([]float64{3, 6, 8, 9, 2})
	t.Log(heap)
	v, err := Pop(&heap)
	if err != nil {
		t.Error(err)
	}
	t.Log(v)
	if !isHeap(heap) {
		t.Errorf("Not a heap: %#v", heap)
	}
	if v != 2 {
		t.Errorf("Did not pop root, popped %f %#v", v, heap)
	}
}

func TestEmpty(t *testing.T) {
	heap := Heapify([]float64{})
	_, err := Pop(&heap)
	if err != EmptyHeapError {
		t.Errorf("Pop failed to return EmptyHeapError, instead got %#v", err)
	}
}
