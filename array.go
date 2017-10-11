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
	javavalue *C.jvalue
	signature string
	globalRef C.jobject
}

func (a *jArray) GoValue() interface{} {
	jobject := *C.jvalue_to_jobject(a.javavalue)
	length := int(C.GetArrayLength(a.jvm.cjvm.env, jobject))
	start := C.jsize(0)
	switch a.Signature()[0:2] {
	case SignatureArray + SignatureBoolean:
		value := make([]bool, length)
		buf := C.calloc_jboolean_array(C.size_t(length))
		defer C.free(unsafe.Pointer(buf))
		jbooleanArray := C.jobject_to_jbooleanArray(jobject)
		C.GetBooleanArrayRegion(a.jvm.cjvm.env, jbooleanArray, start, C.jsize(length), buf)
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(buf), C.size_t(length*SizeOf[SignatureBoolean]))
		return value
	case SignatureArray + SignatureByte:
		value := make([]byte, length)
		buf := C.calloc_jbyte_array(C.size_t(length))
		defer C.free(unsafe.Pointer(buf))
		jbyteArray := C.jobject_to_jbyteArray(jobject)
		C.GetByteArrayRegion(a.jvm.cjvm.env, jbyteArray, start, C.jsize(length), buf)
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(buf), C.size_t(length*SizeOf[SignatureByte]))
		return value
	case SignatureArray + SignatureChar:
		value := make([]uint16, length)
		buf := C.calloc_jchar_array(C.size_t(length))
		defer C.free(unsafe.Pointer(buf))
		jcharArray := C.jobject_to_jcharArray(jobject)
		C.GetCharArrayRegion(a.jvm.cjvm.env, jcharArray, start, C.jsize(length), buf)
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(buf), C.size_t(length*SizeOf[SignatureChar]))
		return value
	case SignatureArray + SignatureShort:
		value := make([]int16, length)
		buf := C.calloc_jshort_array(C.size_t(length))
		defer C.free(unsafe.Pointer(buf))
		jshortArray := C.jobject_to_jshortArray(jobject)
		C.GetShortArrayRegion(a.jvm.cjvm.env, jshortArray, start, C.jsize(length), buf)
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(buf), C.size_t(length*SizeOf[SignatureShort]))
		return value
	case SignatureArray + SignatureInt:
		value := make([]int32, length)
		buf := C.calloc_jint_array(C.size_t(length))
		defer C.free(unsafe.Pointer(buf))
		jintArray := C.jobject_to_jintArray(jobject)
		C.GetIntArrayRegion(a.jvm.cjvm.env, jintArray, start, C.jsize(length), buf)
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(buf), C.size_t(length*SizeOf[SignatureInt]))
		return value
	case SignatureArray + SignatureLong:
		value := make([]int64, length)
		buf := C.calloc_jlong_array(C.size_t(length))
		defer C.free(unsafe.Pointer(buf))
		jlongArray := C.jobject_to_jlongArray(jobject)
		C.GetLongArrayRegion(a.jvm.cjvm.env, jlongArray, start, C.jsize(length), buf)
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(buf), C.size_t(length*SizeOf[SignatureLong]))
		return value
	case SignatureArray + SignatureFloat:
		value := make([]float32, length)
		buf := C.calloc_jfloat_array(C.size_t(length))
		defer C.free(unsafe.Pointer(buf))
		jfloatArray := C.jobject_to_jfloatArray(jobject)
		C.GetFloatArrayRegion(a.jvm.cjvm.env, jfloatArray, start, C.jsize(length), buf)
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(buf), C.size_t(length*SizeOf[SignatureFloat]))
		return value
	case SignatureArray + SignatureDouble:
		value := make([]float64, length)
		buf := C.calloc_jdouble_array(C.size_t(length))
		defer C.free(unsafe.Pointer(buf))
		jdoubleArray := C.jobject_to_jdoubleArray(jobject)
		C.GetDoubleArrayRegion(a.jvm.cjvm.env, jdoubleArray, start, C.jsize(length), buf)
		C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(buf), C.size_t(length*SizeOf[SignatureDouble]))
		return value
	case SignatureArray + SignatureClass:
		value := make([]JObject, 0)
		jobjectArray := C.jobject_to_jobjectArray(jobject)
		for i := 0; i < length; i++ {
			jobject := C.GetObjectArrayElement(a.jvm.cjvm.env, jobjectArray, C.jsize(i))
			jclass, _ := a.jvm.NewJClassFromJava(jobject, a.Signature()[1:len(a.Signature())])
			value = append(value, jclass)
		}
		return value
	default:
		return nil
	}
}

func (a *jArray) JavaValue() C.jvalue {
	return *a.javavalue
}

func (a *jArray) String() string {
	return fmt.Sprintf("0x%x", a.JavaValue())
}

func (a *jArray) Signature() string {
	return a.signature
}

func (jvm *JVM) newJArrayFromJava(jobject *C.jobject, sig string) (*jArray, error) {
	ret := &jArray{
		jvm:       jvm,
		javavalue: C.calloc_jvalue_jobject(jobject),
		signature: sig,
		globalRef: C.NewGlobalRef(jvm.cjvm.env, *jobject),
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
		value := C.NewBooleanArray(jvm.cjvm.env, length)
		buf := C.calloc_jboolean_array(C.size_t(len(t)))
		defer C.free(unsafe.Pointer(buf))
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&t[0]), C.size_t(len(t)*SizeOf[SignatureBoolean]))
		C.SetBooleanArrayRegion(jvm.cjvm.env, value, start, length, buf)
		array = C.jbooleanArray_to_jobject(value)
		sig = SignatureArray + SignatureBoolean
	case []byte:
		length := C.jsize(len(t))
		value := C.NewByteArray(jvm.cjvm.env, length)
		buf := C.calloc_jbyte_array(C.size_t(len(t)))
		defer C.free(unsafe.Pointer(buf))
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&t[0]), C.size_t(len(t)*SizeOf[SignatureByte]))
		C.SetByteArrayRegion(jvm.cjvm.env, value, start, length, buf)
		array = C.jbyteArray_to_jobject(value)
		sig = SignatureArray + SignatureByte
	case []uint16:
		length := C.jsize(len(t))
		value := C.NewCharArray(jvm.cjvm.env, length)
		buf := C.calloc_jchar_array(C.size_t(len(t)))
		defer C.free(unsafe.Pointer(buf))
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&t[0]), C.size_t(len(t)*SizeOf[SignatureChar]))
		C.SetCharArrayRegion(jvm.cjvm.env, value, start, length, buf)
		array = C.jcharArray_to_jobject(value)
		sig = SignatureArray + SignatureChar
	case []int16:
		length := C.jsize(len(t))
		value := C.NewShortArray(jvm.cjvm.env, length)
		buf := C.calloc_jshort_array(C.size_t(len(t)))
		defer C.free(unsafe.Pointer(buf))
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&t[0]), C.size_t(len(t)*SizeOf[SignatureShort]))
		C.SetShortArrayRegion(jvm.cjvm.env, value, start, length, buf)
		array = C.jshortArray_to_jobject(value)
		sig = SignatureArray + SignatureShort
	case []int32:
		length := C.jsize(len(t))
		value := C.NewIntArray(jvm.cjvm.env, length)
		buf := C.calloc_jint_array(C.size_t(len(t)))
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&t[0]), C.size_t(len(t)*SizeOf[SignatureInt]))
		defer C.free(unsafe.Pointer(buf))
		C.SetIntArrayRegion(jvm.cjvm.env, value, start, length, buf)
		array = C.jintArray_to_jobject(value)
		sig = SignatureArray + SignatureInt
	case []int64:
		length := C.jsize(len(t))
		value := C.NewLongArray(jvm.cjvm.env, length)
		buf := C.calloc_jlong_array(C.size_t(len(t)))
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&t[0]), C.size_t(len(t)*SizeOf[SignatureLong]))
		defer C.free(unsafe.Pointer(buf))
		C.SetLongArrayRegion(jvm.cjvm.env, value, start, length, buf)
		array = C.jlongArray_to_jobject(value)
		sig = SignatureArray + SignatureLong
	case []float32:
		length := C.jsize(len(t))
		value := C.NewFloatArray(jvm.cjvm.env, length)
		buf := C.calloc_jfloat_array(C.size_t(len(t)))
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&t[0]), C.size_t(len(t)*SizeOf[SignatureFloat]))
		defer C.free(unsafe.Pointer(buf))
		C.SetFloatArrayRegion(jvm.cjvm.env, value, start, length, buf)
		array = C.jfloatArray_to_jobject(value)
		sig = SignatureArray + SignatureFloat
	case []float64:
		length := C.jsize(len(t))
		value := C.NewDoubleArray(jvm.cjvm.env, length)
		buf := C.calloc_jdouble_array(C.size_t(len(t)))
		C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&t[0]), C.size_t(len(t)*SizeOf[SignatureDouble]))
		defer C.free(unsafe.Pointer(buf))
		C.SetDoubleArrayRegion(jvm.cjvm.env, value, start, length, buf)
		array = C.jdoubleArray_to_jobject(value)
		sig = SignatureArray + SignatureDouble
	case []JObject:
		length := C.jsize(len(t))
		if length == 0 {
			panic("not implemented")
		}
		jclass, ok := t[0].(*JClass)
		if !ok {
			return nil, errors.New("unsupported type")
		}
		value := C.NewObjectArray(jvm.cjvm.env, length, jclass.clazz, nil)
		for i, val := range t {
			C.SetObjectArrayElement(jvm.cjvm.env, value, C.jsize(i), *C.jvalue_to_jobject(val.(*JClass).javavalue))
		}
		array = C.jobjectArray_to_jobject(value)
		sig = SignatureArray + t[0].Signature()
	default:
		return nil, errors.New("unsupported type")
	}

	ret := &jArray{
		jvm:       jvm,
		javavalue: C.calloc_jvalue_jobject(&array),
		signature: sig,
		globalRef: C.NewGlobalRef(jvm.cjvm.env, array),
	}

	runtime.SetFinalizer(ret, jvm.destroyjArray)
	return ret, nil
}

func (jvm *JVM) destroyjArray(jobject *jArray) {
	C.DeleteGlobalRef(jvm.cjvm.env, jobject.globalRef)
	C.free(unsafe.Pointer(jobject.javavalue))
}
