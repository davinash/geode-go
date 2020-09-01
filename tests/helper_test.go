package tests

import (
	"fmt"
	client "github.com/davinash/geode-go"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

type GeodeTestSuite struct {
	suite.Suite
	GeodeHome string
	Client    *client.GeodeClient
}

func (suite *GeodeTestSuite) SetupTest() {
	log.Println("SetupTest")
	suite.GeodeHome = os.Getenv("GEODE_HOME")
	if suite.GeodeHome == "" {
		suite.T().Fatalf("Define Environment variable GEODE_HOME")
	}
	err := suite.startLocator()
	if err != nil {
		suite.Fail("Failed to Start Locator %v", err)
	}

	geodeClient, err := client.NewClient(100)
	if err != nil {
		suite.Fail("Failed to create new Geode Client %v", err)
	}

	suite.Client = geodeClient
	err = suite.startServers(3)
	if err != nil {
		suite.Fail("Failed to Start Cache servers %v", err)
	}
}

func (suite *GeodeTestSuite) TearDownTest() {
	log.Println("TearDownTest")
	suite.stopCluster()
}

func TestGeodeTestSuite(t *testing.T) {
	suite.Run(t, new(GeodeTestSuite))
}

func (suite *GeodeTestSuite) startLocator() error {
	d, err := ioutil.TempDir("", "Locator")
	if err != nil {
		return err
	}
	cmd := exec.Command(filepath.Join(suite.GeodeHome, "bin", "gfsh"),
		"-e",
		fmt.Sprintf("start locator --name=locator --bind-address=localhost --port=10334 --dir=%s",
			d))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("Running Command = %s\n", cmd.String())
	err = cmd.Run()
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

func (suite *GeodeTestSuite) startServers(num int) error {
	port := 40404
	for i := 0; i < num; i++ {
		d, err := ioutil.TempDir("", "Server")
		if err != nil {
			return err
		}
		c := fmt.Sprintf("start server --name=server-%d --bind-address=localhost --server-port=%d "+
			"--J=-Dgeode.feature-protobuf-protocol=true --dir=%s", i, port, d)
		cmd := exec.Command(filepath.Join(suite.GeodeHome, "bin", "gfsh"),
			"-e",
			"connect",
			"-e",
			c)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Printf("Running Command = %s\n", cmd.String())
		err = cmd.Run()
		if err != nil {
			return err
		}
		err = suite.Client.AddServer("localhost", port)
		if err != nil {
			return err
		}
		port = port + 1
	}
	return nil
}

type RegionType string

const (
	Replicate RegionType = "REPLICATE"
	Partition RegionType = "PARTITION"
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
