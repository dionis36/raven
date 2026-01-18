# Architecture & File Structure

## System Architecture

Raven follows a layered architecture to separate CLI concerns from logic and git integration.

```mermaid
graph TD
    User[User] --> CLI[CLI Layer (Cobra + Bubble Tea)]
    CLI --> Analysis[Analysis Engine]
    CLI --> Stats[Statistics Engine]
    Analysis --> Git[Git Integration Layer]
    Stats --> Git
    Git --> OS[OS / Git Binary]
```

### Components
1.  **CLI Layer**: Handles user input, commands, and rendering the TUI.
2.  **Analysis Engine**: Parses diffs and applies heuristics to generate commit messages.
3.  **Statistics Engine**: Aggregates git history for visualization.
4.  **Git Integration**: Wrapper functions to execute git commands safely.

## Proposed File Structure

```text
raven/
├── cmd/                # Entry points for the application
│   └── raven/          # Main package
│       └── main.go
├── internal/           # Private application code
│   ├── app/            # Application logic wiring
│   ├── cli/            # Cobra commands & Bubbletea models
│   │   ├── root.go
│   │   ├── commit.go
│   │   └── stats.go
│   ├── git/            # Git wrapper implementation
│   │   └── git.go
│   ├── analysis/       # Commit message generation logic
│   │   └── analyzer.go
│   └── stats/          # Contribution heatmap logic
│       └── heatmap.go
├── pkg/                # Library code (if any reusable packages)
├── docs/               # Documentation
├── go.mod              # Module definition
└── go.sum
```

This structure ensures maintainability and follows standard Go project layouts.
