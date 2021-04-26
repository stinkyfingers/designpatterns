package designpatterns

import (
	"fmt"
	"strings"
)

/*
  The Decorator pattern is a function that takes another function of a specific type
  and returns a functino of the same type. It is useful for adding & chaining functionality.
  In Go, http middleware is arguably considered Decorator.
*/

type StringModifier func(string) string

func ToLower(s StringModifier) StringModifier {
	return func(s string) string {
		return strings.ToLower(s)
	}
}

func SpaceOut(s StringModifier) StringModifier {
	return func(s string) string {
		var out string
		for _, r := range s {
			out += fmt.Sprintf("%s ", string(r))
		}
		return out
	}
}
