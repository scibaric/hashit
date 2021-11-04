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

	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), 15)

	if err != nil {
		fmt.Errorf("[ERROR] Hash could not be generated", err)
		return
	}

	fmt.Println(string(hashed))
}
