-include .env

GO_CMD=go
GO_MOBILE_CMD=gomobile
PROJECT_NAME := $(shell basename "$(PWD)")
VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
LD_FLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD) -X=main.BuildTimestamp=$(shell date -u "+%Y-%m-%d %H:%M:%S")"
BUILD_FOLDER = build
MOBILE_EMULATOR_BIN = $(BUILD_FOLDER)/$(PROJECT_NAME)-desktop
# MAKEFLAGS += --silent

all: tools build-folder build-cli build-android build-web build-mobile-emulator

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
	#TODO: specify version
	$(GO_CMD) install golang.org/x/mobile/cmd/gomobile \
		github.com/sqs/goreturns \
		github.com/cheekybits/genny \
		github.com/codegangsta/gin \

dev: pre-build tools
	gin -i --build src/cmd/dev --bin build/dev  --buildArgs="-tags dev" run tetris.go

coverage: pre-build
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

clear:
	rm -rf $(BUILD_FOLDER)
