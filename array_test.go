package jnigo

import (
	"fmt"
	"testing"
)

func TestJArray(t *testing.T) {
	jvm := CreateJVM()

	//clazz := "TestClass"
	//gojclass, _ := jvm.NewJClass(clazz, []JObject{})

	testArray := [][]interface{}{
		[]interface{}{[]bool{false, false}, SignatureArray + SignatureBoolean},
		[]interface{}{[]bool{true, false}, SignatureArray + SignatureBoolean},
		[]interface{}{[]byte{1, 1}, SignatureArray + SignatureByte},
		[]interface{}{[]byte{100, 100}, SignatureArray + SignatureByte},
		[]interface{}{[]uint16{1, 1}, SignatureArray + SignatureChar},
		[]interface{}{[]uint16{10000, 10000}, SignatureArray + SignatureChar},
		[]interface{}{[]int16{1, 1}, SignatureArray + SignatureShort},
		[]interface{}{[]int16{10000, 10000}, SignatureArray + SignatureShort},
		[]interface{}{[]int32{1, 1}, SignatureArray + SignatureInt},
		[]interface{}{[]int32{10000, 10000}, SignatureArray + SignatureInt},
		[]interface{}{[]int64{1, 1}, SignatureArray + SignatureLong},
		[]interface{}{[]int64{10000, 10000}, SignatureArray + SignatureLong},
		[]interface{}{[]float32{1.0, 1.0}, SignatureArray + SignatureFloat},
		[]interface{}{[]float32{1000.0, 1000.0}, SignatureArray + SignatureFloat},
		[]interface{}{[]float64{1.0, 1.0}, SignatureArray + SignatureDouble},
		[]interface{}{[]float64{1000.0, 1000.0}, SignatureArray + SignatureDouble},
		//[]interface{}{[]JObject{gojclass, gojclass}, SignatureArray + gojclass.Signature()},
	}

	for _, test := range testArray {
		fmt.Println(test)
		value, err := jvm.NewJArray(test[0])
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(value)

		orig := fmt.Sprintln(test[0])
		goval := fmt.Sprintln(value.GoValue())

		if orig != goval {
			t.Fatal(orig, goval)
		}

		if value.Signature() != test[1] {
			t.Fatal(value.GoValue())
		}
	}
}
