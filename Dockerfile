FROM golang:1.12.4 as builder

COPY . /go/src/github.com/darwinia-network/link

ENV GO111MODULE=off

WORKDIR /go/src/github.com/darwinia-network/link

RUN go build -o link

FROM buildpack-deps:buster-scm

WORKDIR /app

COPY --from=builder go/src/github.com/darwinia-network/link/link /app/link

EXPOSE 5333

ENV TINI_VERSION v0.19.0

ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini

RUN chmod +x /tini

ENTRYPOINT ["/tini", "--"]

CMD ["/app/link"]