package main

import (
	"fmt"
	"reflect"
)

func direction(data interface{}) (bool, interface{}) {
	var value interface{}
	var err bool = false

	if reflect.TypeOf(data).Name() == "int" {
		switch data {
		case 0:
			value = "TOP"
		case 1:
			value = "BOTTOM"
		case 2:
			value = "LEFT"
		case 3:
			value = "RIGHT"
		default:
			err = true
			value = " "
		}

	} else if reflect.TypeOf(data).Name() == "string" {
		switch data {
		case "TOP":
			value = 0
		case "BOTTOM":
			value = 1
		case "LEFT":
			value = 2
		case "RIGHT":
			value = 3
		default:
			err = true
			value = " "
		}
	}

	return err, value

}

func main() {
	err, value := direction("BOTTOM")

	if err {
		fmt.Println(err, value)
	} else {
		fmt.Println(err, value)
	}
}
