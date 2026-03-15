package analysis

import (
	"testing"
)

func TestAnalyzeDiff(t *testing.T) {
	tests := []struct {
		name     string
		diff     string
		wantType string
	}{
		{
			name: "Go file change",
			diff: `diff --git a/main.go b/main.go
index 8f23..123 100644
--- a/main.go
+++ b/main.go
@@ -1 +1 @@
 package main`,
			wantType: "feat",
		},
		{
			name: "Docs change",
			diff: `diff --git a/docs/readme.md b/docs/readme.md
index 8f23..123 100644
--- a/docs/readme.md
+++ b/docs/readme.md`,
			wantType: "docs",
		},
		{
			name:     "Test file change",
			diff:     `diff --git a/main_test.go b/main_test.go`,
			wantType: "test",
		},
		{
			name:     "Go Mod change",
			diff:     `diff --git a/go.mod b/go.mod`,
			wantType: "chore",
		},
		{
			name:     "Go Sum change",
			diff:     `diff --git a/go.sum b/go.sum`,
			wantType: "chore",
		},
		{
			name:     "Dotfile change",
			diff:     `diff --git a/.gitignore b/.gitignore`,
			wantType: "chore",
		},
		{
			name:     "Not diff git header",
			diff:     `some random text`,
			wantType: "feat",
		},
		{
			name:     "Invalid diff command",
			diff:     `diff --git a/something`,
			wantType: "feat",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnalyzeDiff(tt.diff)
			if got.Type != tt.wantType {
				t.Errorf("AnalyzeDiff() type = %v, want %v", got.Type, tt.wantType)
			}
		})
	}
}
