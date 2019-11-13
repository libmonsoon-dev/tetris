-include .env

GO_CMD=go
GO_MOBILE_CMD=gomobile
PROJECT_NAME := $(shell basename "$(PWD)")
VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
LD_FLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
BUILD_FOLDER = build
MOBILE_EMULATOR_BIN = $(BUILD_FOLDER)/$(PROJECT_NAME)-desktop
# MAKEFLAGS += --silent

all: build-folder build-cli build-android build-web build-desktop

build-folder:
	@mkdir -p $(BUILD_FOLDER)

dependency:
	$(GO_CMD) mod download

generate:
	go generate ./...

pre-build: build-folder dependency generate

build-cli: pre-build
	@echo 'Building console'

build-android: pre-build
	@echo 'Building Android'

build-web: pre-build
	@echo 'Building web'

build-mobile-emulator: pre-build
	@echo 'Building mobile version for desktop'
	$(GO_CMD) build $(LD_FLAGS) -o $(MOBILE_EMULATOR_BIN) entrypoint/mobile/tetris.go

run-mobile-emulator: build-mobile-emulator
	$(MOBILE_EMULATOR_BIN)

tools:
	$(GO_CMD) get golang.org/x/mobile/cmd/gomobile \
		github.com/sqs/goreturns
dev:
	$(GO_CMD) run entrypoint/dev/tetris.go

clear:
	rm -rf $(BUILD_FOLDER)
