package bridge

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T) {
	sa := SoftwareA{Software{"a"}}
	sb := SoftwareB{Software{"b"}}

	pa := NewPhoneA("pa")
	pb := NewPhoneB("pb")

	pa.setSoft(&sa)
	pa.Run()

	pb.setSoft(&sb)
	pb.Run()

	fmt.Println()
	p := TSoftware{&sb}
	p.Run()
}
