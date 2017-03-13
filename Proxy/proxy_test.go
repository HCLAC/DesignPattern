package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {

	girl := Girl{}
	girl.SetName("HCLAC")

	p := NewProxy(girl)
	p.giveDolls()
	p.giveFlowers()
	p.giveChocolate()
}
