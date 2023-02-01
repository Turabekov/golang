package main

import (
	"fmt"
	"strings"
)

type Human struct {
	name     string
	age      uint8
	birthday string
}

type Humans struct {
	humans []Human
}

func (h Humans) ageFilter(min, max uint8) ([]Human, string) {
	var newHuman []Human
	var message string

	if max > min {
		message = "Please enter correct credentials"
		return newHuman, message
	}

	for _, human := range h.humans {
		if human.age >= min && human.age <= max {
			newHuman = append(newHuman, human)
		}
	}

	message = "Succes"

	return newHuman, message
}

func (h Humans) nameFilter(str string) ([]Human, string) {
	var newHuman []Human
	message := ""

	str = strings.ToLower(str)

	for _, human := range h.humans {

		if strings.Contains(strings.ToLower(string(human.name)), str) {
			newHuman = append(newHuman, human)
		}
	}

	if len(message) == 0 {
		message = "user does not exist"
	} else {
		message = "Success"
	}

	return newHuman, message
}

func main() {
	var humans Humans = Humans{
		humans: []Human{
			{
				name:     "Asadbek",
				age:      23,
				birthday: "12-02-2000",
			},
			{
				name:     "Shohruh",
				age:      19,
				birthday: "12-02-2001",
			},
			{
				name:     "Asliddin",
				age:      10,
				birthday: "12-02-2009",
			},
			{
				name:     "Xumoyun",
				age:      20,
				birthday: "12-06-2002",
			},
			{
				name:     "Shaxzod",
				age:      17,
				birthday: "12-06-2006",
			},
			{
				name:     "Mirazam",
				age:      17,
				birthday: "12-06-2006",
			},
			{
				name:     "Atabek",
				age:      37,
				birthday: "12-06-1990",
			},
			{
				name:     "Baxrom",
				age:      42,
				birthday: "12-06-1970",
			},
			{
				name:     "Nodir",
				age:      21,
				birthday: "12-06-1999",
			},
			{
				name:     "Nurmuhammad",
				age:      21,
				birthday: "12-06-1999",
			},
		},
	}

	fmt.Println(humans.ageFilter(20, 40))
	fmt.Println(humans.nameFilter("dfsdfds"))

}
