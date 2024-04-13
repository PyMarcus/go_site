# Variáveis
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=go_booking

CMD_DIR=./cmd

all: test build

build:
    $(GOBUILD) -o $(BINARY_NAME) $(CMD_DIR)/*.go

test:
    $(GOTEST) -v ./...

clean:
    $(GOCLEAN)
    rm -f $(BINARY_NAME)

run:
    $(GOBUILD) -o $(BINARY_NAME) $(CMD_DIR)/*.go
    ./$(BINARY_NAME)

.PHONY: all build test clean run
