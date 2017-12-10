package jnigo

import "regexp"

const (
	SignatureBoolean      = "Z"
	SignatureByte         = "B"
	SignatureChar         = "C"
	SignatureShort        = "S"
	SignatureInt          = "I"
	SignatureLong         = "J"
	SignatureFloat        = "F"
	SignatureDouble       = "D"
	SignatureArray        = "["
	SignatureVoid         = "V"
	SignatureClass        = "L"
	SignatureBooleanArray = SignatureArray + SignatureBoolean
	SignatureByteArray    = SignatureArray + SignatureByte
	SignatureCharArray    = SignatureArray + SignatureChar
	SignatureShortArray   = SignatureArray + SignatureShort
	SignatureIntArray     = SignatureArray + SignatureInt
	SignatureLongArray    = SignatureArray + SignatureLong
	SignatureFloatArray   = SignatureArray + SignatureFloat
	SignatureDoubleArray  = SignatureArray + SignatureDouble
	SignatureClassArray   = SignatureArray + SignatureClass
)

var SizeOf = map[string]int{
	SignatureBoolean: 1,
	SignatureByte:    1,
	SignatureChar:    2,
	SignatureShort:   2,
	SignatureInt:     4,
	SignatureLong:    8,
	SignatureFloat:   4,
	SignatureDouble:  8,
	SignatureArray:   8,
	SignatureVoid:    0,
	SignatureClass:   8,
}

type JObject interface {
	Signature() string
	GoValue() interface{}
	JavaValue() CJvalue
}

var funcSignagure = regexp.MustCompile(`\((.*)\)((.).*)`)
