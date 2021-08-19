package Vistor

import `testing`

func TestVisitor(t *testing.T) {
    computer := newComputer("Dell")
    visitor := ComputerVisitor()
    computer.accept(visitor)
}
