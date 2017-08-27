SHELL := /bin/bash

# The name of the executable (default is current directory name)
TARGET := $(shell echo $${PWD\#\#*/})

# These will be provided to the target
VERSION := 1.0.0
BUILD := `git rev-parse HEAD`
BUILD_TIME=`date +%FT%T%z`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD) -X=main.BuildTime=$BUILD_TIME)"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

all: clean build

$(TARGET): $(SRC)
	@go build $(LDFLAGS) -o $(TARGET)
	@strip $(TARGET)

build: $(TARGET)
	@true

clean:
	@rm -f $(TARGET)

check:
	@gofmt -s -l -w .
	@golint .
	@go tool vet ${SRC}

run: build
	@./$(TARGET)