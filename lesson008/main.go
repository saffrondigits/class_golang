package main

import "fmt"

func main() {
	DBusername := "ravi123"
	DBpassword := "Secret01#"

	var usernameInput string
	var passwordInput string

	fmt.Println("Enter the username:")
	fmt.Scanln(&usernameInput)

	fmt.Println("Enter the password:")
	fmt.Scanln(&passwordInput)

	if usernameInput == DBusername && passwordInput == DBpassword {
		fmt.Println("The user is successfully authenticated")
	} else {
		fmt.Println("Username/Password wrong")
	}

}
