.PHONY: install-pre-commit, install-golangci-lint, install-linter, lint, lint-all

install-pre-commit:
	pip install pre-commit

install-golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.47.2

install-lint: install-pre-commit install-golangci-lint ## Install dependencies for linting

lint: ## Run lint on changed files
	pre-commit run

lint-all: ## Run lint on all files
	pre-commit run --all-files
