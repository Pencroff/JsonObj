package JsonStruct

import "errors"

var NotObjectError = errors.New("JsonStruct: not an object, set explicitly")
var NotArrayError = errors.New("JsonStruct: not an array, set explicitly")
var IndexOutOfRangeError = errors.New("JsonStruct: index out of range")
var UnsupportedTypeError = errors.New("JsonStruct: unsupported value type, resolved as null")
var InvalidJsonError = errors.New("JsonStruct: invalid json format")
