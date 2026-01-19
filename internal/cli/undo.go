package cli

import (
	"fmt"
	"os"
	"os/exec"

	"raven/internal/git"

	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo the last commit (keeps changes staged)",
	Long:  "Executes 'git reset --soft HEAD~1', effectively un-committing the last commit while keeping the changes staged.",
	Run: func(cmd *cobra.Command, args []string) {
		if !git.IsRepository() {
			fmt.Println("Error: This is not a git repository.")
			os.Exit(1)
		}

		// Run git reset --soft HEAD~1
		c := exec.Command("git", "reset", "--soft", "HEAD~1")
		// We capture stderr in case of error (e.g., no commits yet)
		output, err := c.CombinedOutput()
		if err != nil {
			fmt.Printf("Error undoing commit (maybe no commits exists?):\n%s\n", string(output))
			os.Exit(1)
		}

		fmt.Println("✔ Undid last commit. Changes are now staged.")
	},
}

func init() {
	rootCmd.AddCommand(undoCmd)
}
