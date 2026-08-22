package texts

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Text struct {
	Value string
	Size  float64
	Font  *text.GoTextFaceSource
	Color color.Color
}

func (t *Text) Render(containerImg *ebiten.Image, pos image.Point) {
	textImg := t.GetTextImg()

	menuImgOp := &ebiten.DrawImageOptions{}
	menuImgOp.GeoM.Translate(float64(pos.X), float64(pos.Y))
	containerImg.DrawImage(textImg, menuImgOp)
}

func (t *Text) GetTextImg() *ebiten.Image {
	img := ebiten.NewImage(200, 50)

	fontOptions := &text.DrawOptions{}
	fontOptions.ColorScale.ScaleWithColor(t.Color)

	text.Draw(
		img,
		t.Value,
		&text.GoTextFace{
			Source: t.Font,
			Size:   t.Size,
		},
		fontOptions,
	)

	return img
}
