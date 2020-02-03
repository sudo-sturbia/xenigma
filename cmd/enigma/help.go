package main

import (
	"flag"
	"fmt"
	"os"
)

// help prints a help message and exits if an option is specified.
func help() {
	// Short help message
	if *shortH || (flag.NArg() == 0 && flag.NFlag() == 0) {
		fmt.Println(
			"Usage\n" +
				"    enigma [options] <message>\n" +
				"For more details use -help.")

		os.Exit(0)
	}

	// Detailed help message
	if *longH {
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
				"    -config                      Print a help message specifying how to\n" +
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

		os.Exit(0)
	}

	// Configuration help message
	if *config {
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

		os.Exit(0)
	}
}
