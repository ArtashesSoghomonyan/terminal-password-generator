package prompts

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/manifoldco/promptui"
)

func YesNo(input string) bool {
	prompt := promptui.Select{
		Label: fmt.Sprintf("%s [Yes/No]", input),
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}

func IntPrompt(input string) int {
	validate := func(input string) error {
		_, err := strconv.ParseInt(input, 10, 0)
		if err != nil {
			return errors.New("invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    input,
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		panic(fmt.Sprintf("Prompt failed %v\n", err))
	}

	returnResult, convErr := strconv.ParseInt(result, 10, 0)

	if convErr != nil {
		panic(convErr)
	} else {
		return int(returnResult)
	}
}
