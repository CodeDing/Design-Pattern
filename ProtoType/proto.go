package ProtoType

import "fmt"

type Shape interface {
	GetType() string
	GetId() string
	SetId(string)
	String() string
	Draw()
}

type base struct {
	id       string
	category string
}

func (b base) GetType() string {
	return b.category
}

func (b base) GetId() string {
	return b.id
}

func (b *base) SetId(id string) {
	if b == nil {
		panic("base cannot be nil")
	}
	b.id = id
}

func (b base) String() string {
	return b.id + ":" + b.category
}

func (b base) Draw() {
	fmt.Println("Base draw")
}

type Rectangle struct {
	base
	w, h int
}

func newRectangle(c string, w, h int) *Rectangle {
	return &Rectangle{
		base{category: c},
		w, h,
	}
}

func (r Rectangle) Draw() {
	fmt.Printf(`
..........................
.                        .
..........................
`)
}

type Square struct {
	base
	edge int
}

func newSquare(c string, e int) *Square {
	return &Square{
		base{category: c},
		e,
	}
}

func (s Square) Draw() {
	fmt.Printf(`
...........
.         .
.         .
.         .
...........
`)
}

type Circle struct {
	base
	radius int
}

func newCircle(c string, radius int) *Circle {
	return &Circle{
		base{category: c},
		radius,
	}
}

func (c Circle) Draw() {
	fmt.Printf(`
      ....
     .    .
    .      .
    .      .
     .    .
      ....
`)
}

type ShapeCache struct {
	shapeMap map[string]Shape
}

func (c *ShapeCache) Load() {
	if c.shapeMap == nil {
		c.shapeMap = make(map[string]Shape)
	}
	rectangle := newRectangle("rectangle", 3, 2)
	rectangle.SetId("1")
	square := newSquare("square", 3)
	square.SetId("2")
	circle := newCircle("circle", 3)
	circle.SetId("3")
	c.shapeMap["1"] = rectangle
	c.shapeMap["2"] = square
	c.shapeMap["3"] = circle
}

func (c ShapeCache) GetShape(n string) Shape {
	if s, ok := c.shapeMap[n]; ok {
		return s
	}
	return nil
}
