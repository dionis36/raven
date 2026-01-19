package git

import (
	"os/exec"
	"strings"
)

type FileStatus struct {
	Path      string
	Status    string // "M ", " M", "A ", "??", etc.
	Staged    bool   // Derived from Status
	Untracked bool   // Derived from Status == "??"
}

// StatusResult holds the full status including branch info.
type StatusResult struct {
	BranchInfo string // "main...origin/main [ahead 1]"
	Files      []FileStatus
}

// GetStatus returns the full status including branch info and changed files.
func GetStatus() (StatusResult, error) {
	// -s: short format
	// -b: show branch info
	// --porcelain: easy parsing
	cmd := exec.Command("git", "status", "-sb", "--porcelain")
	out, err := cmd.Output()
	if err != nil {
		return StatusResult{}, err
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	var result StatusResult
	var files []FileStatus

	for _, line := range lines {
		if len(line) < 3 {
			continue
		}

		// Check for Branch Header
		if strings.HasPrefix(line, "##") {
			// Format: "## main...origin/main [ahead 1]"
			result.BranchInfo = strings.TrimSpace(line[3:])
			continue
		}

		// File Lines: "XY Path"
		status := line[:2]
		path := strings.TrimSpace(line[3:])

		// Determine staged status
		// "M " -> Modified (Staged)
		// "MM" -> Modified (Staged & Unstaged)
		// " M" -> Modified (Unstaged)
		// "??" -> Untracked
		isStaged := status[0] != ' ' && status[0] != '?'

		isUntracked := status == "??"

		files = append(files, FileStatus{
			Path:      path,
			Status:    status,
			Staged:    isStaged,
			Untracked: isUntracked,
		})
	}

	result.Files = files
	return result, nil
}

// StageFile stages a specific file (git add).
func StageFile(path string) error {
	cmd := exec.Command("git", "add", path)
	return cmd.Run()
}

// UnstageFile unstages a file (git restore --staged).
func UnstageFile(path string) error {
	cmd := exec.Command("git", "restore", "--staged", path)
	return cmd.Run()
}
