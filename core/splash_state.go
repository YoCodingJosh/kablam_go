package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SplashState struct {
	game *Game
	timer float64 // how long to show the splash screen
	show bool // TEMP: until the menu is implemented
}

func (s *SplashState) Update(deltaTime float64) error {
	if !s.show {
		s.game.SetState(NewMenuState(s.game))
	}

	if s.timer >= SplashScreenDuration {
		s.show = false
	} else {
		s.timer += deltaTime
	}

	return nil
}

func (s *SplashState) Draw(screen *ebiten.Image) {
	if s.show {
		logoImage := s.game.Assets.Images["hypeworks_logo"]

		// Center the logo on the screen
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(640 - float64(logoImage.Bounds().Dx())/2, 480 - float64(logoImage.Bounds().Dy())/2)

		screen.DrawImage(s.game.Assets.Images["hypeworks_logo"], op)
	}
}

func (s *SplashState) Name() string {
	return "Splash"
}

func NewSplashState(g *Game) *SplashState {
	return &SplashState{
		game: g,
		timer: 0.0,
		show: true,
	}
}
