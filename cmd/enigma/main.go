package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sudo-sturbia/enigma/pkg/helper"
	"github.com/sudo-sturbia/enigma/pkg/machine"
)

// Command line flags.
var (
	write     = flag.String("write", "", "write encrypted message to file at given path")
	read      = flag.String("read", "", "encrypt contents of file at given path")
	load      = flag.String("load", "", "load and use config at given path")
	update    = flag.Bool("update", false, "save updated config before exiting")
	defaults  = flag.Bool("default-rotors", false, "use default values for rotor-related fields")
	shortH    = flag.Bool("h", false, "print a short help message")
	longH     = flag.Bool("help", false, "print a detailed help message")
	config    = flag.Bool("config", false, "print configuration help message")
	correct   = flag.Int("correct", -1, "load configs, generate a new machine if incorrect")
	generate  = flag.Int("generate", -1, "generate a machine with specified number of rotors")
	generateW = flag.Int("gen-w", -1, "generate a machine and save it's configs")
)

func main() {
	flag.Parse()

	m := getMachine()
	message := getMessage()

	encrypted, _ := m.Encrypt(message)

	if *write != "" { // Write message to file at given path
		helper.WriteStringToFile(encrypted, *write)
	} else {
		fmt.Println(encrypted)
	}

	if *update { // Write updated configs to ~/.config/enigma.json
		err := m.Write(os.Getenv("HOME") + "/.config/enigma.json")
		if err != nil {
			log.Fatal(err)
		}
	}
}

// getMachine evaluates flags and returns a machine to use.
func getMachine() *machine.Machine {
	m := new(machine.Machine)
	switch {
	case *correct > 0: // Load ~/.config/enigma.json, change if wrong
		loaded, err := machine.Load(*correct, true)
		if err != nil {
			if loaded == nil {
				log.Fatal(err)
			} else {
				fmt.Println(err)
			}
		}

		m = loaded
	case *load != "": // Load machine at given path
		loaded, err := machine.Read(*load)
		if err != nil {
			log.Fatal(err)
		}

		m = loaded
	case *generate > 0: // Generate a machine to use
		m = machine.Generate(*generate)
	case *generateW > 0: // Generate and write configs to ~/.config/enigma.json
		m = machine.Generate(*generateW)

		err := m.Write(os.Getenv("HOME") + "/.config/enigma.json")
		if err != nil {
			log.Fatal(err)
		}
	default: // Load ~/.config/enigma.json, exit if wrong
		loaded, err := machine.Load(-1, false)
		if err != nil {
			log.Fatal(err)
		}

		m = loaded
	}

	if *defaults { // Use default values for rotors
		m.UseRotorDefaults()
	}

	return m
}

// getMessage returns a message to use.
func getMessage() string {
	// Collect command line arguments
	builder := new(strings.Builder)
	for _, argument := range flag.Args() {
		builder.WriteString(argument + " ")
	}

	message := builder.String()

	if *read != "" { // Read contents of file at given path
		readMessage := helper.ReadStringFromFile(*read)

		switch {
		case message != "" && readMessage != "":
			message = fmt.Sprintf("%s\n%s", readMessage, message)
		case readMessage != "":
			message = readMessage
		}
	}

	return message
}
