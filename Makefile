SHELL := /bin/bash

# Git stuff
REV := $(shell git rev-parse HEAD)
TAG := 123


# Target
TARGET := http-sender
VERSION := $(shell cat VERSION)

OS := linux
ARCH := amd64

# LDFLAGS
LDFLAGS := -s -w -extldflags "-static"
LDFLAGS += 	-X "main.tag=$(TAG)" \
			-X "main.utcTime=$(shell date -u '+%Y/%m/%d %H:%M:%S')" \
			-X "main.rev=$(REV)" \
			-X "main.version=$(VERSION)"

build:
	go build -v \
		-ldflags '$(LDFLAGS)' \
	   	-o "$(TARGET)" .

clean:
	go clean
	rm -rf $(TARGET)
