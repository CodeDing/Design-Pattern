package State

import `testing`

func TestState(t *testing.T) {
    ctx := newContext("init")
    t.Logf("Init state: %s", ctx.GetState())

    var s State
    s = StartState{}
    s.DoAction(ctx)
    t.Logf("Start state: %s", s)
    s = StopState{}
    s.DoAction(ctx)
    t.Logf("Stop state: %s", s)
}
