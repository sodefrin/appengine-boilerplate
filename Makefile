LINT_PACKAGES = $(shell go list ./app/... )

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
