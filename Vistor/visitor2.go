package Vistor

import `fmt`

type Action string

const (
    DELETE Action = "delete"
    GET    Action = "get"
    PUT    Action = "put"
    POST   Action = "POST"
)

type VisitorFunc func(*Info, error) error

type Visitor interface {
    Visit(visitorFunc VisitorFunc) error
}

type Info struct {
    Namespace string
    Name      string
    Action    Action
}

func (i *Info) Visit(fn VisitorFunc) error {
    return fn(i, nil)
}

type NameVisitor struct {
    visitor Visitor
}

func (n NameVisitor) Visit(fn VisitorFunc) error {
    return n.visitor.Visit(func(info *Info, err error) error {
        fmt.Println("NameVisitor call before...")
        err = fn(info, err)
        if err == nil {
            fmt.Printf("=====NameVisitor=====, namespace=%s, name=%s, action=%s\n", info.Namespace, info.Name, info.Action)
        }
        fmt.Println("NameVisitor call after...")
        return err
    })
}

func NameVisitorFunc() VisitorFunc {
    return func(info *Info, err error) error {
        fmt.Println("NameVisitor call before...")
        if err == nil {
            fmt.Printf("=====NameVisitor=====, namespace=%s, name=%s, action=%s\n", info.Namespace, info.Name, info.Action)
        }
        fmt.Println("NameVisitor call after...")
        return err
    }
}

type LogVisitor struct {
    visitor Visitor
}

func (l LogVisitor) Visit(fn VisitorFunc) error {
    return l.visitor.Visit(func(info *Info, err error) error {
        fmt.Println("LogVisitor call before...")
        err = fn(info, err)
        if err == nil {
            fmt.Printf("========LogVisitor====, namespace=%s, name=%s, action=%s\n", info.Namespace, info.Name, info.Action)
        }
        fmt.Println("LogVisitor call after...")
        return err
    })
}

func LogVisitorFunc() VisitorFunc {
    return func(info *Info, err error) error {
        fmt.Println("LogVisitor call before...")
        if err == nil {
            fmt.Printf("##########LogVisitorFunc########, namespace=%s, name=%s, action=%s\n", info.Namespace, info.Name, info.Action)
        }
        fmt.Println("LogVisitor call after...")
        return err
    }
}

type OtherVisitor struct {
    visitor Visitor
}

func (o OtherVisitor) Visit(fn VisitorFunc) error {
    return o.visitor.Visit(func(info *Info, err error) error {
        fmt.Println("OtherVisitor call before...")
        err = fn(info, err)
        if err == nil {
            fmt.Printf("======OtherVisitor======, namespace=%s, name=%s, action=%s\n", info.Namespace, info.Name, info.Action)
        }
        fmt.Println("OtherVisitor call after...")
        return err
    })
}

type DecoratorVisitor struct {
    visitor   Visitor
    decorator []VisitorFunc
}

func NewDecorator(v Visitor, fn ...VisitorFunc) *DecoratorVisitor {
    return &DecoratorVisitor{v, fn}
}

func (d DecoratorVisitor) Visit(fn VisitorFunc) error {
    return d.visitor.Visit(func(info *Info, err error) error {
        if err != nil {
            return err
        }
        err = fn(info, err)
        if err != nil {
            return err
        }
        for i := range d.decorator {
            err = d.decorator[i](info, err)
            if err != nil {
                return err
            }
        }
        return nil
    })
}
