package designpatterns

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	var previousAddr *SingletonObject

	for i := 0; i < 10; i++ {
		instance = ThreadSafeSingleton()

		if i > 0 && instance != previousAddr {
			t.Errorf("expected instance addr to equal last instance's address ")
		}

		fmt.Printf("%p %p\n", instance, previousAddr)
		previousAddr = instance
	}
}
