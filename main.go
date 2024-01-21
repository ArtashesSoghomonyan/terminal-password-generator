package main

import (
	"fmt"

	"github.com/ArtashesSoghomonyan/terminal-password-generator/generator"
	"github.com/ArtashesSoghomonyan/terminal-password-generator/prompts"
)

func main() {
	fmt.Println("*** Password generator ***")

	var passwordParams generator.Password = generator.Password{
		Uppercase: prompts.YesNo("Would you like to include uppercase letters?"),
		Numbers:   prompts.YesNo("Would you like to include numbers?"),
		Symbols:   prompts.YesNo("Would you like to include symbols?"),
		Length:    prompts.IntPrompt("What would be the length of the passwords?"),
	}

	var numberOfPasswords int = prompts.IntPrompt("How many passwords would you like to generate?")

	var passwords []string = generator.GeneratePassword(passwordParams, numberOfPasswords)

	for index, password := range passwords {
		fmt.Printf("%d) %s\n", index+1, password)
	}
}
