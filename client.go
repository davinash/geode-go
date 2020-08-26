package client

import (
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
