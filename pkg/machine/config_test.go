package machine

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// TestRead reads several config files and verifies the output.
func TestRead(t *testing.T) {
	const correctCount = 3
	const incorrectCount = 7

	for i := 1; i <= correctCount; i++ {
		_, err := Read(fmt.Sprintf("../../test/data/config-%d.json", i))
		if err != nil {
			t.Errorf("correct %d: produced error %w", i, err)
		}
	}

	for i := 1; i <= incorrectCount; i++ {
		_, err := Read(fmt.Sprintf("../../test/data/wrong-config-%d.json", i))
		if err == nil {
			t.Errorf("incorrect %d: didn't produce error", i)
		}
	}
}

// TestWrite tests writing a machine to a file.
func TestWrite(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	err := os.MkdirAll("../../test/generate", os.ModePerm)
	if err != nil {
		t.Fatal("failed to create test/generate")
	}

	for i := 0; i < 10; i++ {
		count := rand.Intn(100) + 3
		err := Write(Generate(count), "../../test/generate/generated.json")
		if err != nil {
			t.Errorf("test %d: failed to write: %w", i, err)
		}

		_, err = os.Open("../../test/generate/generated.json")
		if err != nil {
			t.Errorf("test %d: written machine doesn't exist", i)
		}
	}
}

// TestReadAndWrite tests Read and Write by generating a machine, writing it
// to a file, reading the file, and comparing written and read.
func TestReadAndWrite(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	err := os.MkdirAll("../../test/generate", os.ModePerm)
	if err != nil {
		t.Fatal("failed to create test/generate")
	}

	unexported := cmp.AllowUnexported(
		Machine{},
		Rotors{},
		Rotor{},
		Plugboard{},
		Reflector{},
	)

	for i := 0; i < 10; i++ {
		m := Generate(rand.Intn(100) + 3)
		err := Write(m, "../../test/generate/generated.json")
		if err != nil {
			t.Errorf("failed to write: %w", err)
		}

		r, err := Read("../../test/generate/generated.json")
		if err != nil {
			t.Errorf("failed to read: %w", err)
		}

		if diff := cmp.Diff(*m, *r, unexported); diff != "" {
			t.Errorf("test %d: mismatch (-want +got):\n%s", i, diff)
		}
	}
}
