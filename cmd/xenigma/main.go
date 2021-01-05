package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sudo-sturbia/xenigma/v3/pkg/machine"
)

var configPath = fmt.Sprintf("%s/.config/xenigma/xenigma.conf", os.Getenv("HOME"))

var (
	config    = flag.Bool("config-h", false, "print configuration help message")
	verify    = flag.String("verify", "", "verifies the correctness of a machine")
	generate  = flag.Int("generate", -1, "generate a machine with specified number of rotors")
	generateW = flag.Int("gen-w", -1, "generate a machine with n rotors and save it")
	genBackup = flag.Int("gen-backup", -1, "generate a machine with n rotors if default is invalid")
	load      = flag.String("load", "", "use the machine at given path")
	read      = flag.String("read", "", "encrypt contents of a file")
	update    = flag.Bool("update", false, "overwrite machine with new settings after encryption")
	defaults  = flag.Bool("defaults", false, "use default values for rotor-related fields")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	if *config {
		configUsage()
	}

	if *verify != "" {
		_, err := machine.Read(*verify)
		if err != nil {
			fmt.Printf("INVALID: %s\n", err.Error())
		} else {
			fmt.Printf("VALID\n")
		}
	} else {
		m, err := newMachine()
		if err != nil {
			exitWith(err)
		}

		if *defaults {
			m.Rotors().UseDefaults()
		}

		message, err := message()
		if err != nil {
			exitWith(err)
		}

		enc, err := m.Encrypt(message)
		if err != nil {
			exitWith(err)
		}
		fmt.Printf("%s", enc)

		if *update {
			err := machine.Write(m, configPath)
			if err != nil {
				exitWith(err)
			}
		}
	}
}

// message retrieves a user message to encrypt from the command line and/or
// a file, and returns an error in case if there's no file, or no message is
// given.
func message() (string, error) {
	builder := new(strings.Builder)
	for _, argument := range flag.Args() {
		builder.WriteString(argument + " ")
	}
	builder.WriteByte('\n')

	if *read != "" {
		file, err := ioutil.ReadFile(*read)
		if err != nil {
			return "", fmt.Errorf("failed to read %s: %s", *read, err.Error())
		}
		builder.Write(file)
	}

	message := builder.String()
	if message == "" {
		return "", fmt.Errorf("no message to encrypt")
	}
	return message, nil
}

// newMachine creates and a machine based on command line flags.
func newMachine() (*machine.Machine, error) {
	if *generate > 0 && *generateW > 0 {
		return nil, fmt.Errorf("can't use both -generate and -gen-w")
	}
	if *generate > 0 {
		return machine.Generate(*generate), nil
	}
	if *generateW > 0 {
		m := machine.Generate(*generateW)
		return m, machine.Write(m, configPath)
	}

	path := path()
	m, err := machine.Read(path)
	if err != nil {
		if *genBackup > 0 {
			return machine.Generate(*genBackup), nil
		}
		return nil, err
	}

	return m, nil
}

// path returns the path of the machine to load.
func path() string {
	switch {
	case *load != "":
		return *load
	default:
		return configPath
	}
}

// exitWith prints an error message and exits..
func exitWith(err error) {
	fmt.Printf("Error: %s\n", err.Error())
	os.Exit(1)
}

// usage prints a user help message.
func usage() {
	fmt.Fprint(os.Stderr,
		"Description\n",
		"    xenigma is a modified version of the enigma encryption machine.\n",
		"\n",
		"Usage\n",
		"    xenigma [options] <message>\n",
		"\n",
		"Options\n",
		"    -help                Print this help message.\n",
		"\n",
		"    -config-h            Print a help guide explaining xenigma.conf.\n",
		"\n",
		"    -verify <path>       Verify the correctness of the the machine at\n",
		"                         path.\n",
		"\n",
		"    -generate <count>    Generate a machine with given number of rotors\n",
		"                         use it for encryption.\n",
		"\n",
		"    -gen-w <count>       Generate a machine with given number of rotors,\n",
		"                         write it to ~/.config/xenigma/xenigma.conf, and\n",
		"                         use it for encryption.\n",
		"\n",
		"    -gen-backup <count>  Generate a new machine, if\n",
		"                         ~/.config/xenigma/xenigma.conf is invalid.\n",
		"\n",
		"    -load <path>         Use machine at path instead of\n",
		"                         ~/.config/xenigma/xenigma.conf.\n",
		"\n",
		"    -read <path>         Read message from path. If both -read is invoked and\n",
		"                         a message is given as argument, both are encrypted\n",
		"                         and printed seperated by a new line.\n",
		"\n",
		"    -update              Save updated machine to ~/.config/xenigma/xenigma.conf\n",
		"                         before exiting.\n",
		"\n",
		"    -defaults            Use default values for rotor-related fields.\n",
		"                         Default values are \"a\"'s for rotor positions,\n",
		"                         1 for step size, and 26 for cycle size.\n",
		"\n",
		"See github.com/sudo-sturbia/xenigma for source code.\n",
	)
}

// configUsage prints a help message explaining xenigma.conf's options
// and exits.
func configUsage() {
	defer os.Exit(0)
}
