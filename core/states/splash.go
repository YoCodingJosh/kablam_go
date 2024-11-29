package states

import (
	"codingjosh.com/kablam/core"
	"github.com/hajimehoshi/ebiten/v2"
)

type SplashState struct{
	core.State
	game *core.Game
}

func (s *SplashState) Update() error {
	return nil
}

func (s *SplashState) Draw(screen *ebiten.Image) {
}
