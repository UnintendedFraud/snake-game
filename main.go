package main

import (
	"github.com/UnintendedFraud/snake-game/components"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := components.InitGame()

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
