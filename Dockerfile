FROM golang:1.18-alpine as builder

WORKDIR /go/build
COPY go.mod .
COPY go.sum .
COPY ./ .

RUN go mod download
RUN go build -o api main.go

FROM alpine as deploy
RUN apk --no-cache add tzdata
COPY --from=builder /go/build/api api
ENTRYPOINT ["/api"]
