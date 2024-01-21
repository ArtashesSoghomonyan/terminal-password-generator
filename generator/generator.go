package generator

import (
	"math/rand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const uppercaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const symbols = "!\";#$%&'()*+,-./:;<=>?@[]^_`{|}~"

type Password struct {
	Uppercase bool
	Numbers   bool
	Symbols   bool
	Length    int
}

func GeneratePassword(params Password, numberOfPasswords int) []string {
	var result string
	var passwords []string = []string{}
	var characters string

	for n := 1; n <= numberOfPasswords; n++ {
		result = ""
		characters = alphabet

		if params.Uppercase {
			characters += uppercaseAlphabet
		}

		if params.Numbers {
			characters += numbers
		}

		if params.Symbols {
			characters += symbols
		}

		for i := 1; i <= params.Length; i++ {
			result += string(characters[rand.Intn(len(characters))])
		}

		passwords = append(passwords, result)
	}

	return passwords
}
