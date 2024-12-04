package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type MenuState struct {
	game *Game
	promptBlinkTimer float64
	promptVisible bool
}

func (s *MenuState) Update(deltaTime float64) error {
	s.promptBlinkTimer += deltaTime

	if s.promptBlinkTimer >= MenuPromptBlinkDuration {
		s.promptBlinkTimer = 0.0
		s.promptVisible = !s.promptVisible
	}
	return nil
}

func (s *MenuState) Draw(screen *ebiten.Image) {
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

	// Draw the blinking prompt
	if s.promptVisible {
		op := &text.DrawOptions{}
		op.GeoM.Translate(640, 480)
		op.ColorScale.ScaleWithColor(color.White)

		text.Draw(screen, MenuPrompt, &text.GoTextFace{
			Source: s.game.Assets.Fonts["menu"],
			Size: 24,
		}, op)
	}
}

func (s *MenuState) Name() string {
	return "Menu"
}

func NewMenuState(g *Game) *MenuState {
	return &MenuState{
		game: g,
		promptBlinkTimer: 0.0,
		promptVisible: false,
	}
}
