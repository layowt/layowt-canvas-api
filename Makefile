# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Main package name
MAIN_PACKAGE = main.go

# Output binary name
BINARY_NAME = myapp

# Build targets
all: clean build
build:
	$(GOBUILD) -o build/$(BINARY_NAME) $(MAIN_PACKAGE)
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
test:
	$(GOTEST) -v ./...
run:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PACKAGE)
	./$(BINARY_NAME)