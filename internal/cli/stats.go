package cli

import (
	"fmt"
	"os"

	"raven/internal/git"
	"raven/internal/stats"
	"raven/internal/ui"

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

		heatmap := ui.RenderHeatmap(counts)
		fmt.Println(heatmap)
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
