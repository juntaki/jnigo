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
	javavalue *C.jvalue
	signature string
	globalRef C.jobject
}

func (a *jString) GoValue() interface{} {
	jstr := C.jobject_to_jstring(*C.jvalue_to_jobject(a.javavalue))
	jlength := C.GetStringLength(a.jvm.cjvm.env, jstr)
	start := C.jsize(0)
	buf := C.calloc_jchar_array(C.size_t(jlength))
	defer C.free(unsafe.Pointer(buf))

	C.GetStringRegion(a.jvm.cjvm.env, jstr, start, jlength, buf)

	return C.GoString((*C.char)(unsafe.Pointer(buf)))
}

func (a *jString) JavaValue() C.jvalue {
	return *a.javavalue
}

func (a *jString) String() string {
	return fmt.Sprintf("0x%x", a.JavaValue())
}

func (a *jString) Signature() string {
	return a.signature
}

func (jvm *JVM) newjString(str string) (*jString, error) {
	cstr := C.CString(str) // will be freed by JNI??
	jstr := C.NewString(jvm.cjvm.env, (*C.jchar)(unsafe.Pointer(cstr)), C.jsize(len(str)))
	jobj := C.jstring_to_jobject(jstr)

	ret := &jString{
		jvm:       jvm,
		javavalue: C.calloc_jvalue_jobject(&jobj),
		signature: "Ljava/lang/String;",
		globalRef: C.NewGlobalRef(jvm.cjvm.env, jstr),
	}
	runtime.SetFinalizer(ret, jvm.destroyjString)
	return ret, nil
}

func (jvm *JVM) destroyjString(jobject *jString) {
	C.DeleteGlobalRef(jvm.cjvm.env, jobject.globalRef)
	C.free(unsafe.Pointer(jobject.javavalue))
}
