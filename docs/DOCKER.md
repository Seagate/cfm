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
docker run --restart unless-stopped --network=host --name <user-defined-container-name> --detach ghcr.io/seagate/cfm -webui -verbosity 4
```

By default, the cfm-service will be hosted at port 8080 and the webui will be hosted at port 3000. The user could change the port by input argument -Port and/or -webuiPort. The webui only works with --network=host mode.

## View cfm-service logs

The cfm-service runtime logs can be viewed using

```bash
docker logs --follow <user-defined-container-name>
```

## Stop and restart cfm-service

```bash
docker restart <user-defined-container-name>
```

## Excute CLI tool

The user can interact with the running cfm docker container (running cfm-service) to utilize the cli tool.

```bash
docker exec -it <user-defined-container-name> ./cfm-cli <args>
```

NOTE: cfm-cli \<args\> usage examples:

```bash
docker exec -it cfm-container ./cfm-cli -h
docker exec -it cfm-container ./cfm-cli list appliances
```

## Customization

The developer could use the [DockerFile](../docker/Dockerfile) as a reference to build a new docker image with local changes

```bash
docker build --no-cache -t <user-defined-container-name> -f docker/Dockerfile .
```
