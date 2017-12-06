package jnigo

// #include"jni_wrapper.h"
import "C"
import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"
)

const intSize = 32 << (^uint(0) >> 63)

type jPrimitive struct {
	JObject
	javavalue CJvalue
	signature string
}

func (p *jPrimitive) GoValue() interface{} {
	switch p.Signature() {
	case SignatureBoolean:
		var value bool
		C.memcpy(unsafe.Pointer(&value), p.javavalue.unsafePointer(), C.size_t(unsafe.Sizeof(value)))
		return value
	case SignatureByte:
		var value byte
		C.memcpy(unsafe.Pointer(&value), p.javavalue.unsafePointer(), C.size_t(unsafe.Sizeof(value)))
		return value
	case SignatureChar:
		var value uint16
		C.memcpy(unsafe.Pointer(&value), p.javavalue.unsafePointer(), C.size_t(unsafe.Sizeof(value)))
		return value
	case SignatureShort:
		var value int16
		C.memcpy(unsafe.Pointer(&value), p.javavalue.unsafePointer(), C.size_t(unsafe.Sizeof(value)))
		return value
	case SignatureInt:
		var value int32
		C.memcpy(unsafe.Pointer(&value), p.javavalue.unsafePointer(), C.size_t(unsafe.Sizeof(value)))
		return value
	case SignatureLong:
		var value int64
		C.memcpy(unsafe.Pointer(&value), p.javavalue.unsafePointer(), C.size_t(unsafe.Sizeof(value)))
		return value
	case SignatureFloat:
		var value float32
		C.memcpy(unsafe.Pointer(&value), p.javavalue.unsafePointer(), C.size_t(unsafe.Sizeof(value)))
		return value
	case SignatureDouble:
		var value float64
		C.memcpy(unsafe.Pointer(&value), p.javavalue.unsafePointer(), C.size_t(unsafe.Sizeof(value)))
		return value
	case SignatureVoid:
		return nil
	}
	return nil
}

func (p *jPrimitive) JavaValue() CJvalue {
	return p.javavalue
}

func (p *jPrimitive) String() string {
	return fmt.Sprint(p.GoValue())
}

func (p *jPrimitive) Signature() string {
	return p.signature
}

func (jvm *JVM) newJPrimitiveFromJava(initialValue interface{}, sig string) (*jPrimitive, error) {
	javavalue := C.calloc_jvalue()
	var src unsafe.Pointer
	src = unsafe.Pointer(&initialValue)
	switch value := initialValue.(type) {
	case C.jboolean:
		src = unsafe.Pointer(&value)
	case C.jbyte:
		src = unsafe.Pointer(&value)
	case C.jchar:
		src = unsafe.Pointer(&value)
	case C.jshort:
		src = unsafe.Pointer(&value)
	case C.jint:
		src = unsafe.Pointer(&value)
	case C.jlong:
		src = unsafe.Pointer(&value)
	case C.jfloat:
		src = unsafe.Pointer(&value)
	case C.jdouble:
		src = unsafe.Pointer(&value)
	default:
		return nil, errors.New("unknown type")
	}
	C.memcpy(unsafe.Pointer(javavalue), src, C.size_t(SizeOf[sig]))

	ret := &jPrimitive{
		signature: sig,
		javavalue: NewCJvalue(javavalue),
	}
	runtime.SetFinalizer(ret, destroyjPrimitive)
	return ret, nil
}

func (jvm *JVM) newJPrimitive(initialValue interface{}) (*jPrimitive, error) {
	javavalue := C.calloc_jvalue()
	var sig string

	switch value := initialValue.(type) {
	case bool:
		sig = SignatureBoolean
		C.memcpy(unsafe.Pointer(javavalue), unsafe.Pointer(&value), C.size_t(unsafe.Sizeof(value)))
	case byte:
		sig = SignatureByte
		C.memcpy(unsafe.Pointer(javavalue), unsafe.Pointer(&value), C.size_t(unsafe.Sizeof(value)))
	case uint16:
		sig = SignatureChar
		C.memcpy(unsafe.Pointer(javavalue), unsafe.Pointer(&value), C.size_t(unsafe.Sizeof(value)))
	case int16:
		sig = SignatureShort
		C.memcpy(unsafe.Pointer(javavalue), unsafe.Pointer(&value), C.size_t(unsafe.Sizeof(value)))
	case int32:
		sig = SignatureInt
		C.memcpy(unsafe.Pointer(javavalue), unsafe.Pointer(&value), C.size_t(unsafe.Sizeof(value)))
	case int64:
		sig = SignatureLong
		C.memcpy(unsafe.Pointer(javavalue), unsafe.Pointer(&value), C.size_t(unsafe.Sizeof(value)))
	case int:
		if intSize == 64 {
			sig = SignatureLong
			C.memcpy(unsafe.Pointer(javavalue), unsafe.Pointer(&value), C.size_t(unsafe.Sizeof(value)))
		} else {
			sig = SignatureInt
			C.memcpy(unsafe.Pointer(javavalue), unsafe.Pointer(&value), C.size_t(unsafe.Sizeof(value)))
		}
	case float32:
		sig = SignatureFloat
		C.memcpy(unsafe.Pointer(javavalue), unsafe.Pointer(&value), C.size_t(unsafe.Sizeof(value)))
	case float64:
		sig = SignatureDouble
		C.memcpy(unsafe.Pointer(javavalue), unsafe.Pointer(&value), C.size_t(unsafe.Sizeof(value)))
	default:
		C.free(unsafe.Pointer(javavalue))
		return nil, errors.New("unsupported type")
	}
	ret := &jPrimitive{
		signature: sig,
		javavalue: NewCJvalue(javavalue),
	}
	runtime.SetFinalizer(ret, destroyjPrimitive)
	return ret, nil
}

func destroyjPrimitive(jprimitive *jPrimitive) {
	jprimitive.javavalue.free()
}
