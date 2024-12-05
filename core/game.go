package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	currentState State
	timeAccumulator float64
	Assets *AssetManager
}

func NewGame() *Game {
	am := NewAssetManager()

	am.LoadFromJSON("resources/assets.json")

	inst := &Game{
		currentState: nil,
		timeAccumulator: 0.0,
		Assets: am,
	}

	inst.SetState(NewSplashState(inst))

	return inst
}

func (g *Game) SetState(s State) {
	g.currentState = s
}

func (g *Game) CurrentState() State {
	return g.currentState
}

func (g *Game) Update() error {
	// tps := ebiten.ActualTPS()
	// if tps == 0 {
		// tps = ebiten.DefaultTPS // we don't want to divide by zero
	// }
	// deltaTime := 1.0 / tps

	// TODO: temp fix lmao
	deltaTime := 0.0166666667

	// log.Printf("TPS: %f", tps)
	// log.Printf("Delta Time: %f", deltaTime)

	return g.currentState.Update(deltaTime)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(color.White)

	g.currentState.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
