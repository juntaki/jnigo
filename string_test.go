package jnigo

import "testing"

func TestJString(t *testing.T) {
	jvm := CreateJVM()

	jstr, err := jvm.newjString("test")
	if err != nil {
		t.Fatal(err)
	}

	if jstr.GoValue() != "test" {
		t.Fatal(jstr.GoValue())
	}
}
