package designpatterns

/*
  See test for implementation
  The Observer is useful for notifying dependant objects in a one-to-many relationship about
  a change to the subject.
*/

type Subject interface {
	Register(Observer)
	Deregister(Observer)
	Notify()
}

type Observer interface {
	Update(string)
	GetName() string
}
