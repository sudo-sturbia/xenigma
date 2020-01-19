package machine

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
	Reflector            [alphabetSize]string                 `json:"reflector"`
	PlugboardConnections [alphabetSize]string                 `json:"plugboards"`
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
	if err := m.isInit(); err != nil {
		return nil, err
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
		for j, connection := range jsonM.PathConnections[i] {
			if num, verify := strToInt(connection); verify {
				m.pathConnections[i][j] = num
			} else {
				return nil, &initError{fmt.Sprintf("pathways contain invalid value %v",
					connection)}
			}

		}
	}

	// Plugboard
	for i, connection := range jsonM.PlugboardConnections {
		if num, verify := strToInt(connection); verify {
			m.plugboardConnections[i] = num
		} else {
			return nil, &initError{fmt.Sprintf("plugboard contains invalid value %v",
				connection)}
		}
	}

	// Reflector
	for i, connection := range jsonM.Reflector {
		if num, verify := strToInt(connection); verify {
			m.reflector[i] = num
		} else {
			return nil, &initError{fmt.Sprintf("plugboard contains invalid value %v",
				connection)}
		}
	}

	// Rotors
	var rotorsPositions [numberOfRotors]int
	for i, position := range jsonM.RotorsPositions {
		if num, verify := strToInt(position); verify {
			rotorsPositions[i] = num
		} else {
			return nil, &initError{fmt.Sprintf("rotorsPositions contains invalid value %v",
				position)}
		}

	}

	m.initRotors(rotorsPositions, 1, alphabetSize)

	return m, nil
}

// write writes configurations of a Machine object to a JSON file.
// returns an error in case of incorrect writing.
func write(m *Machine, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not write to %s, %s", path, err.Error())
	}
	defer file.Close()

	var jsonM jsonMachine

	// Electric pathways
	for i := 0; i < numberOfRotors; i++ {
		for j := 0; j < alphabetSize; j++ {
			jsonM.PathConnections[i][j] = intToStr(m.pathConnections[i][j])
		}
	}

	for i := 0; i < alphabetSize; i++ {
		// Plugboard
		jsonM.PlugboardConnections[i] = intToStr(m.plugboardConnections[i])

		// Reflector
		jsonM.Reflector[i] = intToStr(m.reflector[i])
	}

	// Rotors
	for i := 0; i < numberOfRotors; i++ {
		jsonM.RotorsPositions[i] = intToStr(m.CurrentRotors()[i])
	}

	contents, err := json.Marshal(jsonM)
	if err != nil {
		return fmt.Errorf("could not create JSON file, %s", err.Error())
	}

	err = ioutil.WriteFile(path, contents, 0744)
	if err != nil {
		return fmt.Errorf("could not create JSON file, %s", err.Error())
	}

	return nil
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

// intToStr returns a one character string representing the ASCII position
// of the given integer.
func intToStr(num int) string {
	return fmt.Sprintf("%v", byte(num)+'a')
}
