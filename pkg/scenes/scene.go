package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update()
	Draw(*ebiten.Image)
}

type Manager struct {
	Scene
}

func NewManager() *Manager {
	return &Manager{
		Scene: NewGameScene(),
	}
}

func (m *Manager) switchTo(scene Scene) {
	m.Scene = scene
}

func (m *Manager) Update() error {
	m.Scene.Update()

	return nil
}

func (m *Manager) Draw(screen *ebiten.Image) {
	m.Scene.Draw(screen)
}

func (g *Manager) Layout(w, h int) (int, int) {
	return w, h
}
