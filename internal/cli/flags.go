package cli

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sudo-sturbia/enigma/pkg/helper"
	"github.com/sudo-sturbia/enigma/pkg/machine"
)

// Command line flags.
var (
	write     = flag.String("write", "", "write encrypted message to file at given path")
	read      = flag.String("read", "", "encrypt contents of file at given path")
	load      = flag.String("load", "", "load and use config at given path")
	verify    = flag.String("verify", "", "verifies the correctness of config at given path")
	update    = flag.Bool("update", false, "save updated config before exiting")
	defaults  = flag.Bool("default-rotors", false, "use default values for rotor-related fields")
	shortH    = flag.Bool("h", false, "print a short help message")
	longH     = flag.Bool("help", false, "print a detailed help message")
	config    = flag.Bool("config-h", false, "print configuration help message")
	correct   = flag.Int("correct", -1, "load configs, generate a new machine if incorrect")
	generate  = flag.Int("generate", -1, "generate a machine with specified number of rotors")
	generateW = flag.Int("gen-w", -1, "generate a machine and save it's configs")
)

// correctIf handles the execution of -correct flag.
// Loads configuration at ~/.config/enigma.json, if configuration is
// incorrect a newly generated machine with specified number of rotors
// is saved to ~/.config/enigma.json and returned.
func correctIf() *machine.Machine {
	if *correct > 0 { // Load ~/.config/enigma.json, change if wrong
		m, err := machine.Load(*correct, true)
		if err != nil {
			if m == nil {
				log.Fatal(err)
			} else {
				fmt.Println(err)
			}
		}
		return m
	}

	return nil
}

// loadIf handles the execution of -load flag.
// Loads and returns machine at given path instead of default config.
func loadIf() *machine.Machine {
	if *load != "" { // Load machine at given path
		m, err := machine.Read(*load)
		if err != nil {
			log.Fatal(err)
		}
		return m
	}

	return nil
}

// generateIf handles the execution of -generate flag.
// Returns a newly generated machine with specified number of rotors.
func generateIf() *machine.Machine {
	if *generate > 0 { // Generate a machine to use
		return machine.Generate(*generate)
	}

	return nil
}

// generatewIf handles the execution of -gen-w flag.
// Returns a newly generated machine with specified number of rotors
// and writes config to ~/.config/enigma.json
func generatewIf() *machine.Machine {
	if *generateW > 0 { // Generate and write to ~/.config/enigma.json
		m := machine.Generate(*generateW)

		err := m.Write(os.Getenv("HOME") + "/.config/enigma.json")
		if err != nil {
			log.Fatal(err)
		}

		return m
	}

	return nil
}

// defaultsIf handles the execution of -default-rotors flag.
func defaultsIf(m *machine.Machine) {
	if *defaults { // Use default values for rotors
		m.UseRotorDefaults()
	}
}

// writeIf handles the execution of -write flag.
// Writes message to file if option is specified.
func writeIf(encrypted string) bool {
	if *write != "" { // Write message to file at given path
		helper.WriteStringToFile(encrypted, *write)
		return true
	}

	return false
}

// readIf handles the execution of -read flag.
// Reads contents of a file and returns a string consisting of message
// + file contents. Returns message without change if can't read.
func readIf(message string) string {
	if *read != "" { // Read contents of file at given path
		readMessage := helper.ReadStringFromFile(*read)

		switch {
		case message == "" && readMessage == "":
			log.Fatal("no message given")
		case readMessage == "":
			return message
		case message == "":
			return readMessage
		default:
			return fmt.Sprintf("%s\n%s", readMessage, message)
		}
	}

	return message
}

// updateIf checks if -update flag was specified, if so updates
// config at ~/.config/enigma.json
func updateIf(m *machine.Machine) {
	if *update { // Write updated configs to ~/.config/enigma.json
		err := m.Write(os.Getenv("HOME") + "/.config/enigma.json")
		if err != nil {
			log.Fatal(err)
		}
	}
}

// verifyIf  handles the execution of -verify flag.
// Verifies config at given path, prints a message accordingly and exits.
func verifyIf() {
	if *verify != "" {
		defer os.Exit(0)

		_, err := machine.Read(*verify)
		if err != nil {
			fmt.Printf("Config is INCORRECT\n%s\n", err.Error())
		} else {
			fmt.Println("Config is CORRECT")
		}
	}
}

// helpIf handles the execution of -help, -h, or -config-h flags.
// Prints specified help message and exits.
func helpIf() {
	// Short help message
	if *shortH || (flag.NArg() == 0 && flag.NFlag() == 0) {
		defer os.Exit(0)
		fmt.Println(
			"Usage\n" +
				"    enigma [options] <message>\n" +
				"For more details use -help.")

	}

	// Detailed help message
	if *longH {
		defer os.Exit(0)
		fmt.Println(
			"Description\n" +
				"    enigma is a modified version of the enigma encryption machine.\n" +
				"\n" +
				"Usage\n" +
				"    enigma [options] <message>\n" +
				"\n" +
				"Options\n" +
				"    -h                           Print a short help message.\n" +
				"\n" +
				"    -help                        Print a detailed help message.\n" +
				"\n" +
				"    -help                        Print a detailed help message.\n" +
				"\n" +
				"    -config-h                    Print a help message specifying how to\n" +
				"                                 configure a machine.\n" +
				"\n" +
				"    -generate <numberofrotors>   Generate a machine with specified number\n" +
				"                                 of rotors and use it for encryption.\n" +
				"\n" +
				"    -gen-w <numberofrotors>      Generate a machine with specified number\n" +
				"                                 of rotors, use it for encryption, and write\n" +
				"                                 generated configs to ~/.config/engima.json\n" +
				"\n" +
				"    -correct <numberofrotors>    Load ~/.config/engima.json, generate a new\n" +
				"                                 machine if configs are incorrect.\n" +
				"\n" +
				"    -load <path>                 Load and use config at given path instead\n" +
				"                                 of ~/.config/engima.json\n" +
				"\n" +
				"    -read <path>                 Read and encrypt contents of file at given\n" +
				"                                 path. If both -read is invoked and a message\n" +
				"                                 is given as argument, both are encrypted and\n" +
				"                                 and printed seperated by a new line.\n" +
				"\n" +
				"    -write <path>                Write encrypted message to file at given path.\n" +
				"\n" +
				"    -update                      Save updated config to ~/.config/engima.json\n" +
				"                                 before exiting. Updated config is config at\n" +
				"                                 ~/.config/engima.json after rotor shifting.\n" +
				"\n" +
				"    -default-rotors              Use default values for rotor-related fields.\n" +
				"                                 Default values are \"a\"'s for rotor positions,\n" +
				"                                 1 for step size, and 26 for cycle size.\n" +
				"\n" +
				"enigma is licensed under MIT license.\n" +
				"For source code check the github repo [github.com/sudo-sturbia/enigma].")
	}

	// Configuration help message
	if *config {
		defer os.Exit(0)
		fmt.Println(
			"Configuration\n" +
				"\n" +
				"    enigma allows for configuration of all machine's componenets through\n" +
				"    JSON. Configurations file should be located at ~/.config/engima.json\n" +
				"\n" +
				"    An example of a ~/.config/enigma.json is the following\n" +
				"\n" +
				"    {\n" +
				"        \"pathways\": [\n" +
				"             [\"a\", \"b\", \"c\", ...],\n" +
				"             [\"a\", \"b\", \"c\", ...],\n" +
				"             [\"a\", \"b\", \"c\", ...]\n" +
				"        ],\n" +
				"        \"reflector\": [\"a\", \"b\", \"c\", ...],\n" +
				"        \"plugboard\": [\"a\", \"b\", \"c\", ...],\n" +
				"        \"rotorPositions\": [\"a\", \"b\", \"c\"],\n" +
				"        \"rotorStep\": 1,\n" +
				"        \"rotorCycle\": 26\n" +
				"    }\n" +
				"\n" +
				"    enigma allows for a variable number of rotors. The number of rotors is\n" +
				"    is decided through the number of electric pathways arrays, or the number\n" +
				"    of rotor positions, which should be equal.\n" +
				"\n" +
				"    Connections, such as pathways, plugboard, and reflector are specified\n" +
				"    through arrays where an element's index represents a character's position\n" +
				"    in the alphabet. For example if element at index 0 of plugboard array is\n" +
				"    \"c\", then \"a\" is connected to \"c\"." +
				"\n" +
				"    Both reflector and plugboard arrays should be symmetric, meaning that if\n" +
				"    \"a\" is connected to \"b\", \"b\" must also be connected to \"a\". Otherwise\n" +
				"    connections are considered incorrect.\n" +
				"\n" +
				"    enigma also allows for configuration of rotors' step and cycle sizes.\n" +
				"    Step size is the number of positions a rotor jumps when shifting. For\n" +
				"    example if a rotor, with step size 2, is at position \"a\", then the rotor\n" +
				"    will jump to \"c\" when shifted once.\n" +
				"    Cycle size is the number of steps a rotor takes to complete a full cycle.\n" +
				"    when a rotor completes a full cycle, the adjacent rotor is shifted. For\n" +
				"    example in a 3-rotor machine if cycle size is 13 then the second rotor\n " +
				"    is shifted once every time the first rotor completes 13 steps, the third\n " +
				"    rotor operates similarly but depends on second rotor's movement, etc.")
	}
}
