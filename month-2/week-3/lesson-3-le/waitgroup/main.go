package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"

	"github.com/bxcodec/faker/v4"
)

type Product struct {
	Name string
	Price float64
}

type User struct {
	Name string
	Products []Product
}

func Cassier_1(wg *sync.WaitGroup, user User) {
	defer wg.Done()

	var totalSum float64 = 0
	for _, product := range user.Products {
		totalSum += product.Price
	}

	fmt.Println(user.Name," Price:", totalSum)
}

func Cassier_2(wg *sync.WaitGroup, user User) {
	defer wg.Done()

	var totalSum float64 = 0
	for _, product := range user.Products {
		totalSum += product.Price
	}

	fmt.Println(user.Name," Price:", totalSum)
}

func Cassier_3(wg *sync.WaitGroup, user User) {
	defer wg.Done()

	var totalSum float64 = 0
	for _, product := range user.Products {
		totalSum += product.Price
	}

	fmt.Println(user.Name," Price:", totalSum)
}

func getRandomProducts() (products []Product) {

	for i := 0; i <= rand.Intn(20); i++ {
		products = append(products, Product{
			Name: faker.FirstName(),
			Price: float64(rand.Intn(100)) * 1000,
		})
	}

	return
}

func main() {

	var users []User

	for i := 0; i < 100; i ++ {

		products := getRandomProducts()
		users = append(users, User{
			Name: faker.FirstName(),
			Products: products,
		})
	}



	length := len(users) - 1
	wg := sync.WaitGroup{}
	now := time.Now()

	for i := 0; i < len(users); i += 3 {
		
		go Cassier_1(&wg, users[i])

		wg.Add(1)
		if i + 1 >= length {
			break
		}

		go Cassier_2(&wg, users[i + 1])

		wg.Add(1)
		if i + 2 >= length {
			break
		}

		go Cassier_3(&wg, users[i + 2])
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println("TIME:", time.Since(now))
}
