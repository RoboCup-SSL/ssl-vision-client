package vision

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
	Stroke      *string  `json:"stroke"`
	StrokeWidth *float32 `json:"strokeWidth"`
	Fill        *string  `json:"fill"`
	FillOpacity *float32 `json:"fillOpacity"`
}

type Line struct {
	P1 Point `json:"p1"`
	P2 Point `json:"p2"`
	Style
}

type Circle struct {
	Center Point   `json:"center"`
	Radius float32 `json:"radius"`
	Style
}

type Path struct {
	D []PathElement `json:"d"`
	Style
}

type PathElement struct {
	Type string    `json:"type"`
	Args []float64 `json:"args"`
}

type Text struct {
	Text       string  `json:"text"`
	P          Point   `json:"p"`
	D          Point   `json:"d"`
	TextLength float32 `json:"textLength"`
	Font       string  `json:"font"`
	Style
}
