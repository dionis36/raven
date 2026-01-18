# RAVEN 🐦  
**Smart Git Commit Assistant (CLI Tool)**

---

## 1. INTRODUCTION

### 1.1 Purpose of the Document
This document is a **combined Software Requirements Specification (SRS)**, **System Design Document (SDD)**, and **Software Development / Project Roadmap** for the project **Raven**.

It is intended to:
- Clearly define what Raven does and why it exists
- Describe the system architecture and design decisions
- Provide a structured development roadmap suitable for academic evaluation and real-world implementation

This document is designed for **Canva presentation or document export** (A4 / slide-friendly sections).

---

### 1.2 Project Overview
**Raven** is a cross-platform, terminal-based CLI tool written in **Go** that analyzes Git staged changes and intelligently generates **Conventional Commit–compliant commit messages**. Raven also provides a **Git contribution heatmap** directly inside the terminal, offering developers insights into their coding activity without relying on web-based tools.

Raven focuses on:
- Developer productivity
- Clean commit history
- Excellent CLI user experience (UX)
- Cross-platform reliability (Windows, Linux, macOS)

---

### 1.3 Scope
**In Scope**:
- Git staged diff analysis
- Commit classification (feat, fix, refactor, docs, test, chore)
- Conventional commit message generation
- Interactive colored CLI interface
- Git contribution statistics visualization (CLI-only)
- Cross-platform binary support

**Out of Scope**:
- Graphical User Interface (GUI)
- Cloud services or accounts
- AI/ML-based natural language models (heuristic-based only)

---

### 1.4 Target Users
- Software developers
- Computer Science students
- Open-source contributors
- DevOps engineers

---

## 2. SOFTWARE REQUIREMENTS SPECIFICATION (SRS)

### 2.1 Functional Requirements

#### FR-1: Git Repository Detection
Raven shall detect whether the current directory is a valid Git repository.

#### FR-2: Staged Change Analysis
Raven shall analyze only **staged files** using Git’s diff mechanism (`git diff --cached`).

#### FR-3: Change Classification
Raven shall classify changes into the following types:
- feat
- fix
- refactor
- docs
- test
- chore

Classification shall be based on:
- File paths
- Diff content
- Keywords and patterns

#### FR-4: Commit Message Generation
Raven shall generate commit messages following the **Conventional Commits specification**:
```
type(scope): short description
```

#### FR-5: Interactive Commit Approval
Raven shall allow the user to:
- Accept the suggested commit
- Edit the commit message
- Cancel the commit

#### FR-6: Commit Execution
Raven shall execute the Git commit using the approved commit message.

#### FR-7: Contribution Statistics
Raven shall analyze Git history and display a **CLI-based contribution heatmap** representing commit frequency over time.

#### FR-8: Configuration Management
Raven shall support basic configuration options such as:
- Default commit scope behavior
- Color theme preferences

---

### 2.2 Non-Functional Requirements

#### NFR-1: Cross-Platform Compatibility
Raven must function correctly on:
- Windows
- Linux
- macOS

#### NFR-2: Performance
Raven shall execute analysis and commit generation in under **500ms** for typical repositories.

#### NFR-3: Usability
- Fully keyboard-driven
- Clear prompts and confirmations
- Readable color contrast

#### NFR-4: Reliability
- Graceful error handling
- Clear error messages for missing Git or invalid repository

#### NFR-5: Maintainability
- Modular Go project structure
- Clear separation of concerns

---

## 3. SYSTEM DESIGN DOCUMENT (SDD)

### 3.1 System Architecture Overview

Raven follows a **layered CLI architecture**:

```
User
 ↓
CLI Interface (Cobra + Bubbletea)
 ↓
Analysis Layer
 ↓
Git Integration Layer
 ↓
Operating System / Git CLI
```

---

### 3.2 High-Level Components

#### 3.2.1 CLI Layer
- Handles commands (`raven commit`, `raven stats`)
- Manages user interaction
- Applies colors and layouts

Libraries:
- Cobra
- Bubbletea
- Lipgloss

---

#### 3.2.2 Git Integration Layer
- Executes Git commands
- Parses Git output

Key Commands Used:
- `git diff --cached`
- `git log --pretty=format:%ad`

---

#### 3.2.3 Analysis Engine
- Applies heuristic rules
- Determines commit type and scope
- Extracts meaningful summaries

---

#### 3.2.4 Commit Formatter
- Builds Conventional Commit messages
- Ensures formatting compliance

---

#### 3.2.5 Statistics Engine
- Parses commit history
- Aggregates commits per day
- Renders ASCII/colored heatmaps

---

### 3.3 Data Flow

1. User invokes Raven command
2. CLI validates Git repository
3. Git data is extracted
4. Analysis engine processes diffs
5. Commit message is generated
6. User confirms or edits
7. Git commit is executed

---

### 3.4 Technology Stack

| Layer | Technology |
|---|---|
| Language | Go |
| CLI Framework | Cobra |
| UI Rendering | Bubbletea, Lipgloss |
| Version Control | Git |
| OS Support | Windows, Linux, macOS |

---

## 4. SOFTWARE DEVELOPMENT & PROJECT ROADMAP

### Phase 1: Planning & Design (Week 1)
- Requirements gathering
- UX flow design
- Architecture definition
- Project naming & branding (Raven)

---

### Phase 2: Core MVP Development (Weeks 2–3)
- Git repository detection
- Staged diff extraction
- Basic heuristic analysis
- Conventional commit generation

Deliverable:
- `raven suggest`

---

### Phase 3: Interactive CLI UX (Weeks 4–5)
- Colored UI
- Keyboard navigation
- Commit approval flow

Deliverable:
- `raven commit`

---

### Phase 4: Contribution Statistics (Week 6)
- Git log parsing
- Commit aggregation
- CLI heatmap rendering

Deliverable:
- `raven stats`

---

### Phase 5: Cross-Platform Packaging (Week 7)
- Windows, Linux, macOS builds
- Binary testing
- Error handling improvements

---

### Phase 6: Documentation & Finalization (Week 8)
- README creation
- Usage examples
- Final testing
- Academic submission / portfolio publishing

---

## 5. RISKS & MITIGATION

| Risk | Mitigation |
|---|---|
| Git output variability | Use stable Git commands |
| Terminal color issues | NO_COLOR fallback |
| Windows compatibility | Native Go binaries |

---

## 6. CONCLUSION

Raven is a **production-quality CLI tool** that combines Git expertise, Go systems programming, and thoughtful developer experience design. The project demonstrates strong software engineering principles and is suitable for:
- Academic evaluation
- Portfolio showcase
- Real-world developer usage

---

**Raven — commits with insight. 🪶**

