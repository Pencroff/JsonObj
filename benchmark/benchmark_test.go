package benchmark

import (
	stdjson "encoding/json"
	"fmt"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/Pencroff/JsonStruct/benchmark/model"
	"github.com/francoispqt/gojay"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"github.com/minio/simdjson-go"
	"testing"
)

func Benchmark_Unmarshal_code(b *testing.B) {
	data, err := ReadData("data/code.json.gz")
	if err != nil {
		b.Fatal(err)
	}
	fmt.Printf("Data size: %.2f Mb\n", float64(len(data))/1024/1024)
	var e error
	b.Run("Std_code", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var o model.CodeModel
			e = stdjson.Unmarshal(data, &o)
		}
		if e != nil {
			b.Fatal(e)
		}
	})
	b.Run("Go_code", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var o model.CodeModel
			e = gojson.Unmarshal(data, &o)
		}
		if e != nil {
			b.Fatal(e)
		}
	})
	b.Run("Iter_code", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var o model.CodeModel
			e = jsoniter.Unmarshal(data, &o)
		}
		if e != nil {
			b.Fatal(e)
		}
	})
	b.Run("JValue_code", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		var v *jsonvalue.V
		for i := 0; i < b.N; i++ {
			v, e = jsonvalue.Unmarshal(data)
		}
		if e != nil {
			b.Fatal(e)
		}
		v.Len()
	})
	b.Run("Jay_code", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var o model.CodeModelJay
			e = gojay.Unmarshal(data, &o)
		}
		if e != nil {
			b.Fatal(e)
		}
	})
	b.Run("Simd_code", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		var pj *simdjson.ParsedJson
		for i := 0; i < b.N; i++ {
			pj, e = simdjson.Parse(data, nil)
		}
		pj.Iter()
		if e != nil {
			b.Fatal(e)
		}
	})
}

func Benchmark_Unmarshal_canada(b *testing.B) {
	data, err := ReadData("data/canada.json.gz")
	if err != nil {
		b.Fatal(err)
	}
	fmt.Printf("Data size: %.2f Mb\n", float64(len(data))/1024/1024)
	var e error
	b.Run("StdJson___canada", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var o model.CanadaModel
			e = stdjson.Unmarshal(data, &o)
		}
		if e != nil {
			b.Fatal(e)
		}
	})
	b.Run("StdSimple_canada", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var o model.CanadaSimpleModel
			e = stdjson.Unmarshal(data, &o)
		}
		if e != nil {
			b.Fatal(e)
		}
	})
	b.Run("Go_canada", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var o model.CanadaModel
			e = gojson.Unmarshal(data, &o)
		}
		if e != nil {
			b.Fatal(e)
		}
	})
	b.Run("GoSimple_canada", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var o model.CanadaSimpleModel
			e = gojson.Unmarshal(data, &o)
		}
		if e != nil {
			b.Fatal(e)
		}
	})
	b.Run("IterSimple_canada", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var o model.CanadaSimpleModel
			e = jsoniter.Unmarshal(data, &o)
		}
		if e != nil {
			b.Fatal(e)
		}
	})
	b.Run("JsonValue_canada", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		var v *jsonvalue.V
		for i := 0; i < b.N; i++ {
			v, e = jsonvalue.Unmarshal(data)
		}
		if e != nil {
			b.Fatal(e)
		}
		v.Len()
	})
	b.Run("Jay_canada", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			var o model.CanadaModel
			e = gojay.Unmarshal(data, &o)
		}
		if e != nil {
			b.Fatal(e)
		}
	})
	b.Run("Simd_canada", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		b.ResetTimer()
		var pj *simdjson.ParsedJson
		for i := 0; i < b.N; i++ {
			pj, e = simdjson.Parse(data, nil)
		}
		pj.Iter()
		if e != nil {
			b.Fatal(e)
		}
	})

}
