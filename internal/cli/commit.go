package cli

import (
	"fmt"
	"os"
	"os/exec"

	"raven/internal/analysis"
	"raven/internal/git"
	"raven/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Interactively generate and apply a commit message",
	Run: func(cmd *cobra.Command, args []string) {
		if !git.IsRepository() {
			fmt.Println("Error: This is not a git repository.")
			os.Exit(1)
		}

		// 1. Check for staged changes
		diff, err := git.GetStagedDiff()
		if err != nil {
			fmt.Printf("Error getting staged changes: %v\n", err)
			os.Exit(1)
		}

		// 2. AUTO-STAGE LOGIC
		if diff == "" {
			fmt.Println("ℹ️  No staged changes found.")
			fmt.Println("Launching interactive staging... (Select files with Space, Enter to Confirm)")

			// Call the ADD flow
			RunInteractiveAdd()

			// Re-check diff after staging
			diff, err = git.GetStagedDiff()
			if err != nil {
				fmt.Printf("Error getting staged changes: %v\n", err)
				os.Exit(1)
			}

			// If still empty, user cancelled staging
			if diff == "" {
				fmt.Println("❌ Nothing staged. Commit aborted.")
				return
			}
		}

		// 3. Proceed with Analysis & Commit
		// Analysis
		suggestion := analysis.AnalyzeDiff(diff)
		msg := ""
		if suggestion.Scope != "" {
			msg = fmt.Sprintf("%s(%s): %s", suggestion.Type, suggestion.Scope, suggestion.Description)
		} else {
			msg = fmt.Sprintf("%s: %s", suggestion.Type, suggestion.Description)
		}

		// Interactive UI
		p := tea.NewProgram(ui.InitialModel(msg))
		m, err := p.Run()
		if err != nil {
			fmt.Println("Error running UI:", err)
			os.Exit(1)
		}

		// Handle Result
		finalModel := m.(ui.Model)
		switch finalModel.Choice {
		case ui.ChoiceApply:
			// git commit -m "msg"
			c := exec.Command("git", "commit", "-m", finalModel.Message)
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			if err := c.Run(); err != nil {
				fmt.Println("Error committing:", err)
			} else {
				fmt.Println("Commit successful! 🚀")
			}

		case ui.ChoiceEdit:
			// git commit -m "msg" --edit
			c := exec.Command("git", "commit", "-m", finalModel.Message, "--edit")
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			if err := c.Run(); err != nil {
				fmt.Println("Error committing:", err)
			} else {
				fmt.Println("Commit successful! 🚀")
			}

		case ui.ChoiceCancel:
			fmt.Println("Commit canceled.")
		case ui.ChoiceNone:
			fmt.Println("Commit canceled.")
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
