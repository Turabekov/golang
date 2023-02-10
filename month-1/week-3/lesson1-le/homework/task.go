package main

import "fmt"

type Engine struct {
	weight    float64
	timeSpeed int
}

type Car struct {
	engine    Engine
	name      string
	run_minut int
}

type CarInterface interface {
	GetEngine() Engine
	GetInfo()
	GetSpeed() int
}

func (c Car) GetEngine() Engine {
	return c.engine
}

func (c Car) getInfo() {
	fmt.Printf("Info:\n\tname: %s; Run in minute: %d\n", c.name, c.run_minut)
	fmt.Printf("Engine:\n\tWeight: %f; Timespeed: %d\n", c.engine.weight, c.engine.timeSpeed)
}

func (c Car) getSpeed() int {
	return c.engine.timeSpeed
}

func main() {
	mers := Car{engine: Engine{weight: 300, timeSpeed: 300}, name: "Mers w 124", run_minut: 1}

	fmt.Println("Engine:", mers.GetEngine())
	mers.getInfo()
	fmt.Println("Speed:", mers.getSpeed())
}
