package pkg

import (
	v1 "github.com/davinash/geode-go/pb/geode/protobuf/v1"
	"log"
)

type Region struct {
	Name string
	Conn *Connection
}

func (r *Region) Put(key interface{}, val interface{}) error {
	log.Println("Doing Put Now")
	entry := v1.Entry{}
	v, err := GetEncodedValue(key)
	if err != nil {
		return err
	}
	entry.Key = v

	v, err = GetEncodedValue(val)
	if err != nil {
		return err
	}
	entry.Value = v
	putRequest := v1.PutRequest{
		RegionName: r.Name,
		Entry:      &entry,
	}
	log.Printf("PutRequest -> %+v", putRequest)

	putResp, err := r.Conn.SendAndReceive(&putRequest)
	if err != nil {
		return err
	}
	log.Println(putResp)

	return nil
}

func (r *Region) Get() error {
	return nil
}
