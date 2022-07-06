package comparison

import (
	"fmt"
	djs "github.com/Pencroff/JsonStruct"
	ejs "github.com/Pencroff/JsonStruct/experiment"
	"github.com/Pencroff/JsonStruct/helper"
	"reflect"
	"strings"
	"testing"
	"time"
)

var intResult int

func BenchmarkPrimitiveOps_Set(b *testing.B) {
	tmStr := "2019-01-01T10:15:20Z"
	tm, _ := time.Parse(time.RFC3339, tmStr)
	vjs := ejs.JsonStructValue{}
	pjs := ejs.JsonStructPtr{}
	PrintSize(&vjs)
	PrintSize(&pjs)

	b.Run("SetBool Value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			vjs.SetBool(true)
			vjs.SetBool(false)
		}
		if vjs.Bool() {
			intResult = 1
		}
	})
	b.Run("SetBool Ptr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pjs.SetBool(true)
			pjs.SetBool(false)
		}
		if pjs.Bool() {
			intResult = 1
		}
	})
	fmt.Println("")
	b.Run("SetInt Value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			vjs.SetInt(helper.MaxInt)
		}
		intResult = int(vjs.Int())
	})
	b.Run("SetInt Ptr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pjs.SetInt(helper.MaxInt)
		}
		intResult = int(pjs.Int())
	})
	fmt.Println("")
	b.Run("SetUint Value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			vjs.SetUint(helper.MaxUint)
		}
		intResult = int(vjs.Int())
	})
	b.Run("SetUint Ptr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pjs.SetUint(helper.MaxUint)
		}
		intResult = int(pjs.Int())
	})
	fmt.Println("")
	b.Run("SetFloat Value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			vjs.SetFloat(3.14159)
		}
		intResult = int(vjs.Int())
	})
	b.Run("SetFloat Ptr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pjs.SetFloat(3.14159)
		}
		intResult = int(pjs.Int())
	})
	fmt.Println("")
	b.Run("SetString Value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			vjs.SetString("3.14159")
		}
		intResult = int(vjs.Int())
	})
	b.Run("SetString Ptr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pjs.SetString("3.14159")
		}
		intResult = int(pjs.Int())
	})
	fmt.Println("")
	b.Run("SetTime Value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			vjs.SetTime(tm)
		}
		intResult = int(vjs.Int())
	})
	b.Run("SetTime Ptr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pjs.SetTime(tm)
		}
		intResult = int(pjs.Int())
	})

}

func BenchmarkPrimitiveOps_Read(b *testing.B) {
	tm, _ := time.Parse(time.RFC3339, "2019-01-01T10:15:20Z")
	vjs := ejs.JsonStructValue{}
	pjs := ejs.JsonStructPtr{}
	tblObj := []struct {
		name string
		o    djs.JStructOps
	}{
		{"Val", &vjs},
		{"Ptr", &pjs},
	}
	tblMethod := []struct {
		method string
		extra  string
		v      interface{}
	}{
		{"SetBool", "true", true},
		{"SetBool", "false", false},
		{"SetInt", "max", helper.MaxInt},
		{"SetInt", "min", helper.MinInt},
		{"SetUint", "", helper.MaxUint},
		{"SetFloat", "pos", 3.14159},
		{"SetFloat", "neg", -3.14159},
		{"SetString", "sm", "Hello World"},
		{"SetString", "md", "One morning, when Gregor Samsa woke from troubled dream"},
		{"SetString", "lg", "The quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymphs grab quick-jived waltz. Brick quiz whangs jumpy veldt fox. Bright vixens jump; dozy fowl quack. Quick wafting zephyrs vex bold Jim. Quick zephyrs blow, vexing daft Jim. Sex-charged fop blew my junk TV quiz. How quickly daft jumping zebras vex. Two driven jocks help fax my big quiz. Quick, Baz, get my woven flax jodhpurs! \"Now fax quiz Jack!\" my b"},
		{"SetString", "fl", "31415926535.897932385"},
		{"SetString", "tm", "2015-05-14T12:34:56+02:00"},
		{"SetTime", "", tm},
	}

	for _, t := range tblObj {
		for _, m := range tblMethod {
			helper.CallMethod(t.o, m.method, m.v)
			nameFn := CreateSetPrimitiveNameFn(t.name, m.method, m.extra)

			b.Run(nameFn("Bool"), func(b *testing.B) {
				n := false
				for i := 0; i < b.N; i++ {
					n = t.o.Bool()
				}
				if n {
					intResult = 1
				}
			})

			b.Run(nameFn("Int"), func(b *testing.B) {
				n := int64(0)
				for i := 0; i < b.N; i++ {
					n = t.o.Int()
				}
				intResult = int(n)
			})

			b.Run(nameFn("Uint"), func(b *testing.B) {
				n := uint64(0)
				for i := 0; i < b.N; i++ {
					n = t.o.Uint()
				}
				intResult = int(n)
			})

			b.Run(nameFn("Float"), func(b *testing.B) {
				n := 0.0
				for i := 0; i < b.N; i++ {
					n = t.o.Float()
				}
				intResult = int(n)
			})

			b.Run(nameFn("String"), func(b *testing.B) {
				n := ""
				for i := 0; i < b.N; i++ {
					n = t.o.String()
				}
				intResult = len(n)
			})

			b.Run(nameFn("Time"), func(b *testing.B) {
				n := time.Time{}
				for i := 0; i < b.N; i++ {
					n = t.o.Time()
				}
				intResult = int(n.Unix())
			})
			fmt.Println("")
		}
	}
}

func BenchmarkPrimitiveOps_Check(b *testing.B) {
	tm, _ := time.Parse(time.RFC3339, "2019-01-01T10:15:20Z")
	vjs := ejs.JsonStructValue{}
	pjs := ejs.JsonStructPtr{}
	tblObj := []struct {
		name string
		o    djs.JStructOps
	}{
		{"Val", &vjs},
		{"Ptr", &pjs},
	}
	tblMethod := []struct {
		method string
		extra  string
		v      interface{}
	}{
		{"SetBool", "true", true},
		{"SetBool", "false", false},
		{"SetInt", "max", helper.MaxInt},
		{"SetInt", "min", helper.MinInt},
		{"SetUint", "", helper.MaxUint},
		{"SetFloat", "pos", 3.14159},
		{"SetFloat", "neg", -3.14159},
		{"SetString", "sm", "Hello World"},
		{"SetString", "md", "One morning, when Gregor Samsa woke from troubled dream"},
		{"SetString", "lg", "The quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymphs grab quick-jived waltz. Brick quiz whangs jumpy veldt fox. Bright vixens jump; dozy fowl quack. Quick wafting zephyrs vex bold Jim. Quick zephyrs blow, vexing daft Jim. Sex-charged fop blew my junk TV quiz. How quickly daft jumping zebras vex. Two driven jocks help fax my big quiz. Quick, Baz, get my woven flax jodhpurs! \"Now fax quiz Jack!\" my b"},
		{"SetString", "fl", "31415926535.897932385"},
		{"SetString", "tm", "2015-05-14T12:34:56+02:00"},
		{"SetTime", "", tm},
	}

	for _, t := range tblObj {
		for _, m := range tblMethod {

			helper.CallMethod(t.o, m.method, m.v)
			nameFn := CreateSetPrimitiveNameFn(t.name, m.method, m.extra)

			b.Run(nameFn("IsBool"), func(b *testing.B) {
				n := false
				for i := 0; i < b.N; i++ {
					n = t.o.IsBool()
				}
				if n {
					intResult = 1
				}
			})

			b.Run(nameFn("IsInt"), func(b *testing.B) {
				n := false
				for i := 0; i < b.N; i++ {
					n = t.o.IsInt()
				}
				if n {
					intResult = 1
				}
			})

			b.Run(nameFn("IsUint"), func(b *testing.B) {
				n := false
				for i := 0; i < b.N; i++ {
					n = t.o.IsUint()
				}
				if n {
					intResult = 1
				}
			})

			b.Run(nameFn("IsFloat"), func(b *testing.B) {
				n := false
				for i := 0; i < b.N; i++ {
					n = t.o.IsFloat()
				}
				if n {
					intResult = 1
				}
			})

			b.Run(nameFn("IsString"), func(b *testing.B) {
				n := false
				for i := 0; i < b.N; i++ {
					n = t.o.IsString()
				}
				if n {
					intResult = 1
				}
			})

			b.Run(nameFn("IsTime"), func(b *testing.B) {
				n := false
				for i := 0; i < b.N; i++ {
					n = t.o.IsTime()
				}
				if n {
					intResult = 1
				}
			})
			fmt.Println("")
		}
	}
}

func PrintSize(v interface{}) {
	fmt.Printf("Size: %v\n", reflect.Indirect(reflect.ValueOf(v)).Type().Size())
}

func CreateSetPrimitiveNameFn(name, setMethod, extra string) func(prefix string) string {
	return func(dataMethod string) string {
		return strings.TrimSpace(fmt.Sprintf("%s--%s_%s => %s", name, setMethod, extra, dataMethod))
	}
}
