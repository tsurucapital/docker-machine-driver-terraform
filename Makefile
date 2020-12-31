VERSION_INFO_FILE = version-info.go
BIN_DIR = bin
BIN_NAME = docker-machine-driver-terraform

REPO_BASE = github.com/krzysztof-miemiec/docker-machine-driver-terraform

default: fmt build test

fmt:
	go fmt $(REPO_BASE)/...

# Peform a development (current-platform-only) build.
dev: version fmt
	go build -o bin/docker-machine-driver-terraform

install: dev
	go install

# Perform a full (all-platforms) build.
build: version build-linux64 build-mac64

build-linux64:
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/linux-amd64/$(BIN_NAME)

build-mac64:
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/darwin-amd64/$(BIN_NAME)

# Produce archives for a GitHub release.
dist: build
	cd $(BIN_DIR)/linux-amd64 && zip -9 ../linux-amd64.zip docker-machine-driver-terraform
	cd $(BIN_DIR)/darwin-amd64 && zip -9 ../darwin-amd64.zip docker-machine-driver-terraform

test: fmt
	go test -v $(REPO_BASE) $(REPO_BASE)/fetch $(REPO_BASE)/terraform

version: $(VERSION_INFO_FILE)

$(VERSION_INFO_FILE): Makefile
	@echo "Update version info: `git describe --tags`."
	@echo "package main" > $(VERSION_INFO_FILE)
	@echo "" >> $(VERSION_INFO_FILE)
	@echo "// DriverVersion is the current version of the Terraform driver for Docker Machine." >> $(VERSION_INFO_FILE)
	@echo "const DriverVersion = \"`git describe --tags` (`git rev-parse HEAD`)\"" >> $(VERSION_INFO_FILE)
