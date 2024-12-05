package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	DefaultShadowColor = color.RGBA{0, 0, 0, 0x80}
)

func DrawTextWithShadow(rt *ebiten.Image, str string, font *text.GoTextFaceSource, fontBaseSize float64, x, y, scale int, clr color.Color, shadowColor color.Color) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x)+1, float64(y)+1)
	op.ColorScale.ScaleWithColor(shadowColor)
	op.LineSpacing = fontBaseSize * float64(scale)
	text.Draw(rt, str, &text.GoTextFace{
		Source: font,
		Size:   fontBaseSize * float64(scale),
	}, op)

	op.GeoM.Reset()
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.Reset()
	op.ColorScale.ScaleWithColor(clr)
	text.Draw(rt, str, &text.GoTextFace{
		Source: font,
		Size:   fontBaseSize * float64(scale),
	}, op)
}
