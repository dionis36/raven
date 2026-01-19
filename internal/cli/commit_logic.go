package cli

import (
	"fmt"
	"os"
	"os/exec"

	"raven/internal/analysis"
	"raven/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

// performCommit handles the analysis, TUI, and final execution of a commit.
// If manualMessage is provided, it skips analysis/TUI and commits directly.
func performCommit(diff string, manualMessage string) {
	var finalMsg string

	if manualMessage != "" {
		// MANUAL MODE: Use the provided flag
		finalMsg = manualMessage
	} else {
		// AI MODE: Analyze and Suggest
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

		finalModel := m.(ui.Model)
		if finalModel.Choice == ui.ChoiceCancel {
			fmt.Println("Commit canceled.")
			return
		}
		finalMsg = finalModel.Message
	}

	// Execute Commit
	c := exec.Command("git", "commit", "-m", finalMsg)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		fmt.Println("Error committing:", err)
	} else {
		fmt.Println("Commit successful! 🚀")
	}
}
