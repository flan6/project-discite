package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/unitoftime/ecs"
)

type System interface {
	Update(GameWorld)
}

type Drawable interface {
	Draw(GameWorld, *ebiten.Image)
}

type GameWorld interface {
	PlayerID() ecs.Id
	World() *ecs.World
}
