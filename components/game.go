package components

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	WINDOW_WIDTH  int = 1200
	WINDOW_HEIGHT int = 800
)

type GameState int

const (
	MainMenu GameState = iota
	Playing
	Dead
)

type Game struct {
	font *text.GoTextFaceSource

	titleImg *ebiten.Image

	state GameState
	snake *Snake
	menu  *Menu
}

func InitGame() *Game {
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Snake")

	fontFile, err := os.Open("fonts/rockers_garage.ttf")
	if err != nil {
		panic(err)
	}

	font, err := text.NewGoTextFaceSource(fontFile)
	if err != nil {
		panic(err)
	}

	return &Game{
		font:     font,
		titleImg: initTitle(font),
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
		g.snake.Eat()
		g.snake.Sprint()
		g.snake.ManageDirection()
		g.snake.Move()

		if g.snake.HasCollided() {
			g.snake.isDead = true
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.titleImg, &ebiten.DrawImageOptions{})

	switch g.state {
	case MainMenu:
		g.menu.Render(screen, g.font)

	case Playing:
		g.snake.Render(screen, g.font)
	}
}

func (g *Game) Layout(outsideWith, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOW_WIDTH, WINDOW_HEIGHT
}

func initTitle(font *text.GoTextFaceSource) *ebiten.Image {
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
