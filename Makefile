.DEFAULT_GOAL := build

.PHONY: fmt vet build clean

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build ./...

clean:
	find . -type f -name '*.out' -delete
	find . -type f -name '*.exe' -delete
