package cli

import (
	"fmt"
	"os"

	"raven/internal/git"
	"raven/internal/stats"
	"raven/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show a heatmap of git contribution history",
	Run: func(cmd *cobra.Command, args []string) {
		if !git.IsRepository() {
			fmt.Println("Error: This is not a git repository.")
			os.Exit(1)
		}

		counts, err := stats.GetCommitCounts()
		if err != nil {
			fmt.Printf("Error getting commit history: %v\n", err)
			os.Exit(1)
		}

		// Interactive Calendar Heatmap
		p := tea.NewProgram(ui.InitialCalendarModel(counts))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
