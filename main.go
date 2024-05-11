package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Name window")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello world")
}

func (g *Game) Layout(outsideWith, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
