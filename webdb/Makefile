APP?=web

.PHONY: build
## build: builds the binaries
build:
	@echo "Building..."
	@go build -o bin/web/${APP} web/main.go

.PHONY: run
## run: runs the binaries
run:
	go run web/main.go

.PHONY: clean
## clean: cleans the binaries
clean:
	@echo "Cleaning"
	rm bin/web/${APP}

.PHONY: setup
## setup: setup go modules
setup:
	@go mod init \
		&& go mod tidy \
		&& go mod vendor

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'