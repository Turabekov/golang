package main

import "fmt"

type Component struct {
	sugar, weight float64
	count         int
}

type Cake struct {
	Component
	name string
}
type Candle struct {
	Component
	name string
}

func (c Component) getSugar() string {
	return fmt.Sprintf("Sugar: %0.2fgr", c.sugar)
}

func (c Component) GetTotalWeight() string {
	return fmt.Sprintf("Total weight: %0.2f gr\n", float64(c.count)*(c.weight))
}

func main() {
	var napaleon = Cake{
		Component: Component{70, 200, 4},
		name:      "Napoleon cake",
	}

	var mars = Candle{
		Component: Component{80, 50, 10},
		name:      "Shokolad",
	}

	fmt.Println(napaleon.name, napaleon.getSugar())
	fmt.Printf(napaleon.GetTotalWeight())

	fmt.Println(mars.name, mars.getSugar())
	fmt.Printf(mars.GetTotalWeight())

}
