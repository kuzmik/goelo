SHELL := /bin/bash

# The name of the executable (default is current directory name)
TARGET := $(shell echo $${PWD\#\#*/})

# These will be provided to the target
VERSION := 0.1.0
BUILD := `git rev-parse HEAD`
BUILD_TIME=`date +%FT%T%z`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD) -X=main.BuildTime=$BUILD_TIME)"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./experiments/*")

all: clean build linux

clean:
	@rm -f $(TARGET) $(TARGET)_linux

check:
	@gofmt -l -s -w .
	@golint .
	@go tool vet ${SRC}

build:
	@go build $(LDFLAGS) -o $(TARGET) $(SRC)
	@strip $(TARGET)

linux:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(TARGET)_linux $(SRC)

.PHONY: build check clean linux