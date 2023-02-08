package main

import (
	"fmt"
)

type Product struct {
	id    int
	name  string
	price int
}

type Category struct {
	id       int
	name     string
	products []Product
}

type Cart struct {
	product  Product
	quantity int
}

func main() {
	categories := []Category{
		{
			id:   1,
			name: "Mevalar",
			products: []Product{
				{
					id:    1,
					name:  "Olma",
					price: 2000,
				},
				{
					id:    2,
					name:  "Banan",
					price: 4000,
				},
				{
					id:    3,
					name:  "Gilos",
					price: 4000,
				},
			},
		},
		{
			id:   2,
			name: "Ichimliklar",
			products: []Product{
				{
					id:    1,
					name:  "Fanta",
					price: 6000,
				},
				{
					id:    2,
					name:  "Coca-Cola",
					price: 6500,
				},
				{
					id:    3,
					name:  "Sprite",
					price: 5000,
				},
			},
		},
	}
	carts := []Cart{}
	// ========================================================================
	summa := 0
	command1 := 0
v1:
	for _, value := range categories {
		fmt.Printf("%d. %s\n", value.id, value.name)

	}
	fmt.Scanf("%d", &command1)
	if command1 > len(categories) {
		fmt.Println("Error Input: not have category")
		goto v1

	}
	idx1 := command1 - 1

v2:
	for _, val := range categories[int(idx1)].products {
		fmt.Printf("%d. %s\n", val.id, val.name)
	}
	fmt.Printf("0. ortga\n")
	fmt.Printf("-1. tamomlash\n")
	command2 := 0
	fmt.Scanf("%d", &command2)
	if command2 > len(categories[idx1].products) {
		fmt.Println("Error Input: not have products")
		goto v2

	}

	if command2 == 0 {
		goto v1
	} else if command2 == -1 {
		goto v3
	} else {
		idx2 := command2 - 1
		qunatity := 0
		product := categories[int(idx1)].products[idx2]

		fmt.Printf("Choose Qunatity of %s product:\n", product.name)
		fmt.Printf("0. ortga\n")
		fmt.Printf("-1. tamomlash\n")

		fmt.Scanf("%d", &qunatity)
		if qunatity == 0 {
			goto v2
		} else if qunatity == -1 {
			goto v3
		}
		carts = append(carts, Cart{
			product:  product,
			quantity: qunatity,
		})
		goto v2
	}

v3:
	fmt.Println("Check:")
	for id, val := range carts {

		fmt.Printf("%d. %s quantity: %dx\n", id+1, val.product.name, val.quantity)
		summa += val.quantity * val.product.price
	}
	fmt.Println("Summa", summa)

}
