help: ## display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup: ## set up go mod dependancies
	go mod download

test: ## run go tests
	go test -race -short -cover ./...
	golint -set_exit_status ./...
	go vet ./...
	go mod tidy
	go mod verify

build: ## build binaries for the project
	gofmt -w -s -d .
	go build -race -v ./pkg/...
	go build -race -v -o bin/cli ./cmd/...
