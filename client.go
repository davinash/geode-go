package client

import (
	"fmt"
	v1 "github.com/davinash/geode-go/pb/geode/protobuf/v1"
	"github.com/davinash/geode-go/pkg"
)

type GeodeClient struct {
	// Connection object
	Conn *pkg.Connection
}

// NewConnection Creates a new connection with Geode Cluster
func NewConnection(host string, port int) (*GeodeClient, error) {
	g := &GeodeClient{}
	c, err := pkg.NewGeodeConnection(host, port)
	if err != nil {
		return nil, err
	}
	g.Conn = c
	return g, nil
}

// Disconnect disconnect current connection
func (g *GeodeClient) Disconnect() {

}

func (g *GeodeClient) Region(s string) *pkg.Region {
	return &pkg.Region{
		Name: s,
		Conn: g.Conn,
	}
}

func (g *GeodeClient) GetRegionNames() ([]string, error) {
	msg := v1.Message{MessageType: &v1.Message_GetRegionNamesRequest{}}
	resp, err := g.Conn.SendAndReceive(&msg)
	if err != nil {
		return nil, err
	}
	if resp.GetErrorResponse() != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Get Failed Message = %s, Error Code = %d",
			resp.GetErrorResponse().GetError().Message,
			resp.GetErrorResponse().GetError().ErrorCode))
	}
	return resp.GetGetRegionNamesResponse().GetRegions(), nil
}
