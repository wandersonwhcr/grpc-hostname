FROM golang:1.24-alpine AS base

RUN apk add --no-cache \
        git \
        protoc \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.5 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

# ------------------------------------------------------------------------------

FROM base AS build

WORKDIR /usr/local/src

COPY . .

RUN go mod download \
    && go build

# ------------------------------------------------------------------------------

FROM scratch AS release

LABEL org.opencontainers.image.title="grpc-hostname" \
      org.opencontainers.image.description="A Simple gRPC Toy Server to Retrieve Hostnames" \
      org.opencontainers.image.source="https://github.com/wandersonwhcr/grpc-hostname"

COPY --from=build /usr/local/src/grpc-hostname /

ENTRYPOINT ["/grpc-hostname"]
