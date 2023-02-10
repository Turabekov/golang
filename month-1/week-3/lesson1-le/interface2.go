package main

import "fmt"

type Abstract interface {
	GetName() string
}

type Audi struct {
	Name   string
	Engine float64
}

type BMW struct {
	Name   string
	Engine float64
}

func (c Audi) GetName() string {
	return "Name: " + c.Name
}

func (m BMW) GetName() string {
	return "Name: " + m.Name
}

func main() {

	var audi_r8 Abstract = Audi{
		Name:   "Audi R8",
		Engine: 6.4,
	}

	var bmwx6 Abstract = BMW{
		Name:   "BMW X6",
		Engine: 4.2,
	}

	fmt.Println(audi_r8.GetName())
	fmt.Println(bmwx6.GetName())
}
