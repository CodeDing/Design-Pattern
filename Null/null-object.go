package Null

type AbstractCustomer interface {
    IsNil() bool
    GetName() string
}

type RealCustomer struct {
    name string
}

func newRealCustomer(name string) *RealCustomer {
    return &RealCustomer{name: name}
}

func (r RealCustomer) IsNil() bool {
    return false
}

func (r RealCustomer) GetName() string {
    return r.name
}

type NullCustomer struct{}

func (NullCustomer) IsNil() bool {
    return true
}

func (NullCustomer) GetName() string {
    return "Not Implemented"
}

type CustomerFactory interface {
    GetCustomer(string) AbstractCustomer
}

type CustomerFunc func(string) AbstractCustomer

func (f CustomerFunc) GetCustomer(name string) AbstractCustomer {
    return f(name)
}

func CustomerWithSpecifiedName(names ...string) CustomerFunc {
    return func(n string) AbstractCustomer {
        for _, name := range names {
            if n == name {
                return newRealCustomer(n)
            }
        }
        return NullCustomer{}
    }
}
