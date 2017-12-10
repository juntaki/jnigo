package jnigo

// #include"jni_wrapper.h"
import "C"
import (
	"fmt"
	"runtime"
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
	buf := C.GetStringUTFChars(a.jvm.env(), jstr, nil)
	defer C.ReleaseStringUTFChars(a.jvm.env(), jstr, buf)

	return C.GoString(buf)
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
	jstr := C.NewStringUTF(jvm.env(), cstr)
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
