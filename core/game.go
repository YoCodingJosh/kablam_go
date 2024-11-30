package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{
	currentState State
	assets *AssetManager
}

func NewGame() *Game {
	am := NewAssetManager()

	am.LoadFromJSON("resources/assets.json")

	return &Game{
		currentState: nil,
		assets: am,
	}
}

func (g *Game) SetState(s State) {
	g.currentState = s
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Loading...")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 960
}
