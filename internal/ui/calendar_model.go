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
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Align(lipgloss.Center).
		Width(42).
		Render(m.ViewingMonth.Format("January 2006")) + "\n"

	// Weekday Headers
	weekdays := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	wHeader := ""
	for _, day := range weekdays {
		// Each box is ~6 chars wide with border?
		// Actually lipgloss border adds to width.
		// If inner width 4, border 1+1, total 6.
		wHeader += lipgloss.NewStyle().Width(6).Align(lipgloss.Center).Bold(true).Render(day)
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
			Foreground(lipgloss.Color("212")).
			Render(fmt.Sprintf("%s: %d commits", m.SelectedDate.Format("Mon Jan 02"), c))
	}

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2).
		Render(header + "\n" + wHeader + gridStr + "\n" + selInfo + "\n" + help)
}

func renderDayBox(day int, count int, selected bool) string {
	// Compact Box Style
	// Width 4: Fits "30 " or " 1 " comfortably
	// Height 2: Line 1 for Day, Line 2 for Dot
	base := lipgloss.NewStyle().
		Width(4).
		Height(2).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("238")) // Darker grey for subtlety

	if selected {
		base = base.
			BorderForeground(lipgloss.Color("212")). // Pink border for selection
			Bold(true)
	}

	// Content
	dayStr := fmt.Sprintf("%d", day)

	// Activity indicator
	indicator := " "
	if count > 0 {
		dotColor := "250" // Light grey default
		if count > 0 {
			dotColor = "46"
		} // Green
		if count > 3 {
			dotColor = "34"
		}
		if count > 6 {
			dotColor = "22"
		}

		indicator = lipgloss.NewStyle().Foreground(lipgloss.Color(dotColor)).Render("●")
	}

	// Layout:
	// "12"  (Align Right for cleaner calendar look?)
	// " ●"
	// Let's try centering both

	l1 := lipgloss.NewStyle().Width(4).Align(lipgloss.Center).Render(dayStr)
	l2 := lipgloss.NewStyle().Width(4).Align(lipgloss.Center).Render(indicator)

	return base.Render(l1 + "\n" + l2)
}

func renderEmptyBox() string {
	return lipgloss.NewStyle().
		Width(4).
		Height(2).
		Border(lipgloss.HiddenBorder()).
		Render(" ")
}

// Helper for days in month
func daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
