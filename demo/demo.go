package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	Width  = 20 // 游戏区域宽度
	Height = 20 // 游戏区域高度
)

type Point struct {
	X int // 横坐标
	Y int // 纵坐标
}

type Snake struct {
	Body      []Point // 蛇的身体
	Direction Point   // 移动方向
}

type Food struct {
	Position Point // 食物位置
	Eaten    bool  // 是否已经被吃掉
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	rand.Seed(time.Now().UnixNano())
	snake := NewSnake()
	food := NewFood(snake.Body)

	for {
		draw(snake, food)

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				snake.Direction = Point{0, -1}
			case termbox.KeyArrowDown:
				snake.Direction = Point{0, 1}
			case termbox.KeyArrowLeft:
				snake.Direction = Point{-1, 0}
			case termbox.KeyArrowRight:
				snake.Direction = Point{1, 0}
			case termbox.KeyEsc:
				return
			}
		case termbox.EventError:
			panic(ev.Err)
		}

		if !snake.Move(food.Position) {
			break
		}

		if food.Eaten {
			food.RandomPosition(snake.Body)
			food.Eaten = false
		}

		time.Sleep(100 * time.Millisecond)
	}

	gameOver()
}

func NewSnake() *Snake {
	snake := &Snake{
		Body: []Point{
			{Width / 2, Height / 2},
		},
		Direction: Point{0, -1},
	}

	for i := 1; i < 3; i++ {
		snake.Body = append(snake.Body, Point{snake.Body[0].X, snake.Body[0].Y + i})
	}

	return snake
}

func (snake *Snake) Move(food Point) bool {
	head := snake.Body[0]
	newHead := Point{head.X + snake.Direction.X, head.Y + snake.Direction.Y}

	if newHead.X < 0 || newHead.X >= Width || newHead.Y < 0 || newHead.Y >= Height {
		return false
	}

	if newHead == food {
		snake.Body = append([]Point{newHead}, snake.Body...)
		snake.Grow()
		return true
	}

	for _, body := range snake.Body {
		if newHead == body {
			return false
		}
	}

	for i := len(snake.Body) - 1; i > 0; i-- {
		snake.Body[i] = snake.Body[i-1]
	}

	snake.Body[0] = newHead

	return true
}

func (snake *Snake) Grow() {
	tail := snake.Body[len(snake.Body)-1]
	snake.Body = append(snake.Body, tail)
}

func NewFood(snake []Point) *Food {
	food := &Food{
		Position: Point{rand.Intn(Width), rand.Intn(Height)},
		Eaten:    false,
	}

	for _, body := range snake {
		if food.Position == body {
			food.RandomPosition(snake)
			return food
		}
	}

	return food
}

func (food *Food) RandomPosition(snake []Point) {
	food.Position = Point{rand.Intn(Width), rand.Intn(Height)}

	for _, body := range snake {
		if food.Position == body {
			food.RandomPosition(snake)
			return
		}
	}
}

func draw(snake *Snake, food *Food) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for _, body := range snake.Body {
		termbox.SetCell(body.X, body.Y, ' ', termbox.ColorGreen, termbox.ColorDefault)
	}

	termbox.SetCell(food.Position.X, food.Position.Y, ' ', termbox.ColorYellow, termbox.ColorDefault)

	termbox.Flush()
}

func gameOver() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	w, h := termbox.Size()

	text := "Game Over"
	x := (w - len(text)) / 2
	y := h / 2

	for _, c := range text {
		termbox.SetCell(x, y, c, termbox.ColorRed, termbox.ColorDefault)
		x++
	}

	termbox.Flush()

	time.Sleep(time.Second)

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}
