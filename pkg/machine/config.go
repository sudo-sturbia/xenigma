package machine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"unicode"
)

// jsonMachine is used for (un)marshalling a Machine from/into a json
// file. Fields in jsonMachine mirror those in a Machine but use string
// arrays instead of int arrays.
type jsonMachine struct {
	Rotors    []*jsonRotor   `json:"rotors"`
	Reflector *jsonReflector `json:"reflector"`
	Plugboard *jsonPlugboard `json:"plugboard"`
}

// jsonRotor mirrors Rotor struct and is used for json (un)marshalling.
type jsonRotor struct {
	Pathways [alphabetSize]string `json:"pathways"`
	Position string               `json:"position"`
	Step     int                  `json:"step"`
	Cycle    int                  `json:"cycle"`
}

// jsonReflector mirrors Reflector struct and is used for json (un)marshalling.
type jsonReflector struct {
	Connections map[string]string `json:"connections"`
}

// jsonPlugboard mirrors Plugboard struct and is used for json (un)marshalling.
type jsonPlugboard struct {
	Connections map[string]string `json:"connections"`
}

// Read loads a machine from a JSON file, verifies its validity, and returns
// an error in case of invalid fields.
func Read(path string) (*Machine, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s", path)
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read contents of %s", path)
	}

	m, err := Parse(fileContents)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal %s: %w", path, err)
	}
	return m, m.Verify()
}

// Write writes a Machine to a file as JSON, and returns an error if Machine
// has invalid fields or writing failed.
func Write(m *Machine, path string) error {
	if err := m.Verify(); err != nil {
		return err
	}

	mToJSON := &jsonMachine{
		Rotors:    marshalRotors(m.rotors),
		Plugboard: marshalPlugboard(m.plugboard),
		Reflector: marshalReflector(m.reflector),
	}

	contents, err := json.MarshalIndent(mToJSON, "", "\t")
	if err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	dir, _ := filepath.Split(path)
	err = os.MkdirAll(dir, 0775)
	if err != nil {
		return fmt.Errorf("failed to create %s: %w", dir, err)
	}

	err = ioutil.WriteFile(path, contents, 0664)
	if err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}
	return nil
}

// Parse parses a given byte array into a Machine, and returns a pointer
// to it, and an error in case of invalid fields.
func Parse(contents []byte) (_ *Machine, err error) {
	var jsonM jsonMachine
	if err := json.Unmarshal(contents, &jsonM); err != nil {
		return nil, err
	}

	m := new(Machine)
	m.rotors, err = parseRotors(jsonM.Rotors)
	if err != nil {
		return nil, err
	}

	m.plugboard, err = parsePlugboard(jsonM.Plugboard)
	if err != nil {
		return nil, err
	}

	m.reflector, err = parseReflector(jsonM.Reflector)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// parseRotors parses a given slice of jsonRotor into a Rotors, and returns an
// error if any of the rotors has invalid fields.
func parseRotors(parse []*jsonRotor) (_ *Rotors, err error) {
	if len(parse) == 0 {
		return nil, fmt.Errorf("no rotors given")
	}

	rotors := make([]*Rotor, len(parse))
	for i, toParse := range parse {
		rotors[i], err = parseRotor(toParse)
		if err != nil {
			return nil, err
		}
	}
	return NewRotors(rotors)
}

// parseRotor parses a given jsonRotor into a Rotor, and returns an error
// if Rotor has invalid fields.
func parseRotor(parse *jsonRotor) (*Rotor, error) {
	var pathways [alphabetSize]int
	for i, connection := range parse.Pathways {
		num, ok := strToInt(connection)
		if !ok {
			return nil, fmt.Errorf("invalid rotor pathway %v", connection)
		}
		pathways[i] = num
	}

	position, ok := strToInt(parse.Position)
	if !ok {
		return nil, fmt.Errorf("invalid rotor position %v", position)
	}

	return NewRotor(pathways, position, parse.Step, parse.Cycle)
}

// parsePlugboard parses a given jsonPlugboard into a Plugboard, and returns
// an error if Plugboard has invalid fields.
func parsePlugboard(parse *jsonPlugboard) (*Plugboard, error) {
	if parse == nil || parse.Connections == nil {
		return nil, fmt.Errorf("no plugboard given")
	}
	if len(parse.Connections) != alphabetSize {
		return nil, fmt.Errorf("invalid plugboard size %v, expected %v", len(parse.Connections), alphabetSize)
	}

	connections := make(map[int]int)
	for key, value := range parse.Connections {
		k, ok := strToInt(key)
		if !ok {
			return nil, fmt.Errorf("invalid plugboard key %v", key)
		}
		v, ok := strToInt(value)
		if !ok {
			return nil, fmt.Errorf("invalid plugboard value %v", value)
		}
		connections[k] = v
	}
	return NewPlugboard(connections)
}

// parseReflector parses a given jsonReflector into a Reflector, and returns
// an error if Reflector has invalid fields.
func parseReflector(parse *jsonReflector) (*Reflector, error) {
	if parse == nil || parse.Connections == nil {
		return nil, fmt.Errorf("no reflector given")
	}
	if len(parse.Connections) != alphabetSize {
		return nil, fmt.Errorf("invalid reflector size %v, expected %v", len(parse.Connections), alphabetSize)
	}

	connections := make(map[int]int)
	for key, value := range parse.Connections {
		k, ok := strToInt(key)
		if !ok {
			return nil, fmt.Errorf("invalid reflector key %v", key)
		}
		v, ok := strToInt(value)
		if !ok {
			return nil, fmt.Errorf("invalid reflector value %v", value)
		}
		connections[k] = v
	}
	return NewReflector(connections)
}

// marshalRotors creates and returns a slice of jsonRotor with the same
// fields as given Rotors.
func marshalRotors(rotors *Rotors) []*jsonRotor {
	marshalled := make([]*jsonRotor, rotors.count)
	for i, r := range rotors.rotors {
		marshalled[i] = marshalRotor(r)
	}
	return marshalled
}

// marshalRotor creates and returns a jsonRotor with the same fields
// as given Rotor.
func marshalRotor(rotor *Rotor) *jsonRotor {
	var pathways [alphabetSize]string
	for i, pathway := range rotor.pathways {
		pathways[i] = intToStr(pathway)
	}

	return &jsonRotor{
		Pathways: pathways,
		Position: intToStr(rotor.position),
		Step:     rotor.step,
		Cycle:    rotor.cycle,
	}
}

// marshalPlugboard creates and returns a jsonPlugboard with the same fields
// as given Plugboard.
func marshalPlugboard(plugboard *Plugboard) *jsonPlugboard {
	connections := make(map[string]string)
	for k, v := range plugboard.connections {
		connections[intToStr(k)] = intToStr(v)
	}

	return &jsonPlugboard{
		Connections: connections,
	}
}

// marshalReflector creates and returns a jsonReflector with the same fields
// as given Reflector.
func marshalReflector(reflector *Reflector) *jsonReflector {
	connections := make(map[string]string)
	for k, v := range reflector.connections {
		connections[intToStr(k)] = intToStr(v)
	}

	return &jsonReflector{
		Connections: connections,
	}
}

// strToInt verifies that a given string contains one alphabetical
// character and returns character's position in the alphabet.
func strToInt(str string) (int, bool) {
	if len(str) == 1 && unicode.IsLetter(rune(str[0])) {
		return int(byte(unicode.ToLower(rune(str[0]))) - 'a'), true
	}
	return -1, false
}

// intToStr returns a one character string representing the ASCII position
// of the given integer.
func intToStr(num int) string {
	return fmt.Sprintf("%c", byte(num)+'a')
}
