package Interpreter

import "strings"

type Expression interface {
	interpret(string) bool
}

type TerminalExpress struct {
	data string
}

func (t TerminalExpress) interpret(c string) bool {
	if strings.Contains(c, t.data) {
		return true
	}
	return false
}

type OrExpression struct {
	exp1, exp2 Expression
}

func (o OrExpression) interpret(c string) bool {
	return o.exp1.interpret(c) || o.exp2.interpret(c)
}

type AndExpression struct {
	exp1, exp2 Expression
}

func (a AndExpression) interpret(c string) bool {
	return a.exp1.interpret(c) && a.exp2.interpret(c)
}


