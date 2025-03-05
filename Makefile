# Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
.PHONY: help clean local run run-defaults validate generate generate-openapi generate-client generate-redfish generate-axios webui-dist docker-image fmt test-go vet-go test-go-backend

APP_NAME := cfm-service
CLIAPP_NAME := cfm-cli
CONF_NAME := cfm-service.conf
OPENAPI_YAML := api/cfm-openapi.yaml
OPENAPI_REDFISH_YAML := api/redfish-openapi.yaml
GO_VERSION := 1.23.7
GO_INSTALL_LOCATION := /usr/local/go/bin/
GOFMT_OPTS := $(GO_INSTALL_LOCATION)gofmt -w /local
GENERATE_USER := $(shell id -u ${USER}):$(shell id -g ${USER})

help:
	@echo ""
	@echo "-----------------------------------------------------------------------------------"
	@echo "make clean-service        - Remove local $(APP_NAME) executable"
	@echo "make clean-cli            - Remove local $(CLIAPP_NAME) executable"
	@echo "make clean-go             - Remove all Go executables"
	@echo "make build-service        - Build local $(APP_NAME) executable"
	@echo "make build-cli            - Build local $(CLIAPP_NAME) executable"
	@echo "make build-go             - Build local $(APP_NAME) and $(CLIAPP_NAME) executables"
	@echo "make build-docker-cfm     - Build a local docker image for the cfm software suite"
	@echo "make run-service          - Build and run a local $(APP_NAME) executable using config file $(CONF_NAME)"
	@echo "make run-service-defaults - Build and run a local $(APP_NAME) executable using its' internal config defaults"
	@echo "make validate             - Validate the openapi specification using openapi-generator-cli"
	@echo "make generate             - Generate supporting code for the whole suite using openapi-generator-cli"
	@echo "make generate-openapi     - Generate go server code for the openapi specification using openapi-generator-cli"
	@echo "make generate-client      - Generate go client code for the openapi specification using openapi-generator-cli"
	@echo "make generate-redfish     - Generate go server code for the redfish specification using openapi-generator-cli"
	@echo "make generate-axios       - Generate axios server code for the openapi specification using openapi-generator-cli"
	@echo "make fmt-go               - Run gofmt"
	@echo "make test-go              - Run all Go tests"
	@echo "make vet-go               - Run go vet"
	@echo "make test-go-backend      - Run Go unit tests on the backend go code"
	@echo ""

clean-service:
	@echo "Clean up cfm-service..."
	go clean
	rm -f $(APP_NAME)

clean-cli:
	@echo "Clean up cfm-cli..."
	go clean
	rm -f $(CLIAPP_NAME)

clean-go:
	@echo "Clean up..."
	go clean
	rm -f $(APP_NAME) $(CLIAPP_NAME)

build-service: clean-service
	@echo "Building cfm-service executable..."
	go build -o $(APP_NAME) -ldflags "-X main.buildTime=`date -u '+%Y-%m-%dT%H:%M:%S'`" ./cmd/$(APP_NAME)/main.go
	ls -lh $(APP_NAME)

build-cli: clean-cli
	@echo "Building cfm-cli executable..."
	go build -o $(CLIAPP_NAME) -ldflags "-X main.buildTime=`date -u '+%Y-%m-%dT%H:%M:%S'`" ./cmd/$(CLIAPP_NAME)/main.go
	ls -lh $(CLIAPP_NAME)

build-go: clean-go
	@echo "Building cfm Go executables..."
	go build -o $(APP_NAME) -ldflags "-X main.buildTime=`date -u '+%Y-%m-%dT%H:%M:%S'`" ./cmd/$(APP_NAME)/main.go
	go build -o $(CLIAPP_NAME) -ldflags "-X main.buildTime=`date -u '+%Y-%m-%dT%H:%M:%S'`" ./cmd/$(CLIAPP_NAME)/main.go
	ls -lh $(APP_NAME) $(CLIAPP_NAME)

build-docker-cfm:
	@echo "Building local docker image of cfm software suite..."
	docker build --no-cache -t cfm -f docker/Dockerfile .

install-webui-dist: build-webui-dist
	@echo "Installing webui distro into $(APP_NAME)..."
	mkdir -p ./services/webui/dist
	cp ./webui/dist ./services/webui/dist

run-service: build-service
	@echo "Running $(APP_NAME) using config file $(CONF_NAME) (Ctrl-C to stop)"
	./$(APP_NAME) -config $(CONF_NAME)

run-service-defaults: build-service
	@echo "Running $(APP_NAME) using config defaults (Ctrl-C to stop)"
	./$(APP_NAME)

validate:
	@echo "Validating $(OPENAPI_YAML) using openapi-generator-cli"
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.0.0 version
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.0.0 validate -i /local/$(OPENAPI_YAML)

generate-openapi:
	@echo "Clean up openapi server generated code (service)"
	rm -rf ./pkg/openapi/
	@echo "Generating $(OPENAPI_YAML) go server using openapi-generator-cli"
	docker run -u $(GENERATE_USER) --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.0.0 generate -i /local/$(OPENAPI_YAML) -g go-server -o /local/pkg/openapi --additional-properties=sourceFolder=,outputAsLibrary=true,router=mux,serverPort=8080,enumClassPrefix=true -t /local/api/templates/go-server --skip-validate-spec
	@echo "Format files after generation to conform to project standard"
	docker run --rm -v ${PWD}:/local golang:$(GO_VERSION) $(GOFMT_OPTS)

generate-client:
	@echo "Clean up openapi client generated code (service)"
	rm -rf ./pkg/client/
	@echo "Generating $(OPENAPI_YAML) go client using openapi-generator-cli"
	docker run -u $(GENERATE_USER) --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.0.0 generate -i /local/$(OPENAPI_YAML) -g go -o /local/pkg/client -p isGoSubmodule=true,withGoMod=false --package-name client --ignore-file-override /local/api/ignore/.openapi-generator-ignore-client -t /local/api/templates/go
	# workaround for withGoMod=false not functioning with openapi-generator
	rm pkg/client/go.mod
	rm pkg/client/go.sum
	rm -rf pkg/client/test # not currently used and have conflict with generated path (github.com/GIT_USER_ID/GIT_REPO_ID/client)

	@echo "Format files after generation to conform to project standard"
	docker run --rm -v ${PWD}:/local golang:$(GO_VERSION) $(GOFMT_OPTS)

generate-redfish:
	@echo "Clean up redfishapi generated code (service)"
	rm -rf ./pkg/redfishapi/
	@echo "Generating $(OPENAPI_YAML) redfish server using openapi-generator-cli"
	docker run -u $(GENERATE_USER) --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.0.0 generate -i /local/$(OPENAPI_REDFISH_YAML) -g go-server -o /local/pkg/redfishapi --package-name redfishapi --additional-properties=sourceFolder=,outputAsLibrary=true,router=mux,serverPort=8080,enumClassPrefix=true -t /local/api/templates/go-server --skip-validate-spec
	@echo "Format files after generation to conform to project standard"
	docker run --rm -v ${PWD}:/local golang:$(GO_VERSION) $(GOFMT_OPTS)
	@echo "Apply local patch for redfish auto generated codes"
	git apply api/patch/*.redfish.patch

generate-axios:
	@echo "Clean up axios generated code (webui)"
	rm -rf ./webui/src/axios/
	@echo "Generating $(OPENAPI_YAML) axios server using openapi-generator-cli"
	docker run -u $(GENERATE_USER) --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.0.0 generate -i /local/$(OPENAPI_YAML) -g typescript-axios -o /local/webui/src/axios --skip-validate-spec

generate: generate-openapi generate-client generate-redfish generate-axios

fmt-go:
	@echo "Format check(Go)"
	docker run --rm -v ${PWD}:/local golang:$(GO_VERSION) $(GOFMT_OPTS)

test-go: | vet-go
	@echo "Running all Go tests"
	go test -v ./...

vet-go:
	@echo "Running go vet"
	go vet ./...

test-go-backend:
	@echo "Running Go tests on pkg/backend"
	go test -v ./pkg/backend
