package SimpleFactory

import "testing"

func TestSimpleFactory(t *testing.T) {
	f := ShapeFactory{}
	rectangle := f.GetShape(RECTANGLE)
	t.Logf("name: %s, area: %f", rectangle.Name(), rectangle.Area())
	rectangle.Draw()
	square := f.GetShape(SQUARE)
	t.Logf("name: %s, area: %f", square.Name(), square.Area())
	square.Draw()

	circle := f.GetShape(CIRCLE)
	t.Logf("name: %s, area: %f", circle, circle.Area())
	circle.Draw()

	for _, shape := range []string{RECTANGLE, SQUARE, CIRCLE} {
		s := f.GetShape(shape)
		t.Logf("Name: %s, Area: %f", s.Name(), s.Area())
		s.Draw()
	}
}
