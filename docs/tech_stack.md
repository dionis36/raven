# Technology Stack

Raven is built using a modern Go ecosystem, focusing on performance, portability, and excellent developer experience.

## Core Language
- **[Go](https://go.dev/)**: Statically typed, compiled language known for simplify and concurrency.
    - **Reason**: Single binary distribution, cross-platform compilation, and strong standard library.

## CLI Framework & Libraries
- **[Cobra](https://github.com/spf13/cobra)**: A library for creating powerful modern CLI applications.
    - **Usage**: Handling commands (`commit`, `stats`), flags, and arguments.
- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)**: A powerful little TUI framework (The Elm Architecture for Go).
    - **Usage**: Managing application state and interactive terminal UIs.
- **[Lip Gloss](https://github.com/charmbracelet/lipgloss)**: Style definitions for nice terminal layouts.
    - **Usage**: Styling text, borders, and colors.

## Integration
- **Git**: Raven acts as a wrapper around Git commands.
    - `git diff --cached`: To analyze staged changes.
    - `git log`: To generate contribution statistics.
    - `git commit`: To execute the final commit.

## Build & CI/CD
- **Go Build**: Native toolchain for compilation.
