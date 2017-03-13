package template

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	b := NewBingA()
	b.g = b
	b.getsomefood()

	g := NewGuo()
	g.g = g
	g.getsomefood()
}
