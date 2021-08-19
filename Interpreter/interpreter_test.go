package Interpreter

import "testing"

func getMaleExpression() Expression {
	robert := TerminalExpress{"Robert"}
	john := TerminalExpress{"John"}
	return OrExpression{robert, john}
}

func getMarriedWomanExpression() Expression {
	julie := TerminalExpress{"Julie"}
	married := TerminalExpress{"Married"}
	return AndExpression{julie, married}
}

func TestInterpreter(t *testing.T) {
	isMale := getMaleExpression()

	isMarried := getMarriedWomanExpression()

	if isMale.interpret("Tom") {
		t.Fatalf("Failed to interpret male, got %t, but want %t", isMale.interpret("Tom"), false)
	}

	if !isMale.interpret("John") {
		t.Fatalf("Failed to interpret male, got %t, but want %t", isMale.interpret("John"), false)
	}

	if !isMarried.interpret("Married Julie") {
		t.Fatalf("Failed to interpret married, got %t, but want true", isMarried.interpret("Julie"))
	}

	if isMarried.interpret("Jack") {
		t.Fatalf("Failed to interpret married, got %t, but want false", isMarried.interpret("Jack"))
	}
}
