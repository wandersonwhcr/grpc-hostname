# grpc-hostname

A Simple gRPC Toy Server to Retrieve Hostnames

## Running

```
docker compose build
```

```
docker compose run --rm server \
  protoc \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    proto/hostname.proto
```
