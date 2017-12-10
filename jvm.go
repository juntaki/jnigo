package jnigo

// #cgo linux CFLAGS: -I/usr/lib/jvm/default-java/include -I/usr/lib/jvm/default-java/include/linux
// #cgo linux LDFLAGS: -L/usr/lib/jvm/default-java/jre/lib/amd64/server/ -ljvm -lpthread
// #include "jni_wrapper.h"
import "C"

import (
	"errors"
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
	C.free(unsafe.Pointer(jvm.cjvm))
}

func (jvm *JVM) env() *C.JNIEnv {
	return jvm.cjvm.env
}

func (jvm *JVM) ExceptionCheck() error {
	errExist := (uint8)(C.ExceptionCheck(jvm.env()))
	if errExist != 0 {
		C.ExceptionClear(jvm.env())
		// TODO: Exception details
		return errors.New("Exception")
	}
	return nil
}

func (jvm *JVM) FindMethodID(clazz C.jobject, method, sig string) (C.jmethodID, error) {
	var methodID C.jmethodID

	cmethod := C.CString(method)
	defer C.free(unsafe.Pointer(cmethod))
	csig := C.CString(sig)
	defer C.free(unsafe.Pointer(csig))

	for targetClass := clazz; ; {
		methodID = C.GetMethodID(jvm.env(), targetClass, cmethod, csig)
		if err := jvm.ExceptionCheck(); err != nil {
			return nil, err
		} else {
			break
		}
	}
	return methodID, nil
}

func (jvm *JVM) CallStaticFunction(classfqcn, method, sig string, argv []JObject) (JObject, error) {
	cname := C.CString(classfqcn)
	defer C.free(unsafe.Pointer(cname))
	clazz := C.FindClass(jvm.env(), cname)
	if err := jvm.ExceptionCheck(); err != nil {
		return nil, err
	}
	defer C.DeleteLocalRef(jvm.env(), clazz)

	cmethod := C.CString(method)
	defer C.free(unsafe.Pointer(cmethod))
	csig := C.CString(sig)
	defer C.free(unsafe.Pointer(csig))
	methodID := C.GetStaticMethodID(jvm.env(), clazz, cmethod, csig)
	if err := jvm.ExceptionCheck(); err != nil {
		return nil, err
	}

	retsig := funcSignagure.FindStringSubmatch(sig)[3]
	retsigFull := funcSignagure.FindStringSubmatch(sig)[2]

	switch retsig {
	case SignatureBoolean:
		ret := C.CallStaticBooleanMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		if err := jvm.ExceptionCheck(); err != nil {
			return nil, err
		}
		return jvm.newJPrimitiveFromJava(ret, SignatureBoolean)
	case SignatureByte:
		ret := C.CallStaticByteMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		if err := jvm.ExceptionCheck(); err != nil {
			return nil, err
		}
		return jvm.newJPrimitiveFromJava(ret, SignatureByte)
	case SignatureChar:
		ret := C.CallStaticCharMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		if err := jvm.ExceptionCheck(); err != nil {
			return nil, err
		}
		return jvm.newJPrimitiveFromJava(ret, SignatureChar)
	case SignatureShort:
		ret := C.CallStaticShortMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		if err := jvm.ExceptionCheck(); err != nil {
			return nil, err
		}
		return jvm.newJPrimitiveFromJava(ret, SignatureShort)
	case SignatureInt:
		ret := C.CallStaticIntMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		if err := jvm.ExceptionCheck(); err != nil {
			return nil, err
		}
		return jvm.newJPrimitiveFromJava(ret, SignatureInt)
	case SignatureLong:
		ret := C.CallStaticLongMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		if err := jvm.ExceptionCheck(); err != nil {
			return nil, err
		}
		return jvm.newJPrimitiveFromJava(ret, SignatureLong)
	case SignatureFloat:
		ret := C.CallStaticFloatMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		if err := jvm.ExceptionCheck(); err != nil {
			return nil, err
		}
		return jvm.newJPrimitiveFromJava(ret, SignatureFloat)
	case SignatureDouble:
		ret := C.CallStaticDoubleMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		if err := jvm.ExceptionCheck(); err != nil {
			return nil, err
		}
		return jvm.newJPrimitiveFromJava(ret, SignatureDouble)
	case SignatureVoid:
		C.CallStaticVoidMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		if err := jvm.ExceptionCheck(); err != nil {
			return nil, err
		}
		return nil, nil
	case SignatureArray:
		ret := C.CallStaticObjectMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		if err := jvm.ExceptionCheck(); err != nil {
			return nil, err
		}
		return jvm.newJArrayFromJava(&ret, retsigFull)
	case SignatureClass:
		ret := C.CallStaticObjectMethodA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(argv))
		if err := jvm.ExceptionCheck(); err != nil {
			return nil, err
		}
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
