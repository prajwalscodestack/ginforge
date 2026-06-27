package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/prajwalscodestack/ginforge/internal/architecture"
	"github.com/prajwalscodestack/ginforge/internal/generator"
	"github.com/prajwalscodestack/ginforge/internal/prompt"
)

var architectureType string

var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new gin project",

	RunE: func(
		cmd *cobra.Command,
		args []string,
	) error {

		var projectName string
		var err error

		if len(args) == 0 {

			input, err :=
				prompt.AskProjectDetails()

			if err != nil {
				return err
			}

			projectName =
				input.Name

			architectureType =
				input.Architecture

		} else {

			projectName =
				args[0]

			if architectureType == "" {

				architectureType,
					err = prompt.AskArchitecture()

				if err != nil {
					return err
				}
			}
		}

		arch, err :=
			architecture.Get(
				architectureType,
			)

		if err != nil {
			return err
		}

		if err := generator.GenerateProject(
			projectName,
			arch,
		); err != nil {
			return err
		}

		fmt.Println()
		fmt.Println("✓ Project created successfully")
		fmt.Println()

		fmt.Printf(
			"ℹ Project: %s\n",
			projectName,
		)

		fmt.Printf(
			"ℹ Architecture: %s\n",
			architectureType,
		)

		fmt.Println()

		fmt.Printf(
			"ℹ Next step: cd %s\n",
			projectName,
		)

		return nil
	},
}

func init() {

	rootCmd.AddCommand(
		newCmd,
	)

	newCmd.Flags().StringVarP(
		&architectureType,
		"architecture",
		"a",
		"",
		"Architecture type (layered|hexagonal)",
	)
}
