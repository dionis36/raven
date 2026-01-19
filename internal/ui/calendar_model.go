package ui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// CalendarModel handles the interactive calendar view.
type CalendarModel struct {
	ViewingMonth time.Time // The first day of the month being viewed
	SelectedDate time.Time // The currently highlighted day
	Counts       map[string]int
	Quitting     bool
}

func InitialCalendarModel(counts map[string]int) CalendarModel {
	now := time.Now()
	// Start viewing current month
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	return CalendarModel{
		ViewingMonth: startOfMonth,
		SelectedDate: now, // Select today initially
		Counts:       counts,
	}
}

func (m CalendarModel) Init() tea.Cmd {
	return nil
}

func (m CalendarModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			m.Quitting = true
			return m, tea.Quit

		case "h", "left":
			m.SelectedDate = m.SelectedDate.AddDate(0, 0, -1)
			// Check if we moved back a month from the viewing window
			if m.SelectedDate.Before(m.ViewingMonth) {
				m.ViewingMonth = m.ViewingMonth.AddDate(0, -1, 0)
			}

		case "l", "right":
			newDate := m.SelectedDate.AddDate(0, 0, 1)
			if !newDate.After(time.Now()) { // Prevent future selection
				m.SelectedDate = newDate
				// Check if we moved forward a month
				nextMonth := m.ViewingMonth.AddDate(0, 1, 0)
				if !m.SelectedDate.Before(nextMonth) {
					m.ViewingMonth = nextMonth
				}
			}

		case "k", "up":
			newDate := m.SelectedDate.AddDate(0, 0, -7)
			if newDate.Before(m.ViewingMonth) {
				// We moved to prev month visuals
				m.ViewingMonth = m.ViewingMonth.AddDate(0, -1, 0)
			}
			m.SelectedDate = newDate

		case "j", "down":
			newDate := m.SelectedDate.AddDate(0, 0, 7)
			if !newDate.After(time.Now()) {
				m.SelectedDate = newDate
				nextMonth := m.ViewingMonth.AddDate(0, 1, 0)
				if !m.SelectedDate.Before(nextMonth) {
					m.ViewingMonth = nextMonth
				}
			}

		case "[", "pgup": // Previous Month
			prev := m.ViewingMonth.AddDate(0, -1, 0)
			m.ViewingMonth = prev
			// Adjust selection to stay in view? Or keep selection?
			// Let's reset selection to 1st of that month
			m.SelectedDate = prev

		case "]", "pgdown": // Next Month
			next := m.ViewingMonth.AddDate(0, 1, 0)
			now := time.Now()
			// Only allow if next month isn't totally in future (e.g. current month is fine)
			if !next.After(time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())) {
				m.ViewingMonth = next
				m.SelectedDate = next
			}
		}
	}
	return m, nil
}

func (m CalendarModel) View() string {
	if m.Quitting {
		return ""
	}

	// Constants
	// Box dimensions
	// width 8 chars, height 4 lines

	// Header: Month Year
	// Grid width approx: 7 columns * (4 width + 2 border) = 42 chars
	// Header: Month Year
	// Grid width approx: 7 columns * (4 wid + 2 bound + 1 margin) = 7 * 7 = 49 chars
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#38BDF8")). // Sky-400
		Align(lipgloss.Center).
		Width(50).
		Render(m.ViewingMonth.Format("January 2006")) + "\n"

	// Weekday Headers
	weekdays := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	wHeader := ""
	for _, day := range weekdays {
		// Each box consumes ~7 chars width (including margin)
		wHeader += lipgloss.NewStyle().Width(7).Align(lipgloss.Center).Bold(true).Render(day)
	}
	wHeader += "\n"

	// Grid Building
	// Start day of week for 1st of month
	startDay := int(m.ViewingMonth.Weekday()) // 0=Sun
	daysInMonth := daysIn(m.ViewingMonth.Month(), m.ViewingMonth.Year())

	currentDayIdx := 1
	gridStr := ""

	// Rows (Allow 6 rows maximum for safety)
	for row := 0; row < 6; row++ {
		var rowBlocks []string
		for col := 0; col < 7; col++ {
			// Determine what to render in this cell
			isDay := false
			dayNum := 0

			// Logic to place the 1st on correct column
			if row == 0 && col < startDay {
				// Empty padding before start of month
				isDay = false
			} else if currentDayIdx <= daysInMonth {
				isDay = true
				dayNum = currentDayIdx
				currentDayIdx++
			}

			// Render Box
			if isDay {
				boxDate := time.Date(m.ViewingMonth.Year(), m.ViewingMonth.Month(), dayNum, 0, 0, 0, 0, m.ViewingMonth.Location())
				count := m.Counts[boxDate.Format("2006-01-02")]

				isSelected := boxDate.Year() == m.SelectedDate.Year() &&
					boxDate.Month() == m.SelectedDate.Month() &&
					boxDate.Day() == m.SelectedDate.Day()

				rowBlocks = append(rowBlocks, renderDayBox(dayNum, count, isSelected))
			} else {
				rowBlocks = append(rowBlocks, renderEmptyBox())
			}
		}
		// Join the blocks horizontally to form a single row string
		gridStr += lipgloss.JoinHorizontal(lipgloss.Top, rowBlocks...) + "\n"

		if currentDayIdx > daysInMonth {
			break
		}
	}

	// Instructions
	help := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		MarginTop(1).
		Render("←/→/↑/↓: navigate  •  [/]: prev/next month  •  q: quit")

	// Selected Info
	selInfo := ""
	if m.SelectedDate.IsZero() {
		selInfo = " "
	} else {
		c := m.Counts[m.SelectedDate.Format("2006-01-02")]
		selInfo = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#38BDF8")). // Sky-400
			Render(fmt.Sprintf("%s: %d commits", m.SelectedDate.Format("Mon Jan 02"), c))
	}

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2).
		Render(header + "\n" + wHeader + gridStr + "\n" + selInfo + "\n" + help)
}

func renderDayBox(day int, count int, selected bool) string {
	// 1. Determine Colors
	// Default: Empty Container
	bgColor := lipgloss.Color("#262626") // Neutral Dark Grey (Distinct from Blue)
	fgColor := lipgloss.Color("250")     // Grey Text

	if count > 0 {
		fgColor = lipgloss.Color("255") // White Text (Active)
		if count <= 2 {
			bgColor = lipgloss.Color("#1E3A5F") // Dark Blue
		} else if count <= 5 {
			bgColor = lipgloss.Color("#0369A1") // Sky-700
		} else if count <= 10 {
			bgColor = lipgloss.Color("#0EA5E9") // Sky-500
		} else if count <= 15 {
			bgColor = lipgloss.Color("#38BDF8") // Sky-400
		} else {
			bgColor = lipgloss.Color("#F59E0B") // Gold (Exceptional)
			fgColor = lipgloss.Color("232")     // Black text on Gold
		}
	}

	// 2. Base Style: UNIFORM 6x3 SOLID BLOCKS
	// No borders. This guarantees identical dimensions for all squares.
	style := lipgloss.NewStyle().
		Width(6).
		Height(3).
		Align(lipgloss.Center, lipgloss.Center).
		Background(bgColor).
		Foreground(fgColor).
		MarginRight(1).
		MarginBottom(1)

	// 3. Selection Indicator
	if selected {
		// Instead of changing size/border, we emphasize the content.
		// 1. Sky Blue Text
		// 2. Bold
		// 3. Underline (to simulate a "cursor" without changing box size)
		style = style.
			Foreground(lipgloss.Color("#38BDF8")). // Sky-400 Text (Matches Theme)
			Bold(true).
			Underline(true)
	}

	return style.Render(fmt.Sprintf("%02d", day))
}

func renderEmptyBox() string {
	// Padding (Pre-month days)
	return lipgloss.NewStyle().
		Width(6).
		Height(3).
		MarginRight(1).
		MarginBottom(1).
		Render(" ")
}

// Helper for days in month
func daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
