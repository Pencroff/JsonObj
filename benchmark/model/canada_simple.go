package model

type CanadaSimpleModel struct {
	Type     string           `json:"type"`
	Features []*FeatureSimple `json:"features"`
}

type FeatureSimple struct {
	Type     string          `json:"type"`
	Props    *PropsSimple    `json:"properties"`
	Geometry *GeometrySimple `json:"geometry"`
}

type PropsSimple struct {
	Name string `json:"name"`
}

type GeometrySimple struct {
	Type        string           `json:"type"`
	Coordinates [][]*PointSimple `json:"coordinates"`
}

type PointSimple [2]float64
