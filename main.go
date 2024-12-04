package main

import (
	"log"

	"codingjosh.com/kablam/core"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	println("Kablam!\nCopyright (C) 2024 Josh Kennedy.")

	game := core.NewGame()

	ebiten.SetWindowSize(core.ScreenWidth, core.ScreenHeight)

	ebiten.SetWindowTitle("Kablam!")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
