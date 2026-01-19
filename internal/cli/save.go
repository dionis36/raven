package cli

import (
	"fmt"
	"os"

	"raven/internal/git"

	"github.com/spf13/cobra"
)

var (
	saveMsgFlag string
)

var saveCmd = &cobra.Command{
	Use:     "save",
	Aliases: []string{"ac", "snap"}, // Common aliases for "Save Point"
	Short:   "Stage all changes and commit (The 'Save Point' command)",
	Long:    "Stages all tracked and untracked changes (git add .) and initiates the commit process.",
	Run: func(cmd *cobra.Command, args []string) {
		if !git.IsRepository() {
			fmt.Println("Error: This is not a git repository.")
			os.Exit(1)
		}

		// 1. Stage All
		if err := git.StageFile("."); err != nil {
			fmt.Println("Error staging files:", err)
			os.Exit(1)
		}

		// 2. Check for staged changes (should be populated now unless directory was clean)
		diff, err := git.GetStagedDiff()
		if err != nil {
			fmt.Printf("Error getting staged changes: %v\n", err)
			os.Exit(1)
		}

		if diff == "" {
			fmt.Println("No changes to save (working tree clean).")
			return
		}

		// 3. Delegate to Shared Commit Logic
		performCommit(diff, saveMsgFlag, "", false)
	},
}

func init() {
	saveCmd.Flags().StringVarP(&saveMsgFlag, "message", "m", "", "Commit message")
	rootCmd.AddCommand(saveCmd)
}
