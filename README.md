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

Check repository status.

- **Alias**: `raven s`

```bash
raven status
# Output:
# On branch main ⬆️ 11
# Changes to be committed / Modified:
#   M internal/cli/commit.go
```

Stage files interactively or instantly.

- **Alias**: `raven a`

```bash
raven add [.]
# 'raven a' invokes interactive list.
# 'raven add .' stages all changes instantly.
```

### 2. Smart Commit

Launch the interactive TUI to auto-analyze changes and suggest a message.

- **Alias**: `raven c`

```bash
raven commit
```

- **Auto-Staging**: If nothing is staged, it prompts you to select files.
- **Inline Editing**: Select [Edit] to modify the message without leaving the CLI.

**Manual Mode**:
Bypass analysis and commit instantly.

```bash
raven commit -m "feat: manual message"
```

### 3. "Super Shorthands" (Efficiency Tools)

Raven includes powerful commands to speed up your workflow.

#### The "Save Point"

Stage **all** files (tracked & untracked) and commit in one go.

- **Alias**: `raven ac`, `raven snap`

```bash
raven save
# Stages all files -> Analyzes -> Opens Interactive TUI

raven save -m "wip: fast save"
# Stages all files -> Commits instantly (No TUI)
```

#### Undo & Amend

- **`raven undo`** (`alias: u`): Instantly "un-commits" the last commit but **keeps your changes staged**. Safe and fast.

#### Quick Fixup

- **`raven fix`** (`alias: f`): Stages all changes and merges them into the last commit **silently** (keeps the old message). Great for fixing typos.
  - _Includes a safety confirmation prompt._

### 4. View Stats

Check your coding activity:

```bash
raven stats
```

### 5. Smart Suggestions

Get a quick AI suggestion printed to stdout.

- **Alias**: `raven sg` (System Command)

```bash
raven suggest
# Output: feat(cli): add simple suggest command
```

- **Smart Feedback**: If nothing is staged, it will check for unstaged files and give you tips.

## License

MIT
