package Builder

import (
	"bytes"
	"strconv"
)

//food item and food packing interface
type Item interface {
	Name() string
	Packing() Packing
	Price() float64
}

type Packing interface {
	Pack() string
}

type Wrapper struct{}

func (w Wrapper) Pack() string {
	return "Wrapper"
}

type Bottle struct{}

func (b Bottle) Pack() string {
	return "Bottle"
}

type Burger interface {
	Item
	Packing
}

type ColdDrink interface {
	Item
	Packing
}

type VegBurger struct {
	name  string
	price float64
}

func (vg VegBurger) Name() string {
	return vg.Name()
}

func (vg VegBurger) Packing() Packing {
	return Wrapper{}
}

func (vg VegBurger) Price() float64 {
	return vg.price
}

type ChickenBurger struct {
	name  string
	price float64
}

func (cb ChickenBurger) Name() string {
	return cb.name
}

func (cb ChickenBurger) Packing() Packing {
	return Wrapper{}
}

func (cb ChickenBurger) Price() float64 {
	return cb.price
}

type Coke struct {
	name  string
	price float64
}

func (c Coke) Name() string {
	return c.name
}

func (c Coke) Packing() Packing {
	return Bottle{}
}

func (c Coke) Price() float64 {
	return c.price
}

type Pepsi struct {
	name  string
	price float64
}

func (c Pepsi) Name() string {
	return c.name
}

func (c Pepsi) Packing() Packing {
	return Bottle{}
}

func (c Pepsi) Price() float64 {
	return c.price
}

type Meal []Item

func NewMeal() *Meal {
	var m Meal
	m = make([]Item, 0)
	return &m
}

func (m *Meal) AddItem(item Item) {
	if m == nil {
		*m = make([]Item, 0)
	}
	*m = append(*m, item)
}

func (m *Meal) GetCost() float64 {
	if m == nil {
		return 0.0
	}
	var cost float64
	for _, item := range *m {
		cost += item.Price()
	}
	return cost
}

func (m *Meal) ShowItems() string {
	if m == nil {
		return "Meal: none"
	}
	var buf bytes.Buffer
	for _, item := range *m {
		buf.WriteString("Item: " + item.Name())
		buf.WriteString(", Packing: " + item.Packing().Pack())
		buf.WriteString(", Price: " + strconv.FormatFloat(item.Price(), 'f', 2, 64))
		buf.Write([]byte{'\n'})
	}
	buf.WriteString("Total: "+ strconv.FormatFloat(m.GetCost(), 'f',2, 64))
	buf.Write([]byte{'\n'})
	return buf.String()
}

type MealFactory struct {
	//VegMeal    Meal
	//NonVegMeal Meal
}

func (f MealFactory) PrepareVegMeal() *Meal {
	 m := NewMeal()
	 m.AddItem(&VegBurger{"vegburger", 10.0})
	 m.AddItem(&Coke{"coke", 11.20})
	 return m
}

func (f MealFactory) PrepareNonVegMeal() *Meal {
	m := NewMeal()
	m.AddItem(&ChickenBurger{"chickenburger", 20.0})
	m.AddItem(&Pepsi{"pepsi", 30.0})
	return m
}
