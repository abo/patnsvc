package patnsvc

// Pattern pattern to extract expected from raw
type Pattern struct {
	Expr  string  `json:"expr"`
	Type  string  `json:"type"`
	Score float64 `json:"score"`
}

// Line train data, pattern apply to raw, return expected
type Line struct {
	Raw      string `json:"raw"`
	Expected string `json:"expected"`
}

// Service interface pattern generator
type Service interface {
	Generate([]Line) ([]Pattern, error)
}

//Request http json request
type Request struct {
	Lines []Line `json:"lines"`
}

//Response http json response
type Response struct {
	Patterns []Pattern `json:"patterns"`
}
