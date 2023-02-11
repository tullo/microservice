# Microservice

Add .proto file. Scaffold Go code. Implement gRPC service. Done.

```sh
1. make
2. make docker
3. docker-compose up -d
4. ./test-service.sh
5. ./test-benchmark-service.sh
```

## Install Drone CI

https://docs.drone.io/cli/install/

```sh
curl -L https://github.com/harness/drone-cli/releases/latest/download/drone_linux_amd64.tar.gz | tar zx
sudo install -t /usr/local/bin drone && rm drone
```

## Build and push dev image

```sh
cd docker/build && make && make push
```

## SQL batch insert

Get latest sqlx module version that supports batch insert.

```sh
go get github.com/jmoiron/sqlx@master
```

## Scaffold and compile using Drone CI

```sh
make
```

## Build Go binaries and Docker service images

```sh
# install figlet
sudo apt install figlet
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
