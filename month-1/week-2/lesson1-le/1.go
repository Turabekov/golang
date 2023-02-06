package main

import (
	"fmt"
)

type User struct {
	Id int
	Name string
}

func main() {

	var users map[int]User = map[int]User{
		1: User{
			Id: 1,
			Name: "Behruz",
		},
		2: User{
			Id: 2,
			Name: "Humoyun",
		},
		3: User{
			Id: 3,
			Name: "Abduqodir",
		},
		4: User{
			Id: 4,
			Name: "Davlatbek",
		},
	}

	fmt.Println(users[1])
	fmt.Println(users[2])
	fmt.Println(users[3])

	// comma ok
	if _, ok := users[4]; !ok {
		fmt.Println("Not found users")
		return
	}

	fmt.Println(users[4])

	// 5 ta dictionary
		// en:uz
		// dict["apple"] // olma
}
