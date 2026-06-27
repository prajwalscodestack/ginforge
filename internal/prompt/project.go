package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

type ProjectInput struct {
	Name         string
	Architecture string
}

func AskProjectDetails() (
	ProjectInput,
	error,
) {

	var input ProjectInput

	namePrompt := &survey.Input{
		Message: "Project name:",
	}

	if err := survey.AskOne(
		namePrompt,
		&input.Name,
		survey.WithValidator(
			survey.Required,
		),
	); err != nil {
		return input, err
	}

	architecture, err :=
		AskArchitecture()

	if err != nil {
		return input, err
	}

	input.Architecture =
		architecture

	var confirm bool

	confirmPrompt := &survey.Confirm{
		Message: fmt.Sprintf(
			"Create '%s' using '%s' architecture?",
			input.Name,
			input.Architecture,
		),
		Default: true,
	}

	if err := survey.AskOne(
		confirmPrompt,
		&confirm,
	); err != nil {
		return input, err
	}

	if !confirm {
		return input, fmt.Errorf(
			"project creation cancelled",
		)
	}

	return input, nil
}
