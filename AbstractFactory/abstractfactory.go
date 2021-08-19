package AbstractFactory

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Draw()
}

type Color interface {
	Fill() string
}

type AbstractFactory interface {
	GetShape(name string) Shape
	GetColor(name string) Color
}

type Rectangle struct {
	name string
	w, h int
}

func newRectangle(w, h int) *Rectangle {
	return &Rectangle{"rectangle", w, h}
}

func (r Rectangle) Area() float64 {
	return float64(r.w * r.h)
}

func (r Rectangle) Draw() {
	fmt.Printf(`
.........................
.                       .
.........................
`)
}

type Square struct {
	name string
	edge int
}

func newSquare(e int) *Square {
	return &Square{"square", e}
}

func (s Square) Area() float64 {
	return float64(s.edge * s.edge)
}

func (s Square) Draw() {
	fmt.Printf(`
............
.          .
.          .
.          .
............
`)
}

type Circle struct {
	name   string
	radius int
}

func newCircle(r int) *Circle {
	return &Circle{"circle", r}
}

func (c Circle) Area() float64 {
	return math.Pi * float64(c.radius*c.radius)
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

type Red struct{}

func (Red) Fill() string {
	return "red"
}

type Green struct{}

func (Green) Fill() string {
	return "green"
}

type Blue struct{}

func (Blue) Fill() string {
	return "blue"
}

type ShapeFactory struct {
}

func (f ShapeFactory) GetShape(name string) Shape {
	switch name {
	case "rectangle":
		return newRectangle(3, 2)
	case "square":
		return newSquare(3)
	case "circle":
		return newCircle(3)
	}
	return nil
}

func (f ShapeFactory) GetColor(name string) Color {
	return nil
}

type ColorFactory struct {
}

func (f ColorFactory) GetShape(name string) Shape {
	return nil
}

func (f ColorFactory) GetColor(name string) Color {
	switch name {
	case "red":
		return Red{}
	case "green":
		return Green{}
	case "blue":
		return Blue{}
	}
	return nil
}

type FactoryProducer interface {
	GetFactory() AbstractFactory
}

type FactoryProducerFunc func() AbstractFactory

func (f FactoryProducerFunc) GetFactory() AbstractFactory {
	return f()
}

var (
	allFactories = map[string]FactoryProducer{
		"shape": FactoryProducerFunc(func() AbstractFactory { return ShapeFactory{} }),
		"color": FactoryProducerFunc(func() AbstractFactory { return ColorFactory{} }),
	}
)
