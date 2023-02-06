package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  uint8
}

func sortUsers(users []User, action string) []User {

	if action == "age" {
		sort.Slice(users, func(i, j int) bool {
			return users[i].Age < users[j].Age
		})

		return users

	} else if action == "name" {
		sort.Slice(users, func(i, j int) bool {
			return users[i].Name < users[j].Name
		})

		return users
	} else if action == "lenname" {
		sort.Slice(users, func(i, j int) bool {
			return len(users[i].Name) < len(users[j].Name)
		})

		return users
	}

	return users
}

func main() {

	users := []User{
		{
			Name: "DD",
			Age:  19,
		},
		{
			Name: "FFFFF",
			Age:  24,
		},
		{
			Name: "EEEEEEE",
			Age:  20,
		},
		{
			Name: "BB",
			Age:  17,
		},
		{
			Name: "A",
			Age:  21,
		},
	}

	users = sortUsers(users, "lenname")
	// users = sortUsers(users, "age")

	for _, user := range users {
		fmt.Println(user)
	}

}
