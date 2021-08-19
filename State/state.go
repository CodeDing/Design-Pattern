package State

import `fmt`

type State interface {
    DoAction(ctx *Context)
}

type Context struct {
    state string
}

func newContext(s string) *Context {
    return &Context{s}
}

func (c *Context) SetState(s string) {
    if c == nil {
        panic("Context can not be a nil")
    }
    c.state = s
}

func (c Context) GetState() string {
    return c.state
}

type StartState struct{}

func (StartState) DoAction(ctx *Context) {
    fmt.Println("Setting start state ...")
    ctx.SetState("start")
}

func (StartState) String() string {
    return "Start state"
}

type StopState struct{}

func (StopState) DoAction(ctx *Context) {
    fmt.Println("Setting stop state ...")
    ctx.SetState("stop")
}

func (StopState) String() string {
    return "Stop state"
}
