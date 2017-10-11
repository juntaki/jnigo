package jnigo

import "reflect"

func (jvm *JVM) Convert(value interface{}) (JObject, error) {
	if jobject, ok := value.(JObject); ok {
		return jobject, nil
	} else if reflect.TypeOf(value).Kind() == reflect.Array || reflect.TypeOf(value).Kind() == reflect.Slice {
		return jvm.newJArray(value)
	} else {
		return jvm.newJPrimitive(value)
	}
}
