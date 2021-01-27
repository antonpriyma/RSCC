ROOT_PATH := $(PWD)
BIN_OUTPUT_PATH := $(ROOT_PATH)/bin

.PHONY: build_api
build_api:
	go build -o $(BIN_OUTPUT_PATH)/api ./cmd


build_docker:
	 docker build -f ./build/docker/Dockerfile . -t antonpriyma/rscc_bot:latest