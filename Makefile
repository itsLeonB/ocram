TEST_DIR := ./internal/test/...
COVER_PKG := ./internal/...

.PHONY: help \
job \
lint \
test \
test-verbose \
test-coverage \
test-coverage-html \
test-clean \
build-job \
install-pre-push-hook \
uninstall-pre-push-hook

help:
	@echo "Makefile commands:"
	@echo "  make job                     - Start the job processor"
	@echo "  make lint                    - Run golangci-lint on the codebase"
	@echo "  make test                    - Run all tests"
	@echo "  make test-verbose            - Run all tests with verbose output"
	@echo "  make test-coverage           - Run all tests with coverage report"
	@echo "  make test-coverage-html      - Run all tests and generate HTML coverage report"
	@echo "  make test-clean              - Clean test cache and run tests"
	@echo "  make build-job               - Build the job application binary"
	@echo "  make install-pre-push-hook   - Install the pre-push git hook"
	@echo "  make uninstall-pre-push-hook - Uninstall the pre-push git hook"

job:
	go run cmd/job/main.go

lint:
	golangci-lint run ./...

test:
	@echo "Running all tests..."
	@if [ -d $(TEST_DIR) ]; then \
		go test $(TEST_DIR); \
	else \
		echo "No tests found in $(TEST_DIR), skipping."; \
	fi

test-verbose:
	@echo "Running all tests with verbose output..."
	@if [ -d $(TEST_DIR) ]; then \
		go test -v $(TEST_DIR); \
	else \
		echo "No tests found in $(TEST_DIR), skipping."; \
	fi

test-coverage:
	@echo "Running all tests with coverage report..."
	@if [ -d $(TEST_DIR) ]; then \
		go test -v -cover -coverprofile=coverage.out -coverpkg=$(COVER_PKG) $(TEST_DIR); \
	else \
		echo "No tests found in $(TEST_DIR), skipping."; \
	fi

test-coverage-html:
	@echo "Running all tests and generating HTML coverage report..."
	@if [ -d $(TEST_DIR) ]; then \
		go test -v -cover -coverprofile=coverage.out -coverpkg=$(COVER_PKG) $(TEST_DIR) && \
		go tool cover -html=coverage.out -o coverage.html && \
		echo "Coverage report generated: coverage.html"; \
	else \
		echo "No tests found in $(TEST_DIR), skipping."; \
	fi

test-clean:
	@echo "Cleaning test cache and running tests..."
	@if [ -d $(TEST_DIR) ]; then \
		go clean -testcache && go test -v $(TEST_DIR); \
	else \
		echo "No tests found in $(TEST_DIR), skipping."; \
	fi

build-job:
	@echo "Building the job..."
	CGO_ENABLED=0 GOOS=linux go build -trimpath -buildvcs=false -ldflags='-w -s' -o bin/job cmd/job/main.go
	@echo "Build success! Binary is located at bin/job"

install-pre-push-hook:
	@echo "Installing pre-push git hook..."
	@mkdir -p .git/hooks
	@cp scripts/git-pre-push.sh .git/hooks/pre-push
	@chmod +x .git/hooks/pre-push
	@echo "Pre-push hook installed successfully!"

uninstall-pre-push-hook:
	@echo "Uninstalling pre-push git hook..."
	@rm -f .git/hooks/pre-push
	@echo "Pre-push hook uninstalled successfully!"
