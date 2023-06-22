package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Dmkk01/kanban-go-svelte/cli/emoji"
)

var mainqs = []*survey.Question{
	{
		Name: "function",
		Prompt: &survey.Select{
			Message: "What do you want to do?",
			Options: []string{"get-emojis", "list-emojis"},
			Default: "get-emojis",
		},
	},
}

func main() {
	fmt.Println("Welcome to the Kanban CLI tool")

	answers := struct {
		Function string
	}{}

	err := survey.Ask(mainqs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	switch answers.Function {
	case "get-emojis":
		emoji.GetEmojis()
	case "list-emojis":
		emoji.ListEmojis()
	default:
		fmt.Printf("invalid option")
	}
}
