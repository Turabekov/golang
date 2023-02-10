package main

import (
	"errors"
	"fmt"
)

type User struct {
	Username string
	Password string
}

type Auth interface {
	Register(user User) error
	Login(username, password string) bool
}

type InMemoryAuth struct {
	users []User
}

func (auth *InMemoryAuth) Register(user User) error {
	for _, val := range auth.users {
		if val.Username == user.Username {
			return errors.New("user like this username already exist")
		}
	}

	auth.users = append(auth.users, user)
	return nil
}

func (auth *InMemoryAuth) Login(username, password string) bool {
	for _, val := range auth.users {
		if val.Username == username {
			if val.Password == password {
				return true
			}
		}
	}

	return false
}

func main() {

	inMemAuth := InMemoryAuth{}

	for {
		fmt.Println("1.Login")
		fmt.Println("2.Register")
		fmt.Println("0.Tamomlash")
		input := 0
		fmt.Scan(&input)

		if input == 1 {
			for {
				var password string
				var username string
				fmt.Printf("username: ")
				fmt.Scanf("%s", &username)
				fmt.Printf("password: ")
				fmt.Scanf("%s", &password)

				boolVar := inMemAuth.Login(username, password)
				if boolVar {
					break
				} else {
					fmt.Println("Credentials are wrong")
					continue
				}
			}
			continue
		}

		if input == 2 {
			for {
				var password string
				var confirmPassword string
				var username string
				fmt.Printf("username: ")
				fmt.Scanf("%s", &username)
				fmt.Printf("password: ")
				fmt.Scanf("%s", &password)
				fmt.Printf("confirm password: ")
				fmt.Scanf("%s", &confirmPassword)

				if password == confirmPassword {
					err := inMemAuth.Register(User{Username: username, Password: password})
					if err != nil {
						fmt.Println(err)
						continue
					}

					fmt.Println(inMemAuth.users)
					break
				} else {
					fmt.Println("Password doesn't matched")
					continue
				}
			}
			continue

		}

		if input == 0 {
			break
		}

	}

}
