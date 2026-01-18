package ui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestModelUpdate(t *testing.T) {
	initial := InitialModel("test commit")

	// Test moving down
	msg := tea.KeyMsg{Type: tea.KeyDown}
	m, _ := initial.Update(msg)
	model := m.(Model)
	if model.cursor != 1 {
		t.Errorf("expected cursor to be 1, got %d", model.cursor)
	}

	// Test selecting "Apply" (cursor 0)
	initial.cursor = 0
	msg = tea.KeyMsg{Type: tea.KeyEnter}
	m, _ = initial.Update(msg)
	model = m.(Model)
	if model.Choice != ChoiceApply {
		t.Errorf("expected ChoiceApply, got %v", model.Choice)
	}
	if !model.Quitting {
		t.Errorf("expected Quitting to be true")
	}
}
