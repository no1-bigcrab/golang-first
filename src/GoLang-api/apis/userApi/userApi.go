package userApi

import (
	"GoLang-api/entities"
	"GoLang-api/models"
	"encoding/json"
	"net/http"
)

//find user
func FindUser(response http.ResponseWriter, request *http.Request) {
	ids, ok := request.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		responseWithError(response, http.StatusBadRequest, "Url Param id is missing")
		return
	}
	user, err := models.FindUser(ids[0])
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
		return
	}
	responseWithJSON(response, http.StatusOK, user)
}

//update user

func CreateUser(response http.ResponseWriter, request *http.Request) {
	var user entities.Users
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		result := models.CreateUser(&user)
		if !result {
			responseWithError(response, http.StatusBadRequest, "Could not create user")
			return
		}
		responseWithJSON(response, http.StatusOK, user)
	}
}

//Get all user
func GetAll(response http.ResponseWriter, request *http.Request) {
	users := models.GetAllUser()
	responseWithJSON(response, http.StatusOK, users)
}

//update user
func UpdateUser(response http.ResponseWriter, request *http.Request) {
	var user entities.Users
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		result := models.UpdateUser(&user)
		if !result {
			responseWithError(response, http.StatusBadRequest, "Could not update user")
			return
		}
		responseWithJSON(response, http.StatusOK, "Update user successfully")
	}
}

//deletet user
func Delete(response http.ResponseWriter, request *http.Request) {
	ids, ok := request.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		responseWithError(response, http.StatusBadRequest, "Url Param id is missing")
		return
	}
	result := models.DeleteUser(ids[0])
	if !result {
		responseWithError(response, http.StatusBadRequest, "Could not delete user")
		return
	}
	responseWithJSON(response, http.StatusOK, "Delete user successfully")
}

//responsive Error
func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJSON(response, statusCode, map[string]string{
		"error": msg,
	})
}

//responsive Json
func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}
