LINT_PACKAGES = $(shell go list ./app/... )
DATABASE_NAME = appengine_boilerplate

.PHONY: install
install: install/api install/web

.PHONY: install/api
install/api:
	./scripts/tools.sh

.PHONY: install/web
install/web:
	npm --prefix ./web install

.PHONY: build
build: build/api build/web

.PHONY: build/api
build/api:
	go build -o ./bin/main ./main.go

.PHONY: build/web
build/web:
	npm --prefix ./web run build

.PHONY: test
test:
	go test -v $(LINT_PACKAGES)

.PHONY: lint
lint: lint/golint lint/staticcheck lint/errcheck lint/vet

.PHONY: lint/golint
lint/golint:
	golint -set_exit_status $(LINT_PACKAGES)

.PHONY: lint/staticcheck
lint/staticcheck:
	staticcheck $(LINT_PACKAGES)

.PHONY: lint/errcheck
lint/errcheck:
	errcheck $(LINT_PACKAGES)

.PHONY: lint/vet
lint/vet:
	go vet $(LINT_PACKAGES)

.PHONY: proto
proto:
	protoc -I/usr/local/include -I. \
  -I$(GOPATH)/src \
  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
	--go_out=plugins=grpc:. \
	./proto/*.proto

.PHONY: db/drop
db/drop:
	mysql -u root -h 127.0.0.1 -e "DROP DATABASE $(DATABASE_NAME)"

.PHONY: db/create
db/create:
	mysql -u root -h 127.0.0.1 -e "CREATE DATABASE $(DATABASE_NAME)"

.PHONY: db/dump
db/dump:
	mysqldump -u root -h 127.0.0.1 --no-data --skip-add-drop-table --compact --set-gtid-purged=OFF --ignore-table=$(DATABASE_NAME).goose_db_version $(DATABASE_NAME) | grep -v '40101 SET' | grep -v 'character_set_client' > db/schema.sql

.PHONY: db/migrate
db/migrate:
	goose -dir ./db/migrations mysql root@/$(DATABASE_NAME) up

.PHONY: db/model
db/model:
	rm -f ./app/model/xo/*.xo.go
	xo mysql://root@127.0.0.1/$(DATABASE_NAME) -o ./app/model/xo
	rm -f ./app/model/xo/goosedbversion.xo.go
