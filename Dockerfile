# Builder
FROM golang:1.11.3-alpine3.8 as builder

LABEL Maintainer="Julien BREUX <julien@golang.fr>"

ARG VERSION="dev"
ARG COMMIT="none"
ARG DATE="unknown"

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE=on

RUN apk --update --no-cache add \
    ca-certificates=20171114-r3 \
    git=2.18.1-r0 \
    openssh-client=7.7_p1-r3

WORKDIR /go/src/github.com/julienbreux/mila/
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN FLAGS=$(printf '-X main.version=%s -X main.commit=%s -X main.date=%s' "$VERSION" "$COMMIT" "$DATE") && \
    export FLAGS && \
    go build \
    -ldflags "$FLAGS" \
    -o mila ./cmd/mila

# Application
FROM alpine:3.8

LABEL maintainer="Julien Breux <julien@golang.fr>"

ENV PATH=/bin

WORKDIR /bin/

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/julienbreux/mila/mila .

ENTRYPOINT [ "mila" ]
