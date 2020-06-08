FROM golang:1.12.4 as builder

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . /app

RUN go build -o link

FROM buildpack-deps:buster-scm

WORKDIR /app

COPY --from=builder /app/link /app/link

EXPOSE 5333

ENV TINI_VERSION v0.19.0

ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini

RUN chmod +x /tini

ENTRYPOINT ["/tini", "--"]

CMD ["/app/link"]