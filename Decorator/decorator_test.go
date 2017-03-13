package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	person := &person{"hcl"}
	person.show()
	fmt.Println()
	ts := new(TShirts)
	ts.SetDecorator(person)
	ts.show()
	fmt.Println()
	bt := new(BigTrouser)
	bt.SetDecorator(ts)
	bt.show()
	fmt.Println()
	sk := new(Sneakers)
	sk.SetDecorator(bt)
	sk.show()
}
