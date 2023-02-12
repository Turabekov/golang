package main

import (
	"fmt"

	"app/controller"
	"app/models"
)

func main() {

	var usersNumber = 109
	controller.GenerateUser(usersNumber)
	overallPages := 0
	if usersNumber/10 != 0 {
		overallPages = usersNumber/10 + 1
	}

	fmt.Println("Pagelar soni:", overallPages)
	fmt.Println("Page raqamini kiriting:")
	pageNumber := 0
	fmt.Scan(&pageNumber)

	pageOffset := pageNumber*10 - 10
	limit := 10
	if pageNumber == overallPages {
		limit = usersNumber % 10
	}

	users, err := controller.GetListUser(models.GetListRequest{
		Offset: pageOffset,
		Limit:  limit,
	})

	if err {
		fmt.Println("users out of range")
		return
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
