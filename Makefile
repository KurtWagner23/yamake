# Define the binary name
BINARY = yamake

TEST_PROJECT_DIR = test_project

SRC_DIR = ./src

build:
	@echo "Building the binary..."
	go build -o $(BINARY) $(SRC_DIR)
	mv $(BINARY) $(TEST_PROJECT_DIR)

run: build
	@echo "Running the binary..."
	cd $(TEST_PROJECT_DIR) && ./$(BINARY)

run_ninja: build
	@echo "Building and running the binary with --build flag..."
	cd $(TEST_PROJECT_DIR) && ./$(BINARY) --build
