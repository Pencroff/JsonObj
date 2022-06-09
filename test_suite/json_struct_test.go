package test_suite

import (
	ejs "github.com/Pencroff/JsonStruct/experiment"
	"github.com/stretchr/testify/suite"
	"testing"
)

// Postpone for final implementation
//func TestJsonStruct_PrimitiveOpsTestSuite(t *testing.T) {
//	s := new(PrimitiveOpsTestSuite)
//	s.SetImplementation(new(js.JsonStruct))
//	suite.Run(t, s)
//}

func TestJsonStructValue_PrimitiveOpsTestSuite(t *testing.T) {
	s := new(PrimitiveOpsTestSuite)
	s.SetImplementation(new(ejs.JsonStructValue))
	suite.Run(t, s)
}

func TestJsonStructPointer_PrimitiveOpsTestSuite(t *testing.T) {
	s := new(PrimitiveOpsTestSuite)
	s.SetImplementation(new(ejs.JsonStructPtr))
	suite.Run(t, s)
}
