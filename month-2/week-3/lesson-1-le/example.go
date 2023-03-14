package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Car struct {
	Name  string
	Speed int
}

func (c Car) GetSpeed() time.Duration {

	return time.Duration(rand.Intn(1000))
}

func Run(car Car, index int, empty chan struct{}) {

	speed := car.GetSpeed()

	fmt.Printf("%d. %s Speed: %d\n", index, car.Name, speed)

	time.Sleep(time.Millisecond * speed)

	empty <- struct{}{}
}

func main() {

	var cars = []Car{
		{Name: "Supra"},
		{Name: "BMW"},
		{Name: "Damas"},
		{Name: "Nexia"},
	}

	empty := make(chan struct{})

	for in, car := range cars {
		go Run(car, in+1, empty)
	}

	<-empty

	// 1. Shohruh imtixon boshlandi
	// 2. Davlatbak imtixon boshlandi

	// 1. Davlatbak imtixon topshirdi
	// 2. Shohruh topshirdi

}
