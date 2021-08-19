package Flyweight

import (
    `fmt`
    `sync`
)

type Shape interface {
    Draw()
}

type Circle struct {
    color        string
    x, y, radius int
}

func newCircle(color string) *Circle {
    return &Circle{color: color}
}

func (c *Circle) SetX(x int) {
    c.x = x
}

func (c *Circle) SetY(y int) {
    c.y = y
}

func (c *Circle) SetRadius(r int) {
    c.radius = r
}

func (c *Circle) Draw() {
    fmt.Printf(`
Circle:
x:%d, y:%d, radius:%d, color: %s
`, c.x, c.y, c.radius, c.color)
}

type ShapeFactory struct {
    sync.RWMutex
    circleMap map[string]*Circle
}

func newShapeFactory() *ShapeFactory {
    return &ShapeFactory{
        circleMap: make(map[string]*Circle),
    }
}

func (s *ShapeFactory) GetCircle(color string) *Circle {
    if s.circleMap == nil {
        s.circleMap = make(map[string]*Circle)
    }
    s.Lock()
    defer s.Unlock()
    if _, ok := s.circleMap[color]; !ok {
        circle := newCircle(color)
        s.circleMap[color] = circle
    }
    return s.circleMap[color]
}
