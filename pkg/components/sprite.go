package components

import "github.com/hajimehoshi/ebiten/v2"

type Sprite ebiten.Image

func NewSprite(img ebiten.Image) Sprite {
	return Sprite(img)
}
