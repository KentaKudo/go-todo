VERSION := v0.1.0

OUT := "bin/goapi-skel"
PKG := "github.com/KentaKudo/goapi-skel"
PKG_BUILD := "${PKG}/cmd/todod"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -extldflags \"-static\""

.PHONY: all
all: build

# Dependency Resolution

.PHONY: deps
deps: ## Resolve the dependencies
	@go mod download

# Binary Build

.PHONY: build
build: deps ## Build the main.go file
	@go build -i -v -o $(OUT) $(PKG_BUILD)

.PHONY: install
install:
	@go install $(LDFLAGS)

.PHONY: clean
clean: ## Remove the previous builds
	@rm -rf bin/*

# Tests

.PHONY: test
test: ## Run the test
	@go test -race -cover -v $(PKG_LIST)

.PHONY: ci-test
ci-test: ## Run the test for CI
	@echo "" > coverage.txt
	@for d in ${PKG_LIST}; do \
		go test -race -coverprofile=profile.out -covermode=atomic $$d; \
		if [ -f profile.out ]; then \
			cat profile.out >> coverage.txt; \
			rm profile.out; \
		fi; \
	done

.PHONY: fmt
fmt: ## Format Go source code
	@go fmt $(PKG_LIST)

.PHONY: vet
vet:
	@go vet $(PKG_LIST)

# Binary Distribution

.PHONY: cross-build
cross-build: deps
	for os in darwin linux windows; do \
		for arch in arm64 386; do \
			# GOOS=$$os GOARCH=$$arch go build
		done;
	done;

.PHONY: dist
dist:

# Docker

DOCKER_REPOSITORY := ""
DOCKER_IMAGE_NAME := "${DOCKER_REPOSITORY}/..."
DOCKER_IMAGE_TAG ?= latest
DOCKER_IMAGE := ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}

.PHONY: docker-build
docker-build:
	@docker build -t $(DOCKER_IMAGE) .

.PHONY: ci-docker-release
ci-docker-release:
	# docker login
	# docker push

# Release

.PHONY: release
release: ## Tag with the current version and push
	git tag $(VERSION)
	git push origin $(VERSION)

# Help

.PHONY: help
help: ## Display the help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'