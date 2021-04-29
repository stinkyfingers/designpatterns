package designpatterns

import (
	"testing"
)

type Widget struct {
	Name string
	Type string
}

type Doodad struct {
	Name string
	Type string
}

func (w *Widget) Update(s string) {
	w.Type = s
}

func (w *Widget) GetName() string {
	return w.Name
}

func (d *Doodad) Update(s string) {
	d.Type = s
}

func (d *Doodad) GetName() string {
	return d.Name
}

// subject
type Color struct {
	Hue       string
	Observers []Observer
}

func (c *Color) Register(o Observer) {
	c.Observers = append(c.Observers, o)
}

func (c *Color) Deregister(o Observer) {
	for i, observer := range c.Observers {
		if observer.GetName() == o.GetName() {
			c.Observers = append(c.Observers[:i], c.Observers[i+1:]...)
			break
		}
	}
}

func (c *Color) Notify() {
	for _, observer := range c.Observers {
		observer.Update(c.Hue)
	}
}

func TestObserver(t *testing.T) {
	w := &Widget{Name: "widg"}
	d := &Doodad{Name: "dood"}
	c := Color{}
	c.Register(w)
	c.Register(d)

	c.Hue = "orange"

	c.Notify()

	if w.Type != "orange" || d.Type != "orange" {
		t.Error("expected orange")
	}

	c.Deregister(w)
	c.Deregister(d)
}
