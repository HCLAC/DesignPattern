package flyweight

import (
	"testing"
)

func TestFlyweight(t *testing.T) {
	ff := NewFlyweightFactory()

	fya := ff.Flyweight("a")
	fya.Operation(1)

	fyb := ff.Flyweight("b")
	fyb.Operation(2)

	fyc := ff.Flyweight("c")
	fyc.Operation(3)

	fyd := ff.Flyweight("d")
	fyd.Operation(4)

	fyu := ff.Flyweight("u")
	fyu.Operation(5)
}
