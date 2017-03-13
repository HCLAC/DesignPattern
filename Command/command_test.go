package command

import (
	"testing"
)

func TestCommand(t *testing.T) {
	invoker := NewInvoker()
	concomA := NewConcreteCommandA()
	receA := NewReceiverA()

	concomA.SetReceiver(*receA)
	invoker.AddCommand(concomA)

	concomB := NewConcreteCommandB()
	receB := NewReceiverB()

	concomB.SetReceiver(*receB)
	invoker.AddCommand(concomB)

	invoker.ExecuteCommand()
}
