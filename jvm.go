package jnigo

// #cgo linux CFLAGS: -I/usr/lib/jvm/default-java/include -I/usr/lib/jvm/default-java/include/linux
// #cgo linux LDFLAGS: -L/usr/lib/jvm/default-java/jre/lib/amd64/server/ -ljvm -lpthread
// #include "jni_wrapper.h"
import "C"

import (
	"errors"
	"fmt"
	"regexp"
	"runtime"
	"unsafe"
)

// Global JVM
var jvm *JVM

type JVM struct {
	cjvm *C.JVM
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
	fmt.Println("JVM freed")
	C.free(unsafe.Pointer(jvm.cjvm))
}

// This may not work
func (jvm *JVM) destroyJVM() {
	C.destroyJVM(jvm.cjvm)
}

const (
	SignatureBoolean = "Z"
	SignatureByte    = "B"
	SignatureChar    = "C"
	SignatureShort   = "S"
	SignatureInt     = "I"
	SignatureLong    = "J"
	SignatureFloat   = "F"
	SignatureDouble  = "D"
	SignatureArray   = "["
	SignatureVoid    = "V"
	SignatureClass   = "L"
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
	JavaValue() C.jvalue
}

var funcSignagure = regexp.MustCompile(`\((.*)\)((.).*)`)

func (jvm *JVM) ExceptionCheck() error {
	errExist := C.jboolean_to_uint8(C.ExceptionCheck(jvm.cjvm.env))
	if errExist != 0 {
		C.ExceptionDescribe(jvm.cjvm.env)
		return errors.New("JNI Exception")
	}
	return nil
}

func (jvm *JVM) GetStaticField(classfqcn, field, sig string) (JObject, error) {
	cname := C.CString(classfqcn)
	defer C.free(unsafe.Pointer(cname))
	clazz := C.FindClass(jvm.cjvm.env, cname)
	if clazz == nil {
		return nil, errors.New("FindClass" + classfqcn)
	}

	cfield := C.CString(field)
	defer C.free(unsafe.Pointer(cfield))
	csig := C.CString(sig)
	defer C.free(unsafe.Pointer(csig))
	fieldID := C.GetStaticFieldID(jvm.cjvm.env, clazz, cfield, csig)
	err := jvm.ExceptionCheck()
	if err != nil {
		return nil, err
	}
	switch string(sig[0]) {
	case SignatureBoolean:
		ret := C.GetStaticBooleanField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureBoolean)
	case SignatureByte:
		ret := C.GetStaticByteField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureByte)
	case SignatureChar:
		ret := C.GetStaticCharField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureChar)
	case SignatureShort:
		ret := C.GetStaticShortField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureShort)
	case SignatureInt:
		ret := C.GetStaticIntField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureInt)
	case SignatureLong:
		ret := C.GetStaticLongField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureLong)
	case SignatureFloat:
		ret := C.GetStaticFloatField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureFloat)
	case SignatureDouble:
		ret := C.GetStaticDoubleField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureDouble)
	case SignatureArray:
		ret := C.GetStaticObjectField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJArrayFromJava(&ret, sig)
	case SignatureClass:
		ret := C.GetStaticObjectField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJClassFromJava(ret, sig)
	default:
		return nil, errors.New("Unknown return signature")
	}
}

func (jvm *JVM) SetField(classfqcn, field string, val JObject) error {
	cname := C.CString(classfqcn)
	defer C.free(unsafe.Pointer(cname))
	clazz := C.FindClass(jvm.cjvm.env, cname)
	if clazz == nil {
		return errors.New("FindClass" + classfqcn)
	}

	cfield := C.CString(field)
	defer C.free(unsafe.Pointer(cfield))
	csig := C.CString(val.Signature())
	defer C.free(unsafe.Pointer(csig))
	fieldID := C.GetFieldID(jvm.cjvm.env, clazz, cfield, csig)

	jvalue := val.JavaValue()

	switch string(val.Signature()[0]) {
	case SignatureBoolean:
		C.SetStaticBooleanField(jvm.cjvm.env, clazz, fieldID,
			*C.jvalue_to_jboolean(&jvalue))
	case SignatureByte:
		C.SetStaticByteField(jvm.cjvm.env, clazz, fieldID,
			*C.jvalue_to_jbyte(&jvalue))
	case SignatureChar:
		C.SetStaticCharField(jvm.cjvm.env, clazz, fieldID,
			*C.jvalue_to_jchar(&jvalue))
	case SignatureShort:
		C.SetStaticShortField(jvm.cjvm.env, clazz, fieldID,
			*C.jvalue_to_jshort(&jvalue))
	case SignatureInt:
		C.SetStaticIntField(jvm.cjvm.env, clazz, fieldID,
			*C.jvalue_to_jint(&jvalue))
	case SignatureLong:
		C.SetStaticLongField(jvm.cjvm.env, clazz, fieldID,
			*C.jvalue_to_jlong(&jvalue))
	case SignatureFloat:
		C.SetStaticFloatField(jvm.cjvm.env, clazz, fieldID,
			*C.jvalue_to_jfloat(&jvalue))
	case SignatureDouble:
		C.SetStaticDoubleField(jvm.cjvm.env, clazz, fieldID,
			*C.jvalue_to_jdouble(&jvalue))
	case SignatureArray:
		C.SetStaticObjectField(jvm.cjvm.env, clazz, fieldID,
			*C.jvalue_to_jobject(&jvalue))
	case SignatureClass:
		C.SetStaticObjectField(jvm.cjvm.env, clazz, fieldID,
			*C.jvalue_to_jobject(&jvalue))
	default:
		return errors.New("Unknown return signature")
	}
	return nil
}

func (jvm *JVM) CallStaticFunction(classfqcn, method, sig string, argv []JObject) (JObject, error) {
	cname := C.CString(classfqcn)
	defer C.free(unsafe.Pointer(cname))
	clazz := C.FindClass(jvm.cjvm.env, cname)
	if clazz == nil {
		return nil, errors.New("FindClass" + classfqcn)
	}

	cmethod := C.CString(method)
	defer C.free(unsafe.Pointer(cmethod))
	csig := C.CString(sig)
	defer C.free(unsafe.Pointer(csig))
	methodID := C.GetStaticMethodID(jvm.cjvm.env, clazz, cmethod, csig)
	C.ExceptionDescribe(jvm.cjvm.env)

	retsig := funcSignagure.FindStringSubmatch(sig)[3]
	retsigFull := funcSignagure.FindStringSubmatch(sig)[2]

	switch retsig {
	case SignatureBoolean:
		ret := C.CallStaticBooleanMethodA(jvm.cjvm.env, clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureBoolean)
	case SignatureByte:
		ret := C.CallStaticByteMethodA(jvm.cjvm.env, clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureByte)
	case SignatureChar:
		ret := C.CallStaticCharMethodA(jvm.cjvm.env, clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureChar)
	case SignatureShort:
		ret := C.CallStaticShortMethodA(jvm.cjvm.env, clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureShort)
	case SignatureInt:
		ret := C.CallStaticIntMethodA(jvm.cjvm.env, clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureInt)
	case SignatureLong:
		ret := C.CallStaticLongMethodA(jvm.cjvm.env, clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureLong)
	case SignatureFloat:
		ret := C.CallStaticFloatMethodA(jvm.cjvm.env, clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureFloat)
	case SignatureDouble:
		ret := C.CallStaticDoubleMethodA(jvm.cjvm.env, clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureDouble)
	case SignatureVoid:
		C.CallStaticVoidMethodA(jvm.cjvm.env, clazz, methodID, jObjectArrayTojvalueArray(argv))
		return nil, nil
	case SignatureArray:
		ret := C.CallStaticObjectMethodA(jvm.cjvm.env, clazz, methodID, jObjectArrayTojvalueArray(argv))
		return jvm.newJArrayFromJava(&ret, retsigFull)
	case SignatureClass:
		ret := C.CallStaticObjectMethodA(jvm.cjvm.env, clazz, methodID, jObjectArrayTojvalueArray(argv))
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
		jvalueArray[i] = arg.JavaValue()
	}
	return (*C.jvalue)(unsafe.Pointer(&jvalueArray[0]))
}
