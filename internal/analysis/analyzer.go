package analysis

import (
	"strings"
)

// Suggestion represents a suggested commit message structure.
type Suggestion struct {
	Type        string
	Scope       string
	Description string
}

// AnalyzeDiff analyzes the staged diff and allows to categorize changes.
// For the MVP, it implements naive heuristics based on file extensions and paths.
func AnalyzeDiff(diff string) Suggestion {
	// Defaults
	suggestion := Suggestion{
		Type:        "feat",
		Scope:       "",
		Description: "update code",
	}

	// Simple heuristic: check changed file paths in the diff header
	lines := strings.Split(diff, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "diff --git") {
			// Extract file path from "a/path/to/file b/path/to/file"
			parts := strings.Fields(line)
			if len(parts) >= 3 {
				// use the 'b' path (new version) usually parts[3] but let's just look at the string
				path := parts[len(parts)-1]

				// Docs
				if strings.Contains(path, "docs/") || strings.HasSuffix(path, ".md") {
					suggestion.Type = "docs"
					suggestion.Description = "update documentation"
					return suggestion // Priority return
				}

				// Go files (Code)
				if strings.HasSuffix(path, ".go") {
					suggestion.Type = "feat" // Default to feat for code
					suggestion.Description = "implement feature"
					// check for tests
					if strings.HasSuffix(path, "_test.go") {
						suggestion.Type = "test"
						suggestion.Description = "add tests"
						return suggestion
					}
				}

				// Config / Chore
				if strings.HasSuffix(path, "go.mod") || strings.HasSuffix(path, "go.sum") || strings.HasPrefix(path, ".") {
					suggestion.Type = "chore"
					suggestion.Description = "update dependencies/config"
				}
			}
		}
	}

	return suggestion
}
