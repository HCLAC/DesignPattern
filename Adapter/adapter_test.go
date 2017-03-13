package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	pF := NewForwards("F")
	pC := NewCenter("c")
	pFC := NewTranslator("FC")

	pF.attack()
	pC.defense()
	pFC.attack()
	pFC.defense()
}
