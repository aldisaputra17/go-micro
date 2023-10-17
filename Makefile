# App Env
GO_APP=go
GO_BINARY=bin/server.sh
GO_MAIN_FILE=cmd/server.go
GO_RUNNER_APP=nodemon
GO_RUNNER_EXT=go,tmpl,html

# Lint Env
LINT_APP=golangci-lint

# Migrate Env
MIGRATE_APP=migrate
MIGRATE_FOLDER=database/migrations
MIGRATE_DB_URL=postgresql://plabs:plabs@127.0.0.1:5432/shoping?sslmode=disable

env:
	cp example.env .env
deps:
	$(GO_APP) mod tidy
	$(GO_APP) mod vendor
run:
	$(GO_APP) run $(GO_MAIN_FILE)
dev:
	$(GO_RUNNER_APP) --exec $(GO_APP) run $(GO_MAIN_FILE) --signal SIGTERM -e $(GO_RUNNER_EXT)
build:
	$(GO_APP) build -o $(GO_BINARY) $(GO_MAIN_FILE)

lint:
	@echo -e "==> start linting..."
	$(LINT_APP) run --fix

migrate.create:
	$(MIGRATE_APP) create -ext sql -dir $(MIGRATE_FOLDER) --seq $(name)
migrate.up:
	$(MIGRATE_APP) -path $(MIGRATE_FOLDER) -database $(MIGRATE_DB_URL) --verbose up
migrate.down:
	$(MIGRATE_APP) -path $(MIGRATE_FOLDER) -database $(MIGRATE_DB_URL) --verbose down
migrate.fix:
	$(MIGRATE_APP) -path $(MIGRATE_FOLDER) -database $(MIGRATE_DB_URL) force $(version)

proto:
	rm -f pb/*.go
	protoc --proto_path=src/domain/proto --go_out=src/domain/pb --go_opt=paths=source_relative \
    --go-grpc_out=src/domain/pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=src/domain/pb --grpc-gateway_opt paths=source_relative \
    src/domain/proto/*.proto