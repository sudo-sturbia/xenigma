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
	PathConnections      [][alphabetSize]string `json:"pathways"`
	Reflector            [alphabetSize]string   `json:"reflector"`
	PlugboardConnections [alphabetSize]string   `json:"plugboard"`
	RotorsPositions      []string               `json:"rotorPositions"`
	Step                 int                    `json:"rotorStep"`
	Cycle                int                    `json:"rotorCycle"`
}

// Read loads a machine from a JSON file and verifies its configurations.
// Returns a pointer to the loaded Machine and an error in case of incorrect configs.
func Read(path string) (*Machine, error) {
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

	// Verify arrays' sizes
	if len(jsonM.PathConnections) != len(jsonM.RotorsPositions) {
		return nil, &initError{"pathways and rotors positions arrays are not of the same size"}
	}

	// Parse jsonM into a Machine
	m := new(Machine)
	m.setNumberOfRotors(len(jsonM.PathConnections))

	// Electric pathways
	m.pathConnections = make([][alphabetSize]int, m.numberOfRotors)

	for i := 0; i < m.numberOfRotors; i++ {
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
			return nil, &initError{fmt.Sprintf("reflector contains invalid value %v",
				connection)}
		}
	}

	// Rotors
	rotorsPositions := make([]int, m.numberOfRotors)
	for i, position := range jsonM.RotorsPositions {
		if num, verify := strToInt(position); verify {
			rotorsPositions[i] = num
		} else {
			return nil, &initError{fmt.Sprintf("rotorsPositions contains invalid value %v",
				position)}
		}

	}

	if err := m.initRotors(rotorsPositions, jsonM.Step, jsonM.Cycle); err != nil {
		return nil, err
	}

	return m, nil
}

// Write writes configurations of a Machine object to a JSON file.
// Returns an error if Machine is not initialized correctly ot
// unable to write to file.
func (m *Machine) Write(path string) error {
	if err := m.isInit(); err != nil {
		return err
	}

	jsonM := new(jsonMachine)

	// Electric pathways
	jsonM.PathConnections = make([][alphabetSize]string, m.numberOfRotors)
	for i := 0; i < m.numberOfRotors; i++ {
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
	jsonM.RotorsPositions = make([]string, m.numberOfRotors)
	for i := 0; i < m.numberOfRotors; i++ {
		jsonM.RotorsPositions[i] = intToStr(m.CurrentRotors()[i])
	}

	jsonM.Step = m.step
	jsonM.Cycle = m.cycle

	contents, err := json.MarshalIndent(jsonM, "", "\t")
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
	return fmt.Sprintf("%c", byte(num)+'a')
}
