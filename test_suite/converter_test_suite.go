package test_suite

import (
	"encoding/json"
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/suite"
)

type JsonConverterTestSuite struct {
	suite.Suite
	factory func() djs.JStructConvertibleOps
	js      djs.JStructConvertibleOps
	mock    *MockedParser
	data    []byte
}

func (s *JsonConverterTestSuite) SetFactory(fn func() djs.JStructConvertibleOps) {
	s.factory = fn
}

func (s *JsonConverterTestSuite) SetupTest() {
	if s.factory == nil {
		panic("factory not provided")
	}
	s.js = s.factory()
	s.mock = new(MockedParser)
	s.data = []byte(`{"someKey":"value"}`)
}

func (s *JsonConverterTestSuite) TestUnmarshaler() {
	djs.UnmarshalJSON = s.mock.JStructParse
	s.mock.On("JStructParse", s.data, s.js).Return(nil)
	json.Unmarshal(s.data, s.js)
	s.mock.AssertExpectations(s.T())
}

func (s *JsonConverterTestSuite) TestMarshaler() {
	djs.MarshalJSON = s.mock.JStructSerialize
	s.mock.On("JStructSerialize", s.js).Return(s.data, nil)
	json.Marshal(s.js)
	s.mock.AssertExpectations(s.T())
}
