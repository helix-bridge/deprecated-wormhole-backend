ARG GO_VERSION=1.12.4

FROM golang:${GO_VERSION} as builder
WORKDIR /go/src/github.com/darwinia-network/link

ENV GO111MODULE=off
# COPY go.mod go.sum ./
# RUN go get -v ./...

COPY . .
RUN go build -o /link

FROM buildpack-deps:buster-scm

WORKDIR /app
COPY ./config ./config
COPY --from=builder /link ./link

EXPOSE 5333
CMD ["/app/link"]
