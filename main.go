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
	}

	fmt.Println(generator.GeneratePassword(passwordParams))
}
