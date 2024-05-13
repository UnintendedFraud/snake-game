package components

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Snake struct {
	speed     int
	direction Direction
	positions []image.Point
}

const (
	START_LENGTH  = 100
	SNAKE_HEIGHT  = 1
	DEFAULT_SPEED = 1
)

func InitSnake(center image.Point) *Snake {
	return &Snake{
		speed:     DEFAULT_SPEED,
		direction: Right,
		positions: getInitPositions(center),
	}
}

func (snake *Snake) Draw(screen *ebiten.Image) {
	if snake == nil {
		return
	}

	for _, point := range snake.positions {
		vector.DrawFilledRect(
			screen,
			float32(point.X),
			float32(point.Y),
			1,
			SNAKE_HEIGHT,
			color.White,
			true,
		)
	}
}

func (snake *Snake) ManageDirection() {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) && snake.direction != Down {
		snake.direction = Up
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) && snake.direction != Left {
		snake.direction = Right
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) && snake.direction != Up {
		snake.direction = Down
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && snake.direction != Right {
		snake.direction = Left
	}
}

func (snake *Snake) HasCollided(maxX, maxY int) bool {
	head := snake.positions[0]

	if head.X <= 0 || head.X >= maxX || head.Y <= 0 || head.Y >= maxY {
		return true
	}

	return snake.hitItself()
}

func (snake *Snake) Move() {
	snake.positions = append(
		[]image.Point{snake.getNextPosition()},
		snake.positions[0:len(snake.positions)-1]...,
	)
}

func (snake *Snake) hitItself() bool {
	head := snake.positions[0]

	for i := 1; i < len(snake.positions); i++ {
		p := snake.positions[i]
		if head.X == p.X && head.Y == p.Y {
			return true
		}
	}

	return false
}

func (snake *Snake) getNextPosition() image.Point {
	nextHead := snake.positions[0]

	switch snake.direction {
	case Up:
		nextHead.Y--
	case Right:
		nextHead.X++
	case Down:
		nextHead.Y++
	case Left:
		nextHead.X--
	}

	return nextHead
}

func getInitPositions(center image.Point) []image.Point {
	points := []image.Point{}

	for i := range START_LENGTH {
		points = append(points, image.Point{X: center.X - i, Y: center.Y})
	}

	return points
}
