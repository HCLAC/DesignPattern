package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	context, err := NewCashContext("正常费")
	if err != nil {
		//t.Error(err)
	} else {
		result := context.acceptCash(10000)
		fmt.Println(result)
	}

	context, err = NewCashContext("满300返100")
	if err != nil {
		t.Error(err)
	} else {
		result := context.acceptCash(350)
		fmt.Println(result)
	}

	context, err = NewCashContext("打8折")
	if err != nil {
		t.Error(err)
	} else {
		result := context.acceptCash(300)
		fmt.Println(result)
	}
}
