package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	password := 123456
	var input int

v1:
	chance := 3
	for {
		fmt.Println("Enter Password:")
		fmt.Scan(&input)

		if input == password {
			fmt.Println("Entered password is correct")
			break
		} else {
			chance--
			if chance <= 0 {
				break
			}
			fmt.Printf("Entered password is wrong, %dta imkoniyat qoldi\n", chance)
		}
	}

	if chance == 0 {
		fmt.Println("You are blocked 30 seconds please wait!")
		for i := 1; i <= 5; i++ {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Println("Your window blocked 5 seconds, please wait Time:", i)
			time.Sleep(time.Second * 1)
		}
		goto v1
	}

}
