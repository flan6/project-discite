package components

import "github.com/hajimehoshi/ebiten/v2"

type Keybinds struct {
	Forward     ebiten.Key
	Backward    ebiten.Key
	StrafeLeft  ebiten.Key
	StrafeRight ebiten.Key

	// TODO : add more keys
}
