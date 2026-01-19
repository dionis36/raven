package cli

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#F25D94")).
			MarginBottom(1)

	subHeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#38BDF8")).
			MarginTop(1).
			MarginBottom(0)

	commandStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF7DB")).
			Width(20)

	descStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("244")) // Grey

	aliasStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")). // Darker Grey
			Italic(true)
)

// CustomHelpFunc renders the help output using Lipgloss
func CustomHelpFunc(cmd *cobra.Command, args []string) {
	// 1. Header (No Emoji)
	fmt.Println(headerStyle.Render("RAVEN - Smart Git Assistant"))
	fmt.Println(descStyle.Render(cmd.Long))
	fmt.Println()

	// 2. Usage
	fmt.Println(subHeaderStyle.Render("USAGE"))
	fmt.Printf("  %s\n", cmd.UseLine())
	fmt.Println(descStyle.Render("  Use 'raven [command] --help' for more info."))

	// 3. Commands Grouping
	workflowCmds := []string{"status", "add", "commit", "save", "undo", "fix", "amend"}
	insightCmds := []string{"stats"}
	systemCmds := []string{"help", "suggest", "completion"}

	renderGroup := func(title string, cmdNames []string) {
		fmt.Println(subHeaderStyle.Render(title))
		for _, name := range cmdNames {
			for _, c := range cmd.Commands() {
				if c.Name() == name {
					// Format aliases explicitly
					aliases := ""
					if len(c.Aliases) > 0 {
						// e.g. "[alias: s]"
						aliases = fmt.Sprintf("[alias: %s]", strings.Join(c.Aliases, ", "))
					}

					// Render
					cmdStr := commandStyle.Render(c.Name())
					aliasStr := aliasStyle.Render(fmt.Sprintf("%-12s", aliases)) // Increased width
					descStr := descStyle.Render(c.Short)

					fmt.Printf("  %s %s %s\n", cmdStr, aliasStr, descStr)
					break
				}
			}
		}
	}

	renderGroup("WORKFLOW", workflowCmds)
	renderGroup("INSIGHTS", insightCmds)
	renderGroup("SYSTEM", systemCmds)

	// 4. Flags (if any)
	if cmd.HasAvailableFlags() {
		fmt.Println(subHeaderStyle.Render("FLAGS"))
		fmt.Println(cmd.Flags().FlagUsages())
	}
	fmt.Println()
}
