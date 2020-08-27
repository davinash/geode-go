package pkg

import (
	"fmt"
	v1 "github.com/davinash/geode-go/pb/geode/protobuf/v1"
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

func (r *Region) Get(key interface{}) (interface{}, error) {
	v, err := GetEncodedValue(key)
	if err != nil {
		return nil, err
	}

	getRequest := v1.GetRequest{
		RegionName: r.Name,
		Key:        v,
	}
	m := v1.Message{MessageType: &v1.Message_GetRequest{GetRequest: &getRequest}}

	resp, err := r.Conn.SendAndReceive(&m)
	if err != nil {
		return nil, err
	}
	if resp.GetErrorResponse() != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Get Failed Message = %s, Error Code = %d",
			resp.GetErrorResponse().GetError().Message,
			resp.GetErrorResponse().GetError().ErrorCode))
	}
	ev := resp.GetGetResponse().GetResult()
	value, err := GetDecodedValue(ev)
	if err != nil {
		return nil, err
	}
	return value, nil
}
