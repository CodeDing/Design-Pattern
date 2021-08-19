package Observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	s := NewSubject(1)
	s.Add(BinaryObserver{}, OctalObserver{}, HexObserver{})
	t.Logf("First to update subject state to 2!")
	s.SetState(2)
	t.Logf("Second to update subject state to 15!")
	s.SetState(15)

}
