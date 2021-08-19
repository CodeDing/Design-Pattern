package SimpleFactory

import (
	"fmt"
	"math"
)

type Shape interface {
	Name() string
	Area() float64
	Draw()
}

type Rectangle struct {
	name string
	h, w int
}

func NewRectangle(w, h int) *Rectangle {
	return &Rectangle{"rectangle", h, w}
}

func (r Rectangle) Name() string {
	return r.name
}

func (r Rectangle) Area() float64 {
	return float64(r.h * r.w)
}

func (r Rectangle) Draw() {
	fmt.Printf(`
************************
*                      *
*                      *
*                      *
************************
`)
}

type Square struct {
	name string
	edge int
}

func NewSquare(e int) *Square {
	return &Square{"square", e}
}

func (s Square) Name() string {
	return s.name
}

func (s Square) Area() float64 {
	return float64(s.edge * s.edge)
}

func (s Square) Draw() {
	fmt.Printf(`
------------
|          |
|          |
|          |
|__________|
`)
}

type Circle struct {
	name   string
	radius int
}

func NewCircle(radius int) *Circle {
	return &Circle{"circle", radius}
}

func (c Circle) Name() string {
	return c.name
}

func (c Circle) Area() float64 {
	return float64(c.radius*c.radius) * math.Pi
}

func (c Circle) Draw() {
	fmt.Printf(`
      @@@@@@@
     @       @
    @         @
   @           @
    @         @
     @       @
      @@@@@@@
`)
}

type NullShape struct{}

func NewNullShape() *NullShape {
	return &NullShape{}
}

func (NullShape) Name() string {
	return "null"
}

func (NullShape) Area() float64 {
	return 0.0
}

func (NullShape) Draw() {
	fmt.Println(`NULL`)
}

const (
	RECTANGLE = "rectangle"
	SQUARE    = "square"
	CIRCLE    = "circle"
)

type ShapeFactory struct{}

func (f ShapeFactory) GetShape(name string) (s Shape) {
	switch name {
	case RECTANGLE:
		s = NewRectangle(3, 2)
	case SQUARE:
		s = NewSquare(3)
	case CIRCLE:
		s = NewCircle(4)
	default:
		s = NewNullShape()
	}
	return
}
