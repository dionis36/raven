package cli

import (
	"fmt"
	"os"

	"raven/internal/git"
	"raven/internal/ui"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the working tree status",
	Run: func(cmd *cobra.Command, args []string) {
		if !git.IsRepository() {
			fmt.Println("Error: This is not a git repository.")
			os.Exit(1)
		}

		result, err := git.GetStatus()
		if err != nil {
			fmt.Println("Error getting status:", err)
			os.Exit(1)
		}

		// Render Static Status (No interaction necessary)
		// This mimics `git status` which prints and exits.
		model := ui.InitialStatusModel(result, ui.StatusModeView)
		model.Static = true
		fmt.Println(model.View())
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
