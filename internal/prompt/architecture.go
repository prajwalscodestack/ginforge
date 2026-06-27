package prompt

import "github.com/AlecAivazis/survey/v2"

func AskArchitecture() (
	string,
	error,
) {

	var choice string

	prompt := &survey.Select{
		Message: "Select architecture:",
		Options: []string{
			"Layered Architecture",
			"Hexagonal Architecture",
		},
		Description: func(
			value string,
			index int,
		) string {

			switch value {

			case "Layered Architecture":
				return "Traditional REST architecture (Handler → Service → Repository)"

			case "Hexagonal Architecture":
				return "Domain → Application → Adapters"
			}

			return ""
		},
	}

	if err := survey.AskOne(
		prompt,
		&choice,
	); err != nil {
		return "", err
	}

	switch choice {

	case "Layered Architecture":
		return "layered", nil

	case "Hexagonal Architecture":
		return "hexagonal", nil
	}

	return "layered", nil
}
