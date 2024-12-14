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

	BombSpriteWidth  = 64
	BombSpriteHeight = 64
	BombAnimationFrameCount = 4
	BombAnimationFPS = 15
	BombAnimationFrameDuration = 1.0 / BombAnimationFPS
	BombVelocity = 64.0

	BadGuySpriteWidth  = 100
	BadGuySpriteHeight = 100

	BadGuySpeed = 3
	BadGuyYPos = 60

	// The rightmost possible position for the bad guy
	MaxBadGuyXPos = ScreenWidth - BadGuySpriteWidth

	// The leftmost possible position for the bad guy
	MinBadGuyXPos = BadGuySpriteWidth

	BadGuyMoveInterval = 1500 // milliseconds, this is arbitrary for now (might be dynamic later)
	BadGuyBombInterval = 1000 // milliseconds, this is arbitrary for now (might be dynamic later)

	ScoreTextSize = 24
	ScoreText = "Score: %d"
)

// too bad Go doesn't have constexpr like C++
var (
	// Copyright text
	GameCopyright = fmt.Sprintf("Copyright %c %d Josh Kennedy", 169, time.Now().Year()) // 169 is the copyright symbol
)
