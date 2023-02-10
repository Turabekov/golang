package main

import "fmt"

type Engine struct {
	weight    float64
	timeSpeed int
}

type Car struct {
	Engine
	name      string
	run_minut int
}

type Plane struct {
	Engine
	name      string
	run_minut int
}

func (e Engine) getSpeed() int {
	return e.timeSpeed
}

// overrading for Car
func (c Car) getSpeed() int {
	return c.timeSpeed * c.Engine.getSpeed()
}

// overrading for Plane
func (p Plane) getSpeed() int {
	return p.timeSpeed * p.Engine.getSpeed()
}

type engineInterface interface {
	getSpeed() int
}

func Run(e engineInterface) int {
	return e.getSpeed()
}

func main() {
	nexia := Car{Engine: Engine{weight: 6.2, timeSpeed: 900}, name: "Nexia", run_minut: 2}
	boing := Plane{Engine: Engine{weight: 8.2, timeSpeed: 1200}, name: "Boing", run_minut: 1}

	fmt.Println(nexia.name, "Speed:", Run(nexia))
	fmt.Println(boing.name, "Speed:", Run(boing))

}
