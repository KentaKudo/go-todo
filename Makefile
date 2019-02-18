OUT := "bin/goapi-skel"
PKG := "github.com/KentaKudo/goapi-skel"
PKG_BUILD := $(PKG)

.PHONY: build

all: build

dep: ## Get the dependencies
	@go get -v -d ./...

build: dep ## Build the main.go file
	@go build -i -v -o $(OUT) $(PKG_BUILD)

clean: ## Remove the previous builds
	@rm $(OUT)

help: ## Display the help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'