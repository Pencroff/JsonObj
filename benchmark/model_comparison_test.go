package benchmark

import (
	stdjson "encoding/json"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/Pencroff/JsonStruct/benchmark/model"
	"github.com/francoispqt/gojay"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestModelComparisonTestSuite(t *testing.T) {
	s := new(ModelComparisonTestSuite)
	suite.Run(t, s)
}

type ModelComparisonTestSuite struct {
	suite.Suite
}

func (s *ModelComparisonTestSuite) SetupTest() {
}

func (s *ModelComparisonTestSuite) TestStdJson_code() {
	data, _ := ReadData("data/code.json.gz")
	var o model.CodeModel
	e := stdjson.Unmarshal(data, &o)
	s.NoError(e)
	s.Equal(3, len(o.Tree.Kids))
	// top
	s.Equal(5, len(o.Tree.Kids[0].Kids))
	s.Equal(4, len(o.Tree.Kids[0].Kids[0].Kids))
	s.Equal(26, len(o.Tree.Kids[0].Kids[0].Kids[0].Kids))
	s.Equal(2, len(o.Tree.Kids[0].Kids[0].Kids[0].Kids[0].Kids))
	s.Equal(10, len(o.Tree.Kids[0].Kids[0].Kids[0].Kids[0].Kids[0].Kids))
	s.Equal(0, len(o.Tree.Kids[0].Kids[0].Kids[0].Kids[0].Kids[0].Kids[0].Kids))
	lf := o.Tree.Kids[0].Kids[0].Kids[0].Kids[0].Kids[0].Kids[0]
	s.Equal(0.1, lf.CLWeight)
	s.Equal(int64(1316289444), lf.MeanT)
	// end
	s.Equal(1, len(o.Tree.Kids[2].Kids))
	s.Equal(22, len(o.Tree.Kids[2].Kids[0].Kids))
	s.Equal(0, len(o.Tree.Kids[2].Kids[0].Kids[21].Kids))

	lv := o.Tree.Kids[2].Kids[0].Kids[21]
	s.Equal("WATCHLISTS", lv.Name)
	s.Equal(1.0063291139240507, lv.CLWeight)
	s.Equal(int64(1247778088), lv.MeanT)
}

func (s *ModelComparisonTestSuite) TestGoJson_code() {
	data, _ := ReadData("data/code.json.gz")
	var o model.CodeModel
	e := gojson.Unmarshal(data, &o)
	s.NoError(e)
	s.Equal(3, len(o.Tree.Kids))
	// top
	s.Equal(5, len(o.Tree.Kids[0].Kids))
	s.Equal(4, len(o.Tree.Kids[0].Kids[0].Kids))
	s.Equal(26, len(o.Tree.Kids[0].Kids[0].Kids[0].Kids))
	s.Equal(2, len(o.Tree.Kids[0].Kids[0].Kids[0].Kids[0].Kids))
	s.Equal(10, len(o.Tree.Kids[0].Kids[0].Kids[0].Kids[0].Kids[0].Kids))
	s.Equal(0, len(o.Tree.Kids[0].Kids[0].Kids[0].Kids[0].Kids[0].Kids[0].Kids))
	lf := o.Tree.Kids[0].Kids[0].Kids[0].Kids[0].Kids[0].Kids[0]
	s.Equal(0.1, lf.CLWeight)
	s.Equal(int64(1316289444), lf.MeanT)

	// end
	s.Equal(1, len(o.Tree.Kids[2].Kids))
	s.Equal(22, len(o.Tree.Kids[2].Kids[0].Kids))
	s.Equal(0, len(o.Tree.Kids[2].Kids[0].Kids[21].Kids))
	lv := o.Tree.Kids[2].Kids[0].Kids[21]
	s.Equal("WATCHLISTS", lv.Name)
	s.Equal(1.0063291139240507, lv.CLWeight)
	s.Equal(int64(1247778088), lv.MeanT)
}

func (s *ModelComparisonTestSuite) TestJsonIter_code() {
	data, _ := ReadData("data/code.json.gz")
	var o model.CodeModel
	e := jsoniter.Unmarshal(data, &o)
	s.NoError(e)
	s.Equal(3, len(o.Tree.Kids))
	// top
	s.Equal(5, len(o.Tree.Kids[0].Kids))
	s.Equal(4, len(o.Tree.Kids[0].Kids[0].Kids))
	s.Equal(26, len(o.Tree.Kids[0].Kids[0].Kids[0].Kids))
	s.Equal(2, len(o.Tree.Kids[0].Kids[0].Kids[0].Kids[0].Kids))
	s.Equal(10, len(o.Tree.Kids[0].Kids[0].Kids[0].Kids[0].Kids[0].Kids))
	s.Equal(0, len(o.Tree.Kids[0].Kids[0].Kids[0].Kids[0].Kids[0].Kids[0].Kids))
	lf := o.Tree.Kids[0].Kids[0].Kids[0].Kids[0].Kids[0].Kids[0]
	s.Equal(0.1, lf.CLWeight)
	s.Equal(int64(1316289444), lf.MeanT)

	// end
	s.Equal(1, len(o.Tree.Kids[2].Kids))
	s.Equal(22, len(o.Tree.Kids[2].Kids[0].Kids))
	s.Equal(0, len(o.Tree.Kids[2].Kids[0].Kids[21].Kids))
	lv := o.Tree.Kids[2].Kids[0].Kids[21]
	s.Equal("WATCHLISTS", lv.Name)
	s.Equal(1.0063291139240507, lv.CLWeight)
	s.Equal(int64(1247778088), lv.MeanT)
}

func (s *ModelComparisonTestSuite) TestJsonValue_code() {
	data, _ := ReadData("data/code.json.gz")

	v, e := jsonvalue.Unmarshal(data)
	s.NoError(e)
	v1, e := v.GetArray("tree", "kids")
	s.Equal(3, v1.Len())
	// top
	v2, e := v1.GetArray(0, "kids")
	s.Equal(5, v2.Len())
	v3, e := v2.GetArray(0, "kids")
	s.Equal(4, v3.Len())
	v4, e := v3.GetArray(0, "kids")
	s.Equal(26, v4.Len())
	v5, e := v4.GetArray(0, "kids")
	s.Equal(2, v5.Len())
	v6, e := v5.GetArray(0, "kids")
	s.Equal(10, v6.Len())
	v7, e := v6.GetArray(0, "kids")
	s.Equal(0, v7.Len())
	lf := v.MustGet("tree", "kids", 0, "kids", 0, "kids", 0, "kids", 0, "kids", 0, "kids", 0)
	s.Equal(0.1, lf.MustGet("cl_weight").Float64())
	s.Equal(int64(1316289444), lf.MustGet("mean_t").Int64())

	// end
	v8, e := v1.GetArray(2, "kids")
	s.Equal(1, v8.Len())
	v9, e := v8.GetArray(0, "kids")
	s.Equal(22, v9.Len())
	v10, e := v9.GetArray(21, "kids")
	s.Equal(0, v10.Len())
	lv := v.MustGet("tree", "kids", 2, "kids", 0, "kids", 21)
	s.Equal("WATCHLISTS", lv.MustGet("name").String())
	s.Equal(1.0063291139240507, lv.MustGet("cl_weight").Float64())
	s.Equal(int64(1247778088), lv.MustGet("mean_t").Int64())
}

func (s *ModelComparisonTestSuite) TestGoJay_code() {
	data, _ := ReadData("data/code.json.gz")
	var o model.CodeModelJay
	e := gojay.Unmarshal(data, &o)
	s.NoError(e)
	s.Equal(3, o.Tree.Kids.Len())
	// top
	s.Equal(5, o.Tree.Kids.Get(0).Kids.Len())
	s.Equal(4, o.Tree.Kids.Get(0).Kids.Get(0).Kids.Len())
	s.Equal(26, o.Tree.Kids.Get(0).Kids.Get(0).Kids.Get(0).Kids.Len())
	s.Equal(2, o.Tree.Kids.Get(0).Kids.Get(0).Kids.Get(0).Kids.Get(0).Kids.Len())
	s.Equal(10, o.Tree.Kids.Get(0).Kids.Get(0).Kids.Get(0).Kids.Get(0).Kids.Get(0).Kids.Len())
	s.Equal(0, o.Tree.Kids.Get(0).Kids.Get(0).Kids.Get(0).Kids.Get(0).Kids.Get(0).Kids.Get(0).Kids.Len())
	lf := o.Tree.Kids.Get(0).Kids.Get(0).Kids.Get(0).Kids.Get(0).Kids.Get(0).Kids.Get(0)
	s.Equal(0.1, lf.CLWeight)
	s.Equal(int64(1316289444), lf.MeanT)

	// end
	s.Equal(1, o.Tree.Kids.Get(2).Kids.Len())
	s.Equal(22, o.Tree.Kids.Get(2).Kids.Get(0).Kids.Len())
	s.Equal(0, o.Tree.Kids.Get(2).Kids.Get(0).Kids.Get(21).Kids.Len())
	lv := o.Tree.Kids.Get(2).Kids.Get(0).Kids.Get(21)
	s.Equal("WATCHLISTS", lv.Name)
	s.InDelta(1.0063291139240507, lv.CLWeight, 3e-16) // Go Jay didn't parse float64s correctly
	s.Equal(int64(1247778088), lv.MeanT)
}

func (s *ModelComparisonTestSuite) TestStdJson_canada() {
	data, _ := ReadData("data/canada.json.gz")
	var o model.CanadaModel
	e := stdjson.Unmarshal(data, &o)
	s.NoError(e)
	s.Equal("Canada", o.Features[0].Props.Name)
	s.Equal("Polygon", o.Features[0].Geometry.Type)
	s.Equal(480, len(o.Features[0].Geometry.Coordinates))
	l1 := o.Features[0].Geometry.Coordinates[479]
	l2 := o.Features[0].Geometry.Coordinates[479].Get(5275)
	s.Equal(5276, l1.Len())
	s.Equal(2, len(l2))
	s.Equal(&model.Point{-59.841667000000029, 43.918602000000021}, o.Features[0].Geometry.Coordinates[1].Get(1))
	s.Equal(&model.Point{-70.111937999999952, 83.109421000000111}, o.Features[0].Geometry.Coordinates[479].Get(5275))
}

func (s *ModelComparisonTestSuite) TestStdJsonSimple_canada() {
	data, _ := ReadData("data/canada.json.gz")
	var o model.CanadaSimpleModel
	e := stdjson.Unmarshal(data, &o)
	s.NoError(e)
	s.Equal("Canada", o.Features[0].Props.Name)
	s.Equal("Polygon", o.Features[0].Geometry.Type)
	s.Equal(480, len(o.Features[0].Geometry.Coordinates))
	l1 := o.Features[0].Geometry.Coordinates[479]
	l2 := o.Features[0].Geometry.Coordinates[479][5275]
	s.Equal(5276, len(l1))
	s.Equal(2, len(l2))
	s.Equal(&model.PointSimple{-59.841667000000029, 43.918602000000021}, o.Features[0].Geometry.Coordinates[1][1])
	s.Equal(&model.PointSimple{-70.111937999999952, 83.109421000000111}, o.Features[0].Geometry.Coordinates[479][5275])
}

func (s ModelComparisonTestSuite) TestGoJson_canada() {
	data, _ := ReadData("data/canada.json.gz")
	var o model.CanadaModel
	e := gojson.Unmarshal(data, &o)
	s.NoError(e)
	s.Equal("Canada", o.Features[0].Props.Name)
	s.Equal("Polygon", o.Features[0].Geometry.Type)
	s.Equal(480, len(o.Features[0].Geometry.Coordinates))
	l1 := o.Features[0].Geometry.Coordinates[479]
	l2 := o.Features[0].Geometry.Coordinates[479].Get(5275)
	s.Equal(5276, l1.Len())
	s.Equal(2, len(l2))
	s.Equal(&model.Point{-59.841667000000029, 43.918602000000021}, o.Features[0].Geometry.Coordinates[1].Get(1))
	s.Equal(&model.Point{-70.111937999999952, 83.109421000000111}, o.Features[0].Geometry.Coordinates[479].Get(5275))
}

func (s *ModelComparisonTestSuite) TestGoJsonSimple_canada() {
	data, _ := ReadData("data/canada.json.gz")
	var o model.CanadaSimpleModel
	e := gojson.Unmarshal(data, &o)
	s.NoError(e)
	s.Equal("Canada", o.Features[0].Props.Name)
	s.Equal("Polygon", o.Features[0].Geometry.Type)
	s.Equal(480, len(o.Features[0].Geometry.Coordinates))
	l1 := o.Features[0].Geometry.Coordinates[479]
	l2 := o.Features[0].Geometry.Coordinates[479][5275]
	s.Equal(5276, len(l1))
	s.Equal(2, len(l2))
	s.Equal(&model.PointSimple{-59.841667000000029, 43.918602000000021}, o.Features[0].Geometry.Coordinates[1][1])
	s.Equal(&model.PointSimple{-70.111937999999952, 83.109421000000111}, o.Features[0].Geometry.Coordinates[479][5275])
}

func (s ModelComparisonTestSuite) TestGoJay_canada() {
	data, _ := ReadData("data/canada.json.gz")
	var o model.CanadaModel
	e := gojay.Unmarshal(data, &o)
	s.NoError(e)
	s.Equal("Canada", o.Features[0].Props.Name)
	s.Equal("Polygon", o.Features[0].Geometry.Type)
	s.Equal(480, len(o.Features[0].Geometry.Coordinates))
	l1 := o.Features[0].Geometry.Coordinates[479]
	l2 := o.Features[0].Geometry.Coordinates[479].Get(5275)
	s.Equal(5276, l1.Len())
	s.Equal(2, len(l2))
	s.Equal(&model.Point{-59.841667000000029, 43.918602000000021}, o.Features[0].Geometry.Coordinates[1].Get(1))
	s.Equal(&model.Point{-70.111937999999952, 83.109421000000111}, o.Features[0].Geometry.Coordinates[479].Get(5275))
}
