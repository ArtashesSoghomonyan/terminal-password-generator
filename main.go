package main

import (
	"fmt"
	"math/rand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func main() {
	fmt.Println(generatePassword())
}

func generatePassword() string {
	var passwordLength int = 6
	var result string = ""
	var characters string = alphabet

	for i := 1; i <= passwordLength; i++ {
		result += string(characters[rand.Intn(len(characters))])
	}

	return result
}
