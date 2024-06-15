package components

import (
	"image/color"
	"os"

	"github.com/UnintendedFraud/snake-game/colors"
	"github.com/UnintendedFraud/snake-game/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	img       ebiten.Image
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

func (m *Menu) UpdateFocus() {
	focusedIdx, err := utils.IndexOf(m.buttons, isButtonFocused)
	if err != nil {
		panic(err)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		if focusedIdx != 0 {
			m.buttons[focusedIdx].isFocused = false
			m.buttons[focusedIdx-1].isFocused = true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		if focusedIdx != len(m.buttons)-1 {
			m.buttons[focusedIdx].isFocused = false
			m.buttons[focusedIdx+1].isFocused = true
		}
	}
}

func (m *Menu) Render(screen *ebiten.Image) {
	m.Clear()

	for idx, b := range m.buttons {
		fontOptions := &text.DrawOptions{}

		menuImg := generateMenuImg(b, m.font, fontOptions)

		x := float64(0)
		y := float64(idx*50 + 50)
		menuImgOp := &ebiten.DrawImageOptions{}
		menuImgOp.GeoM.Translate(x, y)
		m.img.DrawImage(menuImg, menuImgOp)
	}

	imgOp := &ebiten.DrawImageOptions{}
	imgOp.GeoM.Translate(150, 200)
	screen.DrawImage(m.img, imgOp)
}

func (m *Menu) Clear() {
	m.img.Clear()
}

func generateMenuImg(
	button Button,
	font *text.GoTextFaceSource,
	fontOptions *text.DrawOptions,
) *ebiten.Image {
	if button.isFocused {
		fontOptions.ColorScale.ScaleWithColor(colors.Red)
	} else {
		fontOptions.ColorScale.ScaleWithColor(color.White)
	}

	img := ebiten.NewImage(200, 50)

	text.Draw(
		img,
		button.text,
		&text.GoTextFace{
			Source: font,
			Size:   24,
		},
		fontOptions,
	)

	return img
}

func isButtonFocused(b Button) bool {
	return b.isFocused
}
