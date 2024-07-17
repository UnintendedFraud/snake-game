package components

import (
	"fmt"
	"os"

	"github.com/UnintendedFraud/snake-game/components/buttons"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Menu struct {
	focusedIdx int
	buttons    []buttons.MenuButton
	img        *ebiten.Image
}

func InitMenu() *Menu {
	return &Menu{
		img:        ebiten.NewImage(200, 200),
		focusedIdx: 0,
		buttons: []buttons.MenuButton{
			{
				Value:  "start game",
				Action: buttons.StartGame,
			},
			{
				Value:  "exit game",
				Action: buttons.ExitGame,
			},
		},
	}
}

func (m *Menu) Click(game *Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) ||
		inpututil.IsKeyJustPressed(ebiten.KeyKPEnter) {
		m.executeAction(game, m.buttons[m.focusedIdx].Action)
	}
}

func (m *Menu) executeAction(game *Game, action buttons.MenuAction) {
	switch action {
	case buttons.StartGame:
		game.snake = InitSnake(WINDOW_WIDTH, WINDOW_HEIGHT)
		game.state = Playing

	case buttons.ExitGame:
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
		menuImg := b.GetMenuButtonImg(font, 24, m.focusedIdx == idx)

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
