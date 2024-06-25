package components

import (
	"image/color"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	// GAME_WIDTH  int = 480
	// GAME_HEIGHT int = 320

	WINDOW_WIDTH  int = 1200
	WINDOW_HEIGHT int = 800

	TICK_RATE = 1
)

type GameState int

const (
	MainMenu GameState = iota
	Playing
	Dead
)

type TickRate struct {
	prevTime time.Time
	currTime time.Time
	rate     int
	diff     time.Duration
}

type Game struct {
	titleImg *ebiten.Image

	tick TickRate

	state GameState
	snake *Snake
	menu  *Menu
}

func InitGame() *Game {
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Snake")

	return &Game{
		titleImg: initTitle(),
		state:    MainMenu,
		snake:    InitSnake(WINDOW_WIDTH, WINDOW_HEIGHT),
		menu:     InitMenu(),
	}
}

func (g *Game) Update() error {
	switch g.state {
	case MainMenu:
		g.menu.UpdateFocus()
		g.menu.Click(g)

	case Playing:
		g.snake.ManageDirection()
		g.snake.Move()

		if g.snake.HasCollided() {
			g.state = Dead
		}

	case Dead:
		g.state = MainMenu
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.titleImg, &ebiten.DrawImageOptions{})

	switch g.state {
	case MainMenu:
		g.menu.Render(screen)

	case Playing:
		g.snake.Render(screen)
	}
}

func (g *Game) Layout(outsideWith, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOW_WIDTH, WINDOW_HEIGHT
}

func initTitle() *ebiten.Image {
	fontFile, err := os.Open("fonts/rockers_garage.ttf")
	if err != nil {
		panic(err)
	}

	font, err := text.NewGoTextFaceSource(fontFile)
	if err != nil {
		panic(err)
	}

	face := &text.GoTextFace{
		Source: font,
		Size:   60,
	}

	titleW, _ := text.Measure("snake", face, 0)

	fontOptions := &text.DrawOptions{}
	fontOptions.GeoM.Translate(float64(WINDOW_WIDTH)/2-titleW/2, 10)
	fontOptions.ColorScale.ScaleWithColor(color.RGBA{194, 169, 60, 1})

	h := float64(WINDOW_HEIGHT) * 0.1

	titleImg := ebiten.NewImage(WINDOW_WIDTH, int(h))

	text.Draw(titleImg, "snake", face, fontOptions)

	return titleImg
}
