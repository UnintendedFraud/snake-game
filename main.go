package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

func main() {
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Name window")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 1, 1, 4, 1, color.White, true)
}

func (g *Game) Layout(outsideWith, outsideHeight int) (screenWidth, screenHeight int) {
	return 10, 10
}
