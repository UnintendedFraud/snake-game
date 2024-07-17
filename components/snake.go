package components

import (
	"image"
	"image/color"
	"math/rand"
	"slices"
	"time"

	"github.com/UnintendedFraud/snake-game/colors"
	"github.com/UnintendedFraud/snake-game/components/buttons"
	"github.com/UnintendedFraud/snake-game/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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
	speed     int64
	direction Direction
	positions []image.Point
	sprint    bool

	ennemies []image.Point

	prevTime time.Time
	currTime time.Time
	diff     time.Duration

	isDead bool
}

const (
	START_LENGTH       = 20
	THICKNESS          = 8
	MIN_SPEED    int64 = 500
	MAX_SPEED    int64 = 50
	MAX_SPRINT   int64 = 20
)

func InitSnake(width, height int) *Snake {
	h := int(float64(height) * 0.9)

	img := ebiten.NewImage(width, h)

	return &Snake{
		img:       img,
		speed:     MIN_SPEED,
		direction: Right,
		positions: getInitPositions(utils.GetCenter(width, height)),
		ennemies:  randomEnnemies(width, h, 10),
		isDead:    false,
	}
}

func (snake *Snake) Render(screen *ebiten.Image, font *text.GoTextFaceSource) {
	if snake == nil || snake.img == nil {
		return
	}

	snake.img.Clear()

	if snake.isDead {
		renderDeadScreen(screen, snake, font)
		return
	}

	snake.img.Fill(colors.LightGray)

	for _, ennemy := range snake.ennemies {
		vector.DrawFilledRect(
			snake.img,
			float32(ennemy.X),
			float32(ennemy.Y),
			THICKNESS,
			THICKNESS,
			colors.DarkYellow,
			true,
		)
	}

	for _, point := range snake.positions {
		vector.DrawFilledRect(
			snake.img,
			float32(point.X),
			float32(point.Y),
			THICKNESS,
			THICKNESS,
			color.White,
			true,
		)
	}

	options := &ebiten.DrawImageOptions{}

	ty := float64(screen.Bounds().Dy()) * 0.1
	options.GeoM.Translate(0, ty)

	screen.DrawImage(snake.img, options)
}

func (snake *Snake) Sprint() {
	if utils.SliceContains(inpututil.AppendPressedKeys([]ebiten.Key{}), ebiten.KeySpace) && !snake.sprint {
		snake.sprint = true
	} else if snake.sprint {
		snake.sprint = false
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

func (snake *Snake) HasCollided() bool {
	head := snake.positions[0]
	maxX := snake.img.Bounds().Dx()
	maxY := snake.img.Bounds().Dy()

	if head.X <= 0 || head.X >= maxX || head.Y <= 0 || head.Y >= maxY {
		return true
	}

	return snake.hitItself()
}

func (snake *Snake) Eat() {
	head := snake.positions[0]

	idx, err := utils.SliceIndexOf(snake.ennemies, func(e image.Point) bool {
		return e.X == head.X && e.Y == head.Y
	})
	if err != nil || idx < 0 {
		return
	}

	snake.ennemies = slices.Delete(snake.ennemies, idx, idx+1)

	snake.positions = append(
		[]image.Point{snake.getNextPosition()},
		snake.positions...,
	)

	if snake.speed > MAX_SPEED {
		snake.speed -= 50
	}
}

func (snake *Snake) Move() {
	snake.prevTime = snake.currTime
	snake.currTime = time.Now()
	snake.diff += snake.currTime.Sub(snake.prevTime)

	var speed int64
	if snake.sprint {
		speed = MAX_SPRINT
	} else {
		speed = snake.speed
	}

	if snake.diff.Milliseconds() < speed {
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
		nextHead.Y = nextHead.Y - THICKNESS
	case Right:
		nextHead.X = nextHead.X + THICKNESS
	case Down:
		nextHead.Y = nextHead.Y + THICKNESS
	case Left:
		nextHead.X = nextHead.X - THICKNESS
	}

	return nextHead
}

func (snake *Snake) Click(game *Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) ||
		inpututil.IsKeyJustPressed(ebiten.KeyKPEnter) {
		game.state = MainMenu
	}
}

func getInitPositions(center image.Point) []image.Point {
	points := []image.Point{}

	for i := range START_LENGTH {
		points = append(points, image.Point{X: center.X - i*THICKNESS, Y: center.Y})
	}

	return points
}

func randomEnnemies(width int, height int, count int) []image.Point {
	ennemies := []image.Point{}

	w := width / THICKNESS
	h := height / THICKNESS

	for range count {
		x := rand.Intn(w) * THICKNESS
		y := rand.Intn(h) * THICKNESS

		ennemies = append(ennemies, image.Point{X: x, Y: y})
	}

	return ennemies
}

func renderDeadScreen(screen *ebiten.Image, snake *Snake, font *text.GoTextFaceSource) {
	snake.img.Fill(colors.Black)

	fontOptions := &text.DrawOptions{}
	fontOptions.ColorScale.ScaleWithColor(colors.Red)

	deadImg := ebiten.NewImage(450, 90)

	text.Draw(
		deadImg,
		"You are DEAD !",
		&text.GoTextFace{
			Source: font,
			Size:   80,
		},
		fontOptions,
	)

	deadImgW := deadImg.Bounds().Dx()
	deadImgH := deadImg.Bounds().Dy()
	x := float64(screen.Bounds().Dx()/2 - deadImgW/2)
	y := float64(screen.Bounds().Dy()/2 - deadImgH/2)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	snake.img.DrawImage(deadImg, op)

	// ---- return to menu button

	returnToMenuButton := buttons.MenuButton{
		Value:  "return to menu",
		Action: buttons.ReturnToMenu,
	}

	rtmbImg := returnToMenuButton.GetMenuButtonImg(font, 24, true)

	rtmOp := &ebiten.DrawImageOptions{}
	rtmOp.GeoM.Translate(x, y+100)
	snake.img.DrawImage(rtmbImg, rtmOp)

	screen.DrawImage(snake.img, &ebiten.DrawImageOptions{})
}
