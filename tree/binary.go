package tree

import (
	"fmt"
)

type Binary struct {
	Parent *Binary
	Right  *Binary
	Left   *Binary
	Value  int
}

func (t *Binary) Insert(v int) {
	if v < t.Value {
		if t.Left == nil {
			t.Left = &Binary{t, nil, nil, v}
			return
		}
		t.Left.Insert(v)
		return
	}
	if t.Right == nil {
		t.Right = &Binary{t, nil, nil, v}
		return
	}
	t.Right.Insert(v)
}

func (t *Binary) Height() (left int, right int) {
	if t.Left != nil {
		ll, lr := t.Left.Height()
		left = max(ll, lr) + 1
	}
	if t.Right != nil {
		rl, rr := t.Right.Height()
		right = max(rl, rr) + 1
	}
	return
}

func (t *Binary) String() string {
	return fmt.Sprintf("(%d %s %s)", t.Value, t.Left, t.Right)
}

// Full returns the number of children the node has.
func (t *Binary) NumChildren() (n int) {
	if t.Left != nil {
		n++
	}
	if t.Right != nil {
		n++
	}
	return
}

func (t *Binary) anyChild() *Binary {
	if t.Left != nil {
		return t.Left
	}
	return t.Right
}

func (t *Binary) Balance() *Binary {
	list := t.ToSortedList()
	return NewBinaryFromList(list)
}

func NewBinaryFromList(list []int) *Binary {
	if len(list) == 0 {
		return nil
	}
	mid := len(list) / 2
	b := &Binary{nil, nil, nil, list[mid]}
	b.Left = NewBinaryFromList(list[:mid])
	if b.Left != nil {
		b.Left.Parent = b
	}
	b.Right = NewBinaryFromList(list[mid+1:])
	if b.Right != nil {
		b.Right.Parent = b
	}
	return b
}

func (t *Binary) ToSortedList() []int {
	lh, rh := t.Height()
	list := make([]int, 0, 1<<uint(max(lh, rh)))
	list = t.MarshalList(list)
	return list
}

// Marshals the values of the tree into a list. The list passed in must
func (t *Binary) MarshalList(list []int) []int {
	if t.Left != nil {
		list = t.Left.MarshalList(list)
	}
	list = append(list, t.Value)
	if t.Right != nil {
		list = t.Right.MarshalList(list)
	}
	return list
}

func (t *Binary) Trickle() *Binary {
	if c := t.anyChild(); t.NumChildren() == 1 {
		if c.NumChildren() == 2 {
			replace(t, c)
			return c.Trickle()
		}
		if c2 := c.anyChild(); c.NumChildren() == 1 {
			if c2.Value > c.Value {
				if c.Value > t.Value {
					replace(t, c)
					return c.Trickle()
				} else {
					replace(c, c2)
					replace(t, c2)
					return c2.Trickle()
				}
			} else {
				if c.Value > t.Value {
					replace(c, c2)
					replace(t, c2)
					return c2.Trickle()
				} else {
					replace(t, c)
					return c.Trickle()
				}
			}
		}
	}
	if t.NumChildren() == 2 {
		t.Left = t.Left.Trickle()
		t.Right = t.Right.Trickle()
	}
	return t
}

// Replaces t in the tree with c.
func replace(t *Binary, c *Binary) {
	if t.isLeftChild() {
		t.Parent.Left = c
	} else if t.Parent != nil {
		t.Parent.Right = c
	}
	c.Parent = t.Parent
	c.Insert(t.Value)
}

func (t *Binary) isLeftChild() bool {
	if t.Parent != nil && t.Parent.Left == t {
		return true
	}
	return false
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
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
