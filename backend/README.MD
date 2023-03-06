# Development / How to run

First, you should have **Go** installed.

To run the project, at the root directory:
```bash
go run .
```

## Linting and error checking

This project uses [golangci-lint](https://github.com/golangci/golangci-lint.git) as a linter, and therefore it must be configured in your system for a better development experience.

[Installation Guide](https://golangci-lint.run/usage/install/#local-installation)

NOTE: on VS Code you might have to change the lint tool in user preferences.

## Project structure

The main file resides in **main.go** and instead of defining every route in there, we should only call functions that create route groups.

These route groups should have it's own folder and function that handles the route creation in that group, as demonstrated in the api group.