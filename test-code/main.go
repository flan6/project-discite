package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	playerX float64
	playerY float64
	screenX float64
	screenY float64
}

func (g *Game) Update() error {
	// Update the position of the player
	g.playerX += 2.0
	g.playerY += 2.0

	// Update the position of the screen to follow the player
	g.screenX = g.playerX - float64(1920)/2
	g.screenY = g.playerY - float64(1080)/2

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Create a new DrawImageOptions object
	opts := &ebiten.DrawImageOptions{}

	// Translate the options object by the negative position of the player
	opts.GeoM.Translate(-g.playerX+float64(1920)/2, -g.playerY+float64(1080)/2)

	// Draw the player as a red square, applying the screen transform
	rectImage := ebiten.NewImage(32, 32)
	rectImage.Fill(color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff})
	screen.DrawImage(rectImage, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	// Create a new game object with the player starting at position (0, 0)
	game := &Game{playerX: 0, playerY: 0}
	ebiten.SetWindowSize(1920, 1080)

	// Start the game loop
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
