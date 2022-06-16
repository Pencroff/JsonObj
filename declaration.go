package JsonStruct

import (
	"errors"
	"time"
)

type Type byte

const (
	Null Type = iota
	False
	True
	Int
	Uint
	Float
	Time
	String
	Object
	Array
)

func (t Type) String() string {
	switch t {
	default:
		return "Null"
	case False:
		return "False"
	case True:
		return "True"
	case Int:
		return "Int"
	case Uint:
		return "Uint"
	case Float:
		return "Float"
	case Time:
		return "Time"
	case String:
		return "String"
	case Object:
		return "Object"
	case Array:
		return "Array"
	}
}

type GeneralOps interface {
	Type() Type
	Value() interface{}

	IsNull() bool
	SetNull()

	IsObject() bool
	AsObject()

	IsArray() bool
	AsArray()

	Size() int
}

type PrimitiveOps interface {
	IsBool() bool
	SetBool(bool)
	Bool() bool

	IsNumber() bool

	IsInt() bool
	SetInt(int64)
	Int() int64

	IsUint() bool
	SetUint(uint64)
	Uint() uint64

	IsFloat() bool
	SetFloat(float64)
	Float() float64

	IsString() bool
	SetString(string)
	String() string

	IsTime() bool
	SetTime(time.Time)
	Time() time.Time
}

type ObjectOps interface {
	SetKey(string, interface{}) error
	GetKey(string) JsonStructOps
	RemoveKey(string) JsonStructOps
	HasKey(string) bool
	Keys() []string
}

type ArrayOps interface {
	Push(interface{}) error
	Pop() JsonStructOps
	Shift() JsonStructOps
	SetIndex(int, interface{}) error
	GetIndex(int) JsonStructOps
}

type JsonStructOps interface {
	GeneralOps
	PrimitiveOps
	ObjectOps
	ArrayOps
}

var UnsupportedTypeError = errors.
	New("unsupported value type, resolved as null")
var NotObjectError = errors.New("not an object, set explicitly")
var NotArrayError = errors.New("not an array, set explicitly")
var IndexOutOfRangeError = errors.New("index out of range")
