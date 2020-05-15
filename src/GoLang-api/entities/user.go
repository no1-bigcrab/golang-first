package entities

import "fmt"

type Users struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user Users) ToString() string {
	return fmt.Sprintf("id: %s\nName %s\nEmail %s\nPassword %s\n", user.Id, user.Name, user.Email, user.Password)
}
