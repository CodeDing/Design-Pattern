package Template

import "testing"

func TestGame(t *testing.T) {
	var g Game
	g = newFootball("football")
	g.Play()
	g = newCricket("cricket")
	g.Play()
}
