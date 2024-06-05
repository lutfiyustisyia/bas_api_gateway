package main

import (
	"fmt"

	"api_gateway/utils"
)

func main() {
	fmt.Println("-----------------login------------------")
	login := &utils.Login{
		CorrectUsername: "admin",
		CorrectPassword: "admin",
	}
	username := "admin"
	password := "admin"

	if utils.Authenticate(login, username, password) {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Login failed!")
	}
}
