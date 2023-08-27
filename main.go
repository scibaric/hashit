package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	size := len(os.Args)

	if size < 2 {
		fmt.Println("Argument is mandatory")
		printExplanationMessage()
		os.Exit(1)
	}

	if size > 2 {
		fmt.Println("Too many arguments")
		printExplanationMessage()
		os.Exit(2)
	}

	pass := os.Args[1]

	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), 10)

	if err != nil {
		fmt.Println("Hash could not be generated", err.Error())
		os.Exit(3)
	}

	fmt.Println(string(hashed))
}

func printExplanationMessage() {
	fmt.Println("Example: hashit mypassword")
	fmt.Println("Produces: $myhashedpassword")
}
