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

		passed := 0
		failed := 0

		for _, result := range results {

			if result.Passed {

				passed++

				fmt.Printf(
					"✓ %s\n",
					result.Name,
				)

			} else {

				failed++

				fmt.Printf(
					"✗ %s\n",
					result.Name,
				)
			}

			fmt.Println()

			for _, message := range result.Messages {

				fmt.Printf(
					"  %s\n",
					message,
				)
			}

			fmt.Println()
		}

		fmt.Println("Summary")
		fmt.Println("-------")

		fmt.Printf(
			"Passed: %d\n",
			passed,
		)

		fmt.Printf(
			"Failed: %d\n",
			failed,
		)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
