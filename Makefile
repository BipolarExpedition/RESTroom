LDFLAGS := -X main.Version=$(VERSION) -X main.Commit=$(COMMIT) -X main.BuildTime=$(BUILDTIME)

CMD_DIRS := $(wildcard ./cmd/*)
MAIN_BIN := ./cmd/restroom

BUILD_BINARIES := build/bin

.PHONY: lint build clean test run help

check-mod:
	@if [ ! -f go.mod ]; then \
		echo "go.mod not found. Please run 'go mod init' to initialize the Go module."; \
		exit 1; \
	fi

lint: lint-gofmt lint-govet lint-gocritic lint-errcheck

lint-gofmt:
	@echo "Running gofmt"
	@gofmt -s -w ./cmd ./pkg ./internal ./test

lint-govet:
	@echo "Running go vet"
	@go vet ./...

lint-gocritic:
	@echo "Running gocritic (external tool)"
	@go run github.com/go-critic/go-critic/cmd/gocritic@latest check ./...

lint-errcheck:
	@echo "Running errcheck (external tool)"
	@go run github.com/kisielk/errcheck@latest ./...

build: check-mod
	@echo "Building binaries"
	@mkdir -p $(BUILD_BINARIES)
	@VERSION=$(shell git describe --tags || echo "dev"); \
	COMMIT=$(shell git rev-parse --short HEAD || echo "unknown"); \
	BUILDTIME=$(shell date -u "+%Y-%m-%dT%H:%M:%SZ"); \
	for dir in $(CMD_DIRS); do \
		binary_name=$$(basename $$dir); \
		echo "Building $$binary_name from $$dir   (version: $$VERSION)"; \
		go build -ldflags "-X main.Version=$$VERSION -X main.Commit=$$COMMIT -X main.BuildTime=$$BUILDTIME" -o $(BUILD_BINARIES)/$$binary_name $$dir  || exit 1; \
	done

clean:
	rm -f $(BUILD_BINARIES)/* $(CMD_DIRS)/*.test

test: check-mod
	go test -v ./...

run:
	go run $(MAIN_BIN)

help:
	@echo "Available targets:"
	@echo "  help           Show this help message."
	@echo "  clean          Remove all built binaries and test files."
	@echo "  lint           Run all linting tasks (gofmt, govet, gocritic, errcheck)."
	@echo "  test           Run all tests in the project."
	@echo "  build          Build all binaries from cmd/* and place them in $(BUILD_BINARIES)."
	@echo "  run            Run the main application (defined as $(MAIN_BIN))."
	@echo "  check-mod      Verify if go.mod exists, and prompt to initialize it if missing."
	@echo "  lint-gofmt     Format Go code using gofmt."
	@echo "  lint-govet     Run go vet to identify code issues."
	@echo "  lint-gocritic  Run gocritic static analysis checks (external tool)."
	@echo "  lint-errcheck  Check for unhandled errors using errcheck (external tool)."

