package designpatterns

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	s := "Hello World"
	identity := func(s string) string {
		return s
	}

	lower := ToLower(identity)
	fmt.Println(lower(s))
	if lower(s) != "hello world" {
		t.Error("ToLower not lowercasing")
	}

	spaceout := SpaceOut(identity)
	fmt.Println(spaceout(s))
	if spaceout(s) != "H e l l o   W o r l d " {
		t.Error("SpaceOut not space outing")
	}
}
