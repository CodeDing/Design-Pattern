package Strategy

import `testing`

func TestStrategy(t *testing.T) {
    m := newMethod(Add{})
    t.Logf("10+5=%d", m.Execute(10, 5))
    m = newMethod(Subtract{})
    t.Logf("10-5=%d", m.Execute(10, 5))
    m = newMethod(Multiply{})
    t.Logf("10*5=%d", m.Execute(10, 5))
}
