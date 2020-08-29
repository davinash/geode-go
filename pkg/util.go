package pkg

import (
	"encoding/json"
	"fmt"
	v1 "github.com/davinash/geode-go/pb/geode/protobuf/v1"
)

func Serialize(v interface{}) (*v1.EncodedValue, error) {
	eV := &v1.EncodedValue{}
	if v == nil {
		eV.Value = &v1.EncodedValue_NullResult{}
	} else {
		switch t := v.(type) {
		case string:
			eV.Value = &v1.EncodedValue_StringResult{StringResult: t}
		case []byte:
			eV.Value = &v1.EncodedValue_BinaryResult{BinaryResult: t}
		case int:
			eV.Value = &v1.EncodedValue_IntResult{IntResult: int32(t)}
		case int16:
			eV.Value = &v1.EncodedValue_ShortResult{ShortResult: int32(t)}
		case int32:
			eV.Value = &v1.EncodedValue_IntResult{IntResult: t}
		case int64:
			eV.Value = &v1.EncodedValue_LongResult{LongResult: t}
		case byte:
			eV.Value = &v1.EncodedValue_ByteResult{ByteResult: int32(t)}
		case bool:
			eV.Value = &v1.EncodedValue_BooleanResult{BooleanResult: t}
		case float64:
			eV.Value = &v1.EncodedValue_DoubleResult{DoubleResult: t}
		case float32:
			eV.Value = &v1.EncodedValue_FloatResult{FloatResult: t}
		default:
			bytes, err := json.Marshal(t)
			if err != nil {
				return nil, err
			}
			eV.Value = &v1.EncodedValue_JsonObjectResult{JsonObjectResult: string(bytes)}
		}
	}
	return eV, nil
}

func Deserialize(ev *v1.EncodedValue) (interface{}, error) {
	var v interface{}

	switch t := ev.GetValue().(type) {
	case *v1.EncodedValue_StringResult:
		v = t.StringResult
	case *v1.EncodedValue_NullResult:
		v = nil
	case *v1.EncodedValue_IntResult:
		v = t.IntResult
	case *v1.EncodedValue_ShortResult:
		v = t.ShortResult
	case *v1.EncodedValue_LongResult:
		v = t.LongResult
	case *v1.EncodedValue_ByteResult:
		v = uint8(t.ByteResult)
	case *v1.EncodedValue_BooleanResult:
		v = t.BooleanResult
	case *v1.EncodedValue_DoubleResult:
		v = t.DoubleResult
	case *v1.EncodedValue_FloatResult:
		v = t.FloatResult
	case *v1.EncodedValue_BinaryResult:
		v = t.BinaryResult
	case *v1.EncodedValue_JsonObjectResult:
		v = []byte(t.JsonObjectResult)
	default:
		return nil, fmt.Errorf("%v Type not supported", t)
	}
	return v, nil
}
