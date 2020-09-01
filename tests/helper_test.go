package tests

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

type GeodeTestSuite struct {
	suite.Suite
	GeodeHome string
}

func (suite *GeodeTestSuite) SetupTest() {
	log.Println("SetupTest")
	suite.GeodeHome = os.Getenv("GEODE_HOME")
	if suite.GeodeHome == "" {
		suite.T().Fatalf("Define Environment variable GEODE_HOME")
	}
	suite.startLocator()
}

func (suite *GeodeTestSuite) TearDownTest() {
	log.Println("TearDownTest")
	suite.stopCluster()
}

func TestGeodeTestSuite(t *testing.T) {
	suite.Run(t, new(GeodeTestSuite))
}

func (suite *GeodeTestSuite) startLocator() error {
	cmd := exec.Command(filepath.Join(suite.GeodeHome, "bin", "gfsh"),
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

func (suite *GeodeTestSuite) stopCluster() error {
	cmd := exec.Command(filepath.Join(suite.GeodeHome, "bin", "gfsh"),
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

func (suite *GeodeTestSuite) startServers(num int) ([]int, error) {
	port := 40404
	ports := make([]int, 0)
	for i := 0; i < num; i++ {
		c := fmt.Sprintf("start server --name=server-%d --bind-address=localhost --server-port=%d "+
			"--J=-Dgeode.feature-protobuf-protocol=true", i, port)
		cmd := exec.Command(filepath.Join(suite.GeodeHome, "bin", "gfsh"),
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

func (suite *GeodeTestSuite) createRegion(name string, regionType RegionType) error {
	switch regionType {
	case Replicate:
		c := fmt.Sprintf("create region --name=%s --type=%s", name, regionType)
		cmd := exec.Command(filepath.Join(suite.GeodeHome, "bin", "gfsh"),
			"-e",
			"connect",
			"-e",
			c)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		suite.T().Logf("Running Command = %s\n", cmd.String())
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
