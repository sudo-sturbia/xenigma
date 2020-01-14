// Package encrypt contains components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
)

// loadTo is a struct to load Json configs to.
// Fields mirror those in a Machine but use bytes.
type loadTo struct {
	pathConnections      [numberOfRotors][alphabetSize]byte `json:"paths"`
	reflector            [alphabetSize]byte                 `json:"reflector"`
	plugboardConnections [alphabetSize]byte                 `json:"plugboards"`
	rotorsPositions      [numberOfRotors]byte               `json:"rotorsPositions"`
}

// load loads engima configurations from a json file.
// Returns a Machine object and an error in case of incorrect loading.
func load(path string) (*Machine, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("could not load configuration file %s", path)
	}
	defer file.Close()

	// Read from file
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("could not read contents of config file %s", path)
	}

	// Load object
	var loadedMachine loadTo
	if err := json.Unmarshal(fileContents, &loadedMachine); err != nil {
		return nil, errors.New("could not load machine. configurations are incorrect")
	}

	// Create Machine components
	var paths [numberOfRotors][alphabetSize]int
	var reflector [alphabetSize]int
	var plugboard [alphabetSize]int
	var rotorsPositions [numberOfRotors]int

	for i := 0; i < alphabetSize; i++ {
		// Electric pathways
		for j := 0; j < numberOfRotors; j++ {
			if char, isAlpha := byteToInt(loadedMachine.pathConnections[j][i]); isAlpha {
				paths[j][i] = char
			} else {
				// Return error
				return nil, fmt.Errorf(
					"paths contain invalid character %v, all characters must be alphabetical",
					loadedMachine.pathConnections[j][i],
				)
			}
		}

		// Reflector
		if char, isAlpha := byteToInt(loadedMachine.reflector[i]); isAlpha {
			reflector[i] = char
		} else {
			return nil, fmt.Errorf(
				"reflector contains invalid character %v, all characters must be alphabetical",
				loadedMachine.reflector[i],
			)
		}

		// Plugboard
		if char, isAlpha := byteToInt(loadedMachine.plugboardConnections[i]); isAlpha {
			plugboard[i] = char
		} else {
			return nil, fmt.Errorf(
				"plugboard contains invalid character %v, all characters must be alphabetical",
				loadedMachine.plugboardConnections[i],
			)
		}
	}

	for i := 0; i < numberOfRotors; i++ {
		if char, isAlpha := byteToInt(loadedMachine.rotorsPositions[i]); isAlpha {
			rotorsPositions[i] = char
		} else {
			return nil, fmt.Errorf(
				"rotorsPositions contains invalid character %v, all characters must be alphabetical",
				loadedMachine.rotorsPositions[i],
			)
		}
	}

	// Create Machine object
	var machine *Machine

	machine.SetPathConnections(paths)
	machine.SetReflector(reflector)
	machine.SetPlugboard(plugboard)
	machine.InitRotors(rotorsPositions, 1, alphabetSize)

	// Check values
	if !machine.isInit() {
		return nil, fmt.Errorf("values contained in configuration file are incorrect")
	}

	return machine, nil
}

// byteToInt takes a character char.
// If char is alphabetical, it returns true and an int indicating
// the position of char in the alphabet, else false and -1 are returned.
func byteToInt(char byte) (int, bool) {
	if unicode.IsLetter(rune(char)) {
		return int(byte(unicode.ToLower(rune(char))) - 'a'), true
	}

	return -1, false
}
