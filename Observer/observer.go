package Observer

import "fmt"

type Subject struct {
	state     int
	observers []Observer
}

func NewSubject(state int) *Subject {
	return &Subject{state: state, observers: make([]Observer, 0)}
}

func (s Subject) GetState() int {
	return s.state
}

func (s *Subject) SetState(state int) {
	s.state = state
	s.notify()
}

func (s *Subject) Add(o ...Observer) {
	if s.observers == nil {
		s.observers = make([]Observer, 0)
	}
	s.observers = append(s.observers, o...)
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

type Observer interface {
	Update(*Subject)
}

type BinaryObserver struct{}

func (b BinaryObserver) Update(s *Subject) {
	fmt.Printf("Binary string: b%032b\n", s.GetState())
}

type OctalObserver struct{}

func (o OctalObserver) Update(s *Subject) {
	fmt.Printf("Octal string: 0%04o\n", s.GetState())
}

type HexObserver struct{}

func (h HexObserver) Update(s *Subject) {
	fmt.Printf("Hex string: 0x%08X\n", s.GetState())
}
