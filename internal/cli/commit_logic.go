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
// performCommit handles the analysis, TUI, and final execution of a commit.
// If manualMessage is provided, it skips analysis/TUI and commits directly.
// If amend is true, it uses `git commit --amend`.
func performCommit(diff string, manualMessage string, overrideMsg string, amend bool) {
	var finalMsg string

	if manualMessage != "" {
		// MANUAL MODE: Use the provided flag
		finalMsg = manualMessage
	} else {
		// AI MODE: Analyze and Suggest
		// For amend, we might want to default to the existing message if not analyzing?
		// But existing logic assumes we pass a suggestion.
		// If amend is true and we found no diff (meaning we are just rewording), analysis might fail or return nothing.
		// Let's rely on the caller to pass the "message" if they want to pre-fill it.
		// Actually, `amend` implies we start with the old message usually.

		// If analysis is skipped (e.g. amend with no diff change), we just start with empty or provided msg.
		// But wait, `performCommit` is currently designed for "Analysis First".

		// Let's stick to the plan: `amend` will reuse this TUI.
		// If `amend` is passed, the caller (amendCmd) effectively decides the "starting message".
		// But `performCommit` calculates `msg` from `analysis.AnalyzeDiff`.

		// We should change `performCommit` to accept an `initialMsg` instead of just `diff`.
		// But `commit` relies on `diff` to generate `initialMsg`.

		// Let's update signature to:
		// performCommit(initialMsg string, diff string, manualMessage string, amend bool)
		// If `initialMsg` is empty, run Analysis on `diff`.

		// Wait, `amend` command needs to fetch the old message.
		// Let's keep it simple: The caller does the fetching/analysis.
		// `performCommit` just takes the PRE-CALCULATED suggestion/starting-point.

		// Refactor: performCommit(suggestedMsg string, manualMessage string, amend bool)
		// Caller (commit/save) calls AnalyzeDiff first.
		// Caller (amend) calls "get last commit msg" first.
	}

	// WAIT. `commit` and `save` share the exact Analysis Logic.
	// If I move Analysis out, I duplicate it in `commit.go` and `save.go`.

	// Let's overload `performCommit`.
	// func performCommit(diff string, manualMessage string, amend bool)
	// If amend=true, `diff` might be irrelevant for suggestion if we want to overwrite it with old message?
	// Actually, `amend` usually combines old message + new changes.
	// The implementation plan said: "Open TUI, pre-filled with the **last commit message**".

	// So `performCommit` needs to optionally accept an override for the "Suggestion".
	// Let's add `overrideMsg string`.

	// func performCommit(diff string, manualMessage string, overrideMsg string, amend bool)

	msg := overrideMsg
	if msg == "" && manualMessage == "" {
		// AI MODE: Analyze
		suggestion := analysis.AnalyzeDiff(diff)
		if suggestion.Scope != "" {
			msg = fmt.Sprintf("%s(%s): %s", suggestion.Type, suggestion.Scope, suggestion.Description)
		} else {
			msg = fmt.Sprintf("%s: %s", suggestion.Type, suggestion.Description)
		}
	} else if manualMessage != "" {
		msg = manualMessage
	}

	// Interactive UI (skips if manualMessage was set, wait logic below...)
	if manualMessage == "" {
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
	} else {
		finalMsg = msg
	}

	// Execute Commit
	args := []string{"commit", "-m", finalMsg}
	if amend {
		args = append(args, "--amend")
	}

	c := exec.Command("git", args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		fmt.Println("Error committing:", err)
	} else {
		if amend {
			fmt.Println("Commit amended successfully! 🚀")
		} else {
			fmt.Println("Commit successful! 🚀")
		}
	}
}
