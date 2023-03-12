package systems

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/unitoftime/ecs"

	"github.com/flan6/ebiten-go/pkg/components"
)

const playerSpeed = 2

type HandlePlayerInput struct{}

func (HandlePlayerInput) Update(gameWorld GameWorld) {
	query := ecs.Query3[components.Direction, components.Position, components.Keybinds](gameWorld.World())

	query.MapId(func(id ecs.Id, direction *components.Direction, position *components.Position, keybinds *components.Keybinds) {
		if id == gameWorld.PlayerID() {
			mouseX, mouseY := ebiten.CursorPosition()

			*direction = *components.NewDirection(
				math.Atan2(
					float64(mouseY)-position.Y,
					float64(mouseX)-position.X))

			if ebiten.IsKeyPressed(keybinds.Forward) {
				position.X += playerSpeed * math.Cos(float64(*direction))
				position.Y += playerSpeed * math.Sin(float64(*direction))
			}

			if ebiten.IsKeyPressed(keybinds.Backward) {
				position.X -= playerSpeed * math.Cos(float64(*direction))
				position.Y -= playerSpeed * math.Sin(float64(*direction))
			}

			if ebiten.IsKeyPressed(keybinds.StrafeLeft) {
				position.X += math.Cos(float64(*direction) + math.Pi/2)
				position.Y += math.Sin(float64(*direction) + math.Pi/2)
			}

			if ebiten.IsKeyPressed(keybinds.StrafeRight) {
				position.X -= math.Cos(float64(*direction) + math.Pi/2)
				position.Y -= math.Sin(float64(*direction) + math.Pi/2)
			}
		}
	})

}
