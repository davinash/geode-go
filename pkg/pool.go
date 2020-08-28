package pkg

import (
	v1 "github.com/davinash/geode-go/pb/geode/protobuf/v1"
	"github.com/golang/protobuf/proto"
	"io"
	"net"
	"sync"
)

type GeodeConnection struct {
	Host  string
	Port  int
	InUse bool
	Conn  net.Conn
	mux   sync.Mutex
}

type Pool struct {
	Servers     map[string]int
	Connections []*GeodeConnection
}

func NewPool() *Pool {
	return &Pool{}
}

func (p *Pool) Disconnect() {
	for _, c := range p.Connections {
		c.Conn.Close()
	}
}

func (p *Pool) AddServer(host string, port int) {
	p.Servers[host] = port
}

func (p *Pool) Send(m proto.Message, conn net.Conn) error {
	buffer := proto.NewBuffer(nil)
	err := buffer.EncodeMessage(m)
	if err != nil {
		return err
	}
	_, err = conn.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (p *Pool) ReceiveRaw(conn net.Conn) ([]byte, error) {
	// receive a response
	data := make([]byte, 4096)
	bytesRead, err := conn.Read(data)
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
		n, err := io.ReadFull(conn, data[bytesRead:messageLength])
		if err != nil {
			return nil, err
		}
		bytesRead += n
	}
	return data[0:bytesRead], err
}

func (p *Pool) Receive(conn net.Conn) (*v1.Message, error) {
	raw, err := p.ReceiveRaw(conn)
	if err != nil {
		return nil, err
	}
	b := proto.NewBuffer(raw)
	var pr v1.Message
	err = b.DecodeMessage(&pr)
	if err != nil {
		return nil, err
	}
	return &pr, nil
}

func (p *Pool) SendAndReceive(m proto.Message) (*v1.Message, error) {
	conn, err := p.GetOrCreateConnection()
	if err != nil {
		return nil, err
	}
	if err := p.Send(m, conn); err != nil {
		return nil, err
	}
	return p.Receive(conn)
}

func (p *Pool) GetOrCreateConnection() (net.Conn, error) {
	return nil, nil
}
