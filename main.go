package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/UnintendedFraud/snake-game/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type GameState int

const (
	Menu GameState = iota
	Playing
	Dead
)

type Game struct {
	titleImg *ebiten.Image

	state GameState
	snake *components.Snake
	menu  *components.Menu
}

const (
	// GAME_WIDTH  int = 480
	// GAME_HEIGHT int = 320

	WINDOW_WIDTH  int = 1024
	WINDOW_HEIGHT int = 768
)

func main() {
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Snake")

	game := &Game{
		titleImg: initTitle(),
		state:    Playing,
		snake:    components.InitSnake(WINDOW_WIDTH, WINDOW_HEIGHT),
		menu:     components.InitMenu(),
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

func (g *Game) Update() error {
	switch g.state {
	case Menu:

	case Playing:
		g.snake.ManageDirection()
		g.snake.Move()

		if g.snake.HasCollided() {
			g.state = Dead
		}

	case Dead:
		fmt.Println("DEAD")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.titleImg, &ebiten.DrawImageOptions{})

	switch g.state {
	case Menu:
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
