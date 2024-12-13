package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameplayState struct {
	game *Game
	badGuy *BadGuy
}

func (s *GameplayState) Update(deltaTime float64) error {
	s.badGuy.Update(deltaTime)

	return nil
}

func (s *GameplayState) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100, 149, 237, 255})

	bg_tile := s.game.Assets.Images["wall"]

	// Draw the background (with offset for skybox)
	for x := 0; x < ScreenWidth; x += WallTileSize {
		for y := SkyboxHeight; y < ScreenHeight; y += WallTileSize {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))

			screen.DrawImage(bg_tile, op)
		}
	}

	s.badGuy.Draw(screen)
}

func (s *GameplayState) Name() string {
	return "Gameplay"
}

func NewGameplayState(g *Game) *GameplayState {
	badGuy := NewBadGuy(1, func (x, y float64) {
		println("Bomb drop at", x, y)
	})

	// TODO: call badGuy.Stop to stop the bad guy's goroutines

	return &GameplayState{
		game: g,
		badGuy: badGuy,
	}
}
