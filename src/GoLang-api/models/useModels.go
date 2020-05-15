package models

import (
	"GoLang-api/entities"
	"errors"
)

var (
	ListUser = make([]*entities.Users, 0)
)

//create user
func CreateUser(user *entities.Users) bool {
	if user.Id != "" && user.Name != "" && user.Email != "" && user.Password != "" {
		if userF, _ := FindUser(user.Id); userF == nil {
			ListUser = append(ListUser, user)
			return true
		}
	}
	return false
}

//Update user
func UpdateUser(eUser *entities.Users) bool {
	for index, user := range ListUser {
		if user.Id == eUser.Id {
			ListUser[index] = eUser
			return true
		}
	}
	return false
}

//delete user
func DeleteUser(id string) bool {
	for index, user := range ListUser {
		if user.Id == id {
			copy(ListUser[index:], ListUser[index+1:])
			ListUser[len(ListUser)-1] = &entities.Users{}
			ListUser = ListUser[:len(ListUser)-1]
			return true
		}
	}
	return false
}

//find user
func FindUser(id string) (*entities.Users, error) {
	for _, user := range ListUser {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, errors.New("User does not exist")
}

//get all user
func GetAllUser() []*entities.Users {
	return ListUser
}
