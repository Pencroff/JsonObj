package test_suite

import (
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/mock"
	"io"
)

type MockedParser struct {
	mock.Mock
}

func (p *MockedParser) JStructParse(bytes []byte, v djs.JStructOps) error {
	args := p.Called(bytes, v)
	return args.Error(0)
}

func (p *MockedParser) JStructSerialize(v djs.JStructOps) ([]byte, error) {
	args := p.Called(v)
	return args.Get(0).([]byte), args.Error(1)
}

func (p *MockedParser) JStructParseReader(rd io.Reader, v djs.JStructOps) error {
	args := p.Called(rd, v)
	return args.Error(0)
}

func (p *MockedParser) JStructSerializeWriter(v djs.JStructOps, wr io.Writer) error {
	args := p.Called(v, wr)
	return args.Error(0)
}
