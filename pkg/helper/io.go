package helper

import (
	"fmt"
	"io/ioutil"
	"os"
)

// WriteStringToFile writes given string str to a file with the
// given path. Handles io errors.
func WriteStringToFile(str string, path string) {
	err := ioutil.WriteFile(path, []byte(str), 0744)
	if err != nil {
		fmt.Printf("could not write to %s, %s", path, err.Error())
	}
}

// ReadStringFromFile returns a string containing all the text in the
// file with the given path. An empty string is returned in case of an io err.
func ReadStringFromFile(path string) string {
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
