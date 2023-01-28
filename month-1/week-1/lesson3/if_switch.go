package main

import "fmt"

func main() {

	var status = "active"
	switch status {
	case "active":
		fmt.Println("Active")
	case "noactive":
		fmt.Println("noacitve")
	default:
		fmt.Println("Not Status")
	}
}
