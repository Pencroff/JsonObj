package benchmark

import (
	"fmt"
	djs "github.com/Pencroff/JsonStruct"
	ejs "github.com/Pencroff/JsonStruct/experiment"
	"testing"
	"time"
)

var res interface{}

func BenchmarkObjArrayOps(b *testing.B) {
	tmStr := "2015-01-01T12:34:56Z"
	tm, _ := time.Parse(time.RFC3339, tmStr)
	vjs := ejs.JsonStructValue{}
	vjs.AsObject()
	pjs := ejs.JsonStructPtr{}
	pjs.AsObject()
	PrintSize(&vjs)
	PrintSize(&pjs)

	tblObj := []struct {
		name string
		o    djs.JsonStructOps
	}{
		{"Val", &vjs},
		{"Ptr", &pjs},
	}
	tblMethod := []struct {
		keyType string
		v       interface{}
	}{
		{"Bool", true},
		{"Int", 10},
		{"Uint", uint64(10)},
		{"Float", 3.14159},
		{"String", "Hello World"},
		{"Time", tm},
	}

	for _, m := range tblMethod {
		for _, t := range tblObj {
			nameFn := CreateObjArrNameFn(t.name, m.keyType)
			b.Run(nameFn("SetKey"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					t.o.SetKey(m.keyType, m.v)
				}
			})
			b.Run(nameFn("GetKey"), func(b *testing.B) {
				t.o.SetKey(m.keyType, m.v)
				var v djs.JsonStructOps
				for i := 0; i < b.N; i++ {
					v = t.o.GetKey(m.keyType)
				}
				res = v
			})
			b.Run(nameFn("HasKey"), func(b *testing.B) {
				t.o.SetKey(m.keyType, m.v)
				var v bool
				for i := 0; i < b.N; i++ {
					v = t.o.HasKey(m.keyType)
				}
				res = v
			})
			b.Run(nameFn("Keys"), func(b *testing.B) {
				t.o.SetKey(m.keyType, m.v)
				var v []string
				for i := 0; i < b.N; i++ {
					v = t.o.Keys()
				}
				res = len(v)
			})
			fmt.Println("")
		}
	}
	fmt.Println("----------------------------")
	vjs.AsArray()
	pjs.AsArray()
	for _, m := range tblMethod {
		for _, t := range tblObj {
			nameFn := CreateObjArrNameFn(t.name, m.keyType)
			b.Run(nameFn("Push"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					t.o.Push(m.v)
				}
			})
			b.Run(nameFn("Pop"), func(b *testing.B) {
				var v djs.JsonStructOps
				for i := 0; i < b.N; i++ {
					v = t.o.Pop()
				}
				res = v
			})
			//b.Run(nameFn("HasKey"), func(b *testing.B) {
			//	t.o.SetKey(m.keyType, m.v)
			//	var v bool
			//	for i := 0; i < b.N; i++ {
			//		v = t.o.HasKey(m.keyType)
			//	}
			//	res = v
			//})
			//b.Run(nameFn("Keys"), func(b *testing.B) {
			//	t.o.SetKey(m.keyType, m.v)
			//	var v []string
			//	for i := 0; i < b.N; i++ {
			//		v = t.o.Keys()
			//	}
			//	res = len(v)
			//})
			//b.Run(nameFn("Keys"), func(b *testing.B) {
			//	t.o.SetKey(m.keyType, m.v)
			//	var v []string
			//	for i := 0; i < b.N; i++ {
			//		v = t.o.Keys()
			//	}
			//	res = len(v)
			//})
			fmt.Println("")
		}
	}
}

func CreateObjArrNameFn(implementation, keyType string) func(name string) string {
	return func(name string) string {
		return fmt.Sprintf("%s.%s(%s)", implementation, name, keyType)
	}

}
