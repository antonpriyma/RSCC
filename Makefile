ROOT_PATH := $(PWD)
BIN_OUTPUT_PATH := $(ROOT_PATH)/bin

.PHONY: build_api
build_api:
	go build -o $(BIN_OUTPUT_PATH)/bot ./cmd/bot


build_docker_ubuntu:
	 docker build --platform=linux/amd64 -f ./build/docker/Dockerfile . -t antonpriyma/rscc_bot:latest