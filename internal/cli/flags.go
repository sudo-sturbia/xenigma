package cli

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sudo-sturbia/xenigma/pkg/helper"
	"github.com/sudo-sturbia/xenigma/pkg/machine"
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
// Loads configuration at ~/.config/xenigma.json, if configuration is
// incorrect a newly generated machine with specified number of rotors
// is saved to ~/.config/xenigma.json and returned.
func correctIf() *machine.Machine {
	if *correct > 0 { // Load ~/.config/xenigma.json, change if wrong
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
// and writes config to ~/.config/xenigma.json
func generatewIf() *machine.Machine {
	if *generateW > 0 { // Generate and write to ~/.config/xenigma.json
		m := machine.Generate(*generateW)

		err := m.Write(os.Getenv("HOME") + "/.config/xenigma.json")
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
		if err := m.UseRotorsDefaults(); err != nil {
			log.Fatal(err)
		}
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
// config at ~/.config/xenigma.json
func updateIf(m *machine.Machine) {
	if *update { // Write updated configs to ~/.config/xenigma.json
		err := m.Write(os.Getenv("HOME") + "/.config/xenigma.json")
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
			fmt.Printf("INVALID: %s\n", err.Error())
		} else {
			fmt.Println("VALID")
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
				"    xenigma [options] <message>\n" +
				"For more details use -help.")

	}

	// Detailed help message
	if *longH {
		defer os.Exit(0)
		fmt.Println(
			"Description\n" +
				"    xenigma is a modified version of the enigma encryption machine.\n" +
				"\n" +
				"Usage\n" +
				"    xenigma [options] <message>\n" +
				"\n" +
				"Options\n" +
				"    -h                           Print a short help message.\n" +
				"\n" +
				"    -help                        Print a detailed help message.\n" +
				"\n" +
				"    -config-h                    Print a help message specifying how to\n" +
				"                                 configure a machine.\n" +
				"\n" +
				"    -verify <path>               Verify the correctness of the configuration\n" +
				"                                 at the given path.\n" +
				"\n" +
				"    -generate <numberofrotors>   Generate a machine with specified number\n" +
				"                                 of rotors and use it for encryption.\n" +
				"\n" +
				"    -gen-w <numberofrotors>      Generate a machine with specified number\n" +
				"                                 of rotors, use it for encryption, and write\n" +
				"                                 generated configs to ~/.config/xenigma.json\n" +
				"\n" +
				"    -correct <numberofrotors>    Load ~/.config/xenigma.json, generate a new\n" +
				"                                 machine if configs are incorrect.\n" +
				"\n" +
				"    -load <path>                 Load and use config at given path instead\n" +
				"                                 of ~/.config/xenigma.json\n" +
				"\n" +
				"    -read <path>                 Read and encrypt contents of file at given\n" +
				"                                 path. If both -read is invoked and a message\n" +
				"                                 is given as argument, both are encrypted and\n" +
				"                                 and printed seperated by a new line.\n" +
				"\n" +
				"    -write <path>                Write encrypted message to file at given path.\n" +
				"\n" +
				"    -update                      Save updated config to ~/.config/xenigma.json\n" +
				"                                 before exiting. Updated config is config at\n" +
				"                                 ~/.config/xenigma.json after rotor shifting.\n" +
				"\n" +
				"    -default-rotors              Use default values for rotor-related fields.\n" +
				"                                 Default values are \"a\"'s for rotor positions,\n" +
				"                                 1 for step size, and 26 for cycle size.\n" +
				"\n" +
				"xenigma is licensed under MIT license.\n" +
				"For source code check the github repo [github.com/sudo-sturbia/xenigma].")
	}

	// Configuration help message
	if *config {
		defer os.Exit(0)
		fmt.Println(
			"Configuration\n" +
				"\n" +
				"    xenigma allows for configuration of all machine's componenets through\n" +
				"    JSON. Configurations file should be located at ~/.config/xenigma.json\n" +
				"\n" +
				"    An example of a ~/.config/xenigma.json is the following\n" +
				"\n" +
				"    {\n" +
				"        \"rotors\": [\n" +
				"            {\n" +
				"                \"pathways\": [\"a\", \"b\", \"c\", ...],\n" +
				"                \"position\": \"a\",\n" +
				"                \"step\": 1,\n" +
				"                \"cycle\": 26\n" +
				"            },\n" +
				"            {\n" +
				"                \"pathways\": [\"a\", \"b\", \"c\", ...],\n" +
				"                \"position\": \"b\",\n" +
				"                \"step\": 1,\n" +
				"                \"cycle\": 26\n" +
				"            },\n" +
				"            {\n" +
				"                \"pathways\": [\"a\", \"b\", \"c\", ...],\n" +
				"                \"position\": \"c\",\n" +
				"                \"step\": 1,\n" +
				"                \"cycle\": 26\n" +
				"            }\n" +
				"        ],\n" +
				"        \n" +
				"        \"reflector\": {\n" +
				"            \"connections\": [\"a\", \"b\", \"c\", ...]\n" +
				"        },\n" +
				"\n" +
				"        \"plugboard\": {\n" +
				"            \"connections\": [\"a\", \"b\", \"c\", ...]\n" +
				"        }\n" +
				"    }\n" +
				"\n" +
				"    Rotors\n" +
				"\n" +
				"    xenigma allows for any number of rotors. The number of rotors is the size\n" +
				"    of \"rotors\" array in ~/.config/xenigma.json\n" +
				"\n" +
				"    Rotor's fields are: pathways, position, step, and cycle.\n" +
				"\n" +
				"    Pathways are represented using a 26 element array where indices represent \n" +
				"    characters and array elements represent the character they are connected\n" +
				"    to. For example if element at index 0 is \"b\", then \"a\" (character 0) is \n" +
				"    connected to \"b\".\n" +
				"\n" +
				"    Position is an integer which represents the current position of the rotor, \n" +
				"    and must be reachable from the starting position (\"a\").\n" +
				"\n" +
				"    Step is the number of positions a rotor shifts when stepping once (the size \n" +
				"    of rotor's jump.) For example if a rotor at position \"a\", with step = 3, \n" +
				"    steps once, then rotor's position changes to \"d\". The default step size is 1.\n" +
				"\n" +
				"    Cycle is the number of rotor steps considered a full cycle, after which the\n" +
				"    following rotor steps (is shifted.) For example, if a rotor has a cycle = 13,\n" +
				"    then the rotor needs to complete 13 steps in order for the following rotor \n" +
				"    to step once. The default cycle size is 26.\n" +
				"\n" +
				"    Reflector\n" +
				"\n" +
				"    Reflector consists of a connections array similar to pathways with a \n" +
				"    condition that it must be symmetric, meaning that if \"a\" is connected to \"b\", \n" +
				"    then \"b\" must also be connected to \"a\".\n" +
				"\n" +
				"    Plugboard\n" +
				"\n" +
				"    Plugboard, also, consists of a connections array exactly the same as a \n" +
				"    reflector.\n" +
				"\n" +
				"    Plugboard's connections are required to have 26 elements, so characters \n" +
				"    not connected to anything should be connected to themselves (in order \n" +
				"    to not be transformed).")
	}
}
