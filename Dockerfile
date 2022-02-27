# arguments for versions
ARG ALPINE_VERSION=3.15
ARG GO_VERSION=1.18
# argument to override with debug
ARG IMAGE_VERSION=latest

# Get wget for healthcheck
FROM busybox AS wget
ARG BUSYBOX_VERSION=1.31.0-i686-uclibc
ADD https://busybox.net/downloads/binaries/$BUSYBOX_VERSION/busybox_WGET /wget
RUN chmod a+x /wget

FROM golang:${GO_VERSION}-rc-alpine${ALPINE_VERSION} AS build
RUN apk add --no-cache git
WORKDIR /src
COPY ./app/go.mod ./
RUN go mod download
COPY ./app/main.go ./
RUN CGO_ENABLED=0 go build -o /app .

FROM gcr.io/distroless/static:${IMAGE_VERSION} AS app
COPY --from=wget /wget /usr/bin/wget
LABEL maintainer="jamesrwarren"
USER nonroot:nonroot

COPY --from=build --chown=nonroot:nonroot /app /app

ENTRYPOINT ["/app"]