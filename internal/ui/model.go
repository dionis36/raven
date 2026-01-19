package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
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

	// Inline Editing State
	IsEditing bool
	Input     textinput.Model
}

// InitialModel returns the initial model state with the suggested message.
func InitialModel(msg string) Model {
	ti := textinput.New()
	ti.Placeholder = "Commit message..."
	ti.SetValue(msg)
	ti.Width = 60
	ti.CharLimit = 100

	return Model{
		Message: msg,
		Choice:  ChoiceNone,
		choices: []string{"Apply", "Edit", "Cancel"},
		cursor:  0,
		Input:   ti,
	}
}

// Init initializes the IO.
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles messages.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// IF EDITING: Handle Text Input
	if m.IsEditing {
		var cmd tea.Cmd
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEnter:
				// Save changes and exit edit mode
				m.Message = m.Input.Value()
				m.IsEditing = false
				m.Choice = ChoiceNone // Reset choice so they can click Apply
				m.cursor = 0          // Focus Apply
				return m, nil
			case tea.KeyEsc:
				// Cancel edit, revert to original message (?)
				// Or just keep current value but correct it.
				// Let's just exit edit mode with current value to be safe.
				m.IsEditing = false
				return m, nil
			}
		}
		m.Input, cmd = m.Input.Update(msg)
		return m, cmd
	}

	// NORMAL NAVIGATION MODE
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
				m.Quitting = true
				return m, tea.Quit
			case 1:
				// ENTER EDIT MODE
				m.Choice = ChoiceEdit
				m.IsEditing = true
				m.Input.SetValue(m.Message) // Reset input to current message
				m.Input.Focus()
				return m, textinput.Blink
			case 2:
				m.Choice = ChoiceCancel
				m.Quitting = true
				return m, tea.Quit
			}
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

	// MSG BOX or INPUT BOX
	var msgContent string
	if m.IsEditing {
		msgContent = m.Input.View()
	} else {
		msgContent = m.Message
	}

	msgBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2).
		MarginBottom(1).
		Width(60).
		Render(msgContent)

	// Button Styles
	btnStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFF7DB")).
		Background(lipgloss.Color("#888B7E")).
		Padding(0, 3).
		MarginRight(1)

	activeBtnStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFF7DB")).
		Background(lipgloss.Color("#F25D94")). // Pinkish focus
		Padding(0, 3).
		MarginRight(1).
		Bold(true)

	// Render choices
	// Don't render buttons if editing
	s := "\n"
	if m.IsEditing {
		s += lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("(Enter to save, Esc to cancel editing)")
	} else {
		for i, choice := range m.choices {
			if m.cursor == i {
				s += activeBtnStyle.Render(choice)
			} else {
				s += btnStyle.Render(choice)
			}
		}
		s += "\n\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("(Use arrows to navigate, Enter to select)")
	}

	return fmt.Sprintf("\n%s\n%s\n%s", header, msgBox, s)
}
