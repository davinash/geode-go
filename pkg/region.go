package pkg

import (
	"fmt"
	v1 "github.com/davinash/geode-go/pb/geode/protobuf/v1"
)

type Region struct {
	Name string
	Conn *Connection
}

// CreateEntry Creates entry
func CreateEntry(key interface{}, val interface{}) (*v1.Entry, error) {
	entry := v1.Entry{}
	v, err := GetEncodedValue(key)
	if err != nil {
		return nil, err
	}
	entry.Key = v

	v, err = GetEncodedValue(val)
	if err != nil {
		return nil, err
	}
	entry.Value = v

	return &entry, nil
}

func (r *Region) Put(key interface{}, val interface{}) error {
	entry, err := CreateEntry(key, val)
	if err != nil {
		return err
	}

	request := v1.PutRequest{
		RegionName: r.Name,
		Entry:      entry,
	}
	msg := v1.Message{
		MessageType: &v1.Message_PutRequest{PutRequest: &request},
	}

	putResp, err := r.Conn.SendAndReceive(&msg)
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

	request := v1.GetRequest{
		RegionName: r.Name,
		Key:        v,
	}
	msg := v1.Message{MessageType: &v1.Message_GetRequest{GetRequest: &request}}

	resp, err := r.Conn.SendAndReceive(&msg)
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

func (r *Region) PutIfAbsent(key interface{}, val interface{}) (interface{}, error) {
	entry, err := CreateEntry(key, val)
	if err != nil {
		return nil, err
	}
	request := v1.PutIfAbsentRequest{
		RegionName: r.Name,
		Entry:      entry,
	}
	msg := v1.Message{MessageType: &v1.Message_PutIfAbsentRequest{PutIfAbsentRequest: &request}}
	resp, err := r.Conn.SendAndReceive(&msg)
	if err != nil {
		return nil, err
	}
	if resp.GetErrorResponse() != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Get Failed Message = %s, Error Code = %d",
			resp.GetErrorResponse().GetError().Message,
			resp.GetErrorResponse().GetError().ErrorCode))
	}
	ev := resp.GetPutIfAbsentResponse().GetOldValue()
	value, err := GetDecodedValue(ev)
	if err != nil {
		return nil, err
	}
	return value, nil
}

type KeyValue struct {
	Key   interface{}
	Value interface{}
}

func (r *Region) PutAll(kvs []*KeyValue) ([]interface{}, error) {
	entries := make([]*v1.Entry, 0)
	for _, kvs := range kvs {
		entry, err := CreateEntry(kvs.Key, kvs.Value)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	request := v1.PutAllRequest{
		RegionName: r.Name,
		Entry:      entries,
	}

	msg := v1.Message{MessageType: &v1.Message_PutAllRequest{PutAllRequest: &request}}
	resp, err := r.Conn.SendAndReceive(&msg)
	if err != nil {
		return nil, err
	}
	if resp.GetErrorResponse() != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Get Failed Message = %s, Error Code = %d",
			resp.GetErrorResponse().GetError().Message,
			resp.GetErrorResponse().GetError().ErrorCode))
	}
	ev := resp.GetPutAllResponse().GetFailedKeys()
	failedKeys := make([]interface{}, 0)
	for _, k := range ev {
		value, err := GetDecodedValue(k.GetKey())
		if err != nil {
			return failedKeys, err
		}
		failedKeys = append(failedKeys, value)
	}
	return failedKeys, nil
}
