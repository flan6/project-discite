package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/unitoftime/ecs"

	"github.com/flan6/ebiten-go/pkg/components"
)

type Render struct{}

func NewRender() Drawable {
	return Render{}
}

func (Render) Draw(gameWorld GameWorld, screen *ebiten.Image) {
	query := ecs.Query3[components.Position, components.Direction, components.Sprite](gameWorld.World())

	query.MapId(func(id ecs.Id, position *components.Position, direction *components.Direction, sprite *components.Sprite) {
		img := ebiten.Image(*sprite) // TODO : improve
		x, y := img.Size()

		options := &ebiten.DrawImageOptions{}

		options.GeoM.Translate(-float64(x)/2, -float64(y)/2)

		options.GeoM.Rotate(float64(*direction))

		options.GeoM.Translate(position.X, position.Y)

		screen.DrawImage(&img, options)
	})
}
