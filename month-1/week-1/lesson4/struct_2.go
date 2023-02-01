package main

import "fmt"

type Color struct {
	R, G, B uint8
}

type Car struct {
	color   Color
	engine  float32
	litr    float64 // overall litr
	km_litr float64 // km in 1 litr
	name    string
}

func (c Car) getColor() int {
	return int(c.color.R) + int(c.color.G) + int(c.color.B)
}

func (c Car) run(kms float64) float64 {
	var ostatokLitr = c.litr - (kms / c.km_litr)
	if ostatokLitr < 0 {
		return 0
	}
	return ostatokLitr
}

func main() {

	// var audi Car = Car{
	// 	color:  Color{0, 255, 255},
	// 	engine: 2.5,
	// }
	var nexia Car = Car{
		color:   Color{0, 255, 255},
		engine:  2.5,
		litr:    50,
		km_litr: 10,
		name:    "Nexia",
	}

	fmt.Println(nexia.run(20))
	// fmt.Println(audi.getColor())
}
