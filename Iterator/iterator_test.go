package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	bg := BookGroup{}
	nb := Book{"a"}
	bg.add(nb)
	nbb := Book{"b"}
	bg.add(nbb)

	it := bg.createIterator()

	for b := it.first(); b != nil; b = it.next() {
		fmt.Println(b)
	}
}
