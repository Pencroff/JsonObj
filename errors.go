package JsonStruct

import (
	"errors"
	"strconv"
)

var NotObjectError = errors.New("JsonStruct: not an object, set explicitly")
var NotArrayError = errors.New("JsonStruct: not an array, set explicitly")
var IndexOutOfRangeError = errors.New("JsonStruct: index out of range")
var UnsupportedTypeError = errors.New("JsonStruct: unsupported value type, resolved as null")
var InvalidHexNumberError = errors.New("JsonStruct: invalid hex number")
var InvalidEscapeCharacterError = errors.New("JsonStruct: invalid escape character")
var InvalidCharacterError = errors.New("JsonStruct: invalid character")

type InvalidJsonError struct {
	Err error
}

func (i InvalidJsonError) Error() string {
	msg := "JsonStruct: invalid json format"
	if i.Err != nil {
		return msg + ": " + i.Err.Error()
	}
	return msg
}

type InvalidJsonPtrError struct {
	Err error
	Pos int
}

func (i InvalidJsonPtrError) Error() string {
	msg := "JsonStruct: invalid json format at position " + strconv.FormatInt(int64(i.Pos), 10)
	if i.Err != nil {
		return msg + ": " + i.Err.Error()
	}
	return msg
}

type InvalidJsonTokenPtrError struct {
	Err error
	Pos int
}

func (i InvalidJsonTokenPtrError) Error() string {
	msg := "JsonStruct: invalid json token at position " + strconv.FormatInt(int64(i.Pos), 10)
	if i.Err != nil {
		return msg + ": " + i.Err.Error()
	}
	return msg
}

var OffsetOutOfRangeError = errors.New("JStructReader: offset out of range")
