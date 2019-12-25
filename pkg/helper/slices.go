// Helper functions used for validation of
// slices and their properties.
package helper

// Return true if a slice of size n
// contains numbers 0 through n - 1.
func AreElementsIndices(slice []int) bool {
	length := len(slice)

	if slice == nil || length == 0 {
		return false
	}

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

	if slice == nil || length == 0 {
		return false
	}

	for i, element := range slice {
		if element < 0 || element > length-1 || slice[element] != i {
			return false
		}
	}

	return true
}

// Return true if a slice of size n
// contains elements 0 through n - 1
// starting with any element and in order.
func AreElementsOrderedIndices(slice []int) bool {
	length := len(slice)

	if slice == nil || length == 0 {
		return false
	}

	iterator := slice[0]

	for i, element := range slice {
		if element != (iterator+i)%length {
			return false
		}
	}

	return true
}
