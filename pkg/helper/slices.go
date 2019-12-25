// Helper functions used for validation of
// slices and their properties.
package helper

// Return true if a slice of size n
// contains numbers 0 through n - 1.
func AreElementsIndices(slice []int) bool {
	length := len(slice)
	checkArr := make([]bool, length)

	for _, element := range slice {
		if element >= 0 && element < length {
			checkArr[element] = true
		}
	}

	for _, element := range checkArr {
		if element == false {
			return false
		}
	}

	return true
}

// Return true if a slice is symmetric,
// symmetric means that each two elements are mapped to each other.
func IsSymmetric(slice []int) bool {
	length := len(slice)

	for i, element := range slice {
		if element < 0 || element > length-1 || slice[element] != i {
			return false
		}
	}

	return true
}
