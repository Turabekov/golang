package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var users []User

func main() {

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {

		// fmt.Fprintf(w, "Hello world")

		if r.Method == "GET" {

			data, err := json.Marshal(users)
			if err != nil {
				log.Println("ioutil err:", err)
				w.WriteHeader(400)
				w.Write([]byte("get error"))
				return
			}

			w.WriteHeader(200)
			w.Write(data)

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

			fmt.Println(user)
			// push data
			users = append(users, user)
			w.WriteHeader(201)
			w.Write([]byte("Created"))
		} else if r.Method == "PUT" {

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

			idx := 0

			for i, _ := range users {
				if user.Id == users[i].Id {
					idx = i
					users[i].Name = user.Name
				}
			}

			fmt.Println(users[idx])
			w.WriteHeader(200)
			w.Write([]byte("Updated"))

		} else if r.Method == "DELETE" {

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println("ioutil err:", err)
				w.WriteHeader(400)
				w.Write([]byte("Incorrect data"))
				return
			}

			var userid User

			err = json.Unmarshal(body, &userid)
			if err != nil {
				log.Println("Post err:", err)
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
				return
			}

			for i, _ := range users {
				if userid.Id == users[i].Id {
					users = append(users[:i], users[i+1:]...)
				}
			}

			w.WriteHeader(200)
			w.Write([]byte("Deleted"))
		} else {
			w.WriteHeader(405)
			w.Write([]byte("Not Allowed Method"))
		}
	})

	fmt.Println("Listent 4000...")

	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
