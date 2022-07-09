package JsonStruct

import (
	"encoding/json"
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
	String
	Time
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
	GetKey(string) JStructOps
	RemoveKey(string) JStructOps
	HasKey(string) bool
	Keys() []string
}

type ArrayOps interface {
	Push(interface{}) error
	Pop() JStructOps
	Shift() JStructOps
	SetIndex(int, interface{}) error
	GetIndex(int) JStructOps
}

type JStructOps interface {
	GeneralOps
	PrimitiveOps
	ObjectOps
	ArrayOps
}

type JStructConvertibleOps interface {
	GeneralOps
	PrimitiveOps
	ObjectOps
	ArrayOps
	json.Unmarshaler
	json.Marshaler
}
