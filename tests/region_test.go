// +build integration

package tests

import (
	"fmt"
	client "github.com/davinash/geode-go"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func startLocator(geodeHome string) error {
	cmd := exec.Command(filepath.Join(geodeHome, "bin", "gfsh"),
		"-e",
		"start locator --name=locator --bind-address=localhost --port=10334")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("Running Command = %s\n", cmd.String())
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func stopCluster(geodeHome string) error {
	cmd := exec.Command(filepath.Join(geodeHome, "bin", "gfsh"),
		"-e",
		"connect",
		"-e",
		"shutdown --include-locators=yes")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("Running Command = %s\n", cmd.String())
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func startServers(geodeHome string, num int) ([]int, error) {
	port := 40404
	ports := make([]int, 0)
	for i := 0; i < num; i++ {
		c := fmt.Sprintf("start server --name=server --bind-address=localhost --server-port=%d "+
			"--J=-Dgeode.feature-protobuf-protocol=true", port)
		cmd := exec.Command(filepath.Join(geodeHome, "bin", "gfsh"),
			"-e",
			"connect",
			"-e",
			c)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Printf("Running Command = %s\n", cmd.String())
		err := cmd.Run()
		if err != nil {
			return nil, err
		}
		ports = append(ports, port)
		port = port + 1
	}
	return ports, nil
}

type RegionType string

const (
	Replicate RegionType = "REPLICATE"
	Partition RegionType = "PARTITIONED"
)

func createRegion(geodeHome string, name string, regionType RegionType) error {
	switch regionType {
	case Replicate:
		c := fmt.Sprintf("create region --name=%s --type=%s", name, regionType)
		cmd := exec.Command(filepath.Join(geodeHome, "bin", "gfsh"),
			"-e",
			"connect",
			"-e",
			c)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Printf("Running Command = %s\n", cmd.String())
		err := cmd.Run()
		if err != nil {
			return err
		}
	case Partition:

	default:
		return fmt.Errorf(fmt.Sprintf("Unknown Region type received = %v", regionType))
	}
	return nil
}

func TestPut(t *testing.T) {
	t.Run("BasicPut", func(t *testing.T) {
		geodeHome := os.Getenv("GEODE_HOME")
		if geodeHome == "" {
			t.Fatalf("Define Environment variable GEODE_HOME")
		}
		err := startLocator(geodeHome)
		if err != nil {
			t.FailNow()
		}
		defer stopCluster(geodeHome)

		ports, err := startServers(geodeHome, 2)
		if err != nil {
			t.FailNow()
		}
		geodeClient, err := client.NewClient(100)
		if err != nil {
			log.Fatalln(err)
		}
		err = createRegion(geodeHome, "SampleData", Replicate)
		if err != nil {
			log.Fatalln(err)
		}

		for _, p := range ports {
			err = geodeClient.AddServer("127.0.0.1", p)
			if err != nil {
				log.Fatalln(err)
			}
		}

		log.Println(ports)

	})
}
