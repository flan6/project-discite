package entity

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/unitoftime/ecs"

	"github.com/flan6/ebiten-go/pkg/components"
)

type Player struct{}

func NewPlayer(world *ecs.World) ecs.Id {
	curDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	img, _, err := ebitenutil.NewImageFromFile(curDir + "/assets/images/pedestrian/pedestrian.png")
	if err != nil {
		panic(err)
	}

	id := world.NewId()
	ecs.Write(
		world,
		id,
		ecs.C(components.Position{X: 100, Y: 100}),
		ecs.C(components.Direction(0)),
		ecs.C(components.Keybinds{
			Forward:     ebiten.KeyW,
			Backward:    ebiten.KeyS,
			StrafeLeft:  ebiten.KeyA,
			StrafeRight: ebiten.KeyD}), // TODO : get from config
		ecs.C(components.NewSprite(*img)),
	)

	return id
}
