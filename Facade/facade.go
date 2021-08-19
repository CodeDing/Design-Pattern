package Facade

type Shape interface {
	Draw() string
}

type Rectangle struct{}

func (Rectangle) Draw() string {
	return "<rectangle>"
}

type Square struct{}

func (Square) Draw() string {
	return "<square>"
}

type Circle struct{}

func (Circle) Draw() string {
	return "<circle>"
}

type ShapeMaker struct {
	rectangle Rectangle
	square    Square
	circle    Circle
}

func NewShapeMaker() *ShapeMaker {
	return &ShapeMaker{Rectangle{}, Square{}, Circle{}}
}

func (s ShapeMaker) DrawRectangle() string {
	return s.rectangle.Draw()
}

func (s ShapeMaker) DrawSquare() string {
	return s.square.Draw()
}

func (s ShapeMaker) DrawCircle() string {
	return s.circle.Draw()
}
