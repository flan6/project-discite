package entity

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/unitoftime/ecs"

	"github.com/flan6/ebiten-go/pkg/components"
)

type Car struct{}

func NewCar(world *ecs.World) {
	curDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// TODO: embed images
	img, _, err := ebitenutil.NewImageFromFile(curDir + "assets/images/car/car_white.png")
	if err != nil {
		panic(err)
	}

	id := world.NewId()
	ecs.Write(
		world,
		id,
		ecs.C(components.NewSprite(*img)),
		ecs.C(components.NewPosition(50, 50)),
	)

}

// const (
// 	carMaxSpeed     = 8
// 	carAcceleration = 0.02
// 	carDeceleration = 0.01
// 	turningRate     = 0.05
// )

// func (c *Car) Update() {
// 	if c.speed > 1 || c.speed < -1 {
// 		if ebiten.IsKeyPressed(ebiten.KeyD) {
// 			c.direction += turningRate
// 		}

// 		if ebiten.IsKeyPressed(ebiten.KeyA) {
// 			c.direction -= turningRate
// 		}
// 	}

// 	if ebiten.IsKeyPressed(ebiten.KeyW) {
// 		if c.speed < carMaxSpeed {
// 			c.acceleration += carAcceleration
// 			c.speed += c.acceleration
// 		}
// 	} else {
// 		if c.speed > 0 {
// 			c.acceleration -= carDeceleration

// 			if c.acceleration < 0 {
// 				c.acceleration = 0
// 			}

// 			c.speed -= c.acceleration
// 		}
// 	}

// 	if ebiten.IsKeyPressed(ebiten.KeyS) {
// 		if c.speed > -carMaxSpeed {
// 			c.acceleration += carAcceleration
// 			c.speed -= c.acceleration
// 		}
// 	} else {
// 		if c.speed < 0 {
// 			c.acceleration -= carDeceleration

// 			if c.acceleration < 0 {
// 				c.acceleration = 0
// 			}

// 			c.speed += c.acceleration
// 		}
// 	}

// 	c.position.x += c.speed * math.Cos(c.direction)
// 	c.position.y += c.speed * math.Sin(c.direction)
// }

// func (c *Car) Draw(screen *ebiten.Image) {
// 	x, y := c.image.Size()

// 	options := &ebiten.DrawImageOptions{}
// 	options.GeoM.Translate(-float64(x)/2, -float64(y)/2)
// 	options.GeoM.Rotate(c.direction)
// 	options.GeoM.Translate(c.position.x, c.position.y)

// 	screen.DrawImage(c.image, options)
// }
