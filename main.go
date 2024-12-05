package main

import (
	"fmt"
	"log"

	"codingjosh.com/kablam/core"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	println(fmt.Sprintf("%s\n%s", core.GameTitle, core.GameCopyright))

	game := core.NewGame()

	ebiten.SetWindowSize(core.ScreenWidth, core.ScreenHeight)

	ebiten.SetWindowTitle(core.GameTitle)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
