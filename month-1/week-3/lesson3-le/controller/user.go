package controller

import (
	"app/models"

	"github.com/bxcodec/faker/v3"
)

var Users []models.User

func CreateUser(data models.User) {
	Users = append(Users, data)
}

// getbyid
func GetUserById(id int) (models.User, bool) {
	for _, val := range Users {
		if val.Id == id {
			return val, false
		}
	}

	return models.User{}, true
}

// update
func UpdateUser(id int) (models.User, bool) {
	for i, val := range Users {
		if val.Id == id {
			Users[i].Name = "changed Name"
			Users[i].Surname = "changed Surname"
			return Users[i], false
		}
	}

	return models.User{}, true
}

// delete
func DeleteUser(id int) bool {
	for i, val := range Users {
		if val.Id == id {
			Users = append(Users[:i], Users[i+1:]...)
			return false
		}
	}

	return true
}

func GetListUser(req models.GetListRequest) (resp []models.User, err bool) {

	if req.Limit+req.Offset > len(Users) {
		return []models.User{}, true
	}

	return Users[req.Offset : req.Limit+req.Offset], false
}

func GenerateUser(count int) {
	for i := 0; i < count; i++ {
		Users = append(Users, models.User{
			Id:      i + 1,
			Name:    faker.FirstName(),
			Surname: faker.LastName(),
		})
	}
}
