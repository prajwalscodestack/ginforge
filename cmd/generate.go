package cmd

import (
	"github.com/prajwalscodestack/ginforge/internal/architecture"
	"github.com/prajwalscodestack/ginforge/internal/generator"
	"github.com/prajwalscodestack/ginforge/internal/prompt"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code components",
}

var moduleCmd = &cobra.Command{
	Use:   "module [name]",
	Short: "Generate a module",
	Args:  cobra.MaximumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {

		var moduleName string
		var err error

		// ----------------------------
		// INTERACTIVE MODE
		// ----------------------------
		if len(args) == 0 {

			input, err := prompt.AskModuleName()
			if err != nil {
				return err
			}

			moduleName = input.Name

		} else {
			moduleName = args[0]
		}

		// ----------------------------
		// ARCHITECTURE DETECTION
		// ----------------------------
		arch, err := architecture.ResolveFromProject(".")
		if err != nil {
			return err
		}

		// ----------------------------
		// GENERATE MODULE
		// ----------------------------
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
