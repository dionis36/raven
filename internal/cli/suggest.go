package cli

import (
	"fmt"
	"os"

	"raven/internal/analysis"
	"raven/internal/git"

	"github.com/charmbracelet/lipgloss"
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
			// Check if we have unstaged files
			status, err := git.GetStatus()
			if err == nil && len(status.Files) > 0 {
				fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render("ℹ️  No staged changes found."))
				fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("💡 Tip: Use 'raven commit' to automatically stage, analyze, and commit."))
				fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("   Or run 'raven add' to stage files manually."))
			} else {
				fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#38BDF8")).Bold(true).Render("✨ Working tree clean. Nothing to commit."))
			}
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

		// Rich UI Output
		headerStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true).
			Padding(0, 1).
			Render("✨ Raven Suggestion")

		msgStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(1, 4).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("205")).
			MarginTop(1)

		fmt.Println(headerStyle)
		fmt.Println(msgStyle.Render(msg))
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)
}
