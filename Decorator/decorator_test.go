package Decorator

import `testing`

func TestShapeDecorator(t *testing.T) {
    var s Shape
    s = newCircle("circle", 3)

    //RedShapeDecorator(s)
    t.Logf("RecShapeDecorator: %s", RedShapeDecorator2(s))
    s = newRectangle("rect", 3,3)
    //RedShapeDecorator(s)
    t.Logf("RecShapeDecorator: %s", RedShapeDecorator2(s))

}
