package pkg

import (
	"fmt"
	v1 "github.com/davinash/geode-go/pb/geode/protobuf/v1"
)

func GetEncodedValue(v interface{}) (*v1.EncodedValue, error) {
	eV := &v1.EncodedValue{}

	switch t := v.(type) {
	case string:
		eV.Value = &v1.EncodedValue_StringResult{StringResult: t}
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Type %v not supported", t))
	}
	return eV, nil
}