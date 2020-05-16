package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func homePage(w http.ResponseWriter, r *http.Request) {
	ValidToken, err := GenerateJWT()

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, ValidToken)
}

func HandleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func main() {
	fmt.Println("My simple Client")

	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("Something went wrong:")
	}
	fmt.Println(tokenString)
	HandleRequests()
}
