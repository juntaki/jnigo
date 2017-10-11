#include <jni.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>

typedef struct JVM {
    JNIEnv *env;
    JavaVM *jvm;
} JVM;

// Helper

jvalue *calloc_jvalue();
jvalue *calloc_jvalue_array(size_t len);
jvalue *calloc_jvalue_jobject(jobject *val);

jboolean *jvalue_to_jboolean(jvalue *value);
jbyte *jvalue_to_jbyte(jvalue *value);
jchar *jvalue_to_jchar(jvalue *value);
jshort *jvalue_to_jshort(jvalue *value);
jint *jvalue_to_jint(jvalue *value);
jlong *jvalue_to_jlong(jvalue *value);
jfloat *jvalue_to_jfloat(jvalue *value);
jdouble *jvalue_to_jdouble(jvalue *value);
jobject *jvalue_to_jobject(jvalue *jvalue);

jobject *jarray_to_jobject(jarray *jarray);
jarray *jobject_to_jarray(jobject *val);

jboolean *calloc_jboolean_array(size_t len);
jbyte *calloc_jbyte_array(size_t len);
jchar *calloc_jchar_array(size_t len);
jshort *calloc_jshort_array(size_t len);
jint *calloc_jint_array(size_t len);
jlong *calloc_jlong_array(size_t len);
jfloat *calloc_jfloat_array(size_t len);
jdouble *calloc_jdouble_array(size_t len);
jobject *calloc_jobject_array(size_t len);

jbooleanArray jobject_to_jbooleanArray(jobject val);
jbyteArray jobject_to_jbyteArray(jobject val);
jcharArray jobject_to_jcharArray(jobject val);
jshortArray jobject_to_jshortArray(jobject val);
jintArray jobject_to_jintArray(jobject val);
jlongArray jobject_to_jlongArray(jobject val);
jfloatArray jobject_to_jfloatArray(jobject val);
jdoubleArray jobject_to_jdoubleArray(jobject val);
jobjectArray jobject_to_jobjectArray(jobject val);
jobject jbooleanArray_to_jobject(jbooleanArray val);
jobject jbyteArray_to_jobject(jbyteArray val);
jobject jcharArray_to_jobject(jcharArray val);
jobject jshortArray_to_jobject(jshortArray val);
jobject jintArray_to_jobject(jintArray val);
jobject jlongArray_to_jobject(jlongArray val);
jobject jfloatArray_to_jobject(jfloatArray val);
jobject jdoubleArray_to_jobject(jdoubleArray val);
jobject jobjectArray_to_jobject(jobjectArray val);

uint8_t jboolean_to_uint8(jboolean val);
int8_t jbyte_to_int8(jbyte val);
uint16_t jchar_to_uint16(jchar val);
int16_t jshort_to_int16(jshort val);
int32_t jint_to_int32(jint val);
int64_t jlong_to_int64(jlong val);
float jfloat_to_float(jfloat val);
double jdouble_to_double(jdouble val);
jboolean uint8_to_jboolean(uint8_t val);
jbyte int8_to_jbyte(int8_t val);
jchar uint16_to_jchar(uint16_t val);
jshort int16_to_jshort(int16_t val);
jint int32_to_jint(int32_t val);
jlong int64_to_jlong(int64_t val);
jfloat float_to_jfloat(float val);
jdouble double_to_jdouble(double val);

// Wrapper functions

JVM* createJVM();
void destroyJVM(JVM *jvm);

// JNI functions

jint GetVersion(JNIEnv *env);

jclass DefineClass(JNIEnv *env, const char *name, jobject loader, const jbyte *buf, jsize len);
jclass FindClass(JNIEnv *env, const char *name);

jmethodID FromReflectedMethod(JNIEnv *env, jobject method);
jfieldID FromReflectedField(JNIEnv *env, jobject field);

jobject ToReflectedMethod(JNIEnv *env, jclass cls, jmethodID methodID, jboolean isStatic);

jclass GetSuperclass(JNIEnv *env, jclass sub);
jboolean IsAssignableFrom(JNIEnv *env, jclass sub, jclass sup);

jobject ToReflectedField(JNIEnv *env, jclass cls, jfieldID fieldID, jboolean isStatic);

jint Throw(JNIEnv *env, jthrowable obj);
jint ThrowNew(JNIEnv *env, jclass clazz, const char *msg);
jthrowable ExceptionOccurred(JNIEnv *env);
void ExceptionDescribe(JNIEnv *env);
void ExceptionClear(JNIEnv *env);
void FatalError(JNIEnv *env, const char *msg);

jint PushLocalFrame(JNIEnv *env, jint capacity);
jobject PopLocalFrame(JNIEnv *env, jobject result);

jobject NewGlobalRef(JNIEnv *env, jobject lobj);
void DeleteGlobalRef(JNIEnv *env, jobject gref);
void DeleteLocalRef(JNIEnv *env, jobject obj);
jboolean IsSameObject(JNIEnv *env, jobject obj1, jobject obj2);
jobject NewLocalRef(JNIEnv *env, jobject ref);
jint EnsureLocalCapacity(JNIEnv *env, jint capacity);

jobject AllocObject(JNIEnv *env, jclass clazz);
jobject NewObjectA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

jclass GetObjectClass(JNIEnv *env, jobject obj);
jboolean IsInstanceOf(JNIEnv *env, jobject obj, jclass clazz);

jmethodID GetMethodID(JNIEnv *env, jclass clazz, const char *name, const char *sig);

jobject CallObjectMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);

jboolean CallBooleanMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);

jbyte CallByteMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

jchar CallCharMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

jshort CallShortMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

jint CallIntMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

jlong CallLongMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

jfloat CallFloatMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

jdouble CallDoubleMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

void CallVoidMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);

jobject CallNonvirtualObjectMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args);

jboolean CallNonvirtualBooleanMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args);

jbyte CallNonvirtualByteMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);

jchar CallNonvirtualCharMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);

jshort CallNonvirtualShortMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);

jint CallNonvirtualIntMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);

jlong CallNonvirtualLongMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);

jfloat CallNonvirtualFloatMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);

jdouble CallNonvirtualDoubleMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);

void CallNonvirtualVoidMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args);

jfieldID GetFieldID(JNIEnv *env, jclass clazz, const char *name, const char *sig);

jobject GetObjectField(JNIEnv *env, jobject obj, jfieldID fieldID);
jboolean GetBooleanField(JNIEnv *env, jobject obj, jfieldID fieldID);
jbyte GetByteField(JNIEnv *env, jobject obj, jfieldID fieldID);
jchar GetCharField(JNIEnv *env, jobject obj, jfieldID fieldID);
jshort GetShortField(JNIEnv *env, jobject obj, jfieldID fieldID);
jint GetIntField(JNIEnv *env, jobject obj, jfieldID fieldID);
jlong GetLongField(JNIEnv *env, jobject obj, jfieldID fieldID);
jfloat GetFloatField(JNIEnv *env, jobject obj, jfieldID fieldID);
jdouble GetDoubleField(JNIEnv *env, jobject obj, jfieldID fieldID);

void SetObjectField(JNIEnv *env, jobject obj, jfieldID fieldID, jobject val);
void SetBooleanField(JNIEnv *env, jobject obj, jfieldID fieldID, jboolean val);
void SetByteField(JNIEnv *env, jobject obj, jfieldID fieldID, jbyte val);
void SetCharField(JNIEnv *env, jobject obj, jfieldID fieldID, jchar val);
void SetShortField(JNIEnv *env, jobject obj, jfieldID fieldID, jshort val);
void SetIntField(JNIEnv *env, jobject obj, jfieldID fieldID, jint val);
void SetLongField(JNIEnv *env, jobject obj, jfieldID fieldID, jlong val);
void SetFloatField(JNIEnv *env, jobject obj, jfieldID fieldID, jfloat val);
void SetDoubleField(JNIEnv *env, jobject obj, jfieldID fieldID, jdouble val);

jmethodID GetStaticMethodID(JNIEnv *env, jclass clazz, const char *name, const char *sig);

jobject CallStaticObjectMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

jboolean CallStaticBooleanMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

jbyte CallStaticByteMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

jchar CallStaticCharMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

jshort CallStaticShortMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

jint CallStaticIntMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

jlong CallStaticLongMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

jfloat CallStaticFloatMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

jdouble CallStaticDoubleMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

void CallStaticVoidMethodA(JNIEnv *env, jclass cls, jmethodID methodID, const jvalue * args);

jfieldID GetStaticFieldID(JNIEnv *env, jclass clazz, const char *name, const char *sig);
jobject GetStaticObjectField(JNIEnv *env, jclass clazz, jfieldID fieldID);
jboolean GetStaticBooleanField(JNIEnv *env, jclass clazz, jfieldID fieldID);
jbyte GetStaticByteField(JNIEnv *env, jclass clazz, jfieldID fieldID);
jchar GetStaticCharField(JNIEnv *env, jclass clazz, jfieldID fieldID);
jshort GetStaticShortField(JNIEnv *env, jclass clazz, jfieldID fieldID);
jint GetStaticIntField(JNIEnv *env, jclass clazz, jfieldID fieldID);
jlong GetStaticLongField(JNIEnv *env, jclass clazz, jfieldID fieldID);
jfloat GetStaticFloatField(JNIEnv *env, jclass clazz, jfieldID fieldID);
jdouble GetStaticDoubleField(JNIEnv *env, jclass clazz, jfieldID fieldID);

void SetStaticObjectField(JNIEnv *env, jclass clazz, jfieldID fieldID, jobject value);
void SetStaticBooleanField(JNIEnv *env, jclass clazz, jfieldID fieldID, jboolean value);
void SetStaticByteField(JNIEnv *env, jclass clazz, jfieldID fieldID, jbyte value);
void SetStaticCharField(JNIEnv *env, jclass clazz, jfieldID fieldID, jchar value);
void SetStaticShortField(JNIEnv *env, jclass clazz, jfieldID fieldID, jshort value);
void SetStaticIntField(JNIEnv *env, jclass clazz, jfieldID fieldID, jint value);
void SetStaticLongField(JNIEnv *env, jclass clazz, jfieldID fieldID, jlong value);
void SetStaticFloatField(JNIEnv *env, jclass clazz, jfieldID fieldID, jfloat value);
void SetStaticDoubleField(JNIEnv *env, jclass clazz, jfieldID fieldID, jdouble value);

jstring NewString(JNIEnv *env, const jchar *unicode, jsize len);
jsize GetStringLength(JNIEnv *env, jstring str);
const jchar *GetStringChars(JNIEnv *env, jstring str, jboolean *isCopy);
void ReleaseStringChars(JNIEnv *env, jstring str, const jchar *chars);

jstring NewStringUTF(JNIEnv *env, const char *utf);
jsize GetStringUTFLength(JNIEnv *env, jstring str);
const char* GetStringUTFChars(JNIEnv *env, jstring str, jboolean *isCopy);
void ReleaseStringUTFChars(JNIEnv *env, jstring str, const char* chars);

jsize GetArrayLength(JNIEnv *env, jarray array);

jobjectArray NewObjectArray(JNIEnv *env, jsize len, jclass clazz, jobject init);
jobject GetObjectArrayElement(JNIEnv *env, jobjectArray array, jsize index);
void SetObjectArrayElement(JNIEnv *env, jobjectArray array, jsize index, jobject val);

jbooleanArray NewBooleanArray(JNIEnv *env, jsize len);
jbyteArray NewByteArray(JNIEnv *env, jsize len);
jcharArray NewCharArray(JNIEnv *env, jsize len);
jshortArray NewShortArray(JNIEnv *env, jsize len);
jintArray NewIntArray(JNIEnv *env, jsize len);
jlongArray NewLongArray(JNIEnv *env, jsize len);
jfloatArray NewFloatArray(JNIEnv *env, jsize len);
jdoubleArray NewDoubleArray(JNIEnv *env, jsize len);

jboolean * GetBooleanArrayElements(JNIEnv *env, jbooleanArray array, jboolean *isCopy);
jbyte * GetByteArrayElements(JNIEnv *env, jbyteArray array, jboolean *isCopy);
jchar * GetCharArrayElements(JNIEnv *env, jcharArray array, jboolean *isCopy);
jshort * GetShortArrayElements(JNIEnv *env, jshortArray array, jboolean *isCopy);
jint * GetIntArrayElements(JNIEnv *env, jintArray array, jboolean *isCopy);
jlong * GetLongArrayElements(JNIEnv *env, jlongArray array, jboolean *isCopy);
jfloat * GetFloatArrayElements(JNIEnv *env, jfloatArray array, jboolean *isCopy);
jdouble * GetDoubleArrayElements(JNIEnv *env, jdoubleArray array, jboolean *isCopy);

void ReleaseBooleanArrayElements(JNIEnv *env, jbooleanArray array, jboolean *elems, jint mode);
void ReleaseByteArrayElements(JNIEnv *env, jbyteArray array, jbyte *elems, jint mode);
void ReleaseCharArrayElements(JNIEnv *env, jcharArray array, jchar *elems, jint mode);
void ReleaseShortArrayElements(JNIEnv *env, jshortArray array, jshort *elems, jint mode);
void ReleaseIntArrayElements(JNIEnv *env, jintArray array, jint *elems, jint mode);
void ReleaseLongArrayElements(JNIEnv *env, jlongArray array, jlong *elems, jint mode);
void ReleaseFloatArrayElements(JNIEnv *env, jfloatArray array, jfloat *elems, jint mode);
void ReleaseDoubleArrayElements(JNIEnv *env, jdoubleArray array, jdouble *elems, jint mode);

void GetBooleanArrayRegion(JNIEnv *env, jbooleanArray array, jsize start, jsize l, jboolean *buf);
void GetByteArrayRegion(JNIEnv *env, jbyteArray array, jsize start, jsize len, jbyte *buf);
void GetCharArrayRegion(JNIEnv *env, jcharArray array, jsize start, jsize len, jchar *buf);
void GetShortArrayRegion(JNIEnv *env, jshortArray array, jsize start, jsize len, jshort *buf);
void GetIntArrayRegion(JNIEnv *env, jintArray array, jsize start, jsize len, jint *buf);
void GetLongArrayRegion(JNIEnv *env, jlongArray array, jsize start, jsize len, jlong *buf);
void GetFloatArrayRegion(JNIEnv *env, jfloatArray array, jsize start, jsize len, jfloat *buf);
void GetDoubleArrayRegion(JNIEnv *env, jdoubleArray array, jsize start, jsize len, jdouble *buf);

void SetBooleanArrayRegion(JNIEnv *env, jbooleanArray array, jsize start, jsize l, const jboolean *buf);
void SetByteArrayRegion(JNIEnv *env, jbyteArray array, jsize start, jsize len, const jbyte *buf);
void SetCharArrayRegion(JNIEnv *env, jcharArray array, jsize start, jsize len, const jchar *buf);
void SetShortArrayRegion(JNIEnv *env, jshortArray array, jsize start, jsize len, const jshort *buf);
void SetIntArrayRegion(JNIEnv *env, jintArray array, jsize start, jsize len, const jint *buf);
void SetLongArrayRegion(JNIEnv *env, jlongArray array, jsize start, jsize len, const jlong *buf);
void SetFloatArrayRegion(JNIEnv *env, jfloatArray array, jsize start, jsize len, const jfloat *buf);
void SetDoubleArrayRegion(JNIEnv *env, jdoubleArray array, jsize start, jsize len, const jdouble *buf);

jint RegisterNatives(JNIEnv *env, jclass clazz, const JNINativeMethod *methods, jint nMethods);
jint UnregisterNatives(JNIEnv *env, jclass clazz);

jint MonitorEnter(JNIEnv *env, jobject obj);
jint MonitorExit(JNIEnv *env, jobject obj);

jint GetJavaVM(JNIEnv *env, JavaVM **vm);

void GetStringRegion(JNIEnv *env, jstring str, jsize start, jsize len, jchar *buf);
void GetStringUTFRegion(JNIEnv *env, jstring str, jsize start, jsize len, char *buf);

void * GetPrimitiveArrayCritical(JNIEnv *env, jarray array, jboolean *isCopy);
void ReleasePrimitiveArrayCritical(JNIEnv *env, jarray array, void *carray, jint mode);

const jchar * GetStringCritical(JNIEnv *env, jstring string, jboolean *isCopy);
void ReleaseStringCritical(JNIEnv *env, jstring string, const jchar *cstring);

jweak NewWeakGlobalRef(JNIEnv *env, jobject obj);
void DeleteWeakGlobalRef(JNIEnv *env, jweak ref);

jboolean ExceptionCheck(JNIEnv *env);

jobject NewDirectByteBuffer(JNIEnv *env, void* address, jlong capacity);
void* GetDirectBufferAddress(JNIEnv *env, jobject buf);
jlong GetDirectBufferCapacity(JNIEnv *env, jobject buf);

/* New JNI 1.6 Features */

jobjectRefType GetObjectRefType(JNIEnv *env, jobject obj);


