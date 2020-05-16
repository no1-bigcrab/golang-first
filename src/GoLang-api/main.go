package main

import (
	"GoLang-api/apis/userApi"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
	Phone:    "49123454322", //this is a random number
}

func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}
func CreateToken(userId uint64) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/user/find", userApi.FindUser).Methods("GET")
	router.HandleFunc("/api/v1/user/getall", userApi.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/user/create", userApi.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/update", userApi.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/user/delete", userApi.Delete).Methods("DELETE")

	//create author jwt
	router.HandleFunc("/login", Login).Methods("POST")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		panic(err)
	}
}
