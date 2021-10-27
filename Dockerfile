FROM golang:1.17.2-alpine3.14

# Creating the `app` directory in which the app will run 
RUN mkdir /app

# Move everything from root to the newly created app directory
ADD . /app

COPY ./scripts/wait-for.sh /wait-for.sh

# Specifying app as our work directory in which
# futher instructions should run into
WORKDIR /app

ENV CGO_ENABLED=0

# Download all neededed project dependencies
RUN go mod download

# Run all unit tests
RUN go test -shuffle=on ./...

# Build the project executable binary
RUN go build -o main ./cmd/joes-warehouse

EXPOSE 7000/tcp

# Run/Starts the app executable binary
# CMD ["/app/main"]