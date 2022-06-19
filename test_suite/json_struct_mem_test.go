package test_suite

import (
	"fmt"
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/suite"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestJsonStructMem_Value(t *testing.T) {
	s := new(MemTestSuite)
	s.SetFactory(JsonStructValueFactory)
	suite.Run(t, s)
}

func TestJsonStructMem_Ptr(t *testing.T) {
	s := new(MemTestSuite)
	s.SetFactory(JsonStructPointerFactory)
	suite.Run(t, s)
}

var tmStr = "2015-01-01T12:34:56Z"
var tm, _ = time.Parse(time.RFC3339, tmStr)
var tblMethod = []struct {
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
var tblMethodLen = len(tblMethod)

type MemTestSuite struct {
	suite.Suite
	factory func() djs.JsonStructOps
	js      djs.JsonStructOps
}

func (s *MemTestSuite) SetFactory(fn func() djs.JsonStructOps) {
	s.factory = fn
}

func (s *MemTestSuite) SetupTest() {
	if s.factory == nil {
		panic("factory not provided")
	}
	s.js = s.factory()
}

func (s *MemTestSuite) TestObjectMem() {
	keyNumber := 1000000 * tblMethodLen
	s.js.AsObject()
	runtime.GC()
	PrintMemUsage("Start")
	for i := 0; i < keyNumber; i++ {
		key := "key" + strconv.Itoa(i)
		v := tblMethod[i%tblMethodLen]
		s.js.SetKey(key, v.v)
	}
	PrintMemUsage("End")
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage(name string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Println(name)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
