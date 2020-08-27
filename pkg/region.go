package pkg

type Region struct {
	Name string
	Conn *Connection
}

func (r *Region) Put() error {
	return nil
}

func (r *Region) Get() error {
	return nil
}
