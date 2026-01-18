# Raven - Project Overview

## Introduction
**Raven** is a cross-platform, terminal-based CLI tool written in **Go**. It is designed to enhance developer productivity by analyzing Git staged changes and intelligently generating **Conventional Commit–compliant** commit messages. Additionally, it provides a CLI-based **Git contribution heatmap**.

## Features

### Core Functionality
- **Git Repository Detection**: Automatically detects if the current directory is a valid Git repository.
- **Staged Change Analysis**: Analyzes `git diff --cached` to understand the context of changes.
- **Intelligent Commit Generation**: Generates commit messages following the Conventional Commits specification (`type(scope): description`).
    - Supports types: `feat`, `fix`, `refactor`, `docs`, `test`, `chore`.
- **Interactive UX**:
    - Approve suggested message.
    - Edit message before committing.
    - Cancel operation.
- **Contribution Stats**: Visualizes coding activity with a CLI-based heatmap.

### Non-Functional
- **Cross-Platform**: Runs on Windows, Linux, and macOS.
- **Fast**: Targeted execution time under 500ms.
- **Zero-Dependency (Runtime)**: Distributed as a single binary.

## Detailed Roadmap

### Phase 1: Planning & Design (Week 1)
**Goal**: Establish the project foundation, architecture, and design the user experience.

*   **Activities**:
    *   Analyze requirements and define the scope of the MVP.
    *   Design the system architecture (CLI layer, Analysis layer, Git layer).
    *   Define the folder structure and initialize the Go module.
    *   Research libraries: Cobra (CLI), Bubble Tea (TUI), Lip Gloss (Styling).
*   **Deliverables**:
    *   `docs/project_overview.md` (This document).
    *   `docs/architecture.md` (System design).
    *   Initialized Git repository with `go.mod`.
*   **Acceptance Tests**:
    *   [x] Repo is created and accessible.
    *   [x] `go mod tidy` runs without errors.
    *   [x] Documentation is approved by stakeholders.

### Phase 2: Core MVP Development (Weeks 2-3)
**Goal**: Implement the core logic for git analysis and commit suggestion (Headless/Basic CLI).

*   **Activities**:
    *   Implement **Git Layer**: Wrappers for `git rev-parse --is-inside-work-tree` and `git diff --cached`.
    *   Implement **Analysis Engine**: Logic to parse diff outputs and identify keywords (e.g., detecting "fix", "feat", "docs").
    *   Implement **Basic CLI**: A simple `raven suggest` command that prints the generated message to stdout.
*   **Deliverables**:
    *   Working binary that outputs a conventional commit string based on staged files.
*   **Acceptance Tests**:
    *   [ ] Run `raven suggest` in a non-git repo -> Returns descriptive error.
    *   [ ] Stage a documentation change -> `raven suggest` returns `docs: ...`.
    *   [ ] Stage a go file change -> `raven suggest` returns `feat: ...` or `fix: ...`.

### Phase 3: Interactive CLI UX (Weeks 4-5)
**Goal**: Create a rich, interactive terminal user interface (TUI).

*   **Activities**:
    *   Integrate **Bubble Tea**: Create a model for the confirmation dialog.
    *   Implement Interactive Flow:
        1.  Show spinner while analyzing.
        2.  Display the generated message with syntax highlighting.
        3.  Prompt: `[Apply] [Edit] [Cancel]`.
    *   Implement `raven commit` command that ties logic and UI together.
*   **Deliverables**:
    *   `raven commit` command with interactive TUI.
*   **Acceptance Tests**:
    *   [ ] User can navigate options with Arrow Keys/Tab.
    *   [ ] Selecting "Edit" opens the default `$EDITOR`.
    *   [ ] Selecting "Apply" runs `git commit`.

### Phase 4: Contribution Statistics (Week 6)
**Goal**: visualize user activity in the terminal.

*   **Activities**:
    *   Implement **Stats Engine**: Parse `git log` to get commit dates and counts.
    *   Implement **Heatmap Renderer**: Logic to map commit counts to colors (like GitHub's green squares).
    *   Create `raven stats` command.
*   **Deliverables**:
    *   `raven stats` command displaying a heatmap.
*   **Acceptance Tests**:
    *   [ ] `raven stats` displays a grid of squares representing the last few months.
    *   [ ] Days with high commits are brighter/different color than empty days.

### Phase 5: Cross-Platform Packaging & Polish (Weeks 7-8)
**Goal**: Prepare the tool for distribution and general use.

*   **Activities**:
    *   **Cross-Compilation**: Build binaries for Linux (amd64/arm64), Windows, and macOS.
    *   **Error Handling**: Polish error messages (e.g., "Git not found", "No staged changes").
    *   **Final Documentation**: Write a comprehensive `README.md` with installation steps.
*   **Deliverables**:
    *   Release v1.0.0 binaries.
    *   Final `README.md`.
*   **Acceptance Tests**:
    *   [ ] Binary runs on a Windows VM/Environment.
    *   [ ] Binary runs on a Fresh Linux Install.
    *   [ ] "No staged changes" message is friendly and helpful.
