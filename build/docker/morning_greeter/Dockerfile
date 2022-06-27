# Base docker image
FROM golang:latest

# Build bin
WORKDIR /rscc_morning_greeter
COPY .. .
RUN make build_morning_greeter

# Command to start container
CMD ["./bin/tasks/morning_greeter"]