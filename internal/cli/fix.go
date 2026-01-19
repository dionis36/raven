package cli

import (
	"fmt"
	"os"
	"os/exec"

	"raven/internal/git"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:     "fix",
	Aliases: []string{"f"},
	Short:   "Quickly fix the last commit (stage all & amend silent)",
	Long:    "Stages all tracked/untracked changes and folds them into the last commit without changing the message.",
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

		// 1.5. Safety Check / Confirmation
		fmt.Printf("This will stage ALL changes and amend the last commit (%s).\n", lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render("no-edit"))
		fmt.Print("Continue? [y/N]: ")

		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" {
			fmt.Println("Aborted.")
			return
		}

		// 2. Commit Amend No-Edit
		c := exec.Command("git", "commit", "--amend", "--no-edit")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			fmt.Println("Error fixing commit:", err)
			os.Exit(1)
		}

		// Success Message
		fmt.Println(lipgloss.NewStyle().
			Foreground(lipgloss.Color("#38BDF8")).
			Bold(true).
			Render("✔ Patched last commit successfully."))
	},
}

func init() {
	rootCmd.AddCommand(fixCmd)
}
