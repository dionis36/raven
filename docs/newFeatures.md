# Planned Features: Super Shorthands

This document outlines future enhancements designed to make the Raven workflow even faster and smoother.

## 1. The "Save Point" Command (`raven save`)
**Goal**: Instantly save all work with minimal friction.

*   **Command**: `raven save` (or `raven ac` / `raven snap`)
*   **Workflow**:
    1.  Executes `git add .` (Stages all tracked and untracked files).
    2.  Runs AI Analysis on the staged changes.
    3.  Opens the `raven commit` TUI immediately with the generated suggestion.
*   **Why**: Turns the common "add everything and commit" loop into a single, smart command.

## 2. Smart Undo & Amend
**Goal**: Simplify Git's confusing history modification commands.

### `raven amend`
*   **Wrapper**: `git commit --amend`
*   **Behavior**: Opens the **last** commit message in the Raven TUI for editing. If new files are staged, they are folded into that commit.

### `raven undo`
*   **Wrapper**: `git reset --soft HEAD~1`
*   **Behavior**: "Un-commits" the last commit. The changes are preserved in your working directory (staged), allowing you to fix mistakes or split the commit.

## 3. One-Letter Aliases
**Goal**: Speed up strict CLI usage without shell configuration.

*   **`raven s`** → `raven status`
*   **`raven a`** → `raven add`
*   **`raven c`** → `raven commit`

## 4. Quick Fixup (`raven fix`)
**Goal**: seamless "typo fixing" workflow.

*   **Command**: `raven fix`
*   **Workflow**:
    1.  Stages current changes.
    2.  Merges them into the previous commit (`git commit --amend --no-edit`).
    3.  Keeps the original commit message.
*   **Why**: Perfect for when you notice a typo right after committing and just want to patch it silently.
