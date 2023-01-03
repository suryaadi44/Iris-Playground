FROM golang:1.19-alpine AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

# System dependencies
RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    && update-ca-certificates

### Development with hot reload and debugger
FROM base AS dev
WORKDIR /app

# Hot reloading mod
RUN go install github.com/cosmtrek/air@latest && go install github.com/go-delve/delve/cmd/dlv@latest
EXPOSE 9000
EXPOSE 2346

ENTRYPOINT ["air", "-c", ".air.grpc.toml"]

### Executable builder
FROM base AS builder
WORKDIR /app

# Application dependencies
COPY . /app
RUN go mod download \
    && go mod verify

RUN go build -o built_program -a .

### Production
FROM alpine:latest

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    && update-ca-certificates

# Copy executable
COPY --from=builder /app/built_program /usr/local/bin/built_program
EXPOSE 9000

ENTRYPOINT ["/usr/local/bin/built_program"]
