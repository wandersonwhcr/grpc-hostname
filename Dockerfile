FROM golang:1.24-alpine AS builder

RUN apk add --no-cache \
        git \
        protoc \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.5 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

WORKDIR /usr/local/src

COPY . .

RUN go mod download \
    && go build

FROM scratch

COPY --from=builder /usr/local/src/grpc-hostname /

ENTRYPOINT ["/grpc-hostname"]
