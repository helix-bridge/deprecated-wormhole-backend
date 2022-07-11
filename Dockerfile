ARG GO_VERSION=1.18.3

FROM golang:${GO_VERSION} as builder
WORKDIR /go/src/github.com/darwinia-network/link

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /link

FROM buildpack-deps:buster-scm

WORKDIR /app
COPY ./config ./config
COPY --from=builder /link ./link

EXPOSE 5333
CMD ["/app/link"]
