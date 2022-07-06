package test_suite

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

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
