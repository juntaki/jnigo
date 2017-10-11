package jnigo

import "reflect"

func (jvm *JVM) Convert(value interface{}) (JObject, error) {
	if reflect.TypeOf(value).Kind() == reflect.Array || reflect.TypeOf(value).Kind() == reflect.Slice {
		return jvm.newJArray(value)
	} else {
		return jvm.newJPrimitive(value)
	}
}
