package client

type Package struct {
	FieldWidth    float32  `json:"fieldWidth"`
	FieldLength   float32  `json:"fieldLength"`
	BoundaryWidth float32  `json:"boundaryWidth"`
	GoalWidth     float32  `json:"goalWidth"`
	GoalDepth     float32  `json:"goalDepth"`
	Lines         []Line   `json:"lines"`
	Circles       []Circle `json:"circles"`
	Paths         []Path   `json:"paths"`
	Texts         []Text   `json:"texts"`
}

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Style struct {
	Stroke      *string  `json:"stroke,omitempty"`
	StrokeWidth *float32 `json:"strokeWidth,omitempty"`
	Fill        *string  `json:"fill,omitempty"`
	FillOpacity *float32 `json:"fillOpacity,omitempty"`
	Font        *string  `json:"font,omitempty"`
}

type Metadata struct {
	SourceId         string   `json:"sourceId,omitempty"`
	Layer            []string `json:"layer,omitempty"`
	VisibleByDefault bool     `json:"visibleByDefault,omitempty"`
	Order            int32    `json:"order,omitempty"`
}

type Line struct {
	P1 Point `json:"p1"`
	P2 Point `json:"p2"`
	Style
	Metadata
}

type Circle struct {
	Center Point   `json:"center"`
	Radius float32 `json:"radius"`
	Style
	Metadata
}

type Path struct {
	D []PathElement `json:"d"`
	Style
	Metadata
}

type PathElement struct {
	Type string    `json:"type"`
	Args []float64 `json:"args"`
}

type Text struct {
	Text string `json:"text"`
	P    Point  `json:"p"`
	Style
	Metadata
}
