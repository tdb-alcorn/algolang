package shuffle

import (
	"crypto/rand"
	"math/big"
)

func Shuffle(arr []int) []int {
	rnd := make([]int64, len(arr))
	for i := 0; i < len(arr); i++ {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(1)<<62))
		rnd[i] = r.Int64()
	}
	indices := sort.SortToIndex(rnd)
	shuffled := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		shuffled[i] = arr[indices[i]]
	}
	return shuffled
}
