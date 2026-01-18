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
