package cmd

import (
	"ginforge/internal/architecture"
	"ginforge/internal/generator"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use: "generate",
}

var moduleCmd = &cobra.Command{
	Use:  "module [name]",
	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {

		moduleName := args[0]

		arch, err := architecture.Get(
			architectureType,
		)
		if err != nil {
			return err
		}

		return generator.GenerateModule(
			".",
			moduleName,
			arch,
		)
	},
}

func init() {

	rootCmd.AddCommand(generateCmd)

	generateCmd.AddCommand(moduleCmd)
}
