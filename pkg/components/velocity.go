package components

type Velocity float64

func NewVelocity(x float64) *Velocity {
	vel := Velocity(x)
	return &vel
}
