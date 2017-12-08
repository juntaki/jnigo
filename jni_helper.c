#include "jni_wrapper.h"

jobject jobject_conv(jobject val){ return val; }

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

