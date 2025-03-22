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

```
docker compose run --rm server \
    go mod download
```

```
docker compose up --detach
```

```
docker compose up client --detach
```

```
docker compose logs client server
```

## Build and Push

```
docker build . \
    --tag ghcr.io/wandersonwhcr/grpc-hostname:`git short` \
    --target release
```

```
docker push ghcr.io/wandersonwhcr/grpc-hostname:`git short`
```

## Environment Variables

* `HOSTNAME_CMD`
* `HOSTNAME_SERVER_ADDR`
