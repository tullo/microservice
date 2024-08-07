#!/bin/bash -eo pipefail
.PHONY: all build build-cli migrate rpc templates

all:
	drone exec --pipeline build .drone.yml
#	drone exec --pipeline build --resume-at tidy

chown:
	@sudo chown -R ${USER}:${USER} build/ client/ cmd/ db/ docs/ docker/dev/ js/ rpc/ server/
	@echo OK.

lint:
	golangci-lint run -E depguard,dupl,errcheck,errorlint,forbidigo,gocritic,godot,gosec,nakedret,nlreturn,misspell,sqlclosecheck,whitespace,wrapcheck ./...
#	golint -set_exit_status ./...


tidy:
	go mod tidy
	go fmt ./...


# rpc generators ============================================ [dynamic targets]

rpc: $(shell ls -d rpc/* | sed -e 's/\//./g')
	@echo OK.

rpc.%: SERVICE=$*
rpc.%:
	@echo '> protoc gen for $(SERVICE)'
	@mkdir -p js
	@protoc --proto_path=rpc/$(SERVICE) --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative rpc/$(SERVICE)/$(SERVICE).proto
	@protoc --proto_path=rpc/$(SERVICE) --twirp_out=paths=source_relative:. rpc/$(SERVICE)/$(SERVICE).proto
	@protoc --proto_path=rpc/$(SERVICE) --twirp_swagger_out=js --twirp_js_out=js --js_out=import_style=commonjs,binary:js $(SERVICE).proto

# build cmd/ go binaries ==================================== [dynamic targets]

build: export GOOS = linux
build: export GOARCH = amd64
build: export CGO_ENABLED = 0
build: $(shell ls -d cmd/* | grep -v "\-cli" | sed -e 's/cmd\//build./')
	@echo OK.

build.%: SERVICE=$*
build.%:
	go build -o build/$(SERVICE)-$(GOOS)-$(GOARCH) ./cmd/$(SERVICE)/*.go


# code generator for client/server/cmd ====================== [dynamic targets]

templates: export MODULE=$(shell grep ^module go.mod | sed -e 's/module //g')
templates: $(shell ls -d rpc/* | sed -e 's/rpc\//templates./g')
	@rm db/schema_*.go db/schema.go
	@./templates/db_schema.go.sh
	@./templates/client_wire.go.sh
	@echo OK.

templates.%: export SERVICE=$*
templates.%: export SERVICE_CAMEL=$(shell echo $(SERVICE) | sed -r 's/(^|_)([a-z])/\U\2/g')
templates.%: export MODULE=$(shell grep ^module go.mod | sed -e 's/module //g')
templates.%:
	@echo templates: $(SERVICE) $(MODULE)
	@mkdir -p cmd/$(SERVICE) client/$(SERVICE) server/$(SERVICE)
	@echo "~ cmd/$(SERVICE)/main.go"
	@envsubst < templates/cmd_main.go.tpl > cmd/$(SERVICE)/main.go
	@echo "~ client/$(SERVICE)/client.go"
	@envsubst < templates/client_client.go.tpl > client/$(SERVICE)/client.go
	@./templates/server_server.go.sh
	@./templates/server_wire.go.sh


# build cli tooling from cmd/ =============================== [dynamic targets]

build-cli: export GOOS = linux
build-cli: export GOARCH = amd64
build-cli: export CGO_ENABLED = 0
build-cli: $(shell ls -d cmd/*-cli | sed -e 's/cmd\//build-cli./')
	@echo OK.

build-cli.%: SERVICE=$*
build-cli.%:
	go build -o build/$(SERVICE)-$(GOOS)-$(GOARCH) ./cmd/$(SERVICE)/*.go




# database migrations ======================================= [dynamic targets]

migrate: $(shell find . -type f -regex ".*migrations.sql" | xargs -n1 -r dirname | sed -e 's/db.schema./migrate./')
	@echo OK.

# We run the migrations twice, so we make sure that our migration status is logged correctly as well.
# All the migrations in the second run must be skipped.

migrate.%: export SERVICE = $*
migrate.%: DSN = "migrations:migrations@tcp(mysql-test:3306)/migrations"
migrate.%:
	@echo migrate.$(SERVICE)
	./build/db-migrate-cli-linux-amd64 -service $(SERVICE) -db-dsn $(DSN) -real=true
	./build/db-migrate-cli-linux-amd64 -service $(SERVICE) -db-dsn $(DSN) -real=true
	@mkdir -p server/$(SERVICE)
	@find server/$(SERVICE) -name types_gen.go -delete
	@rm -rf docs/schema/$(SERVICE)
	./build/db-schema-cli-linux-amd64 -service $(SERVICE) -schema migrations -db-dsn $(DSN) -format go -output server/$(SERVICE)
	./build/db-schema-cli-linux-amd64 -service $(SERVICE) -schema migrations -db-dsn $(DSN) -format markdown -output docs/schema/$(SERVICE)
	./build/db-schema-cli-linux-amd64 -schema migrations -db-dsn $(DSN) -drop=true



# docker image build ======================================== [dynamic targets]

IMAGE_PREFIX := tullo/service-

docker: $(shell ls -d cmd/* | sed -e 's/cmd\//docker./')
	@echo IMAGE_PREFIX=$(IMAGE_PREFIX) > .env
	@echo OK.

docker.%: export SERVICE = $(shell basename $*)
docker.%:
	@figlet $(SERVICE)
	docker build --rm --no-cache -t $(IMAGE_PREFIX)$(SERVICE) --build-arg service_name=$(SERVICE) -f docker/serve/Dockerfile .

# docker image push========================================== [dynamic targets]

push: $(shell ls -d cmd/* | sed -e 's/cmd\//push./')
	@echo OK.

push.%: export SERVICE = $(shell basename $*)
push.%:
	@figlet $(SERVICE)
	docker push $(IMAGE_PREFIX)$(SERVICE)
