package cmd

import (
	"fmt"

	routescanner "github.com/prajwalscodestack/ginforge/internal/scanner/routes"

	"github.com/spf13/cobra"
)

var (
	jsonOutput bool
	csvOutput  bool
	mdOutput   bool

	methodFilter string
	pathFilter   string
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

		// -------------------------
		// FILTER (non-breaking)
		// -------------------------
		if methodFilter != "" || pathFilter != "" {

			filtered := make([]routescanner.Route, 0)

			for _, r := range routes {

				if methodFilter != "" && r.Method != methodFilter {
					continue
				}

				if pathFilter != "" && !contains(r.Path, pathFilter) {
					continue
				}

				filtered = append(filtered, r)
			}

			routes = filtered
		}

		// -------------------------
		// OUTPUT MODES (unchanged)
		// -------------------------
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
		}

		// -------------------------
		// ENHANCEMENT: SUMMARY (only table mode)
		// -------------------------
		fmt.Println()
		fmt.Printf("Routes Found: %d\n\n", len(routes))

		methodCount := map[string]int{}

		for _, r := range routes {
			methodCount[r.Method]++
		}

		fmt.Println("Summary")
		fmt.Println("-------")

		for method, count := range methodCount {
			fmt.Printf("%-6s %d\n", method, count)
		}

		return nil
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

	// -------------------------
	// NEW OPTIONAL FILTER FLAGS
	// -------------------------
	routesCmd.Flags().StringVar(
		&methodFilter,
		"method",
		"",
		"Filter by HTTP method",
	)

	routesCmd.Flags().
		StringVar(
			&pathFilter,
			"path",
			"",
			"Filter routes by path substring",
		)

	rootCmd.AddCommand(routesCmd)
}

// helper (safe, local)
func contains(str, substr string) bool {
	return len(substr) == 0 || (len(str) > 0 && stringContains(str, substr))
}

// simple wrapper (avoids importing strings if you want minimal change)
func stringContains(s, sub string) bool {
	return len(s) >= len(sub) && (len(sub) == 0 || indexOf(s, sub) >= 0)
}

// naive index function replacement
func indexOf(s, sub string) int {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
