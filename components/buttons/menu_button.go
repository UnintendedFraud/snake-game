package buttons

import (
	"image/color"

	"github.com/UnintendedFraud/snake-game/colors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type MenuAction int

const (
	StartGame MenuAction = iota
	ExitGame
	ReturnToMenu
)

type MenuButton struct {
	Value  string
	Action MenuAction
	Img    *ebiten.Image
}

func (b *MenuButton) Execute() {
}

func (b *MenuButton) GetMenuButtonImg(
	font *text.GoTextFaceSource,
	fontSize float64,
	isFocused bool,
) *ebiten.Image {
	img := ebiten.NewImage(200, 50)

	fontOptions := &text.DrawOptions{}
	if isFocused {
		fontOptions.ColorScale.ScaleWithColor(colors.Red)
	} else {
		fontOptions.ColorScale.ScaleWithColor(color.White)
	}

	text.Draw(
		img,
		b.Value,
		&text.GoTextFace{
			Source: font,
			Size:   fontSize,
		},
		fontOptions,
	)

	return img
}

func (b *MenuButton) Clear() {
	b.Img.Clear()
}
