package JsonStruct

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
	"unsafe"
)

func TestJsonStruct_GetHas(t *testing.T) {
	js := JsonStruct{}
	v := js.Get("key")
	assert.Equal(t, true, v == nil)
	assert.Equal(t, false, js.Has("key"))
}

func TestSetKeyValue(t *testing.T) {
	js := JsonStruct{}
	assert.Equal(t, true, js.IsNull())
	js.Set("key", "value")
	assert.Equal(t, true, js.IsObject())
	assert.Equal(t, false, js.IsArray())
	v := js.Get("key")
	assert.Equal(t, true, v.IsString())
	assert.Equal(t, "value", v.String())
}

func TestJsonStruct_Set(t *testing.T) {
	js := JsonStruct{}
	js.Set("keyStr", "value")
	assert.Equal(t, true, js.Has("keyStr"))
	assert.Equal(t, true, js.Get("keyStr").IsString())
	assert.Equal(t, "value", js.Get("keyStr").String())
	js.Set("keyInt", 10)
	assert.Equal(t, true, js.Has("keyInt"))
	assert.Equal(t, true, js.Get("keyInt").IsInt())
	assert.Equal(t, 10, js.Get("keyInt").Int())
	js.Set("keyFloat", 10.1)
	assert.Equal(t, true, js.Has("keyFloat"))
	assert.Equal(t, true, js.Get("keyFloat").IsFloat())
	assert.Equal(t, 10.1, js.Get("keyFloat").Float())
	js.Set("keyBool", false)
	assert.Equal(t, true, js.Has("keyBool"))
	assert.Equal(t, true, js.Get("keyBool").IsBool())
	assert.Equal(t, false, js.Get("keyBool").Bool())
	tm, _ := time.Parse(time.RFC3339, "2022-02-24T04:59:59Z")
	js.Set("keyTime", tm)
	assert.Equal(t, true, js.Has("keyTime"))
	assert.Equal(t, true, js.Get("keyTime").IsTime())
	assert.Equal(t, tm, js.Get("keyTime").Time())
	vMap := map[string]string{"a": "b"}
	e := js.Set("keyMap", vMap)
	assert.EqualError(t, e, "unsupported value type, resolved as null")
	assert.Equal(t, false, js.Has("keyMap"))
	e = js.Set("key", nil)
	assert.NoError(t, e)
	assert.Equal(t, true, js.Has("key"))
	assert.Equal(t, true, js.Get("key").IsNull())
	vJs := JsonStruct{}
	pJs := &JsonStruct{}
	vJs.SetString("a")
	js.Set("keyValue", vJs)
	pJs.SetString("b")
	js.Set("keyPointer", pJs)
	assert.Equal(t, true, js.Has("keyValue"))
	assert.Equal(t, true, js.Has("keyPointer"))
	v := js.Get("keyValue")
	assert.Equal(t, "a", v.String())
	v = js.Get("keyPointer")
	assert.Equal(t, "b", v.String())
	assert.Equal(t, false, vJs.IsInt())
	js.Set("key", pJs)
	js.Set("key", 10)
	assert.Equal(t, 10, js.Get("key").Int())
	assert.Equal(t, 10, pJs.Int())
	//js.Set("key", vJs)
	//js.Set("key", 20)
	//assert.Equal(t, 20, js.Get("key").Int())
	//assert.Equal(t, 20, vJs.Int())
}

func TestAsObjectAsArray(t *testing.T) {
	js := JsonStruct{}
	assert.Equal(t, true, js.IsNull())
	js.AsObject()
	assert.Equal(t, true, js.IsObject())
	assert.Equal(t, false, js.IsArray())
	js.AsArray()
	assert.Equal(t, false, js.IsObject())
	assert.Equal(t, true, js.IsArray())
}

func TestJsonStruct_Delete(t *testing.T) {
	js := JsonStruct{}
	assert.Equal(t, false, js.Delete("no-key"))
	js.Set("key", "value")
	assert.Equal(t, true, js.Has("key"))
	assert.Equal(t, false, js.Delete("empty-key"))
	assert.Equal(t, true, js.Has("key"))
	assert.Equal(t, true, js.Delete("key"))
	assert.Equal(t, false, js.Has("key"))
}

func TestJsonStruct_Has(t *testing.T) {
	js := JsonStruct{}
	js.SetNull()
	assert.Equal(t, false, js.Has("key"))
	js.Set("key", "value")
	assert.Equal(t, true, js.Has("key"))
	assert.Equal(t, true, js.Get("key").IsString())
	js.Set("keyInt", 10)
	js.Set("keyFloat", 10)
	js.SetNull()
	assert.Equal(t, false, js.Has("keyInt"))
	assert.Equal(t, false, js.Has("keyFloat"))
}

func TestJsonStructSize(t *testing.T) {
	js := JsonStruct{}
	assert.Equal(t, uintptr(0x60), unsafe.Sizeof(js))
	js.SetNull()
	assert.Equal(t, uintptr(0x60), unsafe.Sizeof(js))
	js.AsObject()
	assert.Equal(t, uintptr(0x60), unsafe.Sizeof(js))
	js.AsArray()
	assert.Equal(t, uintptr(0x60), unsafe.Sizeof(js))
}

func TestJsonStruct_Keys(t *testing.T) {
	js := JsonStruct{}
	keys := js.Keys()
	assert.Equal(t, 0, len(keys))
	js.Set("key1", "value1")
	js.Set("key2", "value2")
	js.Set("key3", "value3")
	keys = js.Keys()
	assert.Equal(t, 3, len(keys))
	assert.Equal(t, true, sort.StringSlice(keys).Search("key1") > -1)
	assert.Equal(t, true, sort.StringSlice(keys).Search("key2") > -1)
	assert.Equal(t, true, sort.StringSlice(keys).Search("key3") > -1)
}
