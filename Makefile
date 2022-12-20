VERSION := $(shell git describe --tags --always)
BUILD := go build -v -ldflags "-s -w -X main.Version=$(VERSION)"

.PHONY : all test

all: test

test:
	go test -v ./perfdata/...
	go test -v ./thresholds/...


