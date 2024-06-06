# CFM Service (cfm-service) Regression Testing -- TO BE UPDATED

## Introduction
The `cfm-service-regression` test executable runs a series of **Ginkgo** expressive tests.

## Automated Regression Using Docker

An automated environment using Docker to build containerized images for running **cxl-host**, **cfm-service**, **memory-appliance**, and **cfm-service-regression** and then uses `docker compose` to run all containers. There are two configurations supported:
- (1) Run **cxl-host**, **cfm-service**, **memory-appliance** locally on your machine once, then run **cfm-service-regression** using local files multiple times
- (2) Run everything inside Docker containers, a standalone regression test suite

## Quick Summary
- (1) Run locally
  - One time setup
  - `docker compose --file docker-compose-cfm.yaml build --build-arg ssh_prv_key="$(cat ~/.ssh/id_rsa)" --progress=plain`
  - `docker compose --file docker-compose-cfm.yaml up --abort-on-container-exit`
  - `docker compose --file docker-compose-cfm.yaml down`
  - Run cfm-regression multiple times
    - `docker build -f Dockerfile.regression.local --build-arg ssh_prv_key="$(cat ~/.ssh/id_rsa)" --tag cxl-service-regression --progress=plain .`
    - `docker run --network="host" cxl-service-regression`
- (2) Run totally in docker containers
  - `docker compose --file docker-compose-regression.yaml build --build-arg ssh_prv_key="$(cat ~/.ssh/id_rsa)" --progress=plain`
  - `docker compose --file docker-compose-regression.yaml up --abort-on-container-exit`
  - `docker compose --file docker-compose-regression.yaml down`

***Notes:***
- It is important to use `docker compose` and NOT the older version of `docker-compose` (with a dash)
- Use `docker --version` to determine which version of docker you are using. Run `docker compose version` to validate your `docker` version is new enough.
- See [DOCKER Notes](DOCKER.md) for additional docker information on building the individual docker files
- When running in your local repo:
  - Start up **cxl-host**, **cfm-service**, **memory-appliance**, which takes about 4 minutes
  - The **cfm-regression** program will communicate with `localhost`
      - `IpAddress: "localhost"`
  - Since the `cfm-service` is talking to **cxl-host** and the **memory-appliance** within the docker compose network, use:
    - `IpAddress: "memory-appliance"`
    - `IpAddress: "cxl-host"`


### Configuration

- `Dockerfile.host` - is the Dockerfile for running `cxl-host` using a git clone of the repo
- `Dockerfile.cfm` - is the Dockerfile for running `cfm-service` using a git clone of the repo
- `Dockerfile.mock-appliance` - is the Dockerfile for running a `Mock OpenBMC Memory Appliance` using QEMU and an existing Memory Technology Device (mtd) file
- `Dockerfile.regression` - is the Dockerfile for running `cfm-service-regression` using a git clone of the repo
- `Dockerfile.regression.local` - is the Dockerfile for running `cfm-service-regression` using your local repo files
- `docker-compose-cfm.yaml` - is the Docker Compose file for for running **cxl-host**, **cfm-service**, and **memory-appliance**.
- `docker-compose-regression.yaml` - is the Docker Compose file for for running **cxl-host**, **cfm-service**, **memory-appliance** and **cfm-regression**.

## Two Step Process

It takes two steps to build and then run the containerized regression test.

### Building

The build step uses local SSH keys to enable closing the repository from GitLab and GitHub. The build step takes several minutes the first time. Since docker can cache images, subsequent builds run very quickly.

- `docker compose --file docker-compose-regression.yaml build --build-arg ssh_prv_key="$(cat ~/.ssh/id_rsa)"`
- `docker image ls` will show you all the versions of local docker images.
- `docker compose --file docker-compose-regression.yaml down` will remove running container images, and no harm in running it before running the composed images.

### Debugging (optional)

This step allows you to see the output from the various steps written to the console, and it does not use cached images so all steps are executed.

- `docker compose --file docker-compose-regression.yaml build --build-arg ssh_prv_key="$(cat ~/.ssh/id_rsa)" --progress=plain`
- `docker compose --file docker-compose-regression.yaml build --build-arg ssh_prv_key="$(cat ~/.ssh/id_rsa)" --progress=plain --no-cache`

### Running

This step uses the built docker images and runs them in proper order.

***Notes:***
- The docker compose environment uses its own internal network to communicate over HTTP from container image to container image. (See https://docs.docker.com/compose/networking/)
- Instead of using `localhost` for talking to the CFM Service, you use the service name defined in `docker-compose.yml`. For example, use `cfm-service` for `CFMService.IpAddress`.
- Instead of using `localhost` for talking to the CXL Host, you use the service name defined in `docker-compose.yml`. For example, use `cxl-host` for `CxlHostEndpoint.IpAddress`.
- Instead of using `localhost` for talking to the Memory Appliance, you use the service name defined in `docker-compose.yml`. For example, use `memory-appliance`.

Example `cfm-service-regression-config.yaml`:
```
CFMService:
  IpAddress: "cfm-service"
  Port:      8080
ApplianceEndpoint:
  Name: 
  Username:  "root"
  Password:  "0penBmc"
  IpAddress: "memory-appliance"
  Port:      7443
  Insecure:  true
  Protocol:  "https"
CxlHostEndpoint: 
  Name:
  Username:  "admin"
  Password:  "admin12345"
  IpAddress: "cxl-host"
  Port:      8082
  Insecure:  true
  Protocol:  "http"
```

Running the regression test suite:
- `docker compose --file docker-compose-regression.yaml up --abort-on-container-exit`

The expected output is to see the `cxl-host`, `cfm-service` and `memory-appliance` run first, then for `cfm-service-regression` to run and complete all test cases. After completion, the containers will exit and the compose job will also exit.

![docker compose up](cfm-regresion-docker-compose-up.jpg)

Stopping the composed regression test suite:
- `docker compose --file docker-compose-regression.yaml down`
