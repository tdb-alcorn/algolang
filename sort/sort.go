package sort

// SortToIndex sorts arr in place and returns an array of indices representing
// the sorted order. It uses quicksort under the hood.
func SortToIndex(arr Swapper) []int {
	monitorArray := make([]int, arr.Len())
	for i := range monitorArray {
		monitorArray[i] = i
	}
	monitor := &ints{monitorArray}
	quicksort(arr, 0, arr.Len(), monitor)
	return monitor.array
}

// quicksort sorts in place, replicating all swaps made to the monitor.
//
// Note: strictly speaking, monitor only needs to implement Swap(i,j int). This
// could be represented by defining and interface Swapper.
func quicksort(arr Swapper, min int, max int, monitor Swapper) {
	pivotIdx := max - 1
	i := min
	if max-min <= 1 {
		return
	}
	for ; i < pivotIdx; i++ {
		if arr.Less(pivotIdx, i) {
			// Swap with something further right.
			swapFailed := true
			// j starts at pivotIdx-1 to avoid swapping the pivot prematurely.
			for j := pivotIdx - 1; j > i; j-- {
				if arr.Less(j, pivotIdx) {
					arr.Swap(i, j)
					monitor.Swap(i, j)
					swapFailed = false
					break
				}
			}
			if swapFailed {
				// If no suitable element can be found to swap then everything from i on is
				// larger than the pivot, and therefore the pivot belongs at i and we are
				// done with this level of quicksort.
				arr.Swap(i, pivotIdx)
				monitor.Swap(i, pivotIdx)
				break
			}
		}
	}
	quicksort(arr, min, i, monitor)
	quicksort(arr, i, max, monitor)

}

type Swapper interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

type ints struct {
	array []int
}

func (t *ints) Len() int {
	return len(t.array)
}

// Swap mutates the array in place.
func (t *ints) Swap(i, j int) {
	t.array[i], t.array[j] = t.array[j], t.array[i]
}

func (t *ints) Less(i, j int) bool {
	return t.array[i] < t.array[j]
}
