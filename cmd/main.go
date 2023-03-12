package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/flan6/ebiten-go/pkg/scenes"
)

func main() {
	game := scenes.NewManager()

	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Test Ebiten")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
