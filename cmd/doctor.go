package cmd

import (
	"fmt"

	"github.com/prajwalscodestack/ginforge/internal/doctor"

	"github.com/spf13/cobra"
)

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

		hasFailures := false

		for _, result := range results {

			if result.Passed {

				fmt.Printf(
					"✓ %s\n",
					result.Name,
				)

			} else {

				hasFailures = true

				fmt.Printf(
					"✗ %s\n",
					result.Name,
				)
			}

			for _, message := range result.Messages {

				fmt.Printf(
					"  %s\n",
					message,
				)
			}

			fmt.Println()
		}

		if hasFailures {

			fmt.Println(
				"Doctor found issues.",
			)

			return nil
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
