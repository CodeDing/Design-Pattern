package Proxy

import "testing"

func TestProxyImage_Display(t *testing.T) {
	p := newProxyImage("file1.txt")
	p.Display()

	p.Display()
}
