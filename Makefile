# Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
.PHONY: help clean local run run-defaults validate generate generate-client generate-redfish generate-axios webui-dist fmt test-go vet-go test-go-backend regression run-regression

APP_NAME := cfm-service
CLIAPP_NAME := cfm-cli
CONF_NAME := cfm-service.conf
TEST_CONF_NAME := test-cfm-service.conf
OPENAPI_YAML := api/cfm-openapi.yaml
OPENAPI_REDFISH_YAML := api/redfish-openapi.yaml
GO_VERSION := 1.22.1
GO_INSTALL_LOCATION := /usr/local/go/bin/
GOFMT_OPTS := $(GO_INSTALL_LOCATION)gofmt -w /local
GENERATE_USER := $(shell id -u ${USER}):$(shell id -g ${USER})

help:
	@echo ""
	@echo "-----------------------------------------------------------------------------------"
	@echo "make clean            - Remove all executables"
	@echo "make local            - Build all executables"
	@echo "make run              - Build a local $(APP_NAME) executable and run it using config file $(CONF_NAME)"
	@echo "make run-defaults     - Build a local $(APP_NAME) executable and run it using its' internal config defaults"
	@echo "make validate         - validate the openapi specification using openapi-generator-cli"
	@echo "make generate         - Generate go server code for the openapi specification using openapi-generator-cli"
	@echo "make generate-client  - Generate go client code for the openapi specification using openapi-generator-cli"
	@echo "make generate-redfish - Generate go server code for the redfish specification using openapi-generator-cli"
	@echo "make generate-axios   - Generate axios server code for the openapi specification using openapi-generator-cli"
	@echo "make webui-dist       - Build the webui distribution package for cfm-service"
	@echo "make fmt              - Run gofmt"
	@echo "make test-go          - Run all Go tests"
	@echo "make vet-go           - Run go vet"
	@echo "make test-go-backend  - Run Go unit tests on the backend go code"
	@echo "make regression       - Build cfm-service-regression"
	@echo "make run-regression   - Run cfm-service-regression"
	@echo ""

clean:
	@echo "Clean up..."
	go clean
	rm -f $(APP_NAME) $(CLIAPP_NAME) cfm-service-regression

local: clean
	@echo "Build local executable..."
	go build -o $(APP_NAME) -ldflags "-X main.buildTime=`date -u '+%Y-%m-%dT%H:%M:%S'`" ./cmd/$(APP_NAME)/main.go
	go build -o $(CLIAPP_NAME) -ldflags "-X main.buildTime=`date -u '+%Y-%m-%dT%H:%M:%S'`" ./cmd/$(CLIAPP_NAME)/main.go
	ls -lh $(APP_NAME) $(CLIAPP_NAME)

run: local
	@echo "Running $(APP_NAME) using config file $(CONF_NAME)"
	./$(APP_NAME) -config $(CONF_NAME)

run-defaults: local
	@echo "Running $(APP_NAME) using config defaults"
	./$(APP_NAME)

validate:
	@echo "Validating $(OPENAPI_YAML) using openapi-generator-cli"
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.0.0 version
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.0.0 validate -i /local/$(OPENAPI_YAML)

generate:
	@echo "Generating $(OPENAPI_YAML) go server using openapi-generator-cli"
	docker run -u $(GENERATE_USER) --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.0.0 generate -i /local/$(OPENAPI_YAML) -g go-server -o /local/pkg/openapi --additional-properties=sourceFolder=,outputAsLibrary=true,router=mux,serverPort=8080,enumClassPrefix=true -t /local/templates/go-server --skip-validate-spec
	@echo "Format files after generation to conform to project standard"
	docker run --rm -v ${PWD}:/local golang:$(GO_VERSION) $(GOFMT_OPTS)

generate-client:
	@echo "Generating $(OPENAPI_YAML) go client using openapi-generator-cli"
	docker run -u $(GENERATE_USER) --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.0.0 generate -i /local/$(OPENAPI_YAML) -g go -o /local/pkg/client -p isGoSubmodule=true --package-name client --ignore-file-override /local/api/.openapi-generator-ignore-client -t /local/templates/go
	@echo "Format files after generation to conform to project standard"
	docker run --rm -v ${PWD}:/local golang:$(GO_VERSION) $(GOFMT_OPTS)

generate-redfish:
	@echo "Generating $(OPENAPI_YAML) redfish server using openapi-generator-cli"
	docker run -u $(GENERATE_USER) --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.0.0 generate -i /local/$(OPENAPI_REDFISH_YAML) -g go-server -o /local/pkg/redfishapi --additional-properties=sourceFolder=,outputAsLibrary=true,router=mux,serverPort=8080,enumClassPrefix=true -t /local/templates/go-server --skip-validate-spec
	sed -i 's/package openapi/package redfishapi/g' ./pkg/redfishapi/*.go
	@echo "Format files after generation to conform to project standard"
	docker run --rm -v ${PWD}:/local golang:$(GO_VERSION) $(GOFMT_OPTS)

generate-axios:
	@echo "Generating $(OPENAPI_YAML) axios server using openapi-generator-cli"
	docker run -u $(GENERATE_USER) --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.0.0 generate -i /local/$(OPENAPI_YAML) -g typescript-axios -o /local/webui/src/axios --skip-validate-spec

webui-dist:
	@echo "Generating webui distribution package"
	cd webui
	npm run build
	cd ..

fmt:
	@echo "Format check"
	docker run --rm -v ${PWD}:/local golang:$(GO_VERSION) $(GOFMT_OPTS)

test: local
	@echo "Testing $(APP_NAME) using config file $(TEST_CONF_NAME)"
	./$(APP_NAME) -config $(TEST_CONF_NAME)

test-go: | vet-go
	@echo "Running all Go tests"
	go test -v ./...

vet-go:
	@echo "Running go vet"
	go vet ./...

test-go-backend:
	@echo "Running Go tests on pkg/backend"
	go test -v ./pkg/backend

regression:
	@echo "Build local cfm-service-regression..."
	go build -o cfm-service-regression cmd/cfm-service-regression/main.go

run-regression:
	./cfm-service-regression -debug 4 --ginkgo.v --ginkgo.fail-fast
