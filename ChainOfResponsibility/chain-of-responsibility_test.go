package ChainOfResponsibility

import `testing`

func newChainOfLoggers() AbstractLogger {
    consoleLogger := ConsoleLogger{level: DEBUG}
    fileLogger := FileLogger{level: INFO, next: consoleLogger}
    errorLogger := ErrorLogger{level: ERROR, next: fileLogger}
    return errorLogger
}

func TestChainOfResponsibility(t *testing.T) {
    l := newChainOfLoggers()
    l.LogMessage(DEBUG, "It is a debug information!")
    l.LogMessage(INFO, "It is a info information!")
    l.LogMessage(ERROR, "It is an error information!")
}
