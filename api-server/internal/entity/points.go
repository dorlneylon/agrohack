package entity

type Points struct {
	Timestamp int     `json:"timestamp"`
	Data      []Point `json:"data"`
}

type Point struct {
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Name string  `json:"name"`
}
