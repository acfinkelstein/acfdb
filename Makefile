MAIN_DIR := cmd/acfdb
BINARY_NAME := acfdb

build:
	go build -o $(MAIN_DIR)/$(BINARY_NAME) ./$(MAIN_DIR)/main.go

run:
	cd $(MAIN_DIR) && ./$(BINARY_NAME)