package interpreter

import (
	"testing"
)

func TestInterpreter(t *testing.T) {
	context := Context{"abc"}

	list := []IAbstractExpression{}

	list = append(list, new(TerminalExpression))
	list = append(list, new(TerminalExpression))
	list = append(list, new(NonterminalExpression))

	for _, val := range list {
		val.Interpret(&context)
	}
}
