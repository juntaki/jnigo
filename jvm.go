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

func (jvm *JVM) GetStaticField(classfqcn, field, sig string) (JObject, error) {
	cname := C.CString(classfqcn)
	defer C.free(unsafe.Pointer(cname))
	clazz := C.FindClass(jvm.env(), cname)
	if err := jvm.ExceptionCheck(); err != nil {
		return nil, errors.New("class not found: " + classfqcn)
	}

	cfield := C.CString(field)
	defer C.free(unsafe.Pointer(cfield))
	csig := C.CString(sig)
	defer C.free(unsafe.Pointer(csig))

	fieldID := C.GetStaticFieldID(jvm.env(), clazz, cfield, csig)
	if err := jvm.ExceptionCheck(); err != nil {
		return nil, errors.New("field not found: " + field + sig)
	}

	switch string(sig[0]) {
	case SignatureBoolean:
		ret := C.GetStaticBooleanField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJPrimitiveFromJava(ret, SignatureBoolean)
	case SignatureByte:
		ret := C.GetStaticByteField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJPrimitiveFromJava(ret, SignatureByte)
	case SignatureChar:
		ret := C.GetStaticCharField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJPrimitiveFromJava(ret, SignatureChar)
	case SignatureShort:
		ret := C.GetStaticShortField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJPrimitiveFromJava(ret, SignatureShort)
	case SignatureInt:
		ret := C.GetStaticIntField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJPrimitiveFromJava(ret, SignatureInt)
	case SignatureLong:
		ret := C.GetStaticLongField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJPrimitiveFromJava(ret, SignatureLong)
	case SignatureFloat:
		ret := C.GetStaticFloatField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJPrimitiveFromJava(ret, SignatureFloat)
	case SignatureDouble:
		ret := C.GetStaticDoubleField(jvm.cjvm.env, clazz, fieldID)
		return jvm.newJPrimitiveFromJava(ret, SignatureDouble)
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
			jvalue.jboolean())
	case SignatureByte:
		C.SetStaticByteField(jvm.cjvm.env, clazz, fieldID,
			jvalue.jbyte())
	case SignatureChar:
		C.SetStaticCharField(jvm.cjvm.env, clazz, fieldID,
			jvalue.jchar())
	case SignatureShort:
		C.SetStaticShortField(jvm.cjvm.env, clazz, fieldID,
			jvalue.jshort())
	case SignatureInt:
		C.SetStaticIntField(jvm.cjvm.env, clazz, fieldID,
			jvalue.jint())
	case SignatureLong:
		C.SetStaticLongField(jvm.cjvm.env, clazz, fieldID,
			jvalue.jlong())
	case SignatureFloat:
		C.SetStaticFloatField(jvm.cjvm.env, clazz, fieldID,
			jvalue.jfloat())
	case SignatureDouble:
		C.SetStaticDoubleField(jvm.cjvm.env, clazz, fieldID,
			jvalue.jdouble())
	case SignatureArray:
		C.SetStaticObjectField(jvm.cjvm.env, clazz, fieldID,
			jvalue.jobject())
	case SignatureClass:
		C.SetStaticObjectField(jvm.cjvm.env, clazz, fieldID,
			jvalue.jobject())
	default:
		return errors.New("Unknown return signature")
	}
	return nil
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
