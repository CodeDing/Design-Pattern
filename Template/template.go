package Template

import "fmt"

type Game interface {
	initialize()
	startPlay()
	endPlay()
	Play()
}

type Football struct {
	name string
}

func newFootball(n string) *Football {
	return &Football{n}
}

func (f Football) initialize() {
	fmt.Printf("initialize ===> %s\n", f.name)
}

func (f Football) startPlay() {
	fmt.Printf("start play ===> %s\n", f.name)
}

func (f Football) endPlay() {
	fmt.Printf("end   play ===> %s\n", f.name)
}

func (f Football) Play() {
	f.initialize()
	f.startPlay()
	f.endPlay()
}

type Cricket struct {
	name string
}

func newCricket(n string) *Cricket {
	return &Cricket{n}
}

func (c Cricket) initialize() {
	fmt.Printf("initialize  ===> %s\n", c.name)
}

func (c Cricket) startPlay() {
	fmt.Printf("start play  ===> %s\n", c.name)
}

func (c Cricket) endPlay() {
	fmt.Printf("end   play  ===> %s\n", c.name)
}

func (c Cricket) Play() {
	c.initialize()
	c.startPlay()
	c.endPlay()
}
