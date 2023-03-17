package main

import (
	"fmt"
	"log"
	// "strings"
	"io/ioutil"
	"net/http"
	"encoding/json"

	"github.com/google/uuid"
)

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

func main() {

	var users []User

	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request){

		// fmt.Fprintf(w, "Hello world")

		if r.Method == "GET" {

			// path := strings.Split(r.URL.Path, "/")

			id := r.URL.Path[len("/user/"):]

			if len(id) > 0 {
				var (
					user User
				)

				for ind, _ := range users {
					if id == users[ind].Id {
						user = users[ind]
						break
					}
				}

				body, err := json.Marshal(user)
				if err != nil {
					log.Println("Unmarshal err:", err)
					w.WriteHeader(500)
					w.Write([]byte("Incorrect data"))
					return
				}

				w.WriteHeader(200)
				w.Write(body)
				return
			} else {
				body, err := json.Marshal(users)
				if err != nil {
					log.Println("Unmarshal err:", err)
					w.WriteHeader(500)
					w.Write([]byte("Incorrect data"))
					return
				}

				w.WriteHeader(200)
				w.Write(body)
				return
			}

		} else if r.Method == "POST" {

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println("ioutil err:", err)
				w.WriteHeader(400)
				w.Write([]byte("Incorrect data"))
				return
			}

			var user User

			err = json.Unmarshal(body, &user)
			if err != nil {
				log.Println("Post err:", err)
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
				return
			}

			user.Id = uuid.New().String()

			body, err = json.Marshal(user)
			if err != nil {
				log.Println("Unmarshal err:", err)
				w.WriteHeader(500)
				w.Write([]byte("Incorrect data"))
				return
			}


			users = append(users, user)
			w.WriteHeader(http.StatusCreated)
			w.Write(body)
		} else {
			w.WriteHeader(405)
			w.Write([]byte("Not Allowed Method"))
		}
	})


	http.HandleFunc("/product", cont.Product)

	fmt.Println("Listent 4000...")

	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
