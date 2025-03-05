FROM openapitools/openapi-generator-cli:v7.0.0 AS api-gen

ENV OPENAPI_YAML=api/cfm-openapi.yaml
ENV OPENAPI_REDFISH_YAML=api/redfish-openapi.yaml
ENV BASEPATH=/cfm

# Copy whole repo into the image
COPY . ${BASEPATH}

# Validating ${OPENAPI_YAML} using openapi-generator-cli
RUN docker-entrypoint.sh validate -i ${BASEPATH}/${OPENAPI_YAML}
RUN docker-entrypoint.sh validate -i ${BASEPATH}/${OPENAPI_REDFISH_YAML}

# Generating ${OPENAPI_YAML} go server using openapi-generator-cli
RUN docker-entrypoint.sh generate -i ${BASEPATH}/${OPENAPI_YAML} -g go-server -o ${BASEPATH}/pkg/openapi --additional-properties=sourceFolder=,outputAsLibrary=true,router=mux,serverPort=8080,enumClassPrefix=true -t ${BASEPATH}/api/templates/go-server --skip-validate-spec

# Generating ${OPENAPI_YAML} go client using openapi-generator-cli
RUN docker-entrypoint.sh generate -i ${BASEPATH}/${OPENAPI_YAML} -g go -o ${BASEPATH}/pkg/client -p isGoSubmodule=true,withGoMod=false --package-name client --ignore-file-override ${BASEPATH}/api/ignore/.openapi-generator-ignore-client -t ${BASEPATH}/api/templates/go

# workaround for withGoMod=false not functioning with openapi-generator
RUN	rm ${BASEPATH}/pkg/client/go.mod
RUN	rm ${BASEPATH}/pkg/client/go.sum

# Generating ${OPENAPI_YAML} redfish server using openapi-generator-cli
RUN docker-entrypoint.sh generate -i ${BASEPATH}/${OPENAPI_REDFISH_YAML} -g go-server -o ${BASEPATH}/pkg/redfishapi --package-name redfishapi --additional-properties=sourceFolder=,outputAsLibrary=true,router=mux,serverPort=8080,enumClassPrefix=true -t ${BASEPATH}/api/templates/go-server --skip-validate-spec

# Generating ${OPENAPI_YAML} axios server using openapi-generator-cli
RUN docker-entrypoint.sh generate -i ${BASEPATH}/${OPENAPI_YAML} -g typescript-axios -o ${BASEPATH}/webui/src/axios --skip-validate-spec


FROM golang:1.23.7-alpine AS go
ENV BASEPATH=/cfm

RUN apk add git make
# copy source code with generated files to go image
COPY --from=api-gen ${BASEPATH} ${BASEPATH}

WORKDIR ${BASEPATH}

# Format files to conform to project standard
RUN gofmt -w ${BASEPATH}

# apply local patch for xml response fix
RUN git apply api/patch/*.redfish.patch

# build the excutable
RUN make build-go

FROM node:alpine AS npm
ENV BASEPATH=/cfm

# copy source code with generated files to go image
COPY --from=go ${BASEPATH}/webui ${BASEPATH}/webui

# setup web UI and build the dist package(generate the certificates for webui and cfm-service during this process)
WORKDIR ${BASEPATH}/webui
RUN npm install
# Set the environment variable
ENV NODE_ENV=docker
RUN npm run build

FROM alpine:latest
ENV BASEPATH=/cfm

# Install packages
RUN apk add --no-cache ca-certificates openssl bash curl jq

# copy source code with generated files to go image
COPY --from=go ${BASEPATH}/cfm-service ${BASEPATH}/cfm-service
COPY --from=go ${BASEPATH}/cfm-cli ${BASEPATH}/cfm-cli
COPY --from=npm ${BASEPATH}/webui/dist ${BASEPATH}/webui/dist

RUN ln -s /local/cfmdatastore.json ${BASEPATH}/cfmdatastore.json

LABEL org.opencontainers.image.source=https://github.com/seagate/cfm

WORKDIR ${BASEPATH}
# Start the service
ENTRYPOINT [ "/cfm/cfm-service" ]
