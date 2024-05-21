package components

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type MenuAction int

const (
	StartGame MenuAction = iota
	ExitGame
)

type Menu struct {
	font    *text.GoTextFaceSource
	buttons []Button
	img     *ebiten.Image
}

type Button struct {
	text      string
	action    MenuAction
	isFocused bool
}

func InitMenu() *Menu {
	fontFile, err := os.Open("fonts/rockers_garage.ttf")
	if err != nil {
		panic(err)
	}

	font, err := text.NewGoTextFaceSource(fontFile)
	if err != nil {
		panic(err)
	}

	return &Menu{
		font: font,
		img:  ebiten.NewImage(200, 200),
		buttons: []Button{
			{
				text:      "start game",
				action:    StartGame,
				isFocused: true,
			},
			{
				text:      "exit game",
				action:    ExitGame,
				isFocused: false,
			},
		},
	}
}

func (m *Menu) Render(screen *ebiten.Image) {
	fontOptions := &text.DrawOptions{}
	fontOptions.ColorScale.ScaleWithColor(color.White)

	for idx, b := range m.buttons {
		x := float64(0)
		y := float64(idx*50 + 50)

		fontOptions.GeoM.Translate(x, y)

		text.Draw(
			m.img,
			b.text,
			&text.GoTextFace{
				Source: m.font,
				Size:   24,
			},
			fontOptions,
		)
	}

	imgOp := &ebiten.DrawImageOptions{}
	screen.DrawImage(m.img, imgOp)
}

func (m *Menu) Clear() {
	m.img.Clear()
}
