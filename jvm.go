package jnigo

// #cgo linux CFLAGS: -I/usr/lib/jvm/default-java/include -I/usr/lib/jvm/default-java/include/linux
// #cgo linux LDFLAGS: -L/usr/lib/jvm/default-java/jre/lib/amd64/server/ -ljvm -lpthread
// #include "jni_wrapper.h"
import "C"

import (
	"errors"
	"regexp"
	"runtime"
	"unsafe"
)

// Global JVM
var jvm *JVM

type JVM struct {
	cjvm *C.JVM
}

func (jvm *JVM) env() *C.JNIEnv {
	return jvm.cjvm.env
}

func CreateJVM() *JVM {
	cjvm := C.createJVM()
	if cjvm == nil {
		panic("Failed to create JVM")
	}

	jvm = &JVM{
		cjvm: cjvm,
	}
	runtime.SetFinalizer(jvm, freeJVM)
	return jvm
}

func freeJVM(jvm *JVM) {
	C.free(unsafe.Pointer(jvm.cjvm))
}

// This may not work
func (jvm *JVM) destroyJVM() {
	C.destroyJVM(jvm.cjvm)
}

const (
	SignatureBoolean      = "Z"
	SignatureByte         = "B"
	SignatureChar         = "C"
	SignatureShort        = "S"
	SignatureInt          = "I"
	SignatureLong         = "J"
	SignatureFloat        = "F"
	SignatureDouble       = "D"
	SignatureArray        = "["
	SignatureVoid         = "V"
	SignatureClass        = "L"
	SignatureBooleanArray = SignatureArray + SignatureBoolean
	SignatureByteArray    = SignatureArray + SignatureByte
	SignatureCharArray    = SignatureArray + SignatureChar
	SignatureShortArray   = SignatureArray + SignatureShort
	SignatureIntArray     = SignatureArray + SignatureInt
	SignatureLongArray    = SignatureArray + SignatureLong
	SignatureFloatArray   = SignatureArray + SignatureFloat
	SignatureDoubleArray  = SignatureArray + SignatureDouble
	SignatureClassArray   = SignatureArray + SignatureClass
)

var SizeOf = map[string]int{
	SignatureBoolean: 1,
	SignatureByte:    1,
	SignatureChar:    2,
	SignatureShort:   2,
	SignatureInt:     4,
	SignatureLong:    8,
	SignatureFloat:   4,
	SignatureDouble:  8,
	SignatureArray:   8,
	SignatureVoid:    0,
	SignatureClass:   8,
}

type JObject interface {
	Signature() string
	GoValue() interface{}
	JavaValue() CJvalue
}

var funcSignagure = regexp.MustCompile(`\((.*)\)((.).*)`)

func (jvm *JVM) ExceptionCheck() error {
	errExist := (uint8)(C.ExceptionCheck(jvm.env()))
	if errExist != 0 {
		C.ExceptionDescribe(jvm.env())
		return errors.New("JNI Exception")
	}
	return nil
}

func (jvm *JVM) CallStaticFunction(classfqcn, method, sig string, argv []JObject) (JObject, error) {
	cname := C.CString(classfqcn)
	defer C.free(unsafe.Pointer(cname))
	clazz := C.FindClass(jvm.env(), cname)
	if clazz == nil {
		return nil, errors.New("FindClass" + classfqcn)
	}
	defer C.DeleteLocalRef(jvm.env(), clazz)

	cmethod := C.CString(method)
	defer C.free(unsafe.Pointer(cmethod))
	csig := C.CString(sig)
	defer C.free(unsafe.Pointer(csig))
	methodID := C.GetStaticMethodID(jvm.env(), clazz, cmethod, csig)
	C.ExceptionDescribe(jvm.env())

	retsig := funcSignagure.FindStringSubmatch(sig)[3]
	retsigFull := funcSignagure.FindStringSubmatch(sig)[2]

	switch retsig {
	case SignatureBoolean:
		ret := C.CallStaticBooleanMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJPrimitiveFromJava(ret, SignatureBoolean)
	case SignatureByte:
		ret := C.CallStaticByteMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJPrimitiveFromJava(ret, SignatureByte)
	case SignatureChar:
		ret := C.CallStaticCharMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJPrimitiveFromJava(ret, SignatureChar)
	case SignatureShort:
		ret := C.CallStaticShortMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJPrimitiveFromJava(ret, SignatureShort)
	case SignatureInt:
		ret := C.CallStaticIntMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJPrimitiveFromJava(ret, SignatureInt)
	case SignatureLong:
		ret := C.CallStaticLongMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJPrimitiveFromJava(ret, SignatureLong)
	case SignatureFloat:
		ret := C.CallStaticFloatMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJPrimitiveFromJava(ret, SignatureFloat)
	case SignatureDouble:
		ret := C.CallStaticDoubleMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJPrimitiveFromJava(ret, SignatureDouble)
	case SignatureVoid:
		C.CallStaticVoidMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		return nil, nil
	case SignatureArray:
		ret := C.CallStaticObjectMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJArrayFromJava(&ret, retsigFull)
	case SignatureClass:
		ret := C.CallStaticObjectMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJClassFromJava(ret, retsigFull)
	default:
		return nil, errors.New("Unknown return signature")
	}
}

// TODO: use C malloc
func jObjectArrayTojvalueArray(args []JObject) *C.jvalue {
	if len(args) == 0 {
		return nil
	}

	jvalueArray := make([]C.jvalue, len(args))

	for i, arg := range args {
		jvalueArray[i] = arg.JavaValue().jvalue()
	}
	return (*C.jvalue)(unsafe.Pointer(&jvalueArray[0]))
}
