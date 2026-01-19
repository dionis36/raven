package ui

import (
	"raven/internal/git"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type StatusMode int

const (
	StatusModeView StatusMode = iota
	StatusModeAdd
)

type StatusModel struct {
	BranchInfo string
	Files      []git.FileStatus
	Cursor     int
	Selected   map[int]bool // Indices selected for staging
	Mode       StatusMode
	Quitting   bool
	Done       bool // User hit Enter
}

func InitialStatusModel(data git.StatusResult, mode StatusMode) StatusModel {
	return StatusModel{
		BranchInfo: data.BranchInfo,
		Files:      data.Files,
		Selected:   make(map[int]bool),
		Mode:       mode,
	}
}

func (m StatusModel) Init() tea.Cmd {
	return nil
}

func (m StatusModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			m.Quitting = true
			return m, tea.Quit

		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down", "j":
			if m.Cursor < len(m.Files)-1 {
				m.Cursor++
			}

		case " ": // Space to toggle select
			if m.Mode == StatusModeAdd {
				m.Selected[m.Cursor] = !m.Selected[m.Cursor]
			}

		case "enter":
			// If in Add mode, finalizing selection
			if m.Mode == StatusModeAdd {
				m.Done = true
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m StatusModel) View() string {
	if m.Quitting {
		return ""
	}
	if m.Done {
		return ""
	}

	var s strings.Builder

	// 1. Branch Header
	// Parse e.g., "main...origin/main [ahead 11]"
	// Make it pretty: "On branch main ⬆️ 11"
	branchStr := m.BranchInfo
	branchColor := lipgloss.NewStyle().Foreground(lipgloss.Color("#38BDF8")).Bold(true) // Sky Blue

	if m.BranchInfo != "" {
		s.WriteString(branchColor.Render("On branch "+branchStr) + "\n\n")
	}

	// 2. Separate Files into Groups
	var tracked []int
	var untracked []int
	for i, f := range m.Files {
		if f.Untracked {
			untracked = append(untracked, i)
		} else {
			tracked = append(tracked, i)
		}
	}

	// Helper to render lists
	renderList := func(title string, indices []int) {
		if len(indices) == 0 {
			return
		}

		// Section Title
		s.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Bold(true).Render(title) + "\n")

		for _, i := range indices {
			file := m.Files[i]
			cursor := "  "
			if m.Cursor == i {
				cursor = "> "
			}

			// Interactive Checkbox
			prefix := ""
			if m.Mode == StatusModeAdd {
				if m.Selected[i] {
					prefix = "[x] "
				} else {
					prefix = "[ ] "
				}
			} else {
				// View Mode Icons
				if file.Untracked {
					prefix = " ? "
				} else if file.Staged {
					prefix = " + "
				} else {
					prefix = " M "
				}
			}

			// Styling
			style := lipgloss.NewStyle()

			// Color Logic
			if file.Untracked {
				style = style.Foreground(lipgloss.Color("#9CA3AF")) // Grey (Untracked)
			} else if file.Staged {
				style = style.Foreground(lipgloss.Color("#38BDF8")) // Sky Blue (Staged)
			} else {
				style = style.Foreground(lipgloss.Color("#F472B6")) // Pink (Modified)
			}

			// Cursor Highlight
			if m.Cursor == i {
				style = style.Bold(true).Underline(true)
			}

			s.WriteString(cursor + style.Render(prefix+file.Path) + "\n")
		}
		s.WriteString("\n") // Gap between sections
	}

	// Render Tracked
	if len(tracked) > 0 {
		renderList("Changes to be committed / Modified:", tracked)
	}

	// Render Untracked
	if len(untracked) > 0 {
		renderList("Untracked files:", untracked)
	}

	if len(tracked) == 0 && len(untracked) == 0 {
		s.WriteString("Working tree clean.\n")
	}

	// Footer Help
	helpMsg := ""
	if m.Mode == StatusModeView {
		helpMsg = "(Use 'raven add' to stage changes • q to quit)"
	} else {
		helpMsg = "(Space to toggle • Enter to stage • q to quit)"
	}
	s.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("240")).MarginTop(1).Render(helpMsg))

	return s.String()
}
