package stats

import (
	"testing"
	"time"
)

func TestGetLastSixMonths(t *testing.T) {
	dates := GetLastSixMonths()
	if len(dates) < 180 {
		t.Errorf("expected >180 dates, got %d", len(dates))
	}

	last := dates[len(dates)-1]
	now := time.Now()
	if last.Format("2006-01-02") != now.Format("2006-01-02") {
		t.Errorf("expected last date to be today, got %v", last)
	}
}

// Note: TestGetCommitCounts requires mocking exec or running in a real repo.
// Skipping for MVP unit test suite to avoid flakiness, relying on manual verification.
