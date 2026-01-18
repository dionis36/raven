package ui

import (
	"raven/internal/stats"

	"github.com/charmbracelet/lipgloss"
)

// RenderHeatmap generates a string representation of the contribution graph.
func RenderHeatmap(counts map[string]int) string {
	dates := stats.GetLastSixMonths()

	// Styles
	emptyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).SetString("□") // Grey
	lowStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("46")).SetString("■")    // Bright Green
	midStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("34")).SetString("■")    // Mid Green
	highStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("22")).SetString("■")   // Dark Green

	// We'll render row by row (7 rows for days of week)
	// But standard heatmap is column by column (weeks).
	// Let's print a simple linear block sequence for MVP or rows?
	// GitHub style: 7 rows (Sun-Sat), many columns.

	// Initialize grid: 7 rows, N columns
	// A simple approach: just print them in a single massive block line for MVP?
	// No, let's try a simplified week grid.

	// Simple MVP: Just distinct blocks for the last N days wrapping?
	// Let's stick to true GitHub style:
	// Calculate weeks
	weeks := len(dates) / 7
	if len(dates)%7 != 0 {
		weeks++
	}

	grid := make([][]string, 7) // 7 rows
	for i := range grid {
		grid[i] = make([]string, weeks)
	}

	currWeek := 0
	for _, d := range dates {
		dateStr := d.Format("2006-01-02")
		count := counts[dateStr]

		weekday := int(d.Weekday()) // 0=Sun, 6=Sat

		// Pick style
		var s lipgloss.Style
		if count == 0 {
			s = emptyStyle
		} else if count < 3 {
			s = lowStyle
		} else if count < 6 {
			s = midStyle
		} else {
			s = highStyle
		}

		grid[weekday][currWeek] = s.String()

		if weekday == 6 {
			currWeek++
		}
	}

	output := ""
	output += lipgloss.NewStyle().
		Foreground(lipgloss.Color("212")).
		Bold(true).
		Render("📊 Contribution Activity") + "\n\n"

	for row := 0; row < 7; row++ {
		for col := 0; col < weeks; col++ {
			if col < len(grid[row]) && grid[row][col] != "" {
				output += grid[row][col] + " "
			} else {
				// if we haven't reached this date yet (start of loop offset?)
				// For simplicity, just skip or print space if empty logic
			}
		}
		output += "\n"
	}

	// Legend
	output += "\n" +
		emptyStyle.String() + " 0  " +
		lowStyle.String() + " 1-2  " +
		midStyle.String() + " 3-5  " +
		highStyle.String() + " 6+"

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1, 2).
		Render(output)
}
