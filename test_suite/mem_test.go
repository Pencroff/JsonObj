package test_suite

import (
	"fmt"
	djs "github.com/Pencroff/JsonStruct"
	exp "github.com/Pencroff/JsonStruct/experiment"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func Test_MemTestSuite(t *testing.T) {
	s := new(MemTestSuite)
	suite.Run(t, s)
}

func JsValueFn() djs.JStructOps {
	return &exp.JsonStructValue{}
}

func JsPtrFn() djs.JStructOps {
	return &exp.JsonStructPtr{}
}

type MemTestSuite struct {
	suite.Suite
}

var escape = func(v interface{}) {}

type tbEl struct {
	name   string
	val    interface{}
	callFn func(v tbEl) djs.JStructOps
	jsFn   func() djs.JStructOps
	size   uint64
}

var tm, _ = time.Parse(time.RFC3339, "2015-01-01T12:34:56Z")
var tbl = []tbEl{
	{"ValNull", nil, func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.SetNull()
		return js
	}, JsValueFn, uint64(112)},
	{"PtrNull", nil, func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.SetNull()
		return js
	}, JsPtrFn, uint64(16)},
	// ------------------------------------------------------------
	{"ValBool", true, func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.SetBool(v.val.(bool))
		return js
	}, JsValueFn, uint64(112)},
	{"PtrBool", true, func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.SetBool(v.val.(bool))
		return js
	}, JsPtrFn, uint64(16)},
	// ------------------------------------------------------------
	{"ValInt ", int64(10), func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.SetInt(v.val.(int64))
		return js
	}, JsValueFn, uint64(112)},
	{"PtrInt ", int64(10), func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.SetInt(v.val.(int64))
		return js
	}, JsPtrFn, uint64(24)},
	// ------------------------------------------------------------
	{"ValUint", uint64(10), func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.SetUint(v.val.(uint64))
		return js
	}, JsValueFn, uint64(112)},
	{"PtrUint", uint64(10), func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.SetUint(v.val.(uint64))
		return js
	}, JsPtrFn, uint64(24)},
	// ------------------------------------------------------------
	{"ValFloat", 3.14159, func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.SetFloat(v.val.(float64))
		return js
	}, JsValueFn, uint64(112)},
	{"PtrFloat", 3.14159, func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.SetFloat(v.val.(float64))
		return js
	}, JsPtrFn, uint64(24)},
	// ------------------------------------------------------------
	{"ValString", "Abc...xyz", func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.SetString(v.val.(string))
		return js
	}, JsValueFn, uint64(112)},
	{"PtrString", "Abc...xyz", func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.SetString(v.val.(string))
		return js
	}, JsPtrFn, uint64(32)},
	// ------------------------------------------------------------
	{"ValTime", tm, func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.SetTime(v.val.(time.Time))
		return js
	}, JsValueFn, uint64(112)},
	{"PtrTime", tm, func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.SetTime(v.val.(time.Time))
		return js
	}, JsPtrFn, uint64(40)},
}
var tlbObjArr = []tbEl{
	{"ValObjEmpty", nil, func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.AsObject()
		return js
	}, JsValueFn, uint64(160)},
	{"PtrObjEmpty", nil, func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.AsObject()
		return js
	}, JsPtrFn, uint64(72)},
	{"ValObj", nil, func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.AsObject()
		for idx, el := range tbl {
			if idx%2 == 0 {
				js.SetKey(el.name, el.callFn(el))
			}
		}
		return js
	}, JsValueFn, uint64(1232)},
	{"PtrObj", nil, func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.AsObject()
		for idx, el := range tbl {
			if idx%2 == 1 {
				js.SetKey(el.name, el.callFn(el))
			}

		}
		return js
	}, JsPtrFn, uint64(536)},
	{"ValArrEmpty", nil, func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.AsArray()
		return js
	}, JsValueFn, uint64(112)},
	{"PtrArrEmpty", nil, func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.AsArray()
		return js
	}, JsPtrFn, uint64(40)},
	{"ValArr", nil, func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.AsArray()
		for idx, el := range tbl {
			if idx%2 == 0 {
				js.Push(el.callFn(el))
			}
		}
		return js
	}, JsValueFn, uint64(1136)},
	{"PtrArr", nil, func(v tbEl) djs.JStructOps {
		js := v.jsFn()
		js.AsArray()
		for idx, el := range tbl {
			if idx%2 == 1 {
				js.Push(el.callFn(el))
			}
		}
		return js
	}, JsPtrFn, uint64(624)},
}

func (s *MemTestSuite) TestPrimitiveValues() {
	n := 1000
	for idx, el := range tbl {
		res := testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < n; i++ {
				js := el.callFn(el)
				escape(js)
			}
		})
		size := res.MemBytes / uint64(n)
		fmt.Printf("\t%s:\t%d bytes\n", el.name, size)
		if idx%2 == 1 {
			fmt.Println("")
		}
		s.Equal(el.size, size, "Size of %s is %d", el.name, size)
	}
}

func (s *MemTestSuite) TestObjArrValues() {
	n := 1000

	for idx, el := range tlbObjArr {
		var keysNum int
		res := testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < n; i++ {
				js := el.callFn(el)
				escape(js)
				keysNum = js.Size()
			}
		})
		size := res.MemBytes / uint64(n)
		fmt.Printf("\t%s:\t%d bytes (%d elements)\n", el.name, size, keysNum)
		if idx%2 == 1 {
			fmt.Println("")
		}
		s.Equal(el.size, size, "Size of %s is %d", el.name, size)
	}
}
