// cli handles parsing of command line arguments and flags
// and execution of application.
package cli

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/sudo-sturbia/xenigma/pkg/machine"
)

// Execute evaluates command line arguments and flags, and executes
// the program accordingly.
func Execute() {
	flag.Parse()

	helpIf()
	verifyIf()

	m := getMachine()
	defaultsIf(m)

	encrypted, _ := m.Encrypt(message())
	if !writeIf(encrypted) {
		fmt.Println(encrypted)
	}

	updateIf(m)
}

// getMachine returns a Machine object based on specified command
// line flags, exits if no Machine is loaded.
func getMachine() *machine.Machine {
	if m := correctIf(); m != nil {
		return m
	} else if m := loadIf(); m != nil {
		return m
	} else if m := generateIf(); m != nil {
		return m
	} else if m := generatewIf(); m != nil {
		return m
	} else {
		m, err := machine.Load(-1, false)
		if err != nil {
			log.Fatal(err)
		}

		return m
	}
}

// message returns a string consisting of command line arguments
// (concatenated) and contents of a file (if a file is specified)
// to be encrypted.
func message() string {
	// Collect command line arguments
	builder := new(strings.Builder)
	for _, argument := range flag.Args() {
		builder.WriteString(argument + " ")
	}
	message := builder.String()

	return readIf(message)
}
