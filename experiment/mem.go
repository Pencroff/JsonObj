package experiment

import (
	"reflect"
)

func MemSize(v interface{}) int {
	size := reflect.Indirect(reflect.ValueOf(v)).Type().Size()
	return int(size)
}
