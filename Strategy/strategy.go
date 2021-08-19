package Strategy

type Strategy interface {
    Operate(int, int) int
}

type Add struct{}

func (a Add) Operate(num1, num2 int) int {
    return num1 + num2
}

type Subtract struct{}

func (s Subtract) Operate(num1, num2 int) int {
    return num1 - num2
}

type Multiply struct{}

func (m Multiply) Operate(num1, num2 int) int {
    return num1 * num2
}

type Method struct {
    Strategy
}

func newMethod(s Strategy) *Method {
    return &Method{s}
}

func (m *Method) Execute(num1, num2 int) int {
    if m == nil {
        panic("method can not be nil")
    }
    return m.Strategy.Operate(num1, num2)
}
