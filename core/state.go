package core

import "github.com/hajimehoshi/ebiten/v2"

type State interface {
	Update(deltaTime float64) error
	Draw(screen *ebiten.Image)

	Name() string
}
