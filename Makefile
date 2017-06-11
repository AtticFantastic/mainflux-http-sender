SHELL := /bin/bash

# Git stuff
REV := $(shell git rev-parse HEAD)
TAG := $(shell git describe --tags --exact-match 2> /dev/null || git rev-parse --short HEAD)

# Target
TARGET := http-sender

# LDFLAGS
LDFLAGS := -s -w -extldflags "-static"
LDFLAGS += 	-X "main.Tag=$(TAG)" \
			-X "main.Time=$(shell date -u '+%Y/%m/%d %H:%M:%S')" \
			-X "main.Revision=$(REV)" \

build:
	CGO_ENABLED=0 go build -v \
		-ldflags '$(LDFLAGS)' \
	   	-o "$(TARGET)" .

clean:
	go clean
	rm -rf $(TARGET)
