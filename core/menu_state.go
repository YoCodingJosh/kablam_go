package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type MenuState struct {
	game *Game
}

func (s *MenuState) Update(deltaTime float64) error {
	return nil
}

func (s *MenuState) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100, 149, 237, 255})

	bg_tile := s.game.Assets.Images["wall"]

	// TODO: account for the score and bad guy at the top of the screen

	for x := 0; x < ScreenWidth; x += WallTileSize {
		for y := 0; y < ScreenHeight; y += WallTileSize {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))

			screen.DrawImage(bg_tile, op)
		}
	}
}

func (s *MenuState) Name() string {
	return "Menu"
}

func NewMenuState(g *Game) *MenuState {
	return &MenuState{
		game: g,
	}
}
