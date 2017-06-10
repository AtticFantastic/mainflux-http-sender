SHELL := /bin/bash

# Git stuff
REV := $(shell git rev-parse HEAD)
CHANGES := $(shell test -n "$$(git status --porcelain)" && echo '+CHANGES' || true)

# Target
TARGET := http-sender
VERSION := $(shell cat VERSION)

OS := linux
ARCH := amd64

# LDFLAGS
LDFLAGS := -s -w -extldflags "-static"
LDFLAGS += 	-X "github.com/mainflux/mainflux-http-sender/main.tag=$(TAG)" \
			-X "github.com/mainflux/mainflux-http-sender/main.utcTime=$(shell date -u '+%Y/%m/%d %H:%M:%S')" \
			-X "github.com/mainflux/mainflux-http-sender/main.rev=$(REV)" \
			-X "github.com/mainflux/mainflux-http-sender/main.version=$(VERSION)"

build:
	go build -v \
		-ldflags '$(LDFLAGS)' \
	   	-o "$(TARGET)" .

clean:
	go clean
	rm -rf $(TARGET)
