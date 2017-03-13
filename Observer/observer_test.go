package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	var s = NewSubjectA(12)
	var oa = NewObserverA(s, 1)
	var ob = NewObserverB(s, 1)
	var oafu update = oa.Update
	s.AddCallFunc(&oafu)
	var obfu update = ob.Update
	s.AddCallFunc(&obfu)
	s.Notify()

	s.SetState(3)
	s.Notify()
}
