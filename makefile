APP_NAME := words
SRC := ./cmd/start/main.go

.PHONY: all build run fmt tidy clean

all: fmt tidy build

fmt:
	go fmt ./...

tidy:
	go mod tidy

build:
	go build -o $(APP_NAME) $(SRC)

run: build
	./$(APP_NAME)

clean:
	rm -f $(APP_NAME)
