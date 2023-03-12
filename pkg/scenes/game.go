package scenes

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/unitoftime/ecs"

	"github.com/flan6/ebiten-go/pkg/entity"
	"github.com/flan6/ebiten-go/pkg/systems"
)

type GameScene struct {
	world     *ecs.World
	playerId  ecs.Id
	systems   []systems.System
	drawables []systems.Drawable
}

func NewGameScene() *GameScene {
	world := ecs.NewWorld()

	// entities
	playerId := entity.NewPlayer(world)

	return &GameScene{
		world:    world,
		playerId: playerId,
		systems: []systems.System{
			systems.HandlePlayerInput{},
		},
		drawables: []systems.Drawable{
			systems.Render{},
		},
	}
}

func (g *GameScene) Update() {
	for _, system := range g.systems {
		system.Update(g)
	}
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()))

	for _, drawable := range g.drawables {
		drawable.Draw(g, screen)
	}
}

func (g *GameScene) PlayerID() ecs.Id {
	return g.playerId
}

func (g *GameScene) World() *ecs.World {
	return g.world
}
