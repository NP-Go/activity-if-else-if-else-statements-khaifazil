package main

import "fmt"

func main() {
	//Insert Code here
	var userName string
	fmt.Println("Enter your Username: ")
	fmt.Scanln(&userName)

	if userName == "Admin" {
		fmt.Println("Welcome, Admin!")
	}else if userName == "Robin" || userName == "John" {
		fmt.Println("Welcome!")
	}else {
		fmt.Println("Merry Men")
	}
}
