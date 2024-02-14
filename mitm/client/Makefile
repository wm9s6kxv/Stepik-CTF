.PHONY: all build clean

APP_NAME=client
BUILD_DIR=build

all: clean build

build:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)_linux_amd64
	GOOS=linux GOARCH=arm go build -o $(BUILD_DIR)/$(APP_NAME)_linux_arm
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)_darwin_amd64
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)_windows_amd64.exe

clean:
	rm -rf $(BUILD_DIR)

