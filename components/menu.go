package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MenuAction int

const (
	StartGame MenuAction = iota
	ExitGame
)

type Menu struct {
	buttons []Button
}

type Button struct {
	text   string
	action MenuAction
}

func InitMenu() *Menu {
	return &Menu{
		buttons: []Button{
			{
				text:   "Start",
				action: StartGame,
			},
			{
				text:   "Exit",
				action: ExitGame,
			},
		},
	}
}

func (m *Menu) ManageSelection() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) ||
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		// select option
	}
}
