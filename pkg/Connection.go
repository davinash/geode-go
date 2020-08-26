package pkg

import (
	"fmt"
	"github.com/davinash/geode-go/pb/geode/protobuf"
	"github.com/golang/protobuf/proto"
	"net"
)

type Connection struct {
	// Host IP address or hostname of the Geode Locator
	Host string
	// Port port on which Geode Locator is running
	Port int
	// Connection
	Conn net.Conn
}

func NewGeodeConnection(host string, port int) (*Connection, error) {
	c := &Connection{
		Host: host,
		Port: port,
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		return nil, err
	}
	c.Conn = conn
	cv := protobuf.NewConnectionClientVersion{
		MajorVersion: uint32(protobuf.MajorVersions_CURRENT_MAJOR_VERSION),
		MinorVersion: uint32(protobuf.MinorVersions_CURRENT_MINOR_VERSION),
	}
	new(protobuf.NewConnectionClientVersion)
	// Write to Geode
	buffer := proto.NewBuffer(nil)
	err = buffer.EncodeMessage(&cv)
	if err != nil {
		return nil, err
	}
	_, err = c.Conn.Write(buffer.Bytes())
	if err != nil {
		return nil, err
	}

	return c, nil
}
