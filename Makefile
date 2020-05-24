
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_LOC=bin
CORTO_BINARY_NAME=corto
CORTOCTL_BINARY_NAME=cortoctl
PROJECT_HOME=$(shell pwd)

all: test build
build: 
	$(GOBUILD) -o ./$(BINARY_LOC)/$(CORTO_BINARY_NAME) -v ./cmd/$(CORTO_BINARY_NAME)/...
	$(GOBUILD) -o ./$(BINARY_LOC)/$(CORTOCTL_BINARY_NAME) -v ./cmd/$(CORTOCTL_BINARY_NAME)/...
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -rf $(BINARY_LOC)
run: build
	./$(BINARY_LOC)/$(CORTO_BINARY_NAME)