package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{
	currentState State
	Assets *AssetManager
}

func NewGame() *Game {
	am := NewAssetManager()

	am.LoadFromJSON("resources/assets.json")

	inst := &Game{
		currentState: nil,
		Assets: am,
	}

	inst.SetState(NewSplashState(inst))

	return inst
}

func (g *Game) SetState(s State) {
	g.currentState = s
}

func (g *Game) Update() error {
	return g.currentState.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Loading...")
	screen.Clear()
	screen.Fill(color.White)

	g.currentState.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 960
}
