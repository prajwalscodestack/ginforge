package cmd

import (
	"fmt"
	"os"

	"github.com/prajwalscodestack/ginforge/internal/doctor"

	"github.com/spf13/cobra"
)

var strict bool

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Analyze and validate a Gin project",

	RunE: func(
		cmd *cobra.Command,
		args []string,
	) error {

		results := doctor.Run(".")

		fmt.Println("GinForge Doctor")
		fmt.Println()

		passed := 0
		failed := 0

		for _, result := range results {

			if result.Passed {

				passed++

				fmt.Printf("✓ %s\n", result.Name)

			} else {

				failed++

				fmt.Printf("✗ %s\n", result.Name)
			}

			fmt.Println()

			for _, message := range result.Messages {
				fmt.Printf("  %s\n", message)
			}

			fmt.Println()
		}

		fmt.Println("Summary")
		fmt.Println("-------")

		fmt.Printf("Passed: %d\n", passed)
		fmt.Printf("Failed: %d\n", failed)

		// -------------------------
		// STRICT MODE (CI SUPPORT)
		// -------------------------
		if strict && failed > 0 {
			fmt.Println()
			fmt.Println("❌ Doctor failed (strict mode enabled)")
			os.Exit(1)
		}

		return nil
	},
}

func init() {

	doctorCmd.Flags().
		BoolVar(
			&strict,
			"strict",
			false,
			"Exit with error code if issues found",
		)

	rootCmd.AddCommand(doctorCmd)
}
