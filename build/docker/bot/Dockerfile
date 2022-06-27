# Base docker image
FROM golang:latest

# Build bin
WORKDIR /rscc_bot
COPY .. .
RUN make build_api

# Command to start container
CMD ["./bin/bot"]