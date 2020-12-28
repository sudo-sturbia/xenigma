package machine

import (
	"fmt"
	"io/ioutil"
	"os"
)

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

// writeStringToFile writes given string str to a file with the
// given path. Handles io errors.
func writeStringToFile(str string, path string) {
	err := ioutil.WriteFile(path, []byte(str), 0744)
	if err != nil {
		fmt.Printf("could not write to %s, %s", path, err.Error())
	}
}

// readStringFromFile returns a string containing all the text in the
// file with the given path. An empty string is returned in case of an io err.
func readStringFromFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("could not open %s, %s", path, err.Error())
		return ""
	}
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("could not read contents of %s, %s", path, err.Error())
		return ""
	}

	return string(contents)
}
