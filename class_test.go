package jnigo

import (
	"fmt"
	"testing"
)

func TestJClass(t *testing.T) {
	jvm := CreateJVM()

	testArray := [][]interface{}{
		[]interface{}{"java/lang/String", []JObject{}},
	}

	for _, test := range testArray {
		fmt.Println(test)
		value, err := jvm.NewJClass(test[0].(string), test[1].([]JObject))
		fmt.Println(value)
		if err != nil {
			t.Fatal(err)
		}

		if value.Signature() != "L"+test[0].(string)+";" {
			t.Fatal(value.GoValue())
		}

		v, err := value.CallFunction("length", "()I", []JObject{})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("length", v.GoValue())
	}
}

func TestJClassMethod(t *testing.T) {
	jvm := CreateJVM()

	clazz := "TestClass"
	testArray := [][]string{
		[]string{"mvboolean", "()Z"},
		[]string{"mvbyte", "()B"},
		[]string{"mvchar", "()C"},
		[]string{"mvshort", "()S"},
		[]string{"mvint", "()I"},
		[]string{"mvlong", "()J"},
		[]string{"mvfloat", "()F"},
		[]string{"mvdouble", "()D"},
		[]string{"mvclass", "()LTestClass;"},

		[]string{"maboolean", "()[Z"},
		[]string{"mabyte", "()[B"},
		[]string{"machar", "()[C"},
		[]string{"mashort", "()[S"},
		[]string{"maint", "()[I"},
		[]string{"malong", "()[J"},
		[]string{"mafloat", "()[F"},
		[]string{"madouble", "()[D"},
		[]string{"maclass", "()[LTestClass;"},
	}

	value, err := jvm.NewJClass(clazz, []JObject{})
	if err != nil {
		t.Fatal(err)
	}

	if value.Signature() != "L"+clazz+";" {
		t.Fatal(value.GoValue())
	}

	for _, test := range testArray {
		v, err := value.CallFunction(test[0], test[1], []JObject{})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("return ", v.GoValue(), v.Signature())
	}
}

func TestJClassStaticMethod(t *testing.T) {
	jvm := CreateJVM()

	clazz := "TestClass"
	testArray := [][]string{
		[]string{"smvboolean", "()Z"},
		[]string{"smvbyte", "()B"},
		[]string{"smvchar", "()C"},
		[]string{"smvshort", "()S"},
		[]string{"smvint", "()I"},
		[]string{"smvlong", "()J"},
		[]string{"smvfloat", "()F"},
		[]string{"smvdouble", "()D"},
		[]string{"smvclass", "()LTestClass;"},

		[]string{"smaboolean", "()[Z"},
		[]string{"smabyte", "()[B"},
		[]string{"smachar", "()[C"},
		[]string{"smashort", "()[S"},
		[]string{"smaint", "()[I"},
		[]string{"smalong", "()[J"},
		[]string{"smafloat", "()[F"},
		[]string{"smadouble", "()[D"},
		[]string{"smaclass", "()[LTestClass;"},
	}

	for _, test := range testArray {
		v, err := jvm.CallStaticFunction(clazz, test[0], test[1], []JObject{})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("return ", v.GoValue(), v.Signature())
	}
}

func TestJClassGetField(t *testing.T) {
	jvm := CreateJVM()

	clazz := "TestClass"
	testArray := [][]string{
		[]string{"vboolean", "Z"},
		[]string{"vbyte", "B"},
		[]string{"vchar", "C"},
		[]string{"vshort", "S"},
		[]string{"vint", "I"},
		[]string{"vlong", "J"},
		[]string{"vfloat", "F"},
		[]string{"vdouble", "D"},
		[]string{"vclass", "LTestClass;"},

		[]string{"aboolean", "[Z"},
		[]string{"abyte", "[B"},
		[]string{"achar", "[C"},
		[]string{"ashort", "[S"},
		[]string{"aint", "[I"},
		[]string{"along", "[J"},
		[]string{"afloat", "[F"},
		[]string{"adouble", "[D"},
		[]string{"aclass", "[LTestClass;"},
	}

	value, err := jvm.NewJClass(clazz, []JObject{})
	if err != nil {
		t.Fatal(err)
	}

	if value.Signature() != "L"+clazz+";" {
		t.Fatal(value.GoValue())
	}

	for _, test := range testArray {
		v, err := value.GetField(test[0], test[1])
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("return ", v.GoValue(), v.Signature())
	}
}

func TestJClassSetField(t *testing.T) {
	jvm := CreateJVM()

	clazz := "TestClass"
	gobool, _ := jvm.NewJPrimitive(false)
	gobyte, _ := jvm.NewJPrimitive(byte(20))
	gochar, _ := jvm.NewJPrimitive(uint16(20))
	goshort, _ := jvm.NewJPrimitive(int16(20))
	goint, _ := jvm.NewJPrimitive(int32(20))
	golong, _ := jvm.NewJPrimitive(int64(20))
	gofloat, _ := jvm.NewJPrimitive(float32(20))
	godouble, _ := jvm.NewJPrimitive(float64(20))
	// gojclass, _ := jvm.NewJClass(clazz, []JObject{})

	// goabool, _ := jvm.NewJArray([]bool{true, false})
	// goabyte, _ := jvm.NewJArray([]byte{100, 100})
	// goachar, _ := jvm.NewJArray([]uint16{10000, 10000})
	// goashort, _ := jvm.NewJArray([]int16{10000, 10000})
	// goaint, _ := jvm.NewJArray([]int32{10000, 10000})
	// goalong, _ := jvm.NewJArray([]int64{10000, 10000})
	// goafloat, _ := jvm.NewJArray([]float32{1000.0, 1000.0})
	// goadouble, _ := jvm.NewJArray([]float64{1000.0, 1000.0})
	//goajclass, _ := jvm.NewJArray([]JClass{1000.0, 1000.0})

	testArray := [][]interface{}{
		[]interface{}{"vboolean", gobool},
		[]interface{}{"vbyte", gobyte},
		[]interface{}{"vchar", gochar},
		[]interface{}{"vshort", goshort},
		[]interface{}{"vint", goint},
		[]interface{}{"vlong", golong},
		[]interface{}{"vfloat", gofloat},
		[]interface{}{"vdouble", godouble},
		//[]interface{}{"vclass", gojclass},

		// []interface{}{"aboolean", goabool},
		// []interface{}{"abyte", goabyte},
		// []interface{}{"achar", goachar},
		// []interface{}{"ashort", goashort},
		// []interface{}{"aint", goaint},
		// []interface{}{"along", goalong},
		// []interface{}{"afloat", goafloat},
		// []interface{}{"adouble", goadouble},
		//[]interface{}{"aclass", goajclass},
	}

	value, err := jvm.NewJClass(clazz, []JObject{})
	if err != nil {
		t.Fatal(err)
	}

	if value.Signature() != "L"+clazz+";" {
		t.Fatal(value.GoValue())
	}

	for _, test := range testArray {
		err := value.SetField(test[0].(string), test[1].(JObject))
		if err != nil {
			t.Fatal(err)
		}
	}
}
