package JsonStruct

import (
	"errors"
	"strconv"
)

var NotObjectError = errors.New("JsonStruct: not an object, set explicitly")
var NotArrayError = errors.New("JsonStruct: not an array, set explicitly")
var IndexOutOfRangeError = errors.New("JsonStruct: index out of range")
var UnsupportedTypeError = errors.New("JsonStruct: unsupported value type, resolved as null")
var InvalidJsonError = errors.New("JsonStruct: invalid json format")

type InvalidJsonPtrError struct {
	Pos uint64
}

func (i *InvalidJsonPtrError) Error() string {
	return "JsonStruct: invalid json format at position " + strconv.FormatUint(i.Pos, 10)
}

var OffsetOutOfRangeError = errors.New("JStructScanner: offset out of range")
