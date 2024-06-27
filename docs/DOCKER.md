# Use CFM software in Docker

Notes regarding the use of Docker for the Composer & Fabric Manager (CFM) Software Suite.

## Obtain released docker image

A ready to use docker image for the cfm-service, the cli tool and the webui is released with the github package feature. One can pull the image from the command line:

```bash
docker pull ghcr.io/seagate/cfm
```

If desired, the user can add `:vX.X.X` to the end of the command to obtain an older CFM release version.

## Start the CFM service and its webUI

To enable the webui launching during cfm-service startup, the user must provide the `-webui` flag in the command below.

```bash
docker run --network=host --name <user-defined-name> --detach ghcr.io/seagate/cfm -webui -verbosity 4
```

By default, the cfm-service will be hosted at port 8080 and the webui will be hosted at port 3000. The user could change the port by input argument -Port and/or -webuiPort. The webui only works with --network=host mode.

## Stop and restart cfm-service

```bash
docker restart <user-defined-name>
```

## Excute CLI tool

The user can start a cfm docker container to use the cli tool to interact with the running cfm-service.

```bash
docker run --network=host --entrypoint "/cfm/cfm-cli"  cfm <args>
```

## Customization

The developer could use the [DockerFile](../docker/Dockerfile) as a reference to build a new docker image with local changes

```bash
docker build --no-cache -t <user-defined-name> -f docker/Dockerfile .
```

#TODO: cxl-host
