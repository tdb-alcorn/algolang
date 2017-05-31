package heap

import (
	"fmt"
)

// Heap is a min-heap.
type Heap struct {
	parent   *Heap
	children []*Heap
	value    int
}

func (h *Heap) siftUp() bool {
	if h.parent != nil && h.parent.value > h.value {
		newChildren, err := removeChild(h.parent.children, h)
		if err != nil {
			fmt.Println(err)
		}
		h.parent.children = newChildren
		h.children = addChild(h.children, h.parent)
		if h.parent.parent != nil {
			h.parent.parent.children = addChild(h.parent.parent.children, h)
		}
		// Finally, swap the nodes.
		h.parent, h.parent.parent = h.parent.parent, h
		return true
	}
	return false
}

func (h *Heap) height() int {
	var tallest int
	var height int
	for _, c := range h.children {
		if c == nil {
			fmt.Println("wtf %d", len(h.children))
		}
		height = c.height()
		if height > tallest {
			tallest = height
		}
	}
	return tallest + 1
}

func (h *Heap) Insert(value int) {
	// binary tree
	ln := h.findLastNode(2)
	n := &Heap{ln, make([]*Heap, 0), value}
	ln.children = addChild(ln.children, n)
	for n.siftUp() {
	}
}

// d is the branching factor (how many children each node may have)
func (h *Heap) findLastNode(d int) *Heap {
	if len(h.children) == 0 {
		return h
	}
	if len(h.children) < d {
		return h
	}
	shortestHeight := h.height()
	var shortest *Heap
	for _, c := range h.children {
		// SLOW!
		if height := c.height(); height < shortestHeight {
			shortest = c
			shortestHeight = height
		}
	}
	if shortest == nil {
		fmt.Println("wtf")
	}
	return shortest.findLastNode(d)
}

func removeChild(children []*Heap, child *Heap) ([]*Heap, error) {
	result := make([]*Heap, len(children)-1)
	for i, c := range children {
		if c == child {
			copy(children[:i], result)
			copy(children[i+1:], result[i:])
			return result, nil
		}
	}
	return children, fmt.Errorf("#removeChild: child was not found in children and so couldn't be removed")
}

func addChild(children []*Heap, child *Heap) []*Heap {
	children = append(children, child)
	return children
}

func (h *Heap) String() string {
	return fmt.Sprintf("(%d %v)", h.value, h.children)
}
