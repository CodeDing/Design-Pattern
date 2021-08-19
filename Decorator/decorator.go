package Decorator

import `fmt`

type Shape interface {
    Draw() string
}

type Circle struct {
    name   string
    radius int
}

func newCircle(name string, r int) *Circle {
    return &Circle{name, r}
}

func (c *Circle) Draw() string {
    return fmt.Sprintf("name:%s, radius=%d", c.name, c.radius)
}

type Rectangle struct {
    name string
    w, h int
}

func newRectangle(name string, w, h int) *Rectangle {
    return &Rectangle{name, w, h}
}

func (r *Rectangle) Draw() string {
    return fmt.Sprintf("name:%s, width:%d, height:%d", r.name, r.w, r.h)
}

func ShapeDecorator2(fn func(Shape) string) func(Shape) string {
    return func(s Shape) string {
        return "Decorator=>" + fn(s)
    }
}

var (
    RedShapeDecorator2 = ShapeDecorator2(func(s Shape) string {
        return "Board color: red|" + s.Draw()
    })
)
