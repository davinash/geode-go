// +build integration

package pkg

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

func startLocator() {
	cmd := exec.Command(filepath.Join(geodeHome, "bin", "gfsh"),
		"-e",
		"start locator --name=locator --bind-address=localhost --port=10334")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func TestPut(t *testing.T) {
	t.Run("BasicPut", func(t *testing.T) {
		geodeHome := os.Getenv("GEODE_HOME")
		if geodeHome == "" {
			t.Fatalf("Define Environment variable GEODE_HOME")
		}

	})
}
