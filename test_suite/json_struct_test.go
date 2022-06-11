package test_suite

import (
	djs "github.com/Pencroff/JsonStruct"
	ejs "github.com/Pencroff/JsonStruct/experiment"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestJsonStruct_PrimitiveOpsTestSuite(t *testing.T) {
	s := new(PrimitiveOpsTestSuite)
	s.SetFactory(func() djs.JsonStructPrimitiveOps {
		return &djs.JsonStruct{}
	})
	suite.Run(t, s)
}

func TestJsonStructValue_PrimitiveOpsTestSuite(t *testing.T) {
	s := new(PrimitiveOpsTestSuite)
	s.SetFactory(func() djs.JsonStructPrimitiveOps {
		return &ejs.JsonStructValue{}
	})
	suite.Run(t, s)
}

func TestJsonStructPointer_PrimitiveOpsTestSuite(t *testing.T) {
	s := new(PrimitiveOpsTestSuite)
	s.SetFactory(func() djs.JsonStructPrimitiveOps {
		return &ejs.JsonStructPtr{}
	})
	suite.Run(t, s)
}
