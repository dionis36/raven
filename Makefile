BINARY_NAME=raven

.PHONY: all build clean test run release

all: build

build:
	go build -o $(BINARY_NAME) cmd/raven/main.go

run:
	go run cmd/raven/main.go

test:
	go test ./...

clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)-linux
	rm -f $(BINARY_NAME)-windows.exe
	rm -f $(BINARY_NAME)-macos

# Cross compilation
release:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux cmd/raven/main.go
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME)-windows.exe cmd/raven/main.go
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-macos cmd/raven/main.go
