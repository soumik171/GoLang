package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/soumik171/podcast/auth"
	"github.com/soumik171/podcast/user"
)

func main() {
	auth.LoginWithCredentials("soumik", "123456  ")

	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Name:  "Soumik",
		Email: "user@gmail.com",
	}

	// fmt.Println(user.Name, user.Email)

	color.Green(user.Name)
	color.Blue(user.Email)

}
