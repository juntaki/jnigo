package jnigo

// #include"jni_wrapper.h"
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

type JClass struct {
	JObject
	jvm       *JVM
	javavalue CJvalue
	signature string
	globalRef C.jobject
	clazz     C.jclass
}

func (c *JClass) GetField(field, sig string) (JObject, error) {
	cfield := C.CString(field)
	defer C.free(unsafe.Pointer(cfield))
	csig := C.CString(sig)
	defer C.free(unsafe.Pointer(csig))
	fieldID := C.GetFieldID(c.jvm.env(), c.clazz, cfield, csig)

	switch string(sig[0]) {
	case SignatureBoolean:
		ret := C.GetBooleanField(c.jvm.env(), c.javavalue.jobject(), fieldID)
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureBoolean)
	case SignatureByte:
		ret := C.GetByteField(c.jvm.env(), c.javavalue.jobject(), fieldID)
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureByte)
	case SignatureChar:
		ret := C.GetCharField(c.jvm.env(), c.javavalue.jobject(), fieldID)
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureChar)
	case SignatureShort:
		ret := C.GetShortField(c.jvm.env(), c.javavalue.jobject(), fieldID)
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureShort)
	case SignatureInt:
		ret := C.GetIntField(c.jvm.env(), c.javavalue.jobject(), fieldID)
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureInt)
	case SignatureLong:
		ret := C.GetLongField(c.jvm.env(), c.javavalue.jobject(), fieldID)
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureLong)
	case SignatureFloat:
		ret := C.GetFloatField(c.jvm.env(), c.javavalue.jobject(), fieldID)
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureFloat)
	case SignatureDouble:
		ret := C.GetDoubleField(c.jvm.env(), c.javavalue.jobject(), fieldID)
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureDouble)
	case SignatureArray:
		ret := C.GetObjectField(c.jvm.env(), c.javavalue.jobject(), fieldID)
		return c.jvm.newJArrayFromJava(&ret, sig)
	case SignatureClass:
		ret := C.GetObjectField(c.jvm.env(), c.javavalue.jobject(), fieldID)
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
	fieldID := C.GetFieldID(c.jvm.env(), c.clazz, cfield, csig)

	jvalue := val.JavaValue()

	switch string(val.Signature()[0]) {
	case SignatureBoolean:
		C.SetBooleanField(c.jvm.env(), c.javavalue.jobject(), fieldID, jvalue.jboolean())
	case SignatureByte:
		C.SetByteField(c.jvm.env(), c.javavalue.jobject(), fieldID, jvalue.jbyte())
	case SignatureChar:
		C.SetCharField(c.jvm.env(), c.javavalue.jobject(), fieldID, jvalue.jchar())
	case SignatureShort:
		C.SetShortField(c.jvm.env(), c.javavalue.jobject(), fieldID, jvalue.jshort())
	case SignatureInt:
		C.SetIntField(c.jvm.env(), c.javavalue.jobject(), fieldID, jvalue.jint())
	case SignatureLong:
		C.SetLongField(c.jvm.env(), c.javavalue.jobject(), fieldID, jvalue.jlong())
	case SignatureFloat:
		C.SetFloatField(c.jvm.env(), c.javavalue.jobject(), fieldID, jvalue.jfloat())
	case SignatureDouble:
		C.SetDoubleField(c.jvm.env(), c.javavalue.jobject(), fieldID, jvalue.jdouble())
	case SignatureArray:
		C.SetObjectField(c.jvm.env(), c.javavalue.jobject(), fieldID, jvalue.jobject())
	case SignatureClass:
		C.SetObjectField(c.jvm.env(), c.javavalue.jobject(), fieldID, jvalue.jobject())
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

	methodID := C.GetMethodID(c.jvm.env(), c.clazz, cmethod, csig)

	C.ExceptionDescribe(c.jvm.env())

	retsig := funcSignagure.FindStringSubmatch(sig)[3]
	retsigFull := funcSignagure.FindStringSubmatch(sig)[2]

	switch retsig {
	case SignatureBoolean:
		ret := C.CallBooleanMethodA(c.jvm.env(), c.javavalue.jobject(),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureBoolean)
	case SignatureByte:
		ret := C.CallByteMethodA(c.jvm.env(), c.javavalue.jobject(),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureByte)
	case SignatureChar:
		ret := C.CallCharMethodA(c.jvm.env(), c.javavalue.jobject(),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureChar)
	case SignatureShort:
		ret := C.CallShortMethodA(c.jvm.env(), c.javavalue.jobject(),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureShort)
	case SignatureInt:
		ret := C.CallIntMethodA(c.jvm.env(), c.javavalue.jobject(),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureInt)
	case SignatureLong:
		ret := C.CallLongMethodA(c.jvm.env(), c.javavalue.jobject(),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureLong)
	case SignatureFloat:
		ret := C.CallFloatMethodA(c.jvm.env(), c.javavalue.jobject(),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureFloat)
	case SignatureDouble:
		ret := C.CallDoubleMethodA(c.jvm.env(), c.javavalue.jobject(),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJPrimitiveFromJava(unsafe.Pointer(&ret), SignatureDouble)
	case SignatureVoid:
		C.CallVoidMethodA(c.jvm.env(), c.javavalue.jobject(),
			methodID, jObjectArrayTojvalueArray(argv))
		return nil, nil
	case SignatureArray:
		ret := C.CallObjectMethodA(c.jvm.env(), c.javavalue.jobject(),
			methodID, jObjectArrayTojvalueArray(argv))
		return c.jvm.newJArrayFromJava(&ret, retsigFull)
	case SignatureClass:
		ret := C.CallObjectMethodA(c.jvm.env(), c.javavalue.jobject(),
			methodID, jObjectArrayTojvalueArray(argv))
		if retsigFull == "Ljava/lang/String;" {
			return c.jvm.newjStringFromJava(ret)
		}
		return c.jvm.newJClassFromJava(ret, retsigFull)
	default:
		return nil, errors.New("Unknown return signature")
	}
}

func (c *JClass) GoValue() interface{} {
	return c
}

func (c *JClass) JavaValue() CJvalue {
	return c.javavalue
}

func (c *JClass) JValue() CJvalue {
	return c.javavalue
}

func (c *JClass) String() string {
	val, err := c.CallFunction("toString", "()Ljava/lang/String;", []JObject{})
	if err != nil {
		return err.Error()
	}
	return val.GoValue().(string)
}

func (c *JClass) Signature() string {
	return c.signature
}

func (jvm *JVM) newJClassFromJava(jobject C.jobject, sig string) (*JClass, error) {
	ret := &JClass{
		jvm:       jvm,
		javavalue: NewCJvalue(C.calloc_jvalue_jobject(&jobject)),
		signature: sig,
		globalRef: C.NewGlobalRef(jvm.env(), jobject),
	}

	fqcn := sig[1 : len(sig)-1]
	cname := C.CString(fqcn)
	defer C.free(unsafe.Pointer(cname))
	clazz := C.FindClass(jvm.env(), cname)
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
	clazz := C.FindClass(jvm.env(), cname)
	if clazz == nil {
		return nil, errors.New("FindClass" + fqcn)
	}
	methodID := C.GetMethodID(jvm.env(), clazz, cinit, csig)
	obj := C.NewObjectA(jvm.env(), clazz, methodID, jObjectArrayTojvalueArray(args))
	C.ExceptionDescribe(jvm.env())
	ret := &JClass{
		jvm:       jvm,
		javavalue: NewCJvalue(C.calloc_jvalue_jobject(&obj)),
		signature: "L" + fqcn + ";",
		globalRef: C.NewGlobalRef(jvm.env(), obj),
		clazz:     clazz,
	}

	runtime.SetFinalizer(ret, jvm.destroyJClass)
	return ret, nil
}

func (jvm *JVM) destroyJClass(jobject *JClass) {
	C.DeleteGlobalRef(jvm.env(), jobject.globalRef)
	C.free(unsafe.Pointer(jobject.javavalue.jvaluep()))
}
