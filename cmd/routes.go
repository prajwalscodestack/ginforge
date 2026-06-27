package cmd

import (
	routescanner "github.com/prajwalscodestack/ginforge/internal/scanner/routes"

	"github.com/spf13/cobra"
)

var (
	jsonOutput bool
	csvOutput  bool
	mdOutput   bool
)

var routesCmd = &cobra.Command{
	Use:   "routes",
	Short: "List all Gin routes",

	RunE: func(
		cmd *cobra.Command,
		args []string,
	) error {

		routes, err := routescanner.Scan(".")
		if err != nil {
			return err
		}

		switch {

		case jsonOutput:
			return routescanner.PrintJSON(routes)

		case csvOutput:
			return routescanner.PrintCSV(routes)

		case mdOutput:
			routescanner.PrintMarkdown(routes)
			return nil

		default:
			routescanner.PrintTable(routes)
			return nil
		}
	},
}

func init() {

	routesCmd.Flags().BoolVar(
		&jsonOutput,
		"json",
		false,
		"Output routes as JSON",
	)

	routesCmd.Flags().BoolVar(
		&csvOutput,
		"csv",
		false,
		"Output routes as CSV",
	)

	routesCmd.Flags().BoolVar(
		&mdOutput,
		"md",
		false,
		"Output routes as Markdown",
	)

	rootCmd.AddCommand(routesCmd)
}
