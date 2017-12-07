package jnigo

// #include"jni_wrapper.h"
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

type jString struct {
	JObject
	jvm       *JVM
	javavalue CJvalue
	signature string
	globalRef C.jobject
}

func (a *jString) GoValue() interface{} {
	jstr := a.javavalue.jstring()
	jlength := C.GetStringLength(a.jvm.env(), jstr)
	start := C.jsize(0)
	buf := C.calloc_jchar_array(C.size_t(jlength))
	defer C.free(unsafe.Pointer(buf))

	C.GetStringRegion(a.jvm.env(), jstr, start, jlength, buf)

	return C.GoString((*C.char)(unsafe.Pointer(buf)))
}

func (a *jString) JavaValue() CJvalue {
	return a.javavalue
}

func (a *jString) String() string {
	return fmt.Sprint(a.GoValue())
}

func (a *jString) Signature() string {
	return a.signature
}

func (jvm *JVM) newjStringFromJava(jstr C.jobject) (*jString, error) {
	defer C.DeleteLocalRef(jvm.env(), jstr)
	ref := C.NewGlobalRef(jvm.env(), jstr)

	ret := &jString{
		jvm:       jvm,
		javavalue: NewCJvalue(C.calloc_jvalue_jobject(&ref)),
		signature: "Ljava/lang/String;",
		globalRef: ref,
	}
	runtime.SetFinalizer(ret, jvm.destroyjString)
	return ret, nil
}

func (jvm *JVM) newjString(str string) (*jString, error) {
	cstr := C.CString(str) // will be freed by JNI??
	jstr := C.NewString(jvm.env(), (*C.jchar)(unsafe.Pointer(cstr)), C.jsize(len(str)))
	defer C.DeleteLocalRef(jvm.env(), jstr)
	ref := C.NewGlobalRef(jvm.env(), jstr)

	ret := &jString{
		jvm:       jvm,
		javavalue: NewCJvalue(C.calloc_jvalue_jobject(&ref)),
		signature: "Ljava/lang/String;",
		globalRef: ref,
	}
	runtime.SetFinalizer(ret, jvm.destroyjString)
	return ret, nil
}

func (jvm *JVM) destroyjString(jobject *jString) {
	C.DeleteGlobalRef(jvm.env(), jobject.globalRef)
	jobject.javavalue.free()
}
