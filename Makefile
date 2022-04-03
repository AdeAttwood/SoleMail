help: ## Show help menu
	@echo "\e[32mSoleMail"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

format: format-js format-go ## Formats all of the code

format-js: ## Formats all of the go code
	./node_modules/.bin/prettier --write ./frontend/**/*.{ts,tsx}

format-go: ## Formats all of the frontend js cod
	gofmt -w .

test: test-js test-go ## Run all of the tests available

test-js: ## Run the js test suite
	./node_modules/.bin/jest

test-go: ## Run the go test suite
	go test ./...

antlr4:
	cd pkg/database/parser && antlr4 -Dlanguage=Go Query.g4

.PHONY: build
build: ## Build the application
	wails generate module
	yarn production
	wails build