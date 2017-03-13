package state

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T) {
	w := NewWork()

	fmt.Println("========================")
	w.writeProgram()
	fmt.Println("========================")
	w.SetHour(12)
	w.writeProgram()
	fmt.Println("========================")
	w.SetHour(14)
	w.writeProgram()
	fmt.Println("========================")
	w.SetHour(21)
	w.writeProgram()
}
