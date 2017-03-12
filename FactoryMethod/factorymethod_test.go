package factorymethod

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	var tmp *OperationAdd = nil
	if val, ok := tmp.Result(); ok == true {
		t.Error(val)
	}

	cAdd := COperationAdd{}
	var of OperationFunc = cAdd.createoperation("+")
	of.SetNumA(10)
	of.SetNumB(110)
	if val, ok := of.Result(); ok == true && val != 120 {
		t.Error("Add Error")
	}

	cSub := COperationSub{}
	of = cSub.createoperation("-")
	of.SetNumA(10)
	of.SetNumB(110)
	if val, ok := of.Result(); ok == true && val != -100 {
		t.Error("Sub Error")
	}
}
