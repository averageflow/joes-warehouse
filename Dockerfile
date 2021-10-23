FROM golang:1.17.2-alpine3.14
RUN mkdir /dist
ADD . /dist
WORKDIR /dist
# Pull in any dependencies
RUN go mod download
# Build project as a binary called app
RUN go build -o app .

CMD ["/dist/app"]
