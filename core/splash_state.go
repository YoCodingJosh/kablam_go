package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SplashState struct{
	game *Game
}

func (s *SplashState) Update() error {
	return nil
}

func (s *SplashState) Draw(screen *ebiten.Image) {
	logoImage := s.game.Assets.Images["hypeworks_logo"]

	// Center the logo on the screen
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(640 - float64(logoImage.Bounds().Dx())/2, 480 - float64(logoImage.Bounds().Dy())/2)

	screen.DrawImage(s.game.Assets.Images["hypeworks_logo"], op)
}

func (s *SplashState) Name() string {
	return "Splash"
}

func NewSplashState(g *Game) *SplashState {
	return &SplashState{
		game: g,
	}
}
