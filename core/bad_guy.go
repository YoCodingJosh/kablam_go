package core

import (
	"image/color"
	"math"
	"math/rand/v2"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type BadGuy struct {
	currentLevel       int     // Current level of the game -- determines difficulty
	currentPosition    float64 // Current position of the bad guy
	nextPosition       float64 // Next position of the bad guy, will be lerped to
	shouldMoveTicker 	 time.Ticker
	velocity           float64 // Speed of the bad guy
	velocityMultiplier float64 // Multiplier for the velocity -- will be calculated by the game somehow TBD lol
	shouldDropBomb     bool    // Should the bad guy drop a bomb? Will be determined by a timer
	bombDropTicker     time.Ticker
}

func (b *BadGuy) shouldMoveCallback() {
	b.nextPosition = rand.Float64() * (MaxBadGuyXPos - MinBadGuyXPos) + MinBadGuyXPos
}

func (b *BadGuy) Update(deltaTime float64) {
	// Update the position of the bad guy
	b.currentPosition = Lerp(b.currentPosition, b.nextPosition, b.velocity * deltaTime)

	// If the bad guy is close to the next position, set the next position to a new random position
	if math.Abs(b.currentPosition - b.nextPosition) < 1 {
		b.currentPosition = b.nextPosition
	}
}

func (b *BadGuy) Draw(screen *ebiten.Image) {
	// Draw the bad guy (a black square for now) at the current position
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.currentPosition, BadGuyYPos)

	// TODO: draw the bad guy sprite from assetmanager
	square := ebiten.NewImage(BadGuySpriteWidth, BadGuySpriteHeight)
	square.Fill(color.Black)

	screen.DrawImage(square, op)
}

func NewBadGuy(level int, bombDropCallback func(x, y float64)) *BadGuy {
	moveTicker := time.NewTicker(time.Millisecond * BadGuyMoveInterval)
	bombTicker := time.NewTicker(time.Millisecond * BadGuyBombInterval)

	badGuy := &BadGuy{
		currentLevel:       level,
		currentPosition:    (ScreenWidth / 2) - (BadGuySpriteWidth / 2),
		nextPosition:       (ScreenWidth / 2) - (BadGuySpriteWidth / 2),
		velocity:           BadGuySpeed,
		velocityMultiplier: 1.0,
		shouldDropBomb:     false,
		shouldMoveTicker:   *moveTicker,
		bombDropTicker:     *bombTicker,
	}

	// TODO: How to handle game pausing?

	go func() {
		for range moveTicker.C {
			badGuy.shouldMoveCallback()
		}
	}()

	go func() {
		for range bombTicker.C {
			badGuy.shouldDropBomb = true
			bombDropCallback(badGuy.currentPosition, BadGuyYPos)
		}
	}()

	return badGuy
}

// Call this to clean up when the bad guy is removed
func (b *BadGuy) Stop() {
	b.shouldMoveTicker.Stop()
	b.bombDropTicker.Stop()
}
