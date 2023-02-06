package main

import "fmt"

type Human struct {
	// field
	name string
	age uint8
}

func main() {

	// human := Human{ age: 23, name: "Asadbek"}

	// fmt.Println(human)

	// var human2 Human

	// human2.name = "Shohruh"
	// human2.age = 19

	// fmt.Println(human2)


	humans := []Human {
		{
			name: "Asadbek",
			age: 23,
		},
		{
			name: "Shohruh",
			age: 19,
		},
		{
			name: "Abduqodir",
			age: 17,
		},
	}

	for _, human := range humans {
		fmt.Println(human.name, human.age)
	}
}
