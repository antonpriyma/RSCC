# Base docker image
FROM golang:latest as builder

# Install all dependency
WORKDIR /rscc_bot
COPY . .
RUN make build_api

# Command to start container
CMD ["./bin/api"]