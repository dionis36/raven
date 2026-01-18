package ui

import (
	"fmt"
	"raven/internal/stats"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type StatsModel struct {
	CursorR int // Row (0-6)
	CursorC int // Column (Week)
	Counts  map[string]int
	Dates   [][]time.Time // Grid mapping [row][col] -> Date
	Weeks   int
}

func InitialStatsModel(counts map[string]int) StatsModel {
	// Get full range
	allDates := stats.GetLastSixMonths()

	// Filter: Find first week with activity?
	// For now, let's stick to full 6 months to keep context, but arguably
	// we could scan for first non-zero index.
	// But "previous empty dated" might mean the 0s at the start.
	// Let's find start index.
	// Optimization: Skip empty weeks at the start?
	// Implementation: Check weeks 0..N. If week has 0 commits, skip?
	// Let's keep it simple first: Full range.
	// Interaction is the key requested feature.

	// Build Grid
	// We need a 7xWeeks grid.
	weeks := (len(allDates) / 7) + 2
	grid := make([][]time.Time, 7)
	for i := range grid {
		grid[i] = make([]time.Time, weeks)
	}

	currWeek := 0
	for _, d := range allDates {
		weekday := int(d.Weekday())
		if currWeek < weeks {
			grid[weekday][currWeek] = d
		}
		if weekday == 6 {
			currWeek++
		}
	}

	// Set cursor to the last day (Today)
	// Find last populated column in a row
	// Rough approximation: last week, last row
	startR := 6
	startC := currWeek
	// Adjust bounds
	if startC >= weeks {
		startC = weeks - 1
	}

	return StatsModel{
		CursorR: startR,
		CursorC: startC,
		Counts:  counts,
		Dates:   grid,
		Weeks:   weeks,
	}
}

func (m StatsModel) Init() tea.Cmd {
	return nil
}

func (m StatsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit

		case "up", "k":
			if m.CursorR > 0 {
				m.CursorR--
			} else {
				// Wrap around? Or stop?
				// Stop is better for grid intuition
			}
		case "down", "j":
			if m.CursorR < 6 {
				m.CursorR++
			}
		case "left", "h":
			if m.CursorC > 0 {
				m.CursorC--
			}
		case "right", "l":
			if m.CursorC < m.Weeks-1 {
				m.CursorC++
			}
		}
	}
	return m, nil
}

func (m StatsModel) View() string {
	// Styles
	// emptyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).SetString("□")
	// lowStyle   := lipgloss.NewStyle().Foreground(lipgloss.Color("46")).SetString("■")
	// midStyle   := lipgloss.NewStyle().Foreground(lipgloss.Color("34")).SetString("■")
	// highStyle  := lipgloss.NewStyle().Foreground(lipgloss.Color("22")).SetString("■")

	// Better Palette (Blue/Purple/Pink?) Users like variance.
	// But let's stick to Green for "GitHub" feel per request, but nicer.
	// Let's use blocks with background colors for "modern" feel?
	// Or sticking to chars is safer for terminals.

	output := ""

	// Title
	output += lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true).Render("📊 Interactive Contribution Map") + "\n\n"

	// Render Grid
	for r := 0; r < 7; r++ {
		for c := 0; c < m.Weeks; c++ {
			date := m.Dates[r][c]
			if date.IsZero() {
				output += "  "
				continue
			}

			count := m.Counts[date.Format("2006-01-02")]

			// Determine Color
			color := "240" // Grey
			if count > 0 {
				color = "46"
			}
			if count > 3 {
				color = "34"
			}
			if count > 6 {
				color = "22"
			}

			char := "■"

			// Cursor Logic
			cell := lipgloss.NewStyle().Foreground(lipgloss.Color(color)).SetString(char)

			if r == m.CursorR && c == m.CursorC {
				// Highlight Cursor: White char, or maybe a background?
				// White block with blinking?
				cell = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Bold(true).SetString("▣")
			}

			output += cell.String() + " "
		}
		output += "\n"
	}

	// Details Footer
	// Get date/count at cursor
	selDate := m.Dates[m.CursorR][m.CursorC]
	footer := ""
	if !selDate.IsZero() {
		selCount := m.Counts[selDate.Format("2006-01-02")]
		dateStr := selDate.Format("Mon, Jan 02 2006")

		footer = fmt.Sprintf("\nSelected: %s | %d Commits",
			lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(dateStr),
			selCount,
		)
	} else {
		footer = "\nNo Data"
	}

	help := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("\n(Arrows to move • q to quit)")

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1, 2).
		Render(output + footer + help)
}
