package jnigo

import (
	"fmt"
	"testing"
)

func TestConvertArray(t *testing.T) {
	jvm := CreateJVM()

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
	}

	for _, test := range testArray {
		value, err := jvm.Convert(test[0])
		if err != nil {
			t.Fatal(err)
		}

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

func TestConvertPrimitive(t *testing.T) {
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
		value, err := jvm.Convert(test[0])
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

func TestConvertJString(t *testing.T) {
	jvm := CreateJVM()
	output, err := jvm.Convert("test")
	if err != nil {
		t.Fatal(err)
	}
	if output.GoValue() != "test" {
		t.Fatal(output.GoValue())
	}
}
func TestConvertJObject(t *testing.T) {
	jvm := CreateJVM()

	input, err := jvm.Convert(1)
	if err != nil {
		t.Fatal(err)
	}
	output, err := jvm.Convert(input)
	if err != nil {
		t.Fatal(err)
	}

	if input != output {
		t.Fatal(input, output)
	}
}
