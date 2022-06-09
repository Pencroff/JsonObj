package benchmark

import (
	"fmt"
	"github.com/Pencroff/JsonStruct/experiment"
	"reflect"
	"testing"
)

var intResult int

func BenchmarkPrimitiveOps(b *testing.B) {
	jsV := experiment.JsonStructValue{}
	jsP := experiment.JsonStructPtr{}
	//tm, _ := time.Parse(time.RFC3339, "2022-02-24T04:59:59Z")
	PrintSize(&jsV)
	PrintSize(&jsP)

	b.Run("Value set int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			jsV.SetInt(i)
		}
	})
	b.Run("Value get int", func(b *testing.B) {
		n := 0
		jsV.SetInt(100)
		for i := 0; i < b.N; i++ {
			n = jsV.Int()
		}
		intResult = n
	})
	b.Run("Ptr set int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			jsP.SetInt(i)
		}
	})
	b.Run("Ptr get int", func(b *testing.B) {
		n := 0
		jsP.SetInt(100)
		for i := 0; i < b.N; i++ {
			n = jsP.Int()
		}
		intResult = n
	})
}

func PrintSize(v interface{}) {
	fmt.Printf("Size: %v\n", reflect.Indirect(reflect.ValueOf(v)).Type().Size())
}
