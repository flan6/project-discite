package components

type Acceleration float64

func NewAcceleration(x float64) *Acceleration {
	acceleration := Acceleration(x)
	return &acceleration
}
