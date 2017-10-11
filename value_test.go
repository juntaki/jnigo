package jnigo

import (
	"fmt"
	"testing"
)

func TestjPrimitive(t *testing.T) {
	jvm := CreateJVM()

	testArray := [][]interface{}{
		[]interface{}{false, SignatureBoolean},
		[]interface{}{true, SignatureBoolean},
		[]interface{}{byte(1), SignatureByte},
		[]interface{}{byte(100), SignatureByte},
		[]interface{}{uint16(1), SignatureChar},
		[]interface{}{uint16(10000), SignatureChar},
		[]interface{}{int16(1), SignatureShort},
		[]interface{}{int16(10000), SignatureShort},
		[]interface{}{int32(1), SignatureInt},
		[]interface{}{int32(10000), SignatureInt},
		[]interface{}{int64(1), SignatureLong},
		[]interface{}{int64(10000), SignatureLong},
		[]interface{}{float32(1.0), SignatureFloat},
		[]interface{}{float32(1000.0), SignatureFloat},
		[]interface{}{float64(1.0), SignatureDouble},
		[]interface{}{float64(1000.0), SignatureDouble},
	}

	for _, test := range testArray {
		fmt.Println(test)
		value, err := jvm.newJPrimitive(test[0])
		fmt.Println(value)
		if err != nil {
			t.Fatal(err)
		}
		if value.GoValue() != test[0] {
			t.Fatal(value.GoValue())
		}
		if value.Signature() != test[1] {
			t.Fatal(value.GoValue())
		}
	}
}
