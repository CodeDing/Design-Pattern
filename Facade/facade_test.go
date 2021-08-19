package Facade

import "testing"

func TestShapeMaker(t *testing.T) {
	s := NewShapeMaker()

	t.Logf("Draw rectangle: %s", s.DrawRectangle())
	t.Logf("Draw square: %s", s.DrawSquare())
	t.Logf("Draw circle: %s", s.DrawCircle())
}
