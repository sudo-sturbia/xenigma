package machine

// zeroToNSlice returns true if a slice of size n contains numbers
// 0 through n - 1.
func zeroToNSlice(slice []int) bool {
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

// zeroToN returns true if mappings has size n, and contains keys
// 0 to n-1, mapped to values 0 to n-1.
func zeroToN(mappings map[int]int, n int) bool {
	if len(mappings) != n {
		return false
	}

	var check [2][]bool
	for i := 0; i < 2; i++ {
		check[i] = make([]bool, n)
	}

	for k, v := range mappings {
		if k >= 0 && k < n {
			check[0][k] = true
		}
		if v >= 0 && v < n {
			check[1][v] = true
		}
	}

	for i := 0; i < n; i++ {
		if !check[0][i] || !check[1][i] {
			return false
		}
	}
	return true
}

// isSymmetric returns true if mappings are symmetric. Symmetric means that for
// any k and v, if map[k]=v, then map[v]=k.
func isSymmetric(mappings map[int]int) bool {
	for k, v := range mappings {
		if mappings[k] != v || mappings[v] != k {
			return false
		}
	}
	return true
}
