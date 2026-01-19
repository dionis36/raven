package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "raven",
	Short: "Raven is a smart git commit assistant",
	Long:  `Raven is a CLI tool that analyzes your staged usage and generates conventional commit messages.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default behavior: show help
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	// Override default help
	rootCmd.SetHelpFunc(CustomHelpFunc)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
