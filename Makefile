.PHONY: build clean test package package-deb ui api statics requirements ui-requirements serve update-vendor internal/statics internal/migrations static/swagger/api.swagger.json
PKGS := $(shell go list ./... | grep -v /vendor |grep -v smart-ac/api | grep -v /migrations | grep -v /static | grep -v /ui)
VERSION := $(shell git describe --always |sed -e "s/^v//")

build: ui/build internal/statics internal/migrations
	mkdir -p build
	go build $(GO_EXTRA_BUILD_ARGS) -ldflags "-s -w -X main.version=$(VERSION)" -o build/smart-ac cmd/main.go

clean:
	@echo "Cleaning up workspace"
	@rm -rf build dist internal/migrations/migrations_gen.go internal/static/static_gen.go ui/build static/static
	@rm -f static/index.html static/icon.png static/manifest.json static/asset-manifest.json static/service-worker.js
	@rm -rf static/logo
	@rm -rf static/js
	@rm -rf docs/public
	@rm -rf dist

test: internal/statics internal/migrations
	@echo "Running tests"
	@rm -f coverage.out
	@for pkg in $(PKGS) ; do \
		golint $$pkg ; \
	done
	@go vet $(PKGS)
	@go test -p 1 -v $(PKGS) -failfast -cover -coverprofile coverage.out -timeout 45m

ui/build:
	@echo "Building ui"
	@cd ui && npm run build
	@mv ui/build/* static

api:
	@echo "Generating API code from .proto files"
	@go mod vendor
	@go generate api/api.go
	@rm -rf vendor/

internal/statics internal/migrations: static/swagger/api.swagger.json
	@echo "Generating static files"
	@go generate internal/migrations/migrations.go
	@go generate internal/static/static.go


static/swagger/api.swagger.json:
	@echo "Generating combined Swagger JSON"
	@GOOS="" GOARCH="" go run api/swagger/main.go api/swagger > static/swagger/api.swagger.json
	@cp api/swagger/*.json static/swagger


# shortcuts for development

dev-requirements:
	go mod download
	go install golang.org/x/lint/golint
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go install github.com/golang/protobuf/protoc-gen-go
	go install github.com/elazarl/go-bindata-assetfs/go-bindata-assetfs
	go install github.com/jteeuwen/go-bindata/go-bindata
	go install golang.org/x/tools/cmd/stringer
	go install github.com/goreleaser/nfpm
	
ui-requirements:
	@echo "Installing UI requirements"
	@cd ui && npm install

serve: build
	@echo "Starting Lora App Server"
	./build/smart-ac

update-vendor:
	@echo "Updating vendored packages"
	@govendor update +external
