package pkg

import (
	"fmt"
	"github.com/davinash/geode-go/pb/geode/protobuf"
	v1 "github.com/davinash/geode-go/pb/geode/protobuf/v1"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"math/rand"
	"net"
	"time"
)

type Server struct {
	Host string
	Port int
	Conn net.Conn
}

type Pool struct {
	Servers []Server
	pool    chan *Server
}

func NewPool(maxConn int) *Pool {
	return &Pool{
		pool: make(chan *Server, maxConn),
	}
}

func (p *Pool) Disconnect() {
	//for _, c := range p.Connections {
	//	c.Conn.Close()
	//}
}

func (p *Pool) AddServer(host string, port int) error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}
	cv := protobuf.NewConnectionClientVersion{
		MajorVersion: uint32(protobuf.MajorVersions_CURRENT_MAJOR_VERSION),
		MinorVersion: uint32(protobuf.MinorVersions_CURRENT_MINOR_VERSION),
	}
	err = p.Send(&cv, conn)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := p.ReceiveRaw(conn)
	if err != nil {
		log.Fatalln(err)
	}
	b := proto.NewBuffer(resp)
	var va protobuf.VersionAcknowledgement
	err = b.DecodeMessage(&va)
	if err != nil {
		return err
	}
	if va.GetVersionAccepted() == false {
		return fmt.Errorf("client version is not compitable with server")
	}

	p.Servers = append(p.Servers, Server{
		Host: host,
		Port: port,
	})
	p.pool <- &Server{
		Host: host,
		Port: port,
		Conn: conn,
	}
	return nil
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

func (p *Pool) Return(s *Server) {
	select {
	case p.pool <- s:
	default:
		// let it go, let it go...
	}
}

func (p *Pool) SendAndReceive(m proto.Message) (*v1.Message, error) {
	s, err := p.Borrow()
	defer p.Return(s)

	if err != nil {
		return nil, err
	}
	if err := p.Send(m, s.Conn); err != nil {
		return nil, err
	}
	return p.Receive(s.Conn)
}

func (p *Pool) Borrow() (*Server, error) {
	var c *Server
	select {
	case c = <-p.pool:
	default:
		rand.Seed(time.Now().UTC().UnixNano())
		min := 0
		max := len(p.Servers) - 1
		idx := min + rand.Intn(max-min)
		p.AddServer(p.Servers[idx].Host, p.Servers[idx].Port)

		c = <-p.pool
	}
	return c, nil
}
