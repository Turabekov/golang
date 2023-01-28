package main

import (
	"fmt"
	"reflect"
)

func input(inp interface{}) interface{} {
	fmt.Println(reflect.TypeOf(inp))
	return "234"
}

func main() {
	var k interface{} = 10

	fmt.Println(input(k))
}
