package test_suite

import (
	"bytes"
	"encoding/json"
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/suite"
	"testing"
)

type JsConverterTestSuite struct {
	suite.Suite
	factory func() djs.JStructOps
	mock    *MockedParser
}

func TestJsConverterTestSuite(t *testing.T) {
	s := new(JsConverterTestSuite)
	s.SetFactory(JsonStructFactory)
	suite.Run(t, s)
}

func (s *JsConverterTestSuite) SetFactory(fn func() djs.JStructOps) {
	s.factory = fn
}

func (s *JsConverterTestSuite) SetupTest() {
	if s.factory == nil {
		panic("factory not provided")
	}
	s.mock = new(MockedParser)
}

func (s *JsConverterTestSuite) TestUnmarshaler() {
	djs.UnmarshalJSON = s.mock.UnmarshalJSONFn
	data := []byte(`{"someKey":"value"}`)
	js := s.factory()
	s.mock.On("UnmarshalJSONFn", data, js).Return(nil)
	e := json.Unmarshal(data, js)
	s.mock.AssertExpectations(s.T())
	s.NoError(e)
	djs.UnmarshalJSON = djs.UnmarshalJSONFn
}

func (s *JsConverterTestSuite) TestUnmarshalUseToParse() {
	djs.JStructParse = s.mock.JStructParseFn
	data := []byte(`{"someKey":"value"}`)
	js := s.factory()
	rd := bytes.NewReader(data)
	s.mock.On("JStructParseFn", rd, js).Return(nil)
	djs.UnmarshalJSON(data, js)
	s.mock.AssertExpectations(s.T())
	djs.JStructParse = djs.JStructParseFn
}

func (s *JsConverterTestSuite) TestMarshaler() {
	djs.MarshalJSON = s.mock.MarshalJSONFn
	data := []byte(`{"someKey":"value"}`)
	js := s.factory()
	s.mock.On("MarshalJSONFn", js).Return(data, nil)
	json.Marshal(js)
	s.mock.AssertExpectations(s.T())
	djs.MarshalJSON = djs.MarshalJSONFn
}

func (s *JsConverterTestSuite) TestMarshalUseSerialize() {
	djs.JStructSerialize = s.mock.JStructSerializeFn
	buf := bytes.NewBuffer([]byte{})
	js := s.factory()
	s.mock.On("JStructSerializeFn", js, buf).Return(nil)
	b, e := djs.MarshalJSON(js)
	s.mock.AssertExpectations(s.T())
	s.Equal(buf.Bytes(), b)
	s.NoError(e)
	djs.JStructSerialize = djs.JStructSerializeFn
}
