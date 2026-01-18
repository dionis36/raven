package git

import (
	"os/exec"
)

// IsRepository checks if the current directory is within a git repository.
func IsRepository() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// GetStagedDiff returns the diff of staged changes.
func GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
