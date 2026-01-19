package cli

import (
	"fmt"
	"os"

	"raven/internal/git"
	"raven/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Interactively stage files for commit",
	Run: func(cmd *cobra.Command, args []string) {
		RunInteractiveAdd()
	},
}

// RunInteractiveAdd is exposed so `commit` can call it too.
func RunInteractiveAdd() {
	if !git.IsRepository() {
		fmt.Println("Error: This is not a git repository.")
		os.Exit(1)
	}

	result, err := git.GetStatus()
	if err != nil {
		fmt.Println("Error getting status:", err)
		os.Exit(1)
	}

	// Filter out fully staged files? Or allow unstaging?
	// For simplicity, just show all changed files.
	// Maybe simple "add" just stages.

	if len(result.Files) == 0 {
		fmt.Println("No changed files to stage.")
		return
	}

	// Start UI in Add Mode
	p := tea.NewProgram(ui.InitialStatusModel(result, ui.StatusModeAdd))
	m, err := p.Run()
	if err != nil {
		fmt.Println("Error running UI:", err)
		os.Exit(1)
	}

	finalModel := m.(ui.StatusModel)
	if finalModel.Done && len(finalModel.Selected) > 0 {
		// Stage selected files
		count := 0
		for idx, selected := range finalModel.Selected {
			if selected {
				path := finalModel.Files[idx].Path
				if err := git.StageFile(path); err != nil {
					fmt.Printf("Error staging %s: %v\n", path, err)
				} else {
					count++
				}
			}
		}
		if count > 0 {
			fmt.Printf("Staged %d files.\n", count)
		}
	} else if finalModel.Done {
		fmt.Println("No files selected.")
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
