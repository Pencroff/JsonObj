package test_suite

import (
	djs "github.com/Pencroff/JsonStruct"
	ejs "github.com/Pencroff/JsonStruct/experiment"
	"github.com/stretchr/testify/suite"
	"testing"
)

//func JsonStructFactory() djs.JsonStructOps {
//	return &djs.JsonStruct{}
//}
//
//func TestJsonStruct_PrimitiveOpsTestSuite(t *testing.T) {
//	s := new(PrimitiveOpsTestSuite)
//	s.SetFactory(JsonStructFactory)
//	suite.Run(t, s)
//}
//
//func JsonStructValueFactory() djs.JsonStructOps {
//	return &ejs.JsonStructValue{}
//}
//
//func TestJsonStructValue_PrimitiveOpsTestSuite(t *testing.T) {
//	s := new(PrimitiveOpsTestSuite)
//	s.SetFactory(JsonStructValueFactory)
//	suite.Run(t, s)
//}

func JsonStructPointerFactory() djs.JsonStructOps {
	return &ejs.JsonStructPtr{}
}

func TestJsonStructPointer_PrimitiveOpsTestSuite(t *testing.T) {
	s := new(PrimitiveOpsTestSuite)
	s.SetFactory(JsonStructPointerFactory)
	suite.Run(t, s)
}

func TestJsonStructPointer_ObjectOps(t *testing.T) {
	s := new(ObjectOpsTestSuite)
	s.SetFactory(JsonStructPointerFactory)
	suite.Run(t, s)
}
