package cmd

import (
	"fmt"

	"github.com/prajwalscodestack/ginforge/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show GinForge version information",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("GinForge CLI")
		fmt.Println("--------------------")
		fmt.Println("Version   :", version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
