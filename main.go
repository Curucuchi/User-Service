package main

import (
	"fmt"
	"strings"
	"users/db"
)

func main() {
	answer := ""
	fmt.Println("Do you wanna signup, login, or delete your account? (s, l, d)")
	fmt.Scanf("%s", &answer)
	begin(answer)
}

func begin(answer string) {
	fName := ""
	lName := ""
	Email := ""
	uName := ""
	pWord := ""
	if strings.ToLower(answer) == "s" {
		fmt.Println("Name: ")
		fmt.Scanf("%s", &fName)
		fmt.Println("Last Name: ")
		fmt.Scanf("%s", &lName)
		fmt.Println("Email: ")
		fmt.Scanf("%s", &Email)
		fmt.Println("Username: ")
		fmt.Scanf("%s", &uName)
		fmt.Println("Password: ")
		fmt.Scanf("%s", &pWord)

		db.CreateUser(fName, lName, Email, uName, pWord)
	} else if strings.ToLower(answer) == "l" {
		fmt.Println("Enter username & password to login: ")
		fmt.Println("Username: ")
		fmt.Scanf("%s", &uName)
		fmt.Println("Password: ")
		fmt.Scanf("%s", &pWord)
		db.SignIn(uName, pWord)
	} else if strings.ToLower(answer) == "d" {
		fmt.Println("Which user do you want to delete? (Type Username & Password for the User you want deleted)")
		fmt.Println("Username: ")
		fmt.Scanf("%s", &uName)
		fmt.Println("Password: ")
		fmt.Scanf("%s", &pWord)
		db.DeleteUser(uName, pWord)
	} else {
		fmt.Println("Make sure you typed s, l, or d", "\n")
	}
}
