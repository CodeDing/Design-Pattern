package Bridge

import "fmt"

type DrawAPI interface {
	DrawCircle(int, int, int) string
}

type RedCircle struct{}

func newRedCircle() *RedCircle {
	return &RedCircle{}
}

func (RedCircle) DrawCircle(x, y, radius int) string {
	return fmt.Sprintf("redcircle=x:%d,y:%d,radius:%d", x, y, radius)
}

type GreenCircle struct{}

func newGreenCircle() *GreenCircle {
	return &GreenCircle{}
}

func (GreenCircle) DrawCircle(x, y, radius int) string {
	return fmt.Sprintf("greencircle=x:%d,y:%d,radius:%d", x, y, radius)
}

type Shape interface {
	Draw() string
}

type Circle struct {
	DrawAPI
	x, y   int
	radius int
}

func NewCircle(x, y, radius int, d DrawAPI) *Circle {
	return &Circle{d, x, y, radius}
}

func (c Circle) Draw() string {
	return c.DrawAPI.DrawCircle(c.x, c.y, c.radius)
}
