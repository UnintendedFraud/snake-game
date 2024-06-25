package components

import (
	"image"
	"image/color"
	"time"

	"github.com/UnintendedFraud/snake-game/utils"
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
	img       *ebiten.Image
	speed     int
	direction Direction
	positions []image.Point

	prevTime time.Time
	currTime time.Time
	diff     time.Duration
}

const (
	START_LENGTH  = 50
	SNAKE_HEIGHT  = 10
	DEFAULT_SPEED = 10
)

func InitSnake(width, height int) *Snake {
	h := float64(height) * 0.9

	img := ebiten.NewImage(width, int(h))

	return &Snake{
		img:       img,
		speed:     DEFAULT_SPEED,
		direction: Right,
		positions: getInitPositions(utils.GetCenter(width, height)),
	}
}

func (snake *Snake) Render(screen *ebiten.Image) {
	if snake == nil || snake.img == nil {
		return
	}

	snake.img.Clear()
	snake.img.Fill(color.Gray16{0x1111})

	for _, point := range snake.positions {
		vector.DrawFilledRect(
			snake.img,
			float32(point.X),
			float32(point.Y),
			SNAKE_HEIGHT,
			SNAKE_HEIGHT,
			color.White,
			true,
		)
	}

	options := &ebiten.DrawImageOptions{}

	ty := float64(screen.Bounds().Dy()) * 0.1
	options.GeoM.Translate(0, ty)

	screen.DrawImage(snake.img, options)
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

func (snake *Snake) HasCollided() bool {
	head := snake.positions[0]
	maxX := snake.img.Bounds().Dx()
	maxY := snake.img.Bounds().Dy()

	if head.X <= 0 || head.X >= maxX || head.Y <= 0 || head.Y >= maxY {
		return true
	}

	return snake.hitItself()
}

func (snake *Snake) Move() {
	snake.prevTime = snake.currTime
	snake.currTime = time.Now()
	snake.diff += snake.currTime.Sub(snake.prevTime)

	if snake.diff < 200*time.Millisecond {
		return
	}

	snake.diff = 0

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
		nextHead.Y = nextHead.Y - SNAKE_HEIGHT
	case Right:
		nextHead.X = nextHead.X + SNAKE_HEIGHT
	case Down:
		nextHead.Y = nextHead.Y + SNAKE_HEIGHT
	case Left:
		nextHead.X = nextHead.X - SNAKE_HEIGHT
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
