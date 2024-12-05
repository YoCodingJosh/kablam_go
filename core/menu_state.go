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

	// Draw title and copyright text
	titleFont := &text.GoTextFace{
		Source: s.game.Assets.Fonts["default"],
		Size: 24,
	}

	copyrightWidth, _ := text.Measure(GameCopyright, titleFont, titleFont.Size)

	titleDrawOptions := &text.DrawOptions{}
	titleDrawOptions.GeoM.Translate(10, 5)
	titleDrawOptions.ColorScale.ScaleWithColor(color.White)

	copyrightDrawOptions := &text.DrawOptions{}
	copyrightDrawOptions.GeoM.Translate((ScreenWidth - copyrightWidth - 10), 5)
	copyrightDrawOptions.ColorScale.ScaleWithColor(color.White)

	// Draw the title
	text.Draw(screen, GameTitle, titleFont, titleDrawOptions)

	// Draw the copyright text
	text.Draw(screen, GameCopyright, titleFont, copyrightDrawOptions)

	// Draw the blinking prompt
	if s.promptVisible {
		promptFont := &text.GoTextFace{
			Source: s.game.Assets.Fonts["menu"],
			Size: 24,
		}

		promptWidth, promptHeight := text.Measure(MenuPrompt, promptFont, promptFont.Size)

		promptDrawOptions := &text.DrawOptions{}
		promptDrawOptions.GeoM.Translate((ScreenWidth / 2) - (promptWidth / 2), ScreenHeight - promptHeight - 10)
		promptDrawOptions.ColorScale.ScaleWithColor(color.White)

		// TODO: Use an outline instead of a shadow like the old version
		DrawTextWithShadow(screen, MenuPrompt, s.game.Assets.Fonts["menu"], 24, int(ScreenWidth / 2) - int(promptWidth / 2), int(ScreenHeight - promptHeight - 10), 1, color.White, color.Black)
	}
}

func (s *MenuState) Name() string {
	return "Menu"
}

func NewMenuState(g *Game) *MenuState {
	return &MenuState{
		game: g,
		promptBlinkTimer: 0.0,
		promptVisible: true,
	}
}
