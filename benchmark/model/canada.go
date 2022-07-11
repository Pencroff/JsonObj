package model

import (
	"github.com/francoispqt/gojay"
)

type CanadaModel struct {
	Type     string      `json:"type"`
	Features FeatureList `json:"features"`
}

func (c *CanadaModel) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "type":
		return dec.String(&c.Type)
	case "features":
		return dec.Array(&c.Features)
	}
	return nil
}

func (c *CanadaModel) NKeys() int {
	return 0
}

type FeatureList []*Feature

func (f *FeatureList) UnmarshalJSONArray(dec *gojay.Decoder) error {
	v := &Feature{
		Props:    &Props{},
		Geometry: &Geometry{},
	}
	e := dec.Object(v)
	if e != nil {
		return e
	}
	*f = append(*f, v)
	return nil
}

type Feature struct {
	Type     string    `json:"type"`
	Props    *Props    `json:"properties"`
	Geometry *Geometry `json:"geometry"`
}

func (f *Feature) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "type":
		return dec.String(&f.Type)
	case "properties":
		return dec.Object(f.Props)
	case "geometry":
		return dec.Object(f.Geometry)
	}
	return nil
}

func (f *Feature) NKeys() int {
	return 0
}

type Props struct {
	Name string `json:"name"`
}

func (p *Props) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "name":
		return dec.String(&p.Name)
	}
	return nil
}

func (p *Props) NKeys() int {
	return 0
}

type Geometry struct {
	Type        string          `json:"type"`
	Coordinates PointSliceGroup `json:"coordinates"`
}

func (g *Geometry) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "type":
		return dec.String(&g.Type)
	case "coordinates":
		return dec.Array(&g.Coordinates)
	}
	return nil
}

func (g *Geometry) NKeys() int {
	return 0
}

type PointSliceGroup []*PointSlice

func (p *PointSliceGroup) UnmarshalJSONArray(dec *gojay.Decoder) error {
	v := &PointSlice{}
	e := dec.Array(v)
	if e != nil {
		return e
	}
	*p = append(*p, v)
	return nil
}

func (p *PointSliceGroup) Len() int {
	return len(*p)
}

type PointSlice []*Point

func (p *PointSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	v := &Point{}
	e := dec.Array(v)
	if e != nil {
		return e
	}
	*p = append(*p, v)
	return nil
}

func (p *PointSlice) Len() int {
	return len(*p)
}
func (p *PointSlice) Get(idx int) *Point {
	return (*p)[idx]
}

type Point [2]float64

func (p *Point) UnmarshalJSONArray(dec *gojay.Decoder) error {
	v := 0.0
	e := dec.Float64(&v)
	if e != nil {
		return e
	}
	p[dec.Index()] = v
	return nil
}
