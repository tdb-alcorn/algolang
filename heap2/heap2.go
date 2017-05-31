package heap2

import (
	"errors"
)

type Heap []float64

var EmptyHeapError = errors.New("Heap is empty")

func Heapify(f []float64) Heap {
	h := Heap(f)
	for i := (len(h) >> 1); i >= 0; i-- {
		siftDown(h, i)
	}
	return h
}

func Insert(h *Heap, value float64) {
	*h = append(*h, value)
	siftUp(*h, len(*h)-1)
}

func Peak(h Heap) (float64, error) {
	if empty(h) {
		return 0, EmptyHeapError
	}
	return h[0], nil
}

func Pop(h *Heap) (float64, error) {
	if empty(*h) {
		return 0, EmptyHeapError
	}
	last := len(*h) - 1
	root, err := Replace(*h, (*h)[last])
	if err != nil {
		return 0, err
	}
	*h = (*h)[:last]
	return root, nil
}

func Replace(h Heap, value float64) (float64, error) {
	if empty(h) {
		return 0, EmptyHeapError
	}
	root := h[0]
	h[0] = value
	siftDown(h, 0)
	return root, nil
}

func empty(h Heap) bool {
	if len(h) == 0 {
		return true
	}
	return false
}

func parent(i int) int {
	return i >> 1
}

func left(i int) int {
	return i << 1
}

func right(i int) int {
	return i<<1 + 1
}

// heap is a min-heap
func siftDown(h Heap, i int) {
	smallest := i
	l := left(i)
	r := right(i)
	if l < len(h) && h[l] < h[smallest] {
		smallest = l
	}
	if r < len(h) && h[r] < h[smallest] {
		smallest = r
	}
	if smallest != i {
		h[i], h[smallest] = h[smallest], h[i]
		siftDown(h, smallest)
	}
}

func siftUp(h Heap, i int) {
	p := parent(i)
	if h[i] < h[p] {
		h[i], h[p] = h[p], h[i]
		siftUp(h, p)
	}
}

// minheap
func isHeap(h Heap) bool {
	for i, _ := range h {
		if h[parent(i)] > h[i] {
			return false
		}
	}
	return true
}
