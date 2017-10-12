package jnigo

import "reflect"

func (jvm *JVM) Convert(value interface{}) (JObject, error) {
	if jobject, ok := value.(JObject); ok {
		return jobject, nil
	} else if reflect.TypeOf(value).Kind() == reflect.Array || reflect.TypeOf(value).Kind() == reflect.Slice {
		return jvm.newJArray(value)
	} else if reflect.TypeOf(value).Kind() == reflect.String {
		return jvm.newjString(value.(string))
	} else {
		return jvm.newJPrimitive(value)
	}
}

func (jvm *JVM) ConvertAll(allValue []interface{}) ([]JObject, error) {
	ret := make([]JObject, len(allValue))
	var err error
	for i, value := range allValue {
		ret[i], err = jvm.Convert(value)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}
