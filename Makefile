all: build

$(shell if [ ! -d bin ]; then mkdir bin; fi )

build:
	go mod tidy
	@GOOS=linux go build -ldflags '-linkmode "external" -extldflags "-static"' -o bin/gofound main.go
