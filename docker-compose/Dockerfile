FROM golang:1.24-alpine

RUN apk add --no-cache \
        git \
        protoc \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.5 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
