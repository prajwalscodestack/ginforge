package cmd

import (
	"fmt"
	"os"

	"github.com/prajwalscodestack/ginforge/internal/version"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without subcommands
var rootCmd = &cobra.Command{
	Use:   "ginforge",
	Short: "Architecture-aware CLI toolkit for Gin applications",
	Long: `GinForge is a developer-focused CLI tool for building, analyzing,
and maintaining production-ready Gin applications.

It provides:
- Project scaffolding (layered & hexagonal architecture)
- Module generation
- Route discovery using Go AST
- Architecture validation and project health checks

Example usage:

  ginforge new myapp --architecture layered
  ginforge generate module user
  ginforge routes
  ginforge doctor --strict
`,
}

// Execute runs the root command
func Execute() {

	// handle global --version flag
	if v, _ := rootCmd.Flags().GetBool("version"); v {
		fmt.Println("GinForge CLI")
		fmt.Println("Version:", version.Version)
		return
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
func init() {
	rootCmd.PersistentFlags().
		BoolP("version", "", false, "Show GinForge version")
	// Remove useless default toggle flag (Cobra boilerplate cleanup)
}
