package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	login()
	updatepassword()
	forgetpassword()
}

func login() {
	fmt.Println("Login")
}

func updatepassword() {
	fmt.Println("Update Password")
}

func forgetpassword() {
	fmt.Println("forget password")
}