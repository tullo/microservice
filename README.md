# Microservice

Add .proto file. Scaffold Go code. Implement gRPC service. Done. 

```console
1. make
2. make docker
3. docker-compose up -d
4. ./test-service.sh
5. ./test-benchmark-service.sh
```

## Build and push dev image

```console
cd docker/build && make && make push
```

## SQL batch insert

Get latest sqlx module version that supports batch insert.

```console
go get github.com/jmoiron/sqlx@master
```

## Scaffold and compile using Drone CI

```console
make
```

## Build Go binaries and Docker service images

```console
# build all docker images
make build && make docker
# build specific docker image
make build && make docker.stats
```

## Elastic APM

```console
sudo sysctl vm.max_map_count=512000
docker-compose build
docker-compose up -d
docker-compose logs -f apm
```
