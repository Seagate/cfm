# Use CFM software in Docker

Notes regarding the use of Docker for the Composer & Fabric Manager (CFM) Software Suite.

## Obtain released docker image

A ready to use docker image for the cfm-service, the cli tool and the webui is released with the github package feature. One can pull the image from the command line:

```bash
docker pull ghcr.io/seagate/cfm:vX.X.X
```
where `X.X.X` is the desired CFM release version.

## Start the CFM service and its webUI

In order for the cfm-service to launch the webui properly, the user need to provide the IP address of the hosting server's IP to cfm-service. Example below:

```bash
docker run --network=host --name <name> --detach cfm -webui -verbosity 4
```

By default, the cfm-service will be hosted at port 8080 and the webui will be hosted at port 3000. The user could change the port by input argument -Port and/or -webuiPort. The webui only works with --network=host mode.

## Stop and restart cfm-service

```bash
docker restart <name>
```

## Excute CLI tool

The user can start a cfm docker container to use the cli tool to interact with the running cfm-service.

```bash
docker run --network=host --entrypoint "/cfm/cfm-cli"  cfm <args>
```

## Customization

The developer could use the [DockerFile](../docker/Dockerfile) as a reference to build a new docker image with local changes

```bash
docker build --no-cache -t <name> -f docker/Dockerfile .
```

#TODO: cxl-host
