package jnigo

// #include"jni_wrapper.h"
import "C"
import "unsafe"

func NewCJvalue(javavalue *C.jvalue) CJvalue {
	return CJvalue{
		javavalue: javavalue,
	}
}

// CJvalue wraps C.jvalue for type conversion
type CJvalue struct {
	javavalue *C.jvalue
}

// unsafe
func (v CJvalue) unsafePointer() unsafe.Pointer {
	return unsafe.Pointer(v.javavalue)
}
func (v CJvalue) free() {
	C.free(v.unsafePointer())
}

// Raw value
func (v CJvalue) jvalue() C.jvalue {
	return *v.javavalue
}

// Type cast
func (v CJvalue) jobject() C.jobject {
	return *(*C.jobject)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jboolean() C.jboolean {
	return *(*C.jboolean)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jbyte() C.jbyte {
	return *(*C.jbyte)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jchar() C.jchar {
	return *(*C.jchar)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jshort() C.jshort {
	return *(*C.jshort)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jint() C.jint {
	return *(*C.jint)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jlong() C.jlong {
	return *(*C.jlong)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jfloat() C.jfloat {
	return *(*C.jfloat)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jdouble() C.jdouble {
	return *(*C.jdouble)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jstring() C.jstring {
	return *(*C.jstring)(unsafe.Pointer(v.javavalue))
}

func (v CJvalue) jobjectArray() C.jobjectArray {
	return *(*C.jobjectArray)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jbooleanArray() C.jbooleanArray {
	return *(*C.jbooleanArray)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jbyteArray() C.jbyteArray {
	return *(*C.jbyteArray)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jcharArray() C.jcharArray {
	return *(*C.jcharArray)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jshortArray() C.jshortArray {
	return *(*C.jshortArray)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jintArray() C.jintArray {
	return *(*C.jintArray)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jlongArray() C.jlongArray {
	return *(*C.jlongArray)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jfloatArray() C.jfloatArray {
	return *(*C.jfloatArray)(unsafe.Pointer(v.javavalue))
}
func (v CJvalue) jdoubleArray() C.jdoubleArray {
	return *(*C.jdoubleArray)(unsafe.Pointer(v.javavalue))
}
