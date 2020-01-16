// Package encrypt contains components of the enigma machine.
// Used for encryption and decryption of messages.
package encrypt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
)

// jsonMachine is a struct used for reading and writing of Machine's configs
// into a json file. Fields in jsonMachine mirror those in a Machine but use
// string arrays instead of int arrays.
type jsonMachine struct {
	PathConnections      [numberOfRotors][alphabetSize]string `json:"pathways"`
	Reflector            [alphabetSize / 2]string             `json:"reflector"`
	PlugboardConnections [alphabetSize / 2]string             `json:"plugboards"`
	RotorsPositions      [numberOfRotors]string               `json:"rotorsPositions"`
}

// read loads and verifies Machine's configurations from a json file.
// It returns a pointer to a Machine, and an error in case of incorrect loading.
func read(path string) (*Machine, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not load config file %s", path)
	}
	defer file.Close()

	// Read file's contents into a byte array
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("could not read contents of config file %s", path)
	}

	m, err := parseMachineJSON(fileContents)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal %s: %s", path, err.Error())
	}

	// Verify correct initialization
	if !m.isInit() {
		return nil, fmt.Errorf("configuration values are incorrect")
	}

	return m, nil
}

// parseMachineJSON parses a given file's byte array into character arrays,
// Creates a Machine object, Sets Machine's fields, and returns a pointer.
// An error is returned in case of invalid configs.
func parseMachineJSON(fileContents []byte) (*Machine, error) {
	var jsonM jsonMachine

	if err := json.Unmarshal(fileContents, &jsonM); err != nil {
		return nil, err
	}

	// Parse jsonM into a Machine
	var m *Machine

	// Electric pathways
	for i := 0; i < numberOfRotors; i++ {
		for j := 0; j < alphabetSize; j++ {
			if num, verify := strToInt(jsonM.PathConnections[i][j]); verify {
				m.pathConnections[i][j] = num
			} else {
				return nil, fmt.Errorf("pathways contain invalid value %v", jsonM.PathConnections[i][j])
			}
		}
	}

	for i := 0; i < alphabetSize/2; i++ {
		// Plugboard
		if num, verify := strToInt(jsonM.PlugboardConnections[i]); verify {
			m.plugboardConnections[i] = num
			m.plugboardConnections[num] = i
		} else {
			return nil, fmt.Errorf("plugboard contains invalid value %v", jsonM.PlugboardConnections[i])
		}

		// Reflector
		if num, verify := strToInt(jsonM.Reflector[i]); verify {
			m.reflector[i] = num
			m.reflector[num] = i
		} else {
			return nil, fmt.Errorf("plugboard contains invalid value %v", jsonM.Reflector[i])
		}
	}

	// Rotors
	var rotorsPositions [numberOfRotors]int
	for i := 0; i < numberOfRotors; i++ {
		if num, verify := strToInt(jsonM.RotorsPositions[i]); verify {
			rotorsPositions[i] = num
		} else {
			return nil, fmt.Errorf("rotorsPositions contains invalid value %v", jsonM.RotorsPositions[i])
		}
	}

	m.initRotors(rotorsPositions, 1, alphabetSize)

	return m, nil
}

// strToInt verifies that a given string contains one alphabetical
// character and returns character's position in the alphabet.
func strToInt(str string) (int, bool) {
	if len(str) != 1 {
		return -1, false
	}

	if unicode.IsLetter(rune(str[0])) {
		return int(byte(unicode.ToLower(rune(str[0]))) - 'a'), true
	}

	return -1, false
}
