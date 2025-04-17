package integration

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Change to the root directory
	if err := os.Chdir("../../.."); err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}
