# Use CFM software in Docker

Notes regarding the use of Docker for the Composer & Fabric Manager (CFM) Software Suite.

## Obtain released docker image

A ready to use docker image for the cfm-service, the cli tool and the webui are released with the github package feature. One can pull the image from command line:

```bash
docker pull ghcr.io/seagate/cfm
```

## Service the CFM and its webUI

In order for the cfm-service to launch the webui properly, the user need to provide the IP address of the hosting server's IP to cfm-service. Example below:

```bash
docker run --rm --network=host --name cfm-suite --detach cfm -webuiIp xxx.xxx.xxx.xxx -verbosity 4
```

By default, the cfm-service will be hosted at port 8080 and the webui will be hosted at port 3000. The user could change the port by input argument -Port and/or -webuiPort. The webui only works with --network=host mode.

## Excute CLI tool

The user can start a cfm docker container to use the cli tool to interact with the running cfm-service.

```bash
docker run --rm --network=host --entrypoint "/cfm/cfm-cli"  cfm <args>
```

## Customization

The developer could use the [DockerFile](../docker/Dockerfile) as reference to build the docker image with local changes

```bash
docker build --no-cache -t cfm -f docker/Dockerfile .
```

#TODO: cxl-host
