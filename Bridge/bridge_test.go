package Bridge

import "testing"

func TestCircle_Draw(t *testing.T) {
	c1 := NewCircle(1,1, 2, newRedCircle())
	c2 := NewCircle(2,2,4, newGreenCircle())
	t.Logf("Draw redcircle, %s", c1.Draw())
	t.Logf("Draw greencircle, %s", c2.Draw())
}
