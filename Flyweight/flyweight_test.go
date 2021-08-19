package Flyweight

import (
    `math/rand`
    `testing`
)

var (
    COLORS = []string{"Red", "Green", "Blue", "White", "Black"}
)

func getRandomColor() string {
    return COLORS[rand.Intn(len(COLORS))]
}

func getRandomX() int {
    return rand.Intn(100)
}

func getRandomY() int {
    return rand.Intn(100)
}

func TestFlyweight(t *testing.T) {
    for i := 0; i < 100; i++ {
        circle := newShapeFactory().GetCircle(getRandomColor())
        circle.SetX(getRandomX())
        circle.SetY(getRandomY())
        circle.SetRadius(10)
        circle.Draw()
    }
}
