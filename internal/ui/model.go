package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Choice represents the user's selection.
type Choice int

const (
	ChoiceNone Choice = iota
	ChoiceApply
	ChoiceEdit
	ChoiceCancel
)

// Model represents the state of the UI.
type Model struct {
	Message  string
	Choice   Choice
	Quitting bool
	cursor   int
	choices  []string
}

// InitialModel returns the initial model state with the suggested message.
func InitialModel(msg string) Model {
	return Model{
		Message: msg,
		Choice:  ChoiceNone,
		choices: []string{"Apply", "Edit", "Cancel"},
		cursor:  0,
	}
}

// Init initializes the IO.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.Choice = ChoiceCancel
			m.Quitting = true
			return m, tea.Quit

		case "up", "k", "left", "h":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.choices) - 1
			}

		case "down", "j", "right", "l", "tab":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}

		case "enter", " ":
			// Map cursor to Choice
			switch m.cursor {
			case 0:
				m.Choice = ChoiceApply
			case 1:
				m.Choice = ChoiceEdit
			case 2:
				m.Choice = ChoiceCancel
			}
			m.Quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

// View pushes the string representation of the UI.
func (m Model) View() string {
	if m.Quitting {
		return ""
	}

	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("212")).
		Render("Raven 🐦 Suggestion:")

	msgBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2).
		MarginBottom(1).
		Render(m.Message)

	// Render choices
	s := ""
	for i, choice := range m.choices {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		style := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
		if m.cursor == i {
			style = lipgloss.NewStyle().
				Foreground(lipgloss.Color("86")).
				Bold(true)
		}

		s += fmt.Sprintf("%s %s  ", cursor, style.Render(choice))
	}

	return fmt.Sprintf("\n%s\n%s\n\n%s\n\n(arrows/tab to move, enter to select)\n", header, msgBox, s)
}
