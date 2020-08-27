package pkg

import (
	"fmt"
	v1 "github.com/davinash/geode-go/pb/geode/protobuf/v1"
	"log"
)

type Region struct {
	Name string
	Conn *Connection
}

func (r *Region) Put(key interface{}, val interface{}) error {
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
	m := v1.Message{
		MessageType: &v1.Message_PutRequest{PutRequest: &putRequest},
	}
	log.Printf("PutRequest -> %+v", putRequest)

	putResp, err := r.Conn.SendAndReceive(&m)
	if err != nil {
		return err
	}
	if putResp.GetErrorResponse() != nil {
		return fmt.Errorf(fmt.Sprintf("Put Failed Message = %s, Error Code = %d",
			putResp.GetErrorResponse().GetError().Message,
			putResp.GetErrorResponse().GetError().ErrorCode))
	}
	return nil
}

func (r *Region) Get() error {
	return nil
}
