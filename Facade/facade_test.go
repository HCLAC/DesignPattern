package facade

import (
	"testing"
)

func TestFacade(t *testing.T) {
	f := NewFacade(1, 2.3, "3.4")
	f.OutOne()
	f.OutTwo()
}
