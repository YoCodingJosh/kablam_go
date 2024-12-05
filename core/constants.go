package core

import (
	"fmt"
	"time"
)

// just a bunch of constants to make my life easier

const (
	// ScreenWidth is the width of the screen
	ScreenWidth = 1280

	// ScreenHeight is the height of the screen
	ScreenHeight = 960

	// WallTileSize is the size of the wall tile
	WallTileSize = 128

	// The gap between the wall and the screen
	SkyboxHeight = 160

	// The prompt to play the game on the menu
	MenuPrompt = "Press Enter or Touch/Click to Play"

	// The duration of the splash screen
	SplashScreenDuration = 1.500

	// The duration of the menu prompt blink
	MenuPromptBlinkDuration = 0.66667

	// Kablam!
	GameTitle = "Kablam!"
)

// too bad Go doesn't have constexpr like C++
var (
	// Copyright text
	GameCopyright = fmt.Sprintf("Copyright %c %d Josh Kennedy", 169, time.Now().Year())
)
