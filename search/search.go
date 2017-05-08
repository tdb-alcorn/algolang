package search

/*
 * Takes a sorted (ascending) list of integers I and an integer to search for
 * N, and returns the index at which N would be inserted into I.
 */
func Binary(arr []int, key int) int {
	switch len(arr) {
	case 0:
		return 0
	case 1:
		if key > arr[0] {
			return 1
		}
		return 0
	default:
		halfway := len(arr) / 2
		if key > arr[halfway] {
			return halfway + Binary(arr[halfway:], key)
		}
		return Binary(arr[:halfway], key)
	}
}
