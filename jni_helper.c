#include "jni_wrapper.h"

jvalue *calloc_jvalue() {
    return (jvalue*) calloc(1, sizeof(jvalue));
}

jvalue *calloc_jvalue_array(size_t len){
    return (jvalue*) calloc(len, sizeof(jvalue));
};

// jvalue conversion
jvalue *calloc_jvalue_jobject(jobject *val) {
    jvalue* ret = (jvalue*) calloc(1, sizeof(jvalue));
    ret->l = *val;
    return ret;
}

jboolean *jvalue_to_jboolean(jvalue *jvalue) {
    return &(jvalue->z);
}

jbyte *jvalue_to_jbyte(jvalue *jvalue) {
    return &(jvalue->b);
}

jchar *jvalue_to_jchar(jvalue *jvalue) {
    return &(jvalue->c);
}

jshort *jvalue_to_jshort(jvalue *jvalue) {
    return &(jvalue->s);
}

jint *jvalue_to_jint(jvalue *jvalue) {
    return &(jvalue->i);
}

jlong *jvalue_to_jlong(jvalue *jvalue) {
    return &(jvalue->j);
}

jfloat *jvalue_to_jfloat(jvalue *jvalue) {
    return &(jvalue->f);
}

jdouble *jvalue_to_jdouble(jvalue *jvalue) {
    return &(jvalue->d);
}

jobject *jvalue_to_jobject(jvalue *jvalue) {
    return &(jvalue->l);
}

// jobject conversion
jobject *jarray_to_jobject(jarray *jarray) {
    return (jobject*)jarray;
}

jarray *jobject_to_jarray(jobject *val){
    return (jarray*)val;
}

// Array
jboolean *calloc_jboolean_array(size_t len) {
    return (jboolean*) calloc(len, sizeof(jboolean));
}

jbyte *calloc_jbyte_array(size_t len) {
    return (jbyte*) calloc(len, sizeof(jbyte));
}

jchar *calloc_jchar_array(size_t len) {
    return (jchar*) calloc(len, sizeof(jchar));
}

jshort *calloc_jshort_array(size_t len) {
    return (jshort*) calloc(len, sizeof(jshort));
}

jint *calloc_jint_array(size_t len) {
    return (jint*) calloc(len, sizeof(jint));
}

jlong *calloc_jlong_array(size_t len) {
    return (jlong*) calloc(len, sizeof(jlong));
}

jfloat *calloc_jfloat_array(size_t len) {
    return (jfloat*) calloc(len, sizeof(jfloat));
}

jdouble *calloc_jdouble_array(size_t len) {
    return (jdouble*) calloc(len, sizeof(jdouble));
}

jobject *calloc_jobject_array(size_t len) {
    return (jobject*) calloc(len, sizeof(jobject));
}

// Array type conversion
jbooleanArray jobject_to_jbooleanArray(jobject val) {
    return (jbooleanArray)val;
}
jbyteArray jobject_to_jbyteArray(jobject val) {
    return (jbyteArray)val;
}
jcharArray jobject_to_jcharArray(jobject val) {
    return (jcharArray)val;
}
jshortArray jobject_to_jshortArray(jobject val) {
    return (jshortArray)val;
}
jintArray jobject_to_jintArray(jobject val) {
    return (jintArray)val;
}
jlongArray jobject_to_jlongArray(jobject val) {
    return (jlongArray)val;
}
jfloatArray jobject_to_jfloatArray(jobject val) {
    return (jfloatArray)val;
}
jdoubleArray jobject_to_jdoubleArray(jobject val) {
    return (jdoubleArray)val;
}
jobjectArray jobject_to_jobjectArray(jobject val) {
    return (jobjectArray)val;
}

jobject jbooleanArray_to_jobject(jbooleanArray val) {
    return (jobject)val;
}
jobject jbyteArray_to_jobject(jbyteArray val) {
    return (jobject)val;
}
jobject jcharArray_to_jobject(jcharArray val) {
    return (jobject)val;
}
jobject jshortArray_to_jobject(jshortArray val) {
    return (jobject)val;
}
jobject jintArray_to_jobject(jintArray val) {
    return (jobject)val;
}
jobject jlongArray_to_jobject(jlongArray val) {
    return (jobject)val;
}
jobject jfloatArray_to_jobject(jfloatArray val) {
    return (jobject)val;
}
jobject jdoubleArray_to_jobject(jdoubleArray val) {
    return (jobject)val;
}
jobject jobjectArray_to_jobject(jobjectArray val) {
    return (jobject)val;
}

jobject jstring_to_jobject(jstring val) {
    return (jobject)val;
}
jstring jobject_to_jstring(jobject val) {
    return (jstring)val;
}


// Type conversion
uint8_t jboolean_to_uint8(jboolean val) {
    return (uint8_t)val;
}
int8_t jbyte_to_int8(jbyte val) {
    return (int8_t)val;
}
uint16_t jchar_to_uint16(jchar val) {
    return (uint16_t)val;
}
int16_t jshort_to_int16(jshort val) {
    return (int16_t)val;
}
int32_t jint_to_int32(jint val) {
    return (int32_t)val;
}
int64_t jlong_to_int64(jlong val) {
    return (int64_t)val;
}
float jfloat_to_float(jfloat val) {
    return (float)val;
}
double jdouble_to_double(jdouble val) {
    return (double)val;
}
jboolean uint8_to_jboolean(uint8_t val) {
    return (jboolean)val;
}
jbyte int8_to_jbyte(int8_t val) {
    return (jbyte)val;
}
jchar uint16_to_jchar(uint16_t val) {
    return (jchar)val;
}
jshort int16_to_jshort(int16_t val) {
    return (jshort)val;
}
jint int32_to_jint(int32_t val) {
    return (jint)val;
}
jlong int64_to_jlong(int64_t val) {
    return (jlong)val;
}
jfloat float_to_jfloat(float val) {
    return (jfloat)val;
}
jdouble double_to_jdouble(double val) {
    return (jdouble)val;
}
