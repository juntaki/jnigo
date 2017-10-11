package jnigo

// #include"jni_wrapper.h"
import "C"
import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"
)

type JClass struct {
	JObject
	jvm       *JVM
	javavalue *C.jvalue
	signature string
	globalRef C.jobject
	clazz     C.jclass
}

func (c *JClass) GetField(field, sig string) (JObject, error) {
	cfield := C.CString(field)
	defer C.free(unsafe.Pointer(cfield))
	csig := C.CString(sig)
	defer C.free(unsafe.Pointer(csig))
	fieldID := C.GetFieldID(c.jvm.cjvm.env, c.clazz, cfield, csig)

	switch string(sig[0]) {
	case SignatureBoolean:
		ret := C.GetBooleanField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID)
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureBoolean)
	case SignatureByte:
		ret := C.GetByteField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID)
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureByte)
	case SignatureChar:
		ret := C.GetCharField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID)
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureChar)
	case SignatureShort:
		ret := C.GetShortField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID)
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureShort)
	case SignatureInt:
		ret := C.GetIntField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID)
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureInt)
	case SignatureLong:
		ret := C.GetLongField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID)
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureLong)
	case SignatureFloat:
		ret := C.GetFloatField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID)
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureFloat)
	case SignatureDouble:
		ret := C.GetDoubleField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID)
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureDouble)
	case SignatureArray:
		ret := C.GetObjectField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID)
		return c.jvm.newJArrayFromJava(&ret, sig)
	case SignatureClass:
		ret := C.GetObjectField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID)
		return c.jvm.newJClassFromJava(ret, sig)
	default:
		return nil, errors.New("Unknown return signature")
	}
}

func (c *JClass) SetField(field string, val JObject) error {
	cfield := C.CString(field)
	defer C.free(unsafe.Pointer(cfield))
	csig := C.CString(val.Signature())
	defer C.free(unsafe.Pointer(csig))
	fieldID := C.GetFieldID(c.jvm.cjvm.env, c.clazz, cfield, csig)

	jvalue := val.JavaValue()

	switch string(val.Signature()[0]) {
	case SignatureBoolean:
		C.SetBooleanField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID,
			*C.jvalue_to_jboolean(&jvalue))
	case SignatureByte:
		C.SetByteField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID,
			*C.jvalue_to_jbyte(&jvalue))
	case SignatureChar:
		C.SetCharField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID,
			*C.jvalue_to_jchar(&jvalue))
	case SignatureShort:
		C.SetShortField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID,
			*C.jvalue_to_jshort(&jvalue))
	case SignatureInt:
		C.SetIntField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID,
			*C.jvalue_to_jint(&jvalue))
	case SignatureLong:
		C.SetLongField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID,
			*C.jvalue_to_jlong(&jvalue))
	case SignatureFloat:
		C.SetFloatField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID,
			*C.jvalue_to_jfloat(&jvalue))
	case SignatureDouble:
		C.SetDoubleField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID,
			*C.jvalue_to_jdouble(&jvalue))
	case SignatureArray:
		C.SetObjectField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID,
			*C.jvalue_to_jobject(&jvalue))
	case SignatureClass:
		C.SetObjectField(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue), fieldID,
			*C.jvalue_to_jobject(&jvalue))
	default:
		return errors.New("Unknown return signature")
	}
	return nil
}

func (c *JClass) CallFunction(method, sig string, argv []JObject) (JObject, error) {
	cmethod := C.CString(method)
	defer C.free(unsafe.Pointer(cmethod))
	csig := C.CString(sig)
	defer C.free(unsafe.Pointer(csig))

	methodID := C.GetMethodID(c.jvm.cjvm.env, c.clazz, cmethod, csig)
	C.ExceptionDescribe(c.jvm.cjvm.env)

	retsig := funcSignagure.FindStringSubmatch(sig)[3]
	retsigFull := funcSignagure.FindStringSubmatch(sig)[2]

	switch retsig {
	case SignatureBoolean:
		ret := C.CallBooleanMethodA(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureBoolean)
	case SignatureByte:
		ret := C.CallByteMethodA(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureByte)
	case SignatureChar:
		ret := C.CallCharMethodA(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureChar)
	case SignatureShort:
		ret := C.CallShortMethodA(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureShort)
	case SignatureInt:
		ret := C.CallIntMethodA(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureInt)
	case SignatureLong:
		ret := C.CallLongMethodA(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureLong)
	case SignatureFloat:
		ret := C.CallFloatMethodA(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureFloat)
	case SignatureDouble:
		ret := C.CallDoubleMethodA(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureDouble)
	case SignatureVoid:
		C.CallVoidMethodA(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue),
			methodID, jObjectArrayTojvalueArray(argv))
		return nil, nil
	case SignatureArray:
		ret := C.CallObjectMethodA(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJArrayFromJava(&ret, retsigFull)
	case SignatureClass:
		ret := C.CallObjectMethodA(c.jvm.cjvm.env, *C.jvalue_to_jobject(c.javavalue),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJClassFromJava(ret, retsigFull)
	default:
		return nil, errors.New("Unknown return signature")
	}
}

func (c *JClass) GoValue() interface{} {
	return c
}

func (c *JClass) JavaValue() C.jvalue {
	return *c.javavalue
}

func (c *JClass) String() string {
	return fmt.Sprintf("0x%x", c.JavaValue())
}

func (c *JClass) Signature() string {
	return c.signature
}

func (jvm *JVM) newJClassFromJava(jobject C.jobject, sig string) (*JClass, error) {
	ret := &JClass{
		jvm:       jvm,
		javavalue: C.calloc_jvalue_jobject(&jobject),
		signature: sig,
		globalRef: C.NewGlobalRef(jvm.cjvm.env, jobject),
	}

	fqcn := sig[1 : len(sig)-1]
	cname := C.CString(fqcn)
	defer C.free(unsafe.Pointer(cname))
	clazz := C.FindClass(jvm.cjvm.env, cname)
	if clazz == nil {
		return nil, errors.New("FindClass" + fqcn)
	}

	runtime.SetFinalizer(ret, jvm.destroyJClass)
	return ret, nil
}

func (jvm *JVM) NewJClass(fqcn string, args []JObject) (*JClass, error) {
	cname := C.CString(fqcn)
	defer C.free(unsafe.Pointer(cname))
	cinit := C.CString("<init>")
	defer C.free(unsafe.Pointer(cinit))
	csig := C.CString("()V")
	defer C.free(unsafe.Pointer(csig))
	clazz := C.FindClass(jvm.cjvm.env, cname)
	if clazz == nil {
		return nil, errors.New("FindClass" + fqcn)
	}
	methodID := C.GetMethodID(jvm.cjvm.env, clazz, cinit, csig)
	obj := C.NewObjectA(jvm.cjvm.env, clazz, methodID, jObjectArrayTojvalueArray(args))
	C.ExceptionDescribe(jvm.cjvm.env)
	ret := &JClass{
		jvm:       jvm,
		javavalue: C.calloc_jvalue_jobject(&obj),
		signature: "L" + fqcn + ";",
		globalRef: C.NewGlobalRef(jvm.cjvm.env, obj),
		clazz:     clazz,
	}

	runtime.SetFinalizer(ret, jvm.destroyJClass)
	return ret, nil
}

func (jvm *JVM) destroyJClass(jobject *JClass) {
	C.DeleteGlobalRef(jvm.cjvm.env, jobject.globalRef)
	C.free(unsafe.Pointer(jobject.javavalue))
}
