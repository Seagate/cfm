# Development environment setup

## Golang setup

- It is recomended to follow the [Download and install](https://go.dev/doc/install) page from go.dev
  - Version 1.22.1 is currently used for development

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

## cfm-webui development setup

cfm-webui is developed with Vue.js 3.

- Install dependencies

```bash
sudo apt update
sudo apt install nodejs npm
npm --version
```

- Install project dependencies

```bash
cd webui
npm install
```

More detail can be found at [README-WEBUI](../webui/README.md)
