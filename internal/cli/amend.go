package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"raven/internal/git"

	"github.com/spf13/cobra"
)

var amendCmd = &cobra.Command{
	Use:   "amend",
	Short: "Amend the last commit message and/or staged files",
	Long:  "Opens the last commit message for editing. If you have staged new files, they will be combined into the last commit.",
	Run: func(cmd *cobra.Command, args []string) {
		if !git.IsRepository() {
			fmt.Println("Error: This is not a git repository.")
			os.Exit(1)
		}

		// 1. Get Last Commit Message
		c := exec.Command("git", "log", "-1", "--pretty=%B")
		output, err := c.Output()
		if err != nil {
			fmt.Println("Error reading last commit:", err)
			os.Exit(1)
		}
		lastMsg := strings.TrimSpace(string(output))

		// 2. Check for staged changes (Optional warning if clean?)
		// Actually, git commit --amend works fine even with no changes (just rewords).

		// 3. Launch TUI in Amend Mode
		// We pass 'lastMsg' as the override.
		// We pass 'true' for amend flag.
		// We pass empty diff (not needed since we provide override).
		// We pass empty manualMessage (unless we want to support -m here too? Nah, interactive default).

		performCommit("", "", lastMsg, true)
	},
}

func init() {
	rootCmd.AddCommand(amendCmd)
}
