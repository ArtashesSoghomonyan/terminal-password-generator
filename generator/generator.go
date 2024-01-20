package generator

import "math/rand"

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const uppercaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const symbols = "!\";#$%&'()*+,-./:;<=>?@[]^_`{|}~"

type Password struct {
	Uppercase bool
	Numbers   bool
	Symbols   bool
}

func GeneratePassword(params Password) string {
	var passwordLength int = 6
	var result string = ""
	var characters string = alphabet

	if params.Uppercase {
		characters += uppercaseAlphabet
	}

	if params.Numbers {
		characters += numbers
	}

	if params.Symbols {
		characters += symbols
	}

	for i := 1; i <= passwordLength; i++ {
		result += string(characters[rand.Intn(len(characters))])
	}

	return result
}
