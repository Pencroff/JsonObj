package helper

import (
	djs "github.com/Pencroff/JsonStruct"
	"reflect"
)

func MemCounter(v djs.JsonStructOps) uint64 {
	v := uint(0)
	switch v.Type() {
	case djs.TypeNull:
		v = 0
	}
	return v
}

func MemSize(v interface{}) uint64 {
	size := reflect.Indirect(reflect.ValueOf(v)).Type().Size()
	return uint64(size)
}
