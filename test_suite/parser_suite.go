package test_suite

import (
	"bytes"
	djs "github.com/Pencroff/JsonStruct"
	"github.com/stretchr/testify/suite"
)

type ParserTestSuite struct {
	suite.Suite
	factory func() djs.JStructOps
	js      djs.JStructOps
}

func (s *ParserTestSuite) SetFactory(fn func() djs.JStructOps) {
	s.factory = fn
}

func (s *ParserTestSuite) SetupTest() {
	if s.factory == nil {
		panic("factory not provided")
	}
	s.js = s.factory()
}

func (s *ParserTestSuite) TestUnmarshalFallDownToParse() {
	mock := new(MockedParser)
	djs.JStructParse = mock.JStructParseReader
	data := []byte(`{"someKey":"value"}`)
	rd := bytes.NewReader(data)
	mock.On("JStructParseReader", rd, s.js).Return(nil)
	djs.UnmarshalJSON(data, s.js)
	mock.AssertExpectations(s.T())
	djs.JStructParse = djs.JStructParseReader
}

func (s *ParserTestSuite) TestParsing_PrimitiveValues() {
	tbl := []struct {
		in  []byte
		out func(v djs.JStructOps) djs.JStructOps
	}{
		{[]byte(`true`),
			func(v djs.JStructOps) djs.JStructOps {
				v.SetBool(true)
				return v
			},
		},
		{[]byte(`false`),
			func(v djs.JStructOps) djs.JStructOps {
				v.SetBool(false)
				return v
			},
		},
		{[]byte(`null`),
			func(v djs.JStructOps) djs.JStructOps {
				v.SetNull()
				return v
			},
		},
		{[]byte(`1`),
			func(v djs.JStructOps) djs.JStructOps {
				v.SetInt(1)
				return v
			},
		},
		{[]byte(`-1`),
			func(v djs.JStructOps) djs.JStructOps {
				v.SetInt(-1)
				return v
			},
		},
	}
	for _, el := range tbl {
		rd := bytes.NewReader(el.in)
		e := djs.JStructParseReader(rd, s.js)
		s.NoError(e)
		s.Equal(el.out(s.factory()), s.js)
	}
}
