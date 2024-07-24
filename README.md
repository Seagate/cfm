# Composable Fabric Manager ( CFM ) Software Suite

---

# Overview

The CFM Software Suite provides a client interface for interacting with a Composable Memory Appliance (CMA). The software suite consists of a few components.

- **cfm-service** is a Go language Linux service running in the background. It contains the main business logic for managing CMA memory. It provides a north-side (frontend) OpenAPI interface for client(s) and a south-side (backend) Redfish interface to manage available Composable Memory Appliances (CMAs) and CXL Hosts.
- **cfm-cli** is a Go language interactive program providing cfm-service command line functionality to a single user. This will interact with the north-side (frontend) OpenAPI interface provided by cfm-service. It is intended to free the users of needing to use curl commands to directly interact with cfm-service's OpenAPI interface.
- **cfm-webui** is a single-page application presenting a web UI using Vue.js 3. This will interact with the north-side (frontend) OpenAPI interface provided by cfm-service. This application is generally auto launched by cfm-service.
- **cxl-host** is a linux service that runs directly on the cxl-host server. This will interact with the south-side (backend) Redfish interface of the cfm-service.

# Quick Start

## Setup

### Docker

- The standard CFM release package runs within a docker container.
- To install docker, one can follow the steps from docker.docs
  - Example: [Install Docker Engine on Ubuntu | Docker Docs](https://docs.docker.com/engine/install/ubuntu/)

NOTE: If uncertain of which install option to use, the “install using the apt repository” option is a suggested option.

## Installation\Operation

- [DOCKER](docs/DOCKER.md) - Information on running the CFM Software Suite from within a Docker container

## Build locally

- User must run 'make generate' to generate all supporting files from the api yaml before building local binaries.

# Additional Project Documentation

- [README-SERVICE](docs/SERVICE.md) - README file for the **cfm-service** component
- [README-WEBUI](webui/README.md) - README file for the **cfm-webui** component
- [README-CLI](cli/README.md) - README file for the **cfm-cli** component
- [SETUP](docs/SETUP.md) - Information on setting up a development environment
- [LOG](docs/LOG.md) - Information on logging level definations
- [TEMPLATES](docs/TEMPLATES.md) - Information on the template file used in openapi-generate for generating the go service

# Known Issues

Here is a list of known issues that will be fixed in a future release:

- CFM\Hardware Re-synchronization
  - Currently, if a registered CMA blade is power-cycled, the cfm-service has no way to automatically detect this action.
  - A `resync` webui button (and cli cmd) have been provided to give the user manual control to force the service to update itself against the hardware (since it is assumed that the user is aware of the power cycle action).
- During startup, possible CFM response delay (via the WebUI\CLI tool)
  - During cfm-service initialization, it re-register(s) with previously registered CMA and CXL-Host hardware.
  - This can take some time, depending on how many devices are involved.
    - A single blade takes ~5-10 sec to register\connect.
  - There is currently no user feedback while this action is occurring (other than checking the docker service logs, `docker logs <container_name>`, in real time), so the user will need to wait before attempting any actions via the WebUI or CLI tools.
  - For the WebUI, use the browser's refresh button when attempting to see if this service inialization is complete.
