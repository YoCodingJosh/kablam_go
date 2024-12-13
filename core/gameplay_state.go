package core

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type GameplayState struct {
	game *Game
	badGuy *BadGuy
	bombs []*Bomb
	score uint64
}

func (s *GameplayState) Update(deltaTime float64) error {
	s.badGuy.Update(deltaTime)

	for _, bomb := range s.bombs {
		bomb.Update(deltaTime)
	}

	return nil
}

func (s *GameplayState) Draw(screen *ebiten.Image) {
	scoreFont := &text.GoTextFace{
		Source: s.game.Assets.Fonts["default"],
		Size: 24,
	}

	scoreDrawOptions := &text.DrawOptions{}
	scoreDrawOptions.GeoM.Translate(10, 5)
	scoreDrawOptions.ColorScale.ScaleWithColor(color.White)

	screen.Fill(color.RGBA{100, 149, 237, 255})

	text.Draw(screen, fmt.Sprintf(ScoreText, s.score), scoreFont, scoreDrawOptions)

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

	for _, bomb := range s.bombs {
		bomb.Draw(screen)
	}
}

func (s* GameplayState) handleBombDrop(x, y float64) {
	bomb := NewBomb(x, y, BombVelocity, s.game.Assets.Images["bomb"])
	s.bombs = append(s.bombs, bomb)
}

func (s *GameplayState) Name() string {
	return "Gameplay"
}


func NewGameplayState(g *Game) *GameplayState {
	state := &GameplayState{
		game: g,
		bombs: make([]*Bomb, 0),
		score: 0,
	}

	// TODO: Call badGuy.Stop() to stop the bad guy's tickers
	badGuy := NewBadGuy(1, state.handleBombDrop)

	state.badGuy = badGuy

	return state
}
