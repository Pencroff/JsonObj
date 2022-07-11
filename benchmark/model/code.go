package model

import (
	"github.com/francoispqt/gojay"
)

type CodeModel struct {
	Tree     *CodeNode `json:"tree"`
	Username string    `json:"username"`
}

type CodeNode struct {
	Name     string      `json:"name"`
	Kids     []*CodeNode `json:"kids"`
	CLWeight float64     `json:"cl_weight"`
	Touches  int         `json:"touches"`
	MinT     int64       `json:"min_t"`
	MaxT     int64       `json:"max_t"`
	MeanT    int64       `json:"mean_t"`
}

type CodeModelJay struct {
	Tree     *CodeNodeJay `json:"tree"`
	Username string       `json:"username"`
}

func (c *CodeModelJay) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "tree":
		if c.Tree == nil {
			c.Tree = &CodeNodeJay{
				Name:     "",
				Kids:     &CodeNodeSlice{},
				CLWeight: 0,
				Touches:  0,
				MinT:     0,
				MaxT:     0,
				MeanT:    0,
			}
		}
		return dec.Object(c.Tree)
	case "username":
		return dec.String(&c.Username)
	}
	return nil
}

func (c *CodeModelJay) NKeys() int {
	return 0
}

type CodeNodeJay struct {
	Name     string         `json:"name"`
	Kids     *CodeNodeSlice `json:"kids"`
	CLWeight float64        `json:"cl_weight"`
	Touches  int            `json:"touches"`
	MinT     int64          `json:"min_t"`
	MaxT     int64          `json:"max_t"`
	MeanT    int64          `json:"mean_t"`
}

func (c *CodeNodeJay) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "name":
		return dec.String(&c.Name)
	case "kids":
		return dec.Array(c.Kids)
	case "cl_weight":
		return dec.Float64(&c.CLWeight)
	case "touches":
		return dec.Int(&c.Touches)
	case "min_t":
		return dec.Int64(&c.MinT)
	case "max_t":
		return dec.Int64(&c.MaxT)
	case "mean_t":
		return dec.Int64(&c.MeanT)
	}
	return nil
}

func (c *CodeNodeJay) NKeys() int {
	return 0
}

type CodeNodeSlice []*CodeNodeJay

func (c *CodeNodeSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	n := &CodeNodeJay{
		Name:     "",
		Kids:     &CodeNodeSlice{},
		CLWeight: 0,
		Touches:  0,
		MinT:     0,
		MaxT:     0,
		MeanT:    0,
	}
	e := dec.Object(n)
	if e != nil {
		return e
	}
	*c = append(*c, n)
	return nil
}

func (c *CodeNodeSlice) Len() int {
	return len(*c)
}

func (c *CodeNodeSlice) Get(idx int) *CodeNodeJay {
	return (*c)[idx]
}
