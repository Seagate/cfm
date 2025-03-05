# Development environment setup

## Golang setup

- It is recomended to follow the [Download and install](https://go.dev/doc/install) page from go.dev
  - Version 1.23.7 is currently used for development

## openapi-generator-cli setup

- Docker is used to run the openapi-generator-cli released via docker image
- To install docker, one can follow the steps from docker.docs
  - Example: [Install Docker Engine on Ubuntu | Docker Docs](https://docs.docker.com/engine/install/ubuntu/)

Now validate that `openapi-generator-cli` works:

- `docker run openapitools/openapi-generator-cli <version>`

```bash
docker run openapitools/openapi-generator-cli version 7.0.0
```

## Server setup

Make sure Port 3000 and 8080 are not blocked by the server's firewall.

## cfm-service development setup

- [README-SERVICE](docs/SERVICE.md) - README file for the **cfm-service** component

## cfm-webui development setup

- [README-WEBUI](webui/README.md) - README file for the **cfm-webui** component
