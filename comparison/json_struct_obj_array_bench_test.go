package comparison

import (
	"fmt"
	djs "github.com/Pencroff/JsonStruct"
	ejs "github.com/Pencroff/JsonStruct/experiment"
	"strconv"
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
	eStruct := make(map[string]interface{})
	eArr := make([]interface{}, 0)
	fmt.Println("Object / Map size")
	PrintSize(&vjs)
	PrintSize(&pjs)
	PrintSize(&eStruct)
	tblObj := []struct {
		name    string
		o       djs.JStructOps
		factory func() djs.JStructOps
	}{
		{"Val", &vjs, func() djs.JStructOps { return &ejs.JsonStructValue{} }},
		{"Ptr", &pjs, func() djs.JStructOps { return &ejs.JsonStructPtr{} }},
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
	var nameFn func(string) string
	for _, m := range tblMethod {
		for _, t := range tblObj {
			nameFn = CreateObjArrNameFn(t.name, m.keyType)
			t.o = t.factory()
			t.o.AsObject()
			b.Run(nameFn("SetKey"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					key := m.keyType + strconv.Itoa(i)
					t.o.SetKey(key, m.v)
				}
			})
		}
		nameFn = CreateObjArrNameFn("GoMap", m.keyType)
		eStruct = make(map[string]interface{})
		b.Run(nameFn("SetKey"), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				key := m.keyType + strconv.Itoa(i)
				eStruct[key] = m.v
			}
		})
		fmt.Println("")
		for _, t := range tblObj {
			nameFn = CreateObjArrNameFn(t.name, m.keyType)
			t.o = t.factory()
			t.o.AsObject()
			b.Run(nameFn("GetKey"), func(b *testing.B) {
				var v djs.JStructOps
				for i := 0; i < b.N; i++ {
					key := m.keyType + strconv.Itoa(i)
					v = t.o.GetKey(key)
				}
				res = v
			})
		}
		nameFn = CreateObjArrNameFn("GoMap", m.keyType)
		eStruct = make(map[string]interface{})
		b.Run(nameFn("GetKey"), func(b *testing.B) {
			var v interface{}
			for i := 0; i < b.N; i++ {
				key := m.keyType + strconv.Itoa(i)
				v = eStruct[key]
			}
			res = v
		})
		fmt.Println("")
		for _, t := range tblObj {
			nameFn = CreateObjArrNameFn(t.name, m.keyType)
			t.o = t.factory()
			t.o.AsObject()
			b.Run(nameFn("HasKey"), func(b *testing.B) {
				var v bool
				for i := 0; i < b.N; i++ {
					key := m.keyType + strconv.Itoa(i)
					v = t.o.HasKey(key)
				}
				res = v
			})
		}
		nameFn = CreateObjArrNameFn("GoMap", m.keyType)
		eStruct = make(map[string]interface{})
		b.Run(nameFn("HasKey"), func(b *testing.B) {
			var v bool
			for i := 0; i < b.N; i++ {
				key := m.keyType + strconv.Itoa(i)
				_, v = eStruct[key]
			}
			res = v
		})
		fmt.Println("")
	}
	fmt.Println("----------------------------")
	fmt.Println("Array / Slice size")
	vjs.AsArray()
	pjs.AsArray()
	PrintSize(vjs)
	PrintSize(pjs)
	PrintSize(eArr)
	prePopulateSize := 1000000
	for _, m := range tblMethod {
		for _, t := range tblObj {
			nameFn = CreateObjArrNameFn(t.name, m.keyType)
			t.o = t.factory()
			t.o.AsArray()
			b.Run(nameFn("Push"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					t.o.Push(m.v)
				}
				res = t.o.Size()
			})
		}
		nameFn = CreateObjArrNameFn("GoArr", m.keyType)
		eArr = make([]interface{}, 0)
		b.Run(nameFn("Push"), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				eArr = append(eArr, m.v)
			}
			res = len(eArr)
		})
		fmt.Println("")
		for _, t := range tblObj {
			nameFn = CreateObjArrNameFn(t.name, m.keyType)
			t.o = t.factory()
			t.o.AsArray()
			for _, inEl := range tblMethod {
				for idx := 0; idx < prePopulateSize; idx++ {
					t.o.Push(inEl.v)
				}
			}
			b.Run(nameFn("Pop"), func(b *testing.B) {
				var v djs.JStructOps
				for i := 0; i < b.N; i++ {
					v = t.o.Pop()
				}
				res = v
			})
		}
		nameFn = CreateObjArrNameFn("GoArr", m.keyType)
		eArr = make([]interface{}, prePopulateSize*500)
		for _, inEl := range tblMethod {
			for idx := 0; idx < prePopulateSize*500; idx++ {
				eArr[idx] = inEl.v
			}
		}
		b.Run(nameFn("Pop"), func(b *testing.B) {
			var v interface{}
			for i := 0; i < b.N; i++ {
				idx := len(eArr) - 1
				if idx < 0 {
					return
				}
				v = eArr[idx]
				eArr[idx] = nil
				eArr = eArr[:idx]
			}
			res = v
		})
		fmt.Println("")
	}
}

func CreateObjArrNameFn(implementation, keyType string) func(name string) string {
	return func(name string) string {
		return fmt.Sprintf("%s.%s(%s)", implementation, name, keyType)
	}

}
