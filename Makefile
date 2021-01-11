.PHONY: all build build-cli migrate rpc templates

all:
	@drone exec

# rpc generators ============================================ [dynamic targets]

rpc: $(shell ls -d rpc/* | sed -e 's/\//./g')
	@echo OK.

rpc.%: SERVICE=$*
rpc.%:
	@echo '> protoc gen for $(SERVICE)'
	@protoc --proto_path=$(GOPATH)/src:. -Irpc/$(SERVICE) --go_out=plugins=grpc,paths=source_relative:. rpc/$(SERVICE)/$(SERVICE).proto
	@protoc --proto_path=$(GOPATH)/src:. -Irpc/$(SERVICE) --twirp_out=paths=source_relative:. rpc/$(SERVICE)/$(SERVICE).proto


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
	@./templates/db_schema.go.sh
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
migrate.%: export MYSQL_ROOT_PASSWORD = default
migrate.%:
	@echo migrate.$(SERVICE)
	mysql -h mysql-test -u root -p$(MYSQL_ROOT_PASSWORD) -e "CREATE DATABASE $(SERVICE);"
	./build/db-migrate-cli-linux-amd64 -service $(SERVICE) -db-dsn "root:$(MYSQL_ROOT_PASSWORD)@tcp(mysql-test:3306)/$(SERVICE)" -real=true
	./build/db-migrate-cli-linux-amd64 -service $(SERVICE) -db-dsn "root:$(MYSQL_ROOT_PASSWORD)@tcp(mysql-test:3306)/$(SERVICE)" -real=true
	./build/db-schema-cli-linux-amd64 -schema $(SERVICE) -db-dsn "root:$(MYSQL_ROOT_PASSWORD)@tcp(mysql-test:3306)/$(SERVICE)" > server/$(SERVICE)/types_db.go
