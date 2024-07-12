package components

import (
	"fmt"
	"image/color"
	"os"

	"github.com/UnintendedFraud/snake-game/colors"
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
	focusedIdx int
	buttons    []Button
	img        *ebiten.Image
}

type Button struct {
	text      string
	action    MenuAction
	isFocused bool
	img       ebiten.Image
}

func InitMenu() *Menu {
	return &Menu{
		img:        ebiten.NewImage(200, 200),
		focusedIdx: 0,
		buttons: []Button{
			{
				text:   "start game",
				action: StartGame,
			},
			{
				text:   "exit game",
				action: ExitGame,
			},
		},
	}
}

func (m *Menu) Click(game *Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) ||
		inpututil.IsKeyJustPressed(ebiten.KeyKPEnter) {
		m.executeAction(game)
	}
}

func (m *Menu) executeAction(game *Game) {
	action := m.buttons[m.focusedIdx].action

	switch action {
	case StartGame:
		game.snake = InitSnake(WINDOW_WIDTH, WINDOW_HEIGHT)
		game.state = Playing

	case ExitGame:
		fmt.Println("Closing the game")
		os.Exit(0)
	}
}

func (m *Menu) UpdateFocus() {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		if m.focusedIdx != 0 {
			m.focusedIdx--
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		if m.focusedIdx != len(m.buttons)-1 {
			m.focusedIdx++
		}
	}
}

func (m *Menu) Render(screen *ebiten.Image, font *text.GoTextFaceSource) {
	m.Clear()

	for idx, b := range m.buttons {
		fontOptions := &text.DrawOptions{}
		isFocused := idx == m.focusedIdx

		menuImg := generateMenuImg(b, font, fontOptions, isFocused)

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
	isFocused bool,
) *ebiten.Image {
	if isFocused {
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
