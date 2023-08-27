package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	size := len(os.Args)

	if size < 2 {
		fmt.Println("[ERROR] Argument is mandatory")
		os.Exit(1)
	}

	pass := os.Args[1]

	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), 10)

	if err != nil {
		fmt.Println("[ERROR] Hash could not be generated", err.Error())
		os.Exit(2)
	}

	fmt.Println(string(hashed))
}
