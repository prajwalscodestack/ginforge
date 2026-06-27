package cmd

import (
	"fmt"

	routescanner "github.com/prajwalscodestack/ginforge/internal/scanner/routes"

	"github.com/spf13/cobra"
)

var routesCmd = &cobra.Command{
	Use:   "routes",
	Short: "List all Gin routes",

	RunE: func(cmd *cobra.Command, args []string) error {

		routes, err := routescanner.Scan(".")
		if err != nil {
			return err
		}

		fmt.Printf("%-8s %s\n", "METHOD", "PATH")

		for _, route := range routes {

			fmt.Printf(
				"%-8s %s\n",
				route.Method,
				route.Path,
			)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(routesCmd)
}
