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

### 1. Suggest a Commit (Headless)
Get a quick suggestion printed to stdout:

```bash
raven suggest
# Output: feat(cli): add simple suggest command
```

### 2. Interactive Commit (Recommended)
Launch the interactive TUI to review and confirm:

```bash
raven commit
```
*   **Arrow Keys/Tab**: Navigate options.
*   **Enter**: Select action (Apply, Edit, Cancel).

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
