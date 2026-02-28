---
name: git-conventional-commit
description: Streamline Git commits with the Conventional Commits specification. Use when you need to commit changes, prepare a pull request, or maintain a clean, standardized commit history.
---

# Git Conventional Commit

This skill helps you create standardized commit messages following the Conventional Commits specification.

## Workflow

1.  **Analyze Changes**: Identify staged and unstaged changes using `git status` and `git diff`.
2.  **Stage Files**: If there are unstaged changes, ask the user if they should be staged.
3.  **Determine Commit Type**: Categorize the changes based on the Conventional Commits types.
4.  **Draft Message**: Compose a concise commit message in the format `<type>(<scope>): <description>`.
5.  **Confirm and Commit**: Present the draft to the user and execute `git commit` upon confirmation.

## Conventional Commit Types

- **feat**: A new feature.
- **fix**: A bug fix.
- **docs**: Documentation only changes.
- **style**: Changes that do not affect the meaning of the code (white-space, formatting, etc.).
- **refactor**: A code change that neither fixes a bug nor adds a feature.
- **perf**: A code change that improves performance.
- **test**: Adding missing tests or correcting existing tests.
- **build**: Changes that affect the build system or external dependencies.
- **ci**: Changes to CI configuration files and scripts.
- **chore**: Other changes that don't modify src or test files.
- **revert**: Reverts a previous commit.

## Commit Message Rules

- **Format**: `<type>(<scope>): <description>` (scope is optional).
- **Description**: Use the imperative, present tense ("change" not "changed" nor "changes").
- **Case**: The description should be lowercase.
- **Punctuation**: No period at the end of the description.

## Example Triggers

- "Commit my changes"
- "Wrap up this feature with a commit"
- "Help me commit these fixes"
- "I'm ready to commit"
