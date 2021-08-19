package Vistor

import (
    `encoding/json`
    `errors`
    `fmt`
)

type ComputerPartVisitor func(part ComputerPart) error

type ComputerPart interface {
    accept(ComputerPartVisitor)
}

type KeyBoard struct {
    Name     string `json:"name"`
    Category string `json:"category"`
}

func newKeyBoard(name string) *KeyBoard {
    return &KeyBoard{name, "Keyboard"}
}

func (k KeyBoard) accept(v ComputerPartVisitor) {
    fmt.Println(k)
    v(k)
}

func (k KeyBoard) String() string {
    return k.Category + ":" + k.Name
}

type Monitor struct {
    Name     string `json:"name"`
    Category string `json:"category"`
}

func newMonitor(name string) *Monitor {
    return &Monitor{name, "Monitor"}
}

func (m Monitor) accept(v ComputerPartVisitor) {
    fmt.Println(m)
    v(m)
}

func (m Monitor) String() string {
    return m.Category + ":" + m.Name
}

type Mouse struct {
    Name     string `json:"name"`
    Category string `json:"category"`
}

func newMouse(name string) *Mouse {
    return &Mouse{name, "Mouse"}
}

func (m Mouse) accept(v ComputerPartVisitor) {
    fmt.Println(m)
    v(m)
}

func (m Mouse) String() string {
    return m.Category + ":" + m.Name
}

type Computer struct {
    Name  string         `json:"name"`
    Parts []ComputerPart `json:"parts"`
}

func newComputer(name string) *Computer {
    var parts []ComputerPart
    parts = append(parts, newKeyBoard("Lenovo"), newMonitor("Sharp"), newMouse("logitech"))
    return &Computer{name, parts}
}

func (c Computer) accept(v ComputerPartVisitor) {
    for i := 0; i < len(c.Parts); i++ {
        c.Parts[i].accept(v)
    }
    v(c)
}

func (c Computer) String() string {
    return "Computer:" + c.Name
}

func KeyBoardVisitor() ComputerPartVisitor {
    return func(part ComputerPart) error {
        v, ok := part.(KeyBoard)
        if !ok {
            return errors.New("Not a keyboard")
        }
        fmt.Printf("Keyboard brand: %s\n", v)
        return nil
    }
}

func MonitorVisitor() ComputerPartVisitor {
    return func(part ComputerPart) error {
        v, ok := part.(Monitor)
        if !ok {
            return errors.New("Not a monitor")
        }
        fmt.Printf("Monitor brand: %s\n", v)
        return nil
    }
}

func MouseVisitor() ComputerPartVisitor {
    return func(part ComputerPart) error {
        v, ok := part.(Mouse)
        if !ok {
            return errors.New("Not a mouse")
        }
        fmt.Printf("Mouse brand: %s\n", v)
        return nil
    }
}

func ComputerVisitor() ComputerPartVisitor {
    return func(part ComputerPart) error {
        fmt.Println("Start computer visit ...")
        v, ok := part.(Computer)
        if ok {
            data, _ := json.Marshal(v)
            fmt.Printf("Computer: %s\n", string(data))
        }

        k, ok := part.(KeyBoard)
        if ok {
            fmt.Printf("Computer visit %s\n", k)
            return nil
        }

        m, ok := part.(Mouse)
        if ok {
            fmt.Printf("Computer visit %s\n", m)
            return nil
        }

        mo, ok := part.(Monitor)
        if ok {
            fmt.Printf("Computer visit %s\n", mo)
            return nil
        }

        fmt.Println("Leave computer visit ...")
        return nil
    }
}
