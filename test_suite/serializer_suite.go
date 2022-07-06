package test_suite

import (
	"bytes"
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/suite"
)

type SerializerTestSuite struct {
	suite.Suite
	factory func() djs.JStructOps
	js      djs.JStructOps
}

func (s *SerializerTestSuite) SetFactory(fn func() djs.JStructOps) {
	s.factory = fn
}

func (s *SerializerTestSuite) SetupTest() {
	if s.factory == nil {
		panic("factory not provided")
	}
	s.js = s.factory()
}

func (s *SerializerTestSuite) TestMarshalFallDownToSerialize() {
	mock := new(MockedParser)
	djs.JStructSerialize = mock.JStructSerializeWriter
	buf := bytes.Buffer{}
	mock.On("JStructSerializeWriter", s.js, &buf).Return(nil)
	b, e := djs.MarshalJSON(s.js)
	mock.AssertExpectations(s.T())
	s.Equal(b, buf.Bytes())
	s.NoError(e)
	djs.JStructSerialize = djs.JStructSerializeWriter
}
