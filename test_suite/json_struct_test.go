package test_suite

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestJsonStruct_GeneralOpsTestSuite(t *testing.T) {
	s := new(GeneralOpsTestSuite)
	s.SetFactory(JsonStructFactory)
	suite.Run(t, s)
}

func TestJsonStruct_PrimitiveOpsTestSuite(t *testing.T) {
	s := new(PrimitiveOpsTestSuite)
	s.SetFactory(JsonStructFactory)
	suite.Run(t, s)
}

func TestJsonStruct_ObjectOps(t *testing.T) {
	s := new(ObjectOpsTestSuite)
	s.SetFactory(JsonStructFactory)
	suite.Run(t, s)
}

func TestJsonStruct_ArrayOps(t *testing.T) {
	s := new(ArrayOpsTestSuite)
	s.SetFactory(JsonStructFactory)
	suite.Run(t, s)
}

func TestJsonStruct_ConverterOps(t *testing.T) {
	s := new(JsonConverterTestSuite)
	s.SetFactory(JsonStructConvertibleFactory)
	suite.Run(t, s)
}

//--------------------------------------------------------------------------------------------------------------------
func TestJsonStructValue_GeneralOpsTestSuite(t *testing.T) {
	s := new(GeneralOpsTestSuite)
	s.SetFactory(JsonStructValueFactory)
	suite.Run(t, s)
}

func TestJsonStructValue_PrimitiveOpsTestSuite(t *testing.T) {
	s := new(PrimitiveOpsTestSuite)
	s.SetFactory(JsonStructValueFactory)
	suite.Run(t, s)
}

func TestJsonStructValue_ObjectOps(t *testing.T) {
	s := new(ObjectOpsTestSuite)
	s.SetFactory(JsonStructValueFactory)
	suite.Run(t, s)
}

func TestJsonStructValue_ArrayOps(t *testing.T) {
	s := new(ArrayOpsTestSuite)
	s.SetFactory(JsonStructValueFactory)
	suite.Run(t, s)
}

//--------------------------------------------------------------------------------------------------------------------
func TestJsonStructPointer_GeneralOpsTestSuite(t *testing.T) {
	s := new(GeneralOpsTestSuite)
	s.SetFactory(JsonStructPointerFactory)
	suite.Run(t, s)
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

func TestJsonStructPointer_ArrayOps(t *testing.T) {
	s := new(ArrayOpsTestSuite)
	s.SetFactory(JsonStructPointerFactory)
	suite.Run(t, s)
}
