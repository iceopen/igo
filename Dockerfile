FROM golang:1.17-alpine as builder

RUN apk add --no-cache \
    wget \
    make \
    git \
    gcc \
    binutils-gold \
    musl-dev

RUN mkdir -p /go/src/github.com/iceopen/igo
WORKDIR /go/src/github.com/iceopen/igo

# Cache dependencies
ENV GOPROXY=https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod tidy

# Build real binaries
COPY . .
RUN make

# Executable image
FROM alpine:3.18.3

RUN apk add --no-cache \
    curl

COPY --from=builder /go/src/github.com/iceopen/igo/bin/igo /igo

WORKDIR /

ENTRYPOINT ["/igo"]
