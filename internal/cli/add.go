package cli

import (
	"fmt"
	"os"

	"raven/internal/git"
	"raven/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add [.]",
	Aliases: []string{"a"},
	Short:   "Interactively stage files for commit",
	Run: func(cmd *cobra.Command, args []string) {
		// Check for "add ." shortcut
		if len(args) > 0 && args[0] == "." {
			stageAll()
			return
		}
		RunInteractiveAdd()
	},
}

func stageAll() {
	if !git.IsRepository() {
		fmt.Println("Error: This is not a git repository.")
		os.Exit(1)
	}
	// "git add ." stages everything
	if err := git.StageFile("."); err != nil {
		fmt.Println("Error staging all files:", err)
		os.Exit(1)
	}
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#38BDF8")).Bold(true).Render("✔ Staged all changes."))
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

	// Filter is now handled inside ui.InitialStatusModel,
	// but we check if we have anything generally.
	// Actually, let's create the model first to see what's filtered.
	model := ui.InitialStatusModel(result, ui.StatusModeAdd)

	if len(model.Files) == 0 {
		fmt.Println("No unstaged files to stage.")
		return
	}

	// Start UI in Add Mode
	p := tea.NewProgram(model)
	m, err := p.Run()
	if err != nil {
		fmt.Println("Error running UI:", err)
		os.Exit(1)
	}

	finalModel := m.(ui.StatusModel)
	if finalModel.Done && len(finalModel.Selected) > 0 {
		// Stage selected files
		count := 0
		var stagedFiles []string

		for idx, selected := range finalModel.Selected {
			if selected {
				path := finalModel.Files[idx].Path
				if err := git.StageFile(path); err != nil {
					fmt.Printf("Error staging %s: %v\n", path, err)
				} else {
					count++
					stagedFiles = append(stagedFiles, path)
				}
			}
		}

		if count > 0 {
			// Better Feedback
			heading := lipgloss.NewStyle().Foreground(lipgloss.Color("#38BDF8")).Bold(true).Render(fmt.Sprintf("✔ Staged %d files:", count))
			fmt.Println(heading)
			for _, f := range stagedFiles {
				fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("  + " + f))
			}
		}
	} else if finalModel.Done {
		fmt.Println("No files selected.")
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
