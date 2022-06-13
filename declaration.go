package JsonStruct

import (
	"errors"
	"time"
)

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

	IsNull() bool
	SetNull()

	Type() Type
}

type ObjectOps interface {
	Set(string, interface{}) error
	Get(string) JsonStructOps
	Remove(string) bool
	Has(string) bool
	Keys() []string
	IsObject() bool
	AsObject()
}

type ArrayOps interface {
	Len() int
	Push(interface{}) error
	Pop() JsonStructOps
	SetIndex(int, interface{}) error
	GetIndex(int) JsonStructOps
	IsArray() bool
	AsArray()
}

type JsonStructOps interface {
	PrimitiveOps
	ObjectOps
	ArrayOps
}

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

var UnsupportedTypeError = errors.
	New("unsupported value type, resolved as null")
var NotObjectError = errors.New("not an object, set explicitly")
var NotArrayError = errors.New("not an array, set explicitly")
