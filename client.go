package client

import (
	"fmt"
	v1 "github.com/davinash/geode-go/pb/geode/protobuf/v1"
	"github.com/davinash/geode-go/pkg"
)

type GeodeClient struct {
	// Connection object
	Pool *pkg.Pool
}

// NewClient Creates a new connection with Geode Cluster
func NewClient(host string, port int) (*GeodeClient, error) {
	g := &GeodeClient{}
	g.Pool = pkg.NewPool()

	return g, nil
}

// Disconnect disconnect current connection
func (g *GeodeClient) Disconnect() {

}

func (g *GeodeClient) Region(s string) *pkg.Region {
	return &pkg.Region{
		Name: s,
		Pool: g.Pool,
	}
}

func (g *GeodeClient) GetRegionNames() ([]string, error) {
	msg := v1.Message{MessageType: &v1.Message_GetRegionNamesRequest{}}
	resp, err := g.Pool.SendAndReceive(&msg)
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

func (g *GeodeClient) AddServer(host string, port int) {
	//g.Servers[host] = port
	g.Pool.AddServer(host, port)
}
