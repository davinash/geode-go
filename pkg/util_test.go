package pkg

import (
	v1 "github.com/davinash/geode-go/pb/geode/protobuf/v1"
	"testing"
)

func TestSerialize(t *testing.T) {
	t.Run("StringSerialize", func(t *testing.T) {
		serValue, err := Serialize("abc")
		if err != nil {
			t.FailNow()
		}
		_, ok := serValue.GetValue().(*v1.EncodedValue_StringResult)
		if !ok {
			t.FailNow()
		}
	})
	t.Run("BinarySerialize", func(t *testing.T) {
		serValue, err := Serialize([]byte("abc"))
		if err != nil {
			t.FailNow()
		}
		_, ok := serValue.GetValue().(*v1.EncodedValue_BinaryResult)
		if !ok {
			t.FailNow()
		}
	})
	t.Run("IntSerialize", func(t *testing.T) {
		serValue, err := Serialize(100)
		if err != nil {
			t.FailNow()
		}
		_, ok := serValue.GetValue().(*v1.EncodedValue_IntResult)
		if !ok {
			t.FailNow()
		}
	})
	t.Run("IntSerialize", func(t *testing.T) {
		serValue, err := Serialize(100)
		if err != nil {
			t.FailNow()
		}
		_, ok := serValue.GetValue().(*v1.EncodedValue_IntResult)
		if !ok {
			t.FailNow()
		}
	})
	t.Run("BooleanSerialize", func(t *testing.T) {
		serValue, err := Serialize(true)
		if err != nil {
			t.FailNow()
		}
		_, ok := serValue.GetValue().(*v1.EncodedValue_BooleanResult)
		if !ok {
			t.FailNow()
		}
	})
}

func TestDeserialize(t *testing.T) {
	t.Run("StringDeSerialize", func(t *testing.T) {

	})
}
