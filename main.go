package main

import (
	"log"

	"codingjosh.com/kablam/core"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	println("Kablam!\nCopyright (C) 2024 Josh Kennedy.")

	game := core.NewGame()

	ebiten.SetWindowSize(1280, 960) // 2x of 640x480

	ebiten.SetWindowTitle("Kablam!")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
