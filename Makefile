ROOT_PATH := $(PWD)
BIN_OUTPUT_PATH := $(ROOT_PATH)/bin

.PHONY: build_api
build_api:
	go build -o $(BIN_OUTPUT_PATH)/bot ./cmd/bot


build_docker_bot:
	 docker build --platform=linux/amd64 -f ./build/docker/bot/Dockerfile . -t antonpriyma/rscc_bot:latest
    docker push antonpriyma/rscc_bot:latest


.PHONY: build_morning_greeter
build_morning_greeter:
	go build -o $(BIN_OUTPUT_PATH)/tasks/morning_greeter ./cmd/tasks/morning_greeter


build_docker_morning_greeter:
	 docker build --platform=linux/amd64 -f ./build/docker/morning_greeter/Dockerfile . -t antonpriyma/rscc_morning_greeter:latest
	 docker push antonpriyma/rscc_morning_greeter:latest