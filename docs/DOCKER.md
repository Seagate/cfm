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
docker run --restart unless-stopped --network=host --name <container-name> --detach ghcr.io/seagate/cfm -webui -verbosity 4
```

By default, the cfm-service will be hosted at port 8080 and the webui will be hosted at port 3000. The user could change the port by input argument -Port and/or -webuiPort. The webui only works with --network=host mode.

## View cfm-service logs

The cfm-service runtime logs can be viewed using

```bash
docker logs --follow <container-name>
```

## Stop and restart cfm-service

```bash
docker restart <container-name>
```

## Execute CLI tool

The user can interact with the running cfm docker container (running cfm-service) to utilize the cli tool.

```bash
docker exec -it <container-name> ./cfm-cli <args>
```

NOTE: cfm-cli \<args\> usage examples:

```bash
docker exec -it cfm-container ./cfm-cli -h
docker exec -it cfm-container ./cfm-cli list appliances
```

NOTE: Currently, every cfm-cli command requires various tcpip options (e.g.: --service-net-ip) regarding the specific cfm-service that is being interacted with.
However, using the docker container, the user can rely on the default cli settings for these cfm-service options since they point to the cfm-service running within the same docker container.

## Customization

The developer could use the [DockerFile](../docker/Dockerfile) as a reference to build a new docker image using a local [cfm](https://github.com/Seagate/cfm) clone...

```bash
cd /path/to/local/cfm
docker build --no-cache -t <new-image-name> -f docker/Dockerfile .
```

...and then run those changes

```bash
docker run --restart unless-stopped --network=host --name <new-container-name> --detach --privileged -v /var/run/dbus/system_bus_socket:/var/run/dbus/system_bus_socket <new-image-name> -webui -verbosity 4
```
