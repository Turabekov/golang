package main

import "fmt"

func drop(chars []byte, index int) (bool, string, []byte) {
	err := false
	message := ""
	data := []byte{}

	if len(chars) == 0 {
		err = true
		message = "Please enter elements"
	} else {
		for i, v := range chars {
			if i == index {
				continue
			}
			data = append(data, v)
		}
	}

	return err, message, data
}

func main() {

	var chars []byte = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'J', 'I', 'H'}

	err, message, data := drop(chars, 1)

	if err {
		fmt.Println(message)
	} else {
		fmt.Println(string(data)) // 'A', 'C', 'D', 'E', 'F', 'J', 'I', 'H'
	}

}
