package ChainOfResponsibility

import `fmt`

type Level int

const (
    DEBUG Level = iota
    INFO
    ERROR
)

type AbstractLogger interface {
    LogMessage(Level, string)
    write(string)
}

type ConsoleLogger struct {
    level Level
    next  AbstractLogger
}

func (c ConsoleLogger) LogMessage(l Level, m string) {
    if c.level <= l {
        c.write(m)
        return
    }
    if c.next != nil {
        c.next.LogMessage(l, m)
    }
}

func (c ConsoleLogger) write(m string) {
    fmt.Println("[ConsoleLogger] ", m)
}

type ErrorLogger struct {
    level Level
    next  AbstractLogger
}

func (e ErrorLogger) LogMessage(level Level, m string) {
    if e.level <= level {
        e.write(m)
        return
    }

    if e.next != nil {
        e.next.LogMessage(level, m)
    }
}

func (e ErrorLogger) write(m string) {
    fmt.Println("[ErrorLogger] ", m)
}

type FileLogger struct {
    level Level
    next  AbstractLogger
}

func (f FileLogger) LogMessage(level Level, m string) {
    if f.level <= level {
        f.write(m)
        return
    }
    if f.next != nil {
        f.next.LogMessage(level, m)
    }
}

func (f FileLogger) write(m string) {
    fmt.Println("[FileLogger] ", m)
}
