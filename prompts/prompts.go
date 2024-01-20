package prompts

import (
	"fmt"
	"log"

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
