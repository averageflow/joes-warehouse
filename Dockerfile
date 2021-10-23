FROM golang:1.17.2-alpine3.14

# Creating the `app` directory in which the app will run 
RUN mkdir /app

# Move everything from root to the newly created app directory
ADD . /app

# Specifying app as our work directory in which
# futher instructions should run into
WORKDIR /app

# Download all neededed project dependencies
RUN go mod download

# Build the project executable binary
RUN go build -o main ./cmd/joes-warehouse

# Run/Starts the app executable binary
CMD ["/app/main"]