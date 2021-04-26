package designpatterns

import "fmt"

/*
  The Factory pattern creates objects without exposing the creation logic.
  The case below is simple and contrived. It is most useful when a number of
  classes' instantiations share complex logic not worth repeating in each constructor (DRY), e.g.
  DB connections, integrations w/ external services, etc.
*/

type Shape interface {
	Draw()
}

func NewShape(shape string) Shape {
	switch shape {
	case "rectangle":
		return &Rectangle{}
	case "square":
		return &Square{}
	case "triangle":
		return &Triangle{}
	}
	return nil
}

type Rectangle struct {
	Shape
}

type Square struct {
	Shape
}

type Triangle struct {
	Shape
}

func (r *Rectangle) Draw() {
	fmt.Print("------\n|    |\n------\n")
}

func (s *Square) Draw() {
	fmt.Print("----\n|  |\n----\n")
}

func (t *Triangle) Draw() {
	fmt.Print(" ^\n/ \\\n___\n")
}
