package pkg

import (
	v1 "github.com/davinash/geode-go/pb/geode/protobuf/v1"
	"github.com/golang/protobuf/proto"
	"io"
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
	return nil
}

func (c *Connection) ReceiveRaw() ([]byte, error) {
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

func (c *Connection) Receive() (*v1.Message, error) {
	raw, err := c.ReceiveRaw()
	if err != nil {
		return nil, err
	}
	p := proto.NewBuffer(raw)
	var pr v1.Message
	err = p.DecodeMessage(&pr)
	if err != nil {
		return nil, err
	}
	return &pr, nil
}

func (c *Connection) SendAndReceive(m proto.Message) (*v1.Message, error) {
	// Send a Message
	if err := c.Send(m); err != nil {
		return nil, err
	}
	return c.Receive()
}
