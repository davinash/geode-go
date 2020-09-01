// Client package provides to create client and
// entry point for this go package
package client

import (
	"fmt"
	v1 "github.com/davinash/geode-go/pb/geode/protobuf/v1"
	"github.com/davinash/geode-go/pkg"
)

type GeodeClient struct {
	Pool *pkg.Pool
}

// NewClient Creates a new client,
// maxConn is maximum number of connection this instance of client
// will open
func NewClient(maxConn int) (*GeodeClient, error) {
	return &GeodeClient{Pool: pkg.NewPool(maxConn)}, nil
}

// Disconnect disconnect current connection
func (g *GeodeClient) Disconnect() {

}

// Region Creates a instance of a new Region
func (g *GeodeClient) Region(s string) *pkg.Region {
	return &pkg.Region{
		Name: s,
		Pool: g.Pool,
	}
}

// GetRegionNames Gets all the regions created in a cluster
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

// AddServer Add new server for this client to connect
func (g *GeodeClient) AddServer(host string, port int) error {
	return g.Pool.AddServer(host, port)
}
