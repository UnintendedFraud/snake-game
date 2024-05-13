package main

import (
	"image"
	"image/color"

	"github.com/UnintendedFraud/snake-game/components"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameState int

const (
	Menu GameState = iota
	Playing
	Dead
)

type Game struct {
	state GameState
	snake *components.Snake
}

const (
	GAME_WIDTH  int = 480
	GAME_HEIGHT int = 320
)

func main() {
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Name window")

	game := &Game{
		state: Playing,
		snake: components.InitSnake(getCenter(GAME_WIDTH, GAME_HEIGHT)),
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

		if g.snake.HasCollided(GAME_WIDTH, GAME_HEIGHT) {
			g.state = Dead
		}

	case Dead:
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray16{0x5050})
	g.snake.Draw(screen)
}

func (g *Game) Layout(outsideWith, outsideHeight int) (screenWidth, screenHeight int) {
	return GAME_WIDTH, GAME_HEIGHT
}

func getCenter(x, y int) image.Point {
	return image.Point{
		X: x / 2,
		Y: y / 2,
	}
}
