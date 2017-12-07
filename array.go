package jnigo

// #include"jni_wrapper.h"
import "C"
import (
	"errors"
	"fmt"
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
		buf := C.calloc_jboolean_array(C.size_t(length))
		defer C.free(unsafe.Pointer(buf))
		jbooleanArray := a.javavalue.jbooleanArray()
		C.GetBooleanArrayRegion(a.jvm.env(), jbooleanArray, start, C.jsize(length), buf)
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(buf), C.size_t(length*SizeOf[SignatureBoolean]))
		return value
	case SignatureByteArray:
		value := make([]byte, length)
		buf := C.calloc_jbyte_array(C.size_t(length))
		defer C.free(unsafe.Pointer(buf))
		jbyteArray := a.javavalue.jbyteArray()
		C.GetByteArrayRegion(a.jvm.env(), jbyteArray, start, C.jsize(length), buf)
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(buf), C.size_t(length*SizeOf[SignatureByte]))
		return value
	case SignatureCharArray:
		value := make([]uint16, length)
		buf := C.calloc_jchar_array(C.size_t(length))
		defer C.free(unsafe.Pointer(buf))
		jcharArray := a.javavalue.jcharArray()
		C.GetCharArrayRegion(a.jvm.env(), jcharArray, start, C.jsize(length), buf)
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(buf), C.size_t(length*SizeOf[SignatureChar]))
		return value
	case SignatureShortArray:
		value := make([]int16, length)
		buf := C.calloc_jshort_array(C.size_t(length))
		defer C.free(unsafe.Pointer(buf))
		jshortArray := a.javavalue.jshortArray()
		C.GetShortArrayRegion(a.jvm.env(), jshortArray, start, C.jsize(length), buf)
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(buf), C.size_t(length*SizeOf[SignatureShort]))
		return value
	case SignatureIntArray:
		value := make([]int32, length)
		buf := C.calloc_jint_array(C.size_t(length))
		defer C.free(unsafe.Pointer(buf))
		jintArray := a.javavalue.jintArray()
		C.GetIntArrayRegion(a.jvm.env(), jintArray, start, C.jsize(length), buf)
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(buf), C.size_t(length*SizeOf[SignatureInt]))
		return value
	case SignatureLongArray:
		value := make([]int64, length)
		buf := C.calloc_jlong_array(C.size_t(length))
		defer C.free(unsafe.Pointer(buf))
		jlongArray := a.javavalue.jlongArray()
		C.GetLongArrayRegion(a.jvm.env(), jlongArray, start, C.jsize(length), buf)
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(buf), C.size_t(length*SizeOf[SignatureLong]))
		return value
	case SignatureFloatArray:
		value := make([]float32, length)
		buf := C.calloc_jfloat_array(C.size_t(length))
		defer C.free(unsafe.Pointer(buf))
		jfloatArray := a.javavalue.jfloatArray()
		C.GetFloatArrayRegion(a.jvm.env(), jfloatArray, start, C.jsize(length), buf)
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(buf), C.size_t(length*SizeOf[SignatureFloat]))
		return value
	case SignatureDoubleArray:
		value := make([]float64, length)
		buf := C.calloc_jdouble_array(C.size_t(length))
		defer C.free(unsafe.Pointer(buf))
		jdoubleArray := a.javavalue.jdoubleArray()
		C.GetDoubleArrayRegion(a.jvm.env(), jdoubleArray, start, C.jsize(length), buf)
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(buf), C.size_t(length*SizeOf[SignatureDouble]))
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

	switch t := goArray.(type) {
	case []bool:
		length := C.jsize(len(t))
		value := C.NewBooleanArray(jvm.env(), length)
		buf := C.calloc_jboolean_array(C.size_t(len(t)))
		defer C.free(unsafe.Pointer(buf))
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&t[0]), C.size_t(len(t)*SizeOf[SignatureBoolean]))
		C.SetBooleanArrayRegion(jvm.env(), value, start, length, buf)
		array = C.jbooleanArray_to_jobject(value)
		sig = SignatureBooleanArray
	case []byte:
		length := C.jsize(len(t))
		value := C.NewByteArray(jvm.env(), length)
		buf := C.calloc_jbyte_array(C.size_t(len(t)))
		defer C.free(unsafe.Pointer(buf))
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&t[0]), C.size_t(len(t)*SizeOf[SignatureByte]))
		C.SetByteArrayRegion(jvm.env(), value, start, length, buf)
		array = C.jbyteArray_to_jobject(value)
		sig = SignatureByteArray
	case []uint16:
		length := C.jsize(len(t))
		value := C.NewCharArray(jvm.env(), length)
		buf := C.calloc_jchar_array(C.size_t(len(t)))
		defer C.free(unsafe.Pointer(buf))
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&t[0]), C.size_t(len(t)*SizeOf[SignatureChar]))
		C.SetCharArrayRegion(jvm.env(), value, start, length, buf)
		array = C.jcharArray_to_jobject(value)
		sig = SignatureCharArray
	case []int16:
		length := C.jsize(len(t))
		value := C.NewShortArray(jvm.env(), length)
		buf := C.calloc_jshort_array(C.size_t(len(t)))
		defer C.free(unsafe.Pointer(buf))
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&t[0]), C.size_t(len(t)*SizeOf[SignatureShort]))
		C.SetShortArrayRegion(jvm.env(), value, start, length, buf)
		array = C.jshortArray_to_jobject(value)
		sig = SignatureShortArray
	case []int32:
		length := C.jsize(len(t))
		value := C.NewIntArray(jvm.env(), length)
		buf := C.calloc_jint_array(C.size_t(len(t)))
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&t[0]), C.size_t(len(t)*SizeOf[SignatureInt]))
		defer C.free(unsafe.Pointer(buf))
		C.SetIntArrayRegion(jvm.env(), value, start, length, buf)
		array = C.jintArray_to_jobject(value)
		sig = SignatureIntArray
	case []int64:
		length := C.jsize(len(t))
		value := C.NewLongArray(jvm.env(), length)
		buf := C.calloc_jlong_array(C.size_t(len(t)))
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&t[0]), C.size_t(len(t)*SizeOf[SignatureLong]))
		defer C.free(unsafe.Pointer(buf))
		C.SetLongArrayRegion(jvm.env(), value, start, length, buf)
		array = C.jlongArray_to_jobject(value)
		sig = SignatureLongArray
	case []float32:
		length := C.jsize(len(t))
		value := C.NewFloatArray(jvm.env(), length)
		buf := C.calloc_jfloat_array(C.size_t(len(t)))
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&t[0]), C.size_t(len(t)*SizeOf[SignatureFloat]))
		defer C.free(unsafe.Pointer(buf))
		C.SetFloatArrayRegion(jvm.env(), value, start, length, buf)
		array = C.jfloatArray_to_jobject(value)
		sig = SignatureFloatArray
	case []float64:
		length := C.jsize(len(t))
		value := C.NewDoubleArray(jvm.env(), length)
		buf := C.calloc_jdouble_array(C.size_t(len(t)))
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&t[0]), C.size_t(len(t)*SizeOf[SignatureDouble]))
		defer C.free(unsafe.Pointer(buf))
		C.SetDoubleArrayRegion(jvm.env(), value, start, length, buf)
		array = C.jdoubleArray_to_jobject(value)
		sig = SignatureDoubleArray
	case []JObject:
		length := C.jsize(len(t))
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
		array = C.jobjectArray_to_jobject(value)
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
