package jnigo

// #include"jni_wrapper.h"
import "C"
import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"
)

type JPrimitive struct {
	JObject
	javavalue *C.jvalue
	signature string
}

func (p *JPrimitive) GoValue() interface{} {
	switch p.Signature() {
	case SignatureBoolean:
		var value bool
		C.memcpy(unsafe.Pointer(&value), unsafe.Pointer(p.javavalue), C.size_t(unsafe.Sizeof(value)))
		return value
	case SignatureByte:
		var value byte
		C.memcpy(unsafe.Pointer(&value), unsafe.Pointer(p.javavalue), C.size_t(unsafe.Sizeof(value)))
		return value
	case SignatureChar:
		var value uint16
		C.memcpy(unsafe.Pointer(&value), unsafe.Pointer(p.javavalue), C.size_t(unsafe.Sizeof(value)))
		return value
	case SignatureShort:
		var value int16
		C.memcpy(unsafe.Pointer(&value), unsafe.Pointer(p.javavalue), C.size_t(unsafe.Sizeof(value)))
		return value
	case SignatureInt:
		var value int32
		C.memcpy(unsafe.Pointer(&value), unsafe.Pointer(p.javavalue), C.size_t(unsafe.Sizeof(value)))
		return value
	case SignatureLong:
		var value int64
		C.memcpy(unsafe.Pointer(&value), unsafe.Pointer(p.javavalue), C.size_t(unsafe.Sizeof(value)))
		return value
	case SignatureFloat:
		var value float32
		C.memcpy(unsafe.Pointer(&value), unsafe.Pointer(p.javavalue), C.size_t(unsafe.Sizeof(value)))
		return value
	case SignatureDouble:
		var value float64
		C.memcpy(unsafe.Pointer(&value), unsafe.Pointer(p.javavalue), C.size_t(unsafe.Sizeof(value)))
		return value
	case SignatureVoid:
		return nil
	}
	return nil
}

func (p *JPrimitive) JavaValue() C.jvalue {
	return *p.javavalue
}

func (p *JPrimitive) String() string {
	return fmt.Sprintf("0x%x", p.JavaValue())
}

func (p *JPrimitive) Signature() string {
	return p.signature
}

func (jvm *JVM) NewJPrimitiveFromJava(jinitialValue unsafe.Pointer, sig string) (*JPrimitive, error) {
	javavalue := C.calloc_jvalue()
	C.memcpy(unsafe.Pointer(javavalue), jinitialValue, C.size_t(SizeOf[sig]))

	ret := &JPrimitive{
		signature: sig,
		javavalue: javavalue,
	}
	runtime.SetFinalizer(ret, destroyJPrimitive)
	return ret, nil
}

func (jvm *JVM) NewJPrimitive(initialValue interface{}) (*JPrimitive, error) {
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
	ret := &JPrimitive{
		signature: sig,
		javavalue: javavalue,
	}
	runtime.SetFinalizer(ret, destroyJPrimitive)
	return ret, nil
}

func destroyJPrimitive(jprimitive *JPrimitive) {
	C.free(unsafe.Pointer(jprimitive.javavalue))
}

func assign(dest, src unsafe.Pointer) {
	*(*unsafe.Pointer)(dest) = *(*unsafe.Pointer)(src)
}
