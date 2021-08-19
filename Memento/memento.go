package Memento

type Memento struct {
    state string
}

func newMemento(s string) *Memento {
    return &Memento{s}
}

func (m Memento) GetState() string {
    return m.state
}

type Originator struct {
    state string
}

func newOriginator(s string) *Originator {
    return &Originator{s}
}

func (o Originator) GetState() string {
    return o.state
}

func (o *Originator) SetState(s string) {
    o.state = s
}

func (o Originator) SaveStateToMemento() *Memento {
    return newMemento(o.state)
}

func (o *Originator) GetStateFromMemento(m *Memento) {
    o.state = m.GetState()
}

type CareTaker []*Memento

func newCareTaker() *CareTaker {
    var c CareTaker
    c = make([]*Memento, 0)
    return &c
}

func (c *CareTaker) Add(m *Memento) {
    if c == nil {
        panic("caretaker is a nil")
    }
    *c = append(*c, m)
}

func (c CareTaker) Get(idx int) *Memento {
    if idx >= len(c) {
        return nil
    }
    return c[idx]
}
