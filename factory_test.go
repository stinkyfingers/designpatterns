package designpatterns

import (
	"reflect"
	"testing"
)

func TestFactory(t *testing.T) {
	rectangle := NewShape("rectangle")
	if reflect.TypeOf(rectangle) != reflect.TypeOf(&Rectangle{}) {
		t.Errorf("expected *designpatterns.Rectangle, got %v", reflect.TypeOf(rectangle))
	}
	rectangle.Draw()

	square := NewShape("square")
	if reflect.TypeOf(square) != reflect.TypeOf(&Square{}) {
		t.Errorf("expected *designpatterns.Square, got %v", reflect.TypeOf(square))
	}
	square.Draw()

	triangle := NewShape("triangle")
	if reflect.TypeOf(triangle) != reflect.TypeOf(&Triangle{}) {
		t.Errorf("expected *designpatterns.Triangle, got %v", reflect.TypeOf(rectangle))
	}
	triangle.Draw()
}
