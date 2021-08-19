package AbstractFactory

import "testing"

func TestAbstractFactory(t *testing.T) {
	shapeFactory, ok := allFactories["shape"]
	if !ok {
		t.Fatalf("Failed to get shape factory!")
	}
	rectangle := shapeFactory.GetFactory().GetShape("rectangle")
	rectangle.Draw()
	t.Logf("Rectangle area: %f", rectangle.Area())
	square := shapeFactory.GetFactory().GetShape("square")
	square.Draw()
	t.Logf("Square area: %f", square.Area())
	circle := shapeFactory.GetFactory().GetShape("circle")
	circle.Draw()
	t.Logf("Circle area: %f", circle.Area())

	colorFactory, ok := allFactories["color"]
	if !ok {
		t.Fatalf("Failed to get color factory!")
	}
	red := colorFactory.GetFactory().GetColor("red")
	t.Logf("Color filled: %s", red.Fill())
	blue := colorFactory.GetFactory().GetColor("blue")
	t.Logf("Color filled: %s", blue.Fill())
	green := colorFactory.GetFactory().GetColor("green")
	t.Logf("Color filled: %s", green.Fill())
}
