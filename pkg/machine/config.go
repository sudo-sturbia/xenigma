package machine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
)

// jsonMachine is a struct used for reading and writing of Machine's configs
// from/into a json file. Fields in jsonMachine mirror those in a Machine but
// use string arrays instead of int arrays.
type jsonMachine struct {
	Rotors    []*jsonRotor   `json:"rotors"`
	Reflector *jsonReflector `json:"reflector"`
	Plugboard *jsonPlugboard `json:"plugboard"`
}

// jsonRotor is used to marshall/unmarshall Rotor configs from/into a json
// file. Fields in jsonRotor mirror those in Rotor but use strings.
type jsonRotor struct {
	Pathways [alphabetSize]string `json:"pathways"`
	Position string               `json:"position"`
	Step     int                  `json:"step"`
	Cycle    int                  `json:"cycle"`
}

// jsonReflector  is used to marshall/unmarshall Reflector configs from/into
// a json file. Fields in jsonReflector mirror those in Reflector but use strings.
type jsonReflector struct {
	Connections [alphabetSize]string `json:"connections"`
}

// jsonPlugboard is used to marshall/unmarshall Plugboard configs from/into
// a json file. Fields in jsonPlugboard mirror those in Plugboard but use strings.
type jsonPlugboard struct {
	Connections [alphabetSize]string `json:"connections"`
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

	m, err := parseMachine(fileContents)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal %s: %s", path, err.Error())
	}

	// Verify correct initialization
	if err := m.IsConfigCorrect(); err != nil {
		return nil, err
	}

	return m, nil
}

// parseMachine parses a given file's byte array into character arrays,
// Creates a Machine object, Sets Machine's fields, and returns a pointer.
// An error is returned in case of invalid configs.
func parseMachine(fileContents []byte) (*Machine, error) {
	var jsonM jsonMachine
	if err := json.Unmarshal(fileContents, &jsonM); err != nil {
		return nil, err
	}

	// Verify rotors' slice
	if jsonM.Rotors == nil || len(jsonM.Rotors) == 0 {
		return nil, &initError{"no rotors given"}
	}

	m := new(Machine)

	// Rotors
	rotorCount := len(jsonM.Rotors)
	rotors := make([]*Rotor, rotorCount)
	for i := 0; i < rotorCount; i++ {
		rotor, err := m.parseRotor(jsonM.Rotors[i])
		if err != nil {
			return nil, err
		}

		rotors[i] = rotor
	}

	err := m.SetRotors(rotors)
	if err != nil {
		return nil, err
	}

	m.plugboard, err = m.parsePlugboard(jsonM.Plugboard)
	if err != nil {
		return nil, err
	}

	m.reflector, err = m.parseReflector(jsonM.Reflector)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// parseRotor parses a given jsonRotor into a Rotor object. Returns
// parsed object and an error in case of incorrect configs.
func (m *Machine) parseRotor(parse *jsonRotor) (*Rotor, error) {
	parsed := new(Rotor)

	var (
		pathways [alphabetSize]int
		position int
	)

	// Pathways
	for i, connection := range parse.Pathways {
		if num, ok := strToInt(connection); ok {
			pathways[i] = num
		} else {
			return nil, &initError{fmt.Sprintf("rotor pathways contain invalid value %v",
				connection)}
		}
	}

	// Position, step, and cycle.
	if num, ok := strToInt(parse.Position); ok {
		position = num
	} else {
		return nil, &initError{fmt.Sprintf("given rotor position %v is incorrect",
			position)}
	}

	if err := parsed.Set(pathways, position, parse.Step, parse.Cycle); err != nil {
		return nil, err
	}

	return parsed, nil
}

// parsePlugboard parses a given jsonPlugboard into a Plugboard object. Returns
// parsed object and an error in case of incorrect configs.
func (m *Machine) parsePlugboard(parse *jsonPlugboard) (*Plugboard, error) {
	if parse == nil {
		return nil, &initError{"non plugboard given"}
	}

	parsed := new(Plugboard)
	for i, connection := range parse.Connections {
		if num, verify := strToInt(connection); verify {
			parsed.connections[i] = num
		} else {
			return nil, &initError{
				fmt.Sprintf("plugboard connections contain invalid value %v", connection),
			}
		}
	}

	return parsed, nil
}

// parseReflector parses a given jsonReflector into a Reflector object. Returns
// parsed object and an error in case of incorrect configs.
func (m *Machine) parseReflector(parse *jsonReflector) (*Reflector, error) {
	if parse == nil {
		return nil, &initError{"non reflector given"}
	}

	parsed := new(Reflector)
	for i, connection := range parse.Connections {
		if num, verify := strToInt(connection); verify {
			parsed.connections[i] = num
		} else {
			return nil, &initError{
				fmt.Sprintf("reflector connections contain invalid value %v", connection),
			}
		}
	}

	return parsed, nil
}

// Write writes configurations of a Machine object to a JSON file.
// Returns an error if Machine is not initialized correctly or
// unable to write to file.
func (m *Machine) Write(path string) error {
	if err := m.IsConfigCorrect(); err != nil {
		return err
	}

	jsonM := new(jsonMachine)

	jsonM.Plugboard = m.marshalPlugboard(m.plugboard)
	jsonM.Reflector = m.marshalReflector(m.reflector)

	jsonM.Rotors = make([]*jsonRotor, m.numberOfRotors)
	for i := 0; i < m.numberOfRotors; i++ {
		jsonM.Rotors[i] = m.marshalRotor(m.rotors[i])
	}

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

// marshalRotor creates and returns a jsonRotor object with properties
// similar to given Rotor object.
func (m *Machine) marshalRotor(rotor *Rotor) *jsonRotor {
	marshalled := new(jsonRotor)

	// Pathways
	for i := 0; i < alphabetSize; i++ {
		marshalled.Pathways[i] = intToStr(rotor.pathways[i])
	}

	// Position, step, and cycle
	marshalled.Position = intToStr(rotor.position)
	marshalled.Step = rotor.step
	marshalled.Cycle = rotor.cycle

	return marshalled
}

// marshalPlugboard creates and returns a jsonPlugboard object with properties
// similar to given Pluboard object.
func (m *Machine) marshalPlugboard(plugboard *Plugboard) *jsonPlugboard {
	marshalled := new(jsonPlugboard)
	for i := 0; i < alphabetSize; i++ {
		marshalled.Connections[i] = intToStr(plugboard.connections[i])
	}

	return marshalled
}

// marshalReflector creates and returns a jsonReflector object with properties
// similar to given Reflector object.
func (m *Machine) marshalReflector(reflector *Reflector) *jsonReflector {
	marshalled := new(jsonReflector)
	for i := 0; i < alphabetSize; i++ {
		marshalled.Connections[i] = intToStr(reflector.connections[i])
	}

	return marshalled
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
