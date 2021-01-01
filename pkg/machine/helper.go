package machine

// areElementsIndices returns true if a slice of size n contains numbers
// 0 through n - 1.
func areElementsIndices(slice []int) bool {
	check := make([]bool, len(slice))
	for _, value := range slice {
		if value >= 0 && value < len(slice) {
			check[value] = true
		}
	}

	for _, value := range check {
		if !value {
			return false
		}
	}
	return true
}

// isSymmetric returns true if a slice is symmetric. Symmetric means that if
// slice[n] = m, then slice[m] = n. An empty slice is considered symmetric.
func isSymmetric(slice []int) bool {
	min, max := 0, len(slice)-1
	for i, value := range slice {
		if value < min || value > max || slice[value] != i {
			return false
		}
	}
	return true
}

// AreElementsOrderedIndices returns true if a slice of size n contains elements
// 0 through n-1 in circular order.
func areElementsOrderedIndices(slice []int) bool {
	start, length := slice[0], len(slice)
	for i, value := range slice {
		if value != (start+i)%length {
			return false
		}
	}
	return true
}
