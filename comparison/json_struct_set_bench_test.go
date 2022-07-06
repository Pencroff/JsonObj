package comparison

import (
	. "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type JsonStructA struct {
	JsonStruct
	m map[string]*JsonStruct
}

func (s *JsonStructA) Set(key string, value interface{}) (err error) {
	if s.m == nil {
		s.m = make(map[string]*JsonStruct)
	}
	vval, ok := value.(JsonStruct)
	if ok {
		s.m[key] = &vval
		return nil
	}
	vptr, ok := value.(*JsonStruct)
	if ok {
		s.m[key] = vptr
		return nil
	}
	jsonStruct, ok := s.m[key]
	if !ok {
		jsonStruct = &JsonStruct{}
	}

	vInt, ok := value.(int64)
	if ok {
		jsonStruct.SetInt(vInt)
	} else {
		vFloat, ok := value.(float64)
		if ok {
			jsonStruct.SetFloat(vFloat)
		} else {
			vBool, ok := value.(bool)
			if ok {
				jsonStruct.SetBool(vBool)
			} else {
				vStr, ok := value.(string)
				if ok {
					jsonStruct.SetString(vStr)
				} else {
					vTime, ok := value.(time.Time)
					if ok {
						jsonStruct.SetTime(vTime)
					} else {
						if value != nil {
							return UnsupportedTypeError
						}
					}
				}
			}
		}
	}
	s.m[key] = jsonStruct
	return nil
}

func (s *JsonStructA) Get(key string) *JsonStruct {
	if s.m != nil {
		return s.m[key]
	}
	return nil
}

func (s *JsonStructA) Has(key string) bool {
	_, ok := s.m[key]
	return ok
}

type JsonStructB struct {
	JsonStruct
	m map[string]*JsonStruct
}

func (s *JsonStructB) Set(key string, value interface{}) (err error) {
	if s.m == nil {
		s.m = make(map[string]*JsonStruct)
	}
	v, ok := s.m[key]
	if !ok {
		v = &JsonStruct{}
	}
	switch data := value.(type) {
	case nil:
	case JsonStruct:
		v = &data
	case *JsonStruct:
		v = data
	case int64:
		v.SetInt(data)
	case float64:
		v.SetFloat(data)
	case bool:
		v.SetBool(data)
	case string:
		v.SetString(data)
	case time.Time:
		v.SetTime(data)
	default:
		return UnsupportedTypeError
	}
	s.m[key] = v
	return nil
}

func (s *JsonStructB) Get(key string) *JsonStruct {
	if s.m != nil {
		return s.m[key]
	}
	return nil
}

func (s *JsonStructB) Has(key string) bool {
	_, ok := s.m[key]
	return ok
}

func TestJsonStructA_Set(t *testing.T) {
	js := &JsonStructA{}
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
}

func TestJsonStructB_Set(t *testing.T) {
	js := &JsonStructB{}
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
}

func BenchmarkRndGen(b *testing.B) {
	jsA := JsonStructA{}
	jsB := JsonStructB{}
	tm, _ := time.Parse(time.RFC3339, "2022-02-24T04:59:59Z")
	v := JsonStruct{}
	vptr := &JsonStruct{}
	m := map[string]string{"a": "b"}
	b.Run("if based", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			jsA.Set("keyStr", "value")
			jsA.Set("keyInt", i)
			jsA.Set("keyFloat", 0.1)
			jsA.Set("keyBool", false)
			jsA.Set("keyTime", tm)
			jsA.Set("keyValue", v)
			jsA.Set("keyPtr", vptr)
			jsA.Set("null", nil)
			jsA.Set("keyMap", m)
		}
	})
	b.Run("switch based", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			jsB.Set("keyStr", "value")
			jsB.Set("keyInt", i)
			jsB.Set("keyFloat", 0.1)
			jsB.Set("keyBool", false)
			jsB.Set("keyTime", tm)
			jsB.Set("keyValue", v)
			jsB.Set("keyPtr", vptr)
			jsB.Set("null", nil)
			jsA.Set("keyMap", m)
		}
	})
}
