package jnigo

// #include"jni_wrapper.h"
import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

type jArray struct {
	JObject
	jvm       *JVM
	javavalue CJvalue
	signature string
	globalRef C.jobject
}

func (a *jArray) GoValue() interface{} {
	length := int(C.GetArrayLength(a.jvm.env(), a.javavalue.jobject()))
	start := C.jsize(0)
	switch a.Signature()[0:2] {
	case SignatureBooleanArray:
		value := make([]bool, length)
		if length != 0 {
			jbooleanArray := a.javavalue.jbooleanArray()
			C.GetBooleanArrayRegion(a.jvm.env(), jbooleanArray, start, C.jsize(length),
				(*C.jboolean)(unsafe.Pointer(&value[0])))
		}
		return value
	case SignatureByteArray:
		value := make([]byte, length)
		if length != 0 {
			jbyteArray := a.javavalue.jbyteArray()
			C.GetByteArrayRegion(a.jvm.env(), jbyteArray, start, C.jsize(length),
				(*C.jbyte)(unsafe.Pointer(&value[0])))
		}
		return value
	case SignatureCharArray:
		value := make([]uint16, length)
		if length != 0 {
			jcharArray := a.javavalue.jcharArray()
			C.GetCharArrayRegion(a.jvm.env(), jcharArray, start, C.jsize(length),
				(*C.jchar)(unsafe.Pointer(&value[0])))
		}
		return value
	case SignatureShortArray:
		value := make([]int16, length)
		if length != 0 {
			jshortArray := a.javavalue.jshortArray()
			C.GetShortArrayRegion(a.jvm.env(), jshortArray, start, C.jsize(length),
				(*C.jshort)(unsafe.Pointer(&value[0])))
		}
		return value
	case SignatureIntArray:
		value := make([]int32, length)
		if length != 0 {
			jintArray := a.javavalue.jintArray()
			C.GetIntArrayRegion(a.jvm.env(), jintArray, start, C.jsize(length),
				(*C.jint)(unsafe.Pointer(&value[0])))
		}
		return value
	case SignatureLongArray:
		value := make([]int64, length)
		if length != 0 {
			jlongArray := a.javavalue.jlongArray()
			C.GetLongArrayRegion(a.jvm.env(), jlongArray, start, C.jsize(length),
				(*C.jlong)(unsafe.Pointer(&value[0])))
		}
		return value
	case SignatureFloatArray:
		value := make([]float32, length)
		if length != 0 {
			jfloatArray := a.javavalue.jfloatArray()
			C.GetFloatArrayRegion(a.jvm.env(), jfloatArray, start, C.jsize(length),
				(*C.jfloat)(unsafe.Pointer(&value[0])))
		}
		return value
	case SignatureDoubleArray:
		value := make([]float64, length)
		if length != 0 {
			jdoubleArray := a.javavalue.jdoubleArray()
			C.GetDoubleArrayRegion(a.jvm.env(), jdoubleArray, start, C.jsize(length),
				(*C.jdouble)(unsafe.Pointer(&value[0])))
		}
		return value
	case SignatureClassArray:
		value := make([]JObject, 0)
		jobjectArray := a.javavalue.jobjectArray()
		for i := 0; i < length; i++ {
			jobject := C.GetObjectArrayElement(a.jvm.env(), jobjectArray, C.jsize(i))
			jclass, _ := a.jvm.newJClassFromJava(jobject, a.Signature()[1:len(a.Signature())])
			value = append(value, jclass)
		}
		return value
	default:
		return nil
	}
}

func (a *jArray) JavaValue() CJvalue {
	return a.javavalue
}

func (a *jArray) String() string {
	return fmt.Sprint(a.GoValue())
}

func (a *jArray) Signature() string {
	return a.signature
}

func (jvm *JVM) newJArrayFromJava(array *C.jobject, sig string) (*jArray, error) {
	defer C.DeleteLocalRef(jvm.env(), *array)
	ref := C.NewGlobalRef(jvm.env(), *array)
	ret := &jArray{
		jvm:       jvm,
		javavalue: NewCJvalue(C.calloc_jvalue_jobject(&ref)),
		signature: sig,
		globalRef: ref,
	}

	runtime.SetFinalizer(ret, jvm.destroyjArray)
	return ret, nil
}

func (jvm *JVM) newJArray(goArray interface{}) (*jArray, error) {
	start := C.jsize(0)
	var array C.jobject
	var sig string

	length := C.jsize(reflect.ValueOf(goArray).Len())
	switch t := goArray.(type) {
	case []bool:
		sig = SignatureBooleanArray
		value := C.NewBooleanArray(jvm.env(), length)
		if length != 0 {
			C.SetBooleanArrayRegion(jvm.env(), value, start, length, (*C.jboolean)(unsafe.Pointer(&t[0])))
		}
		array = C.jobject_conv(value)
	case []byte:
		sig = SignatureByteArray
		value := C.NewByteArray(jvm.env(), length)
		if length != 0 {
			C.SetByteArrayRegion(jvm.env(), value, start, length, (*C.jbyte)(unsafe.Pointer(&t[0])))
		}
		array = C.jobject_conv(value)
	case []uint16:
		sig = SignatureCharArray
		value := C.NewCharArray(jvm.env(), length)
		if length != 0 {
			C.SetCharArrayRegion(jvm.env(), value, start, length, (*C.jchar)(unsafe.Pointer(&t[0])))
		}
		array = C.jobject_conv(value)
	case []int16:
		sig = SignatureShortArray
		value := C.NewShortArray(jvm.env(), length)
		if length != 0 {
			C.SetShortArrayRegion(jvm.env(), value, start, length, (*C.jshort)(unsafe.Pointer(&t[0])))
		}
		array = C.jobject_conv(value)
	case []int32:
		sig = SignatureIntArray
		value := C.NewIntArray(jvm.env(), length)
		if length != 0 {
			C.SetIntArrayRegion(jvm.env(), value, start, length, (*C.jint)(unsafe.Pointer(&t[0])))
		}
		array = C.jobject_conv(value)
	case []int64:
		sig = SignatureLongArray
		value := C.NewLongArray(jvm.env(), length)
		if length != 0 {
			C.SetLongArrayRegion(jvm.env(), value, start, length, (*C.jlong)(unsafe.Pointer(&t[0])))
		}
		array = C.jobject_conv(value)
	case []float32:
		sig = SignatureFloatArray
		value := C.NewFloatArray(jvm.env(), length)
		if length != 0 {
			C.SetFloatArrayRegion(jvm.env(), value, start, length, (*C.jfloat)(unsafe.Pointer(&t[0])))
		}
		array = C.jobject_conv(value)
	case []float64:
		sig = SignatureDoubleArray
		value := C.NewDoubleArray(jvm.env(), length)
		if length != 0 {
			C.SetDoubleArrayRegion(jvm.env(), value, start, length, (*C.jdouble)(unsafe.Pointer(&t[0])))
		}
		array = C.jobject_conv(value)
	case []JObject:
		if length == 0 {
			panic("not implemented")
		}
		jclass, ok := t[0].(*JClass)
		if !ok {
			return nil, errors.New("unsupported type")
		}
		value := C.NewObjectArray(jvm.env(), length, jclass.clazz, nil)
		for i, val := range t {
			C.SetObjectArrayElement(jvm.env(), value, C.jsize(i), val.JavaValue().jobject())
		}
		array = C.jobject_conv(value)
		sig = SignatureArray + t[0].Signature()
	default:
		return nil, errors.New("unsupported type")
	}

	defer C.DeleteLocalRef(jvm.env(), array)
	ref := C.NewGlobalRef(jvm.env(), array)
	ret := &jArray{
		jvm:       jvm,
		javavalue: NewCJvalue(C.calloc_jvalue_jobject(&ref)),
		signature: sig,
		globalRef: ref,
	}
	runtime.SetFinalizer(ret, jvm.destroyjArray)
	return ret, nil
}

func (jvm *JVM) destroyjArray(jobject *jArray) {
	C.DeleteGlobalRef(jvm.env(), jobject.globalRef)
	jobject.javavalue.free()
}
