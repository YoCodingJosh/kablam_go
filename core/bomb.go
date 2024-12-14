package core

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bomb struct {
	x            float64
	y            float64
	velocity     float64
	active       bool
	currentFrame int
	currentFrameTime float64
	spriteSheet  *ebiten.Image
}

func (b *Bomb) Update(deltaTime float64) {
	b.currentFrameTime += deltaTime

	if b.currentFrameTime >= BombAnimationFrameDuration {
		b.currentFrame++
		b.currentFrameTime = 0
		if b.currentFrame >= BombAnimationFrameCount {
			b.currentFrame = 0
		}
	}

	b.y += b.velocity * deltaTime

	// TODO: this is just temporary: this should be handled by the game state and also it should check if it's past the player
	if b.y > ScreenHeight {
		b.active = false
	}
}

func (b *Bomb) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.x, b.y)

	sx := b.currentFrame * BombSpriteWidth

	screen.DrawImage(b.spriteSheet.SubImage(image.Rect(sx, 0, sx+BombSpriteWidth, BombSpriteHeight)).(*ebiten.Image), op)
}

func NewBomb(x, y, velocity float64, spriteSheet *ebiten.Image) *Bomb {
	return &Bomb{
		x:           x,
		y:           y,
		velocity:    velocity,
		active:      true,
		spriteSheet: spriteSheet,
	}
}
