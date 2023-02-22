ARG GO_VERSION=1.19
ARG ALPINE_VERSION=3.16
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder
RUN apk update && apk add --no-cache git
ENV GOPROXY=direct GOSUMDB=off
WORKDIR /go/src/project
ADD . .
RUN env CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o tranches .

FROM alpine:${ALPINE_VERSION}
COPY --from=builder /go/src/project/tranches /
ENTRYPOINT ["/tranches"]
