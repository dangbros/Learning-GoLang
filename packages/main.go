package main

import (
	"fmt"

	"github.com/dangbros/Go-package/auth"
	"github.com/dangbros/Go-package/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("John Doe", "secret")
	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "souryaroy30@gmail.com",
		Name:  "Sourya Roy",
	}

	// fmt.Println(user.Email, user.Name)
	//

	color.Green(user.Email)
	color.Red(user.Name)

}
