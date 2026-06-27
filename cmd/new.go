package cmd

import (
	"github.com/spf13/cobra"

	"ginforge/internal/architecture"
	"ginforge/internal/generator"
)

var architectureType string

var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new gin project",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {

		projectName := args[0]

		arch, err := architecture.Get(architectureType)
		if err != nil {
			return err
		}

		return generator.GenerateProject(
			projectName,
			arch,
		)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().
		StringVarP(
			&architectureType,
			"architecture",
			"a",
			"layered",
			"Architecture type",
		)
}
