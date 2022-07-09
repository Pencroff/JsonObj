package test_suite

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestJStruct_Scanner(t *testing.T) {
	s := new(ScannerTestSuite)
	suite.Run(t, s)
}

func TestJsonStructConverter_ParserSuite(t *testing.T) {
	s := new(ParserTestSuite)
	s.SetFactory(JsonStructFactory)
	suite.Run(t, s)
}

func TestJsonStructConverter_SerializerSuite(t *testing.T) {
	s := new(SerializerTestSuite)
	s.SetFactory(JsonStructFactory)
	suite.Run(t, s)
}
