FROM golang:1.18-alpine

WORKDIR /go/build
COPY go.mod .
COPY go.sum .
COPY ./ .

RUN go mod download
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]
