# Contributing
Thank you for your interest in contributing to gfind!
There are 2 ways of contributing: opening an Issue and opening a Pull request.
## Opening an Issue
Before opening a new Issue, please make sure it isn't a duplicate of an already existing Issue.
1. Go to the ["Issues" tab](https://github.com/player-372/gfind/issues) and [create a new Issue](https://github.com/player-372/gfind/issues/new/choose)
2. Choose a template for your Issue
3. Fill it in, and submit by pressing "Create" button
## Opening a Pull request
> [!note]
> Go programming language knowledge required for opening a Pull request
1. [Open a new Issue](#opening-an-issue) with description of changes you want to make and check the "I am willing to work on this myself" checkbox
2. Make sure you have Go version 1.26.5+ installed
3. Fork the repository and create your branch from `dev`.
4. Make your changes and ensure the code is formatted:
```bash
go fmt ./...
go vet ./...
```
5. Open a Pull request to the `dev` branch with a link to the Issue you made in the first step
## Branching model
- `main` - stable releases only
- `dev` - development branch
- `feature/*` - new features (merge to `dev`)
## Code style
- Use `go fmt` before committing
- Write clear commit messages in English
- Use [Conventional Commits](https://conventionalcommits.org)
# License
By contributing, you agree that your contributions will be licensed under the [MIT License](https://github.com/player-372/gfind/blob/main/LICENSE)