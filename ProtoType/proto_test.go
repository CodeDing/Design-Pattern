package ProtoType

import "testing"

func TestProto(t *testing.T) {
	cache := &ShapeCache{}
	cache.Load()
	shape1 := cache.GetShape("1")
	shape1.Draw()
	t.Logf("Shape1 should be a rectangle, %s", shape1.String())
	shape2 := cache.GetShape("2")
	shape2.Draw()
	t.Logf("Shape2 should be a square, %s", shape2.String())
	shape3 := cache.GetShape("3")
	shape3.Draw()
	t.Logf("Shape3 should be a circle, %s", shape3.String())
}
