.PHONY: all clean build clean-url

all: build
	@echo "Build complete."

build:
	go clean
	go build -ldflags="-H windowsgui -w -s -buildid=" -trimpath -o bin/cookie.exe ./cmd/main.go