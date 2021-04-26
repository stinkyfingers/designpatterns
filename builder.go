package designpatterns

/*
  Builder is useful when the object constructed is complex (unlike these examples) and/or
  when different versions of the same product need to be created. Concrete types implement
  the HouseBuilder interface. The Director has a Builder (interface) field and a BuildHouse method.
*/

type House struct {
	Foundation string
	Walls      string
	Roof       string
}

type HouseBuilder interface {
	BuildFoundation()
	BuildWalls()
	BuildRoof()
	GetHouse() *House
}

type Director struct {
	Builder HouseBuilder
}

func NewDirector(h HouseBuilder) *Director {
	return &Director{Builder: h}
}

func (d *Director) BuildHouse() *House {
	d.Builder.BuildFoundation()
	d.Builder.BuildWalls()
	d.Builder.BuildRoof()
	return d.Builder.GetHouse()
}

/* Concrete Types, implement HouseBuilder interface */
type CapeCod struct {
	Foundation string
	Walls      string
	Roof       string
}

func (c *CapeCod) BuildFoundation() {
	c.Foundation = "concrete"
}

func (c *CapeCod) BuildWalls() {
	c.Walls = "drywall"
}

func (c *CapeCod) BuildRoof() {
	c.Roof = "shingles"
}

func (c *CapeCod) GetHouse() *House {
	return &House{
		Foundation: c.Foundation,
		Walls:      c.Walls,
		Roof:       c.Roof,
	}
}

type Igloo struct {
	Foundation string
	Walls      string
	Roof       string
}

func (i *Igloo) BuildFoundation() {
	i.Foundation = "snow"
}

func (i *Igloo) BuildWalls() {
	i.Walls = "ice blocks"
}

func (i *Igloo) BuildRoof() {
	i.Roof = "more ice"
}

func (i *Igloo) GetHouse() *House {
	return &House{
		Foundation: i.Foundation,
		Walls:      i.Walls,
		Roof:       i.Roof,
	}
}
