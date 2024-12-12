package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameplayState struct {
	game *Game
}

func (s *GameplayState) Update(deltaTime float64) error {
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

	// TODO: draw score, player, bombs, etc.
}

func (s *GameplayState) Name() string {
	return "Gameplay"
}

func NewGameplayState(g *Game) *GameplayState {
	return &GameplayState{
		game: g,
	}
}
