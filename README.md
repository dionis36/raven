# Raven 🐦  
**Smart Git Commit Assistant (CLI Tool)**

Raven is a cross-platform, terminal-based tool written in Go that analyzes your staged changes and suggests **Conventional Commit** messages. It also tracks your coding activity with a built-in contribution heatmap.

![Raven Demo](https://placehold.co/600x400?text=Raven+CLI+Demo) (Placeholder)

## Features

- 🧠 **Smart Analysis**: Automatically analyzes staged diffs to suggest commit types (`feat`, `fix`, `docs`, etc.).
- 🚀 **Interactive UX**: Review, edit, or apply suggest messages using a beautiful TUI (Terminal User Interface).
- 📊 **Contribution Stats**: View your local git contribution heatmap directly in the terminal.
- ⚡ **Fast & Lightweight**: Built with Go, compiles to a single native binary.
- 🐧 **Cross-Platform**: Works on Linux, macOS, and Windows.

## Installation

### From Source
Requires **Go 1.25+**.

```bash
git clone https://github.com/dionis36/raven.git
cd raven
go build -o raven cmd/raven/main.go
sudo mv raven /usr/local/bin/
```

## Usage

### 1. View & Manage Status
Check the status of your repository and stage files interactively:

```bash
raven status
# Output:
# On branch main ⬆️ 11
# Changes to be committed / Modified:
#   M internal/cli/commit.go
# Untracked files:
#   ? new_feature.go
```

Stage files easily:

```bash
raven add [.]
# Launch interactive staging list (Space to toggle, 'a' for all)
# Or use '.' to stage all changes instantly
```

### 2. Smart Commit
Launch the interactive TUI to auto-analyze changes and suggest a message:

```bash
raven commit
```
*   **Auto-Staging**: If nothing is staged, it prompts you to select files.
*   **Inline Editing**: Select [Edit] to modify the message without leaving the CLI.

**Manual Mode**:
Bypass analysis and commit directly (still runs auto-staging check):

```bash
raven commit -m "feat: manual message"
```

### 3. View Stats
Check your coding activity:

```bash
raven stats
```

## Contributing
Pull requests are welcome! Please ensure you run tests before submitting.

```bash
go test ./...
```

## License
MIT
