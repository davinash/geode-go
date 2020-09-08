package pkg

import (
	"fmt"
	v1 "github.com/davinash/geode-go/pb/geode/protobuf/v1"
)

type Function struct {
	Pool *Pool
}

func (f *Function) ExecuteFunctionOnRegion(functionId string, regionName string,
	arguments interface{}, keyFilter []interface{}) ([]interface{}, error) {
	request := v1.ExecuteFunctionOnRegionRequest{
		FunctionID: functionId,
		Region:     regionName,
	}
	args, err := Serialize(arguments)
	if err != nil {
		return nil, fmt.Errorf("failed to Serialized the arguments")
	}
	request.Arguments = args
	kfs := make([]*v1.EncodedValue, 0)
	for _, kf := range keyFilter {
		v, err := Serialize(kf)
		if err != nil {
			return nil, fmt.Errorf("failed to Serialized the Key Filters")
		}
		kfs = append(kfs, v)
	}
	request.KeyFilter = kfs
	msg := v1.Message{
		MessageType: &v1.Message_ExecuteFunctionOnRegionRequest{
			ExecuteFunctionOnRegionRequest: &request,
		},
	}
	resp, err := f.Pool.SendAndReceive(&msg)
	if err != nil {
		return nil, err
	}
	if resp.GetErrorResponse() != nil {
		return nil, fmt.Errorf(fmt.Sprintf("ExecuteFunctionOnRegion Failed Message = %s, Error Code = %d",
			resp.GetErrorResponse().GetError().Message,
			resp.GetErrorResponse().GetError().ErrorCode))
	}
	ers := resp.GetExecuteFunctionOnRegionResponse().GetResults()
	result := make([]interface{}, 0)
	for _, er := range ers {
		v, err := Deserialize(er)
		if err != nil {
			return nil, fmt.Errorf("failed to Serialized result")
		}
		result = append(result, v)
	}
	return result, nil
}

func (f *Function) ExecuteFunctionOnMember(functionId string, members []string, arguments interface{}) ([]interface{}, error) {
	request := v1.ExecuteFunctionOnMemberRequest{
		FunctionID: functionId,
	}
	request.MemberName = append(request.MemberName, members...)
	args, err := Serialize(arguments)
	if err != nil {
		return nil, fmt.Errorf("failed to Serialized the arguments")
	}
	request.Arguments = args

	msg := v1.Message{
		MessageType: &v1.Message_ExecuteFunctionOnMemberRequest{
			ExecuteFunctionOnMemberRequest: &request,
		},
	}
	resp, err := f.Pool.SendAndReceive(&msg)
	if err != nil {
		return nil, err
	}
	if resp.GetErrorResponse() != nil {
		return nil, fmt.Errorf(fmt.Sprintf("ExecuteFunctionOnMember Failed Message = %s, Error Code = %d",
			resp.GetErrorResponse().GetError().Message,
			resp.GetErrorResponse().GetError().ErrorCode))
	}
	ers := resp.GetExecuteFunctionOnMemberResponse().GetResults()
	result := make([]interface{}, 0)
	for _, er := range ers {
		v, err := Deserialize(er)
		if err != nil {
			return nil, fmt.Errorf("failed to Serialized result")
		}
		result = append(result, v)
	}
	return result, nil
}

func (f *Function) ExecuteFunctionOnGroup(functionId string, groupName string) error {
	return nil
}
