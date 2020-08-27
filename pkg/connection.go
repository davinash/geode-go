package pkg

import (
	"fmt"
	"github.com/davinash/geode-go/pb/geode/protobuf"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
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

func (c *Connection) Send(m proto.Message) error {
	buffer := proto.NewBuffer(nil)
	err := buffer.EncodeMessage(m)
	if err != nil {
		return err
	}
	_, err = c.Conn.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	log.Println("Send Complete")
	return nil
}

func (c *Connection) Receive() ([]byte, error) {
	// receive a response
	data := make([]byte, 4096)
	bytesRead, err := c.Conn.Read(data)
	if err != nil {
		return nil, err
	}

	l, n := proto.DecodeVarint(data)
	messageLength := int(l) + n

	if messageLength > len(data) {
		t := make([]byte, len(data), messageLength)
		copy(t, data)
		data = t
	}
	for bytesRead < messageLength {
		n, err := io.ReadFull(c.Conn, data[bytesRead:messageLength])
		if err != nil {
			return nil, err
		}
		bytesRead += n
	}
	return data[0:bytesRead], err
}

func (c *Connection) SendAndReceive(m proto.Message) ([]byte, error) {
	// Send a Message
	if err := c.Send(m); err != nil {
		return nil, err
	}
	return c.Receive()
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
	resp, err := c.SendAndReceive(&cv)
	if err != nil {
		log.Fatalln(err)
	}
	p := proto.NewBuffer(resp)
	var va protobuf.VersionAcknowledgement
	err = p.DecodeMessage(&va)
	if err != nil {
		return nil, err
	}
	if va.GetVersionAccepted() == false {
		return nil, fmt.Errorf("client version is not compitable with server")
	} else {
		log.Println("Connection established")
	}
	return c, nil
}
