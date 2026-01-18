package stats

import (
	"os/exec"
	"strings"
	"time"
)

// GetCommitCounts returns a map of date (YYYY-MM-DD) to commit count.
func GetCommitCounts() (map[string]int, error) {
	// git log --pretty=format:%ad --date=short
	cmd := exec.Command("git", "log", "--pretty=format:%ad", "--date=short")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	counts := make(map[string]int)
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		date := strings.TrimSpace(line)
		if date == "" {
			continue
		}
		counts[date]++
	}

	return counts, nil
}

// GetLastSixMonths returns a slice of dates for the last 6 months (approx 180 days).
func GetLastSixMonths() []time.Time {
	var dates []time.Time
	now := time.Now()
	// Start from 6 months ago
	start := now.AddDate(0, -6, 0)

	// Iterate until now
	for d := start; !d.After(now); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d)
	}
	return dates
}
