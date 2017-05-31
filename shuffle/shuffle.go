package shuffle

import (
	"crypto/rand"
	"math/big"

	"algolang/sort"
)

func Shuffle(arr []int) []int {
	rnd := make([]int64, len(arr))
	for i := 0; i < len(arr); i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(int64(1)<<62))
		rnd[i] = r.Int64()
	}
	indices := sort.SortToIndex(&int64Swapper{rnd})
	shuffled := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		shuffled[i] = arr[indices[i]]
	}
	return shuffled
}

type int64Swapper struct {
	array []int64
}

func (i *int64Swapper) Len() int {
	return len(i.array)
}

func (i *int64Swapper) Less(a, b int) bool {
	if i.array[a] < i.array[b] {
		return true
	}
	return false
}

func (i *int64Swapper) Swap(a, b int) {
	i.array[a], i.array[b] = i.array[b], i.array[a]
}
