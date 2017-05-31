package filter

import (
	"crypto/rand"
)

func hash() ([]byte, int, error) {
	b := make([]byte, 32)
	n, err := rand.Read(b)
	return b, n, err
}

// Assuming k hash functions and n items to insert, compute the optimal size of
// the filter. The magic number 1.4427 is a hardcoded approximation for 1/ln(2).
//
// Reference: https://en.wikipedia.org/wiki/Bloom_filter
func optimalSize(k int, n int) (m int) {
	return int(1.4427 * float64(k) * float64(n))
}

func optimalNumHashFunctions(m int, n int) (k int) {
	return int(float64(m)/float64(n)*0.69315) + 1
}

type Hasher interface {
	// Hash returns an array of k indices into the Bloom filter.
	Hash(k int) ([]int, error)
}

type StringHasher string

func (s StringHasher) Hash(k int) ([]int, error) {

}

type Bloom struct {
	// k is the number of hashes to be computed for each check
	k        int
	bitArray []bool
}

func NewBloomFilter(numHashFunctions int, numItemsIntended int) *Bloom {
	size := optimalSize(numHashFunctions, numItemsIntended)
	b := make([]bool, size)
	return &Bloom{numHashFunctions, b}
}

func (b *Bloom) Insert(h Hasher) error {
	idx, err := h.Hash(b.k)
	if err != nil {
		return err
	}
	m := b.Size()
	for _, v := range idx {
		b.bitArray[v%m] = true
	}
	return nil
}

func (b *Bloom) Check(h Hasher) bool {
	idx, err := h.Hash(b.k)
	if err != nil {
		return false
	}
	m := b.Size()
	contains := true
	for _, v := range idx {
		contains = contains && b.bitArray[v%m]
	}
	return contains
}

func (b *Bloom) Size() int {
	return len(b.bitArray)
}
