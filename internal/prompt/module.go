package prompt

import (
	"github.com/AlecAivazis/survey/v2"
)

type ModuleInput struct {
	Name string
}

func AskModuleName() (ModuleInput, error) {

	var input ModuleInput

	prompt := &survey.Input{
		Message: "Module name:",
	}

	err := survey.AskOne(
		prompt,
		&input.Name,
		survey.WithValidator(survey.Required),
	)

	return input, err
}
