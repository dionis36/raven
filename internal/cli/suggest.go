package cli

import (
	"fmt"
	"os"

	"raven/internal/analysis"
	"raven/internal/git"

	"github.com/spf13/cobra"
)

var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "Suggest a commit message for staged changes",
	Run: func(cmd *cobra.Command, args []string) {
		if !git.IsRepository() {
			fmt.Println("Error: This is not a git repository.")
			os.Exit(1)
		}

		diff, err := git.GetStagedDiff()
		if err != nil {
			fmt.Printf("Error getting staged changes: %v\n", err)
			os.Exit(1)
		}

		if diff == "" {
			fmt.Println("No staged changes found.")
			return
		}

		suggestion := analysis.AnalyzeDiff(diff)

		// Format: type(scope): description
		// If scope is empty, omit parens
		msg := ""
		if suggestion.Scope != "" {
			msg = fmt.Sprintf("%s(%s): %s", suggestion.Type, suggestion.Scope, suggestion.Description)
		} else {
			msg = fmt.Sprintf("%s: %s", suggestion.Type, suggestion.Description)
		}

		fmt.Println(msg)
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)
}
