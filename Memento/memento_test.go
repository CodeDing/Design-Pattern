package Memento

import `testing`

func TestMemento(t *testing.T) {
    o := newOriginator("State #1")
    o.SetState("State #2")
    c := newCareTaker()
    c.Add(o.SaveStateToMemento())
    o.SetState("State #3")
    c.Add(o.SaveStateToMemento())
    o.SetState("State #4")

    t.Logf("Current state: %s", o.GetState())
    o.GetStateFromMemento(c.Get(0))
    t.Logf("First saved state: %s", o.GetState())
    o.GetStateFromMemento(c.Get(1))
    t.Logf("Second saved state: %s", o.GetState())
}
