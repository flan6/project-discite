package components

type Direction float64

func NewDirection(x float64) *Direction {
	dir := Direction(x)
	return &dir
}
