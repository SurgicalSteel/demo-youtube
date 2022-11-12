package unit_test_basic

// FilterOdd returns only odd elements from a slice
func FilterOdd(slice []int) []int {
	newSlice := make([]int, 0)
	for _, i := range slice {
		if i%2 == 1 {
			newSlice = append(newSlice, i)
		}
	}
	return newSlice
}
