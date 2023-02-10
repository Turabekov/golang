package main

import "fmt"

func main() {

	sport := struct {
		id   int
		name string
	}{
		id:   1,
		name: "sport",
	}

	fmt.Println(sport)
}
