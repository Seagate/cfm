# Composable Fabric Manager ( CFM ) Service for Linux

**cfm-service** is a Go language Linux service running in the background. It contains the main business logic for managing CMA memory. It provides a north-side (frontend) OpenAPI interface for client(s) and a south-side (backend) Redfish interface to manage available Composable Memory Appliances (CMAs) and CXL Hosts.

## Configuring cfm-service

The `cfm-service` application supports a number of configuration parameters that can be set in one of three ways:

- passed on the command line,
- set as ENV variables, or
- stored in a configuration file (EASIEST).

| Type   | Flag                     | Environment             | File                    |
| ------ | ------------------------ | ----------------------- | ----------------------- |
| string | -config cfm-service.conf | CONFIG=cfm-service.conf | config cfm-service.conf |
| bool   | -version true\|false     | VERSION=true\|false     | version true\|false     |
| string | -verbosity `<level>`     | VERBOSITY=`<level>`     | verbosity `<level>`     |
| int    | -Port `<servicePort>`    | PORT=`<servicePort>`    | Port `<servicePort>`    |
| string | -webuiIp `<serverIP>`    | WEBUIIP=`<serverIP>`    | webuiIp `<serverIP>`    |
| int    | -webuiPort `<webuiPort>` | WEBUIPORT=`<webuiPort>` | webuiPort `<webuiPort>` |

Order of precedence:

1. Command line options
2. Environment variables
3. Configuration file
4. Default values

## Running cfm-service

For development, follow this steps to easily start testing your code changes:

- Edit `cfm-service.conf` as needed
  - NOTE: If `webuiIp` is not provided, the webui service within cfm-service will **not** start
- Run `make build-service` to build cfm-service

### Example 1

```bash
./cfm-service -verbosity 4
```

### Example 2

```bash
./cfm-service -config cfm-service.conf -webuiIp xxx.xxx.xxx.xxx
```

## OpenAPI Documentation

The REST API for this service that is used by user interfaces is defined in an OpenAPI specification document. That document was used to generate a Go language service.

- [api/cfm-openapi.yaml](api/cfm-openapi.yaml) The OpenAPI document
- Run `make validate` to verify that any changes made to the OpenAPI document are valid using openapi-generator-cli.
- Run `make generate` to generate a Go language service from the OpenAPI document using openapi-generator-cli.
  - Generating new go files requires manual steps to incorporate the new code.

## Code Standards

Code must pass a gofmt test to be acc epted.

- For VSCode, in settings search for "go fmt" and change the language server from the default of gopls to gofmt.
- For VIM, add "au BufWritePost \*.go !gofmt -w %" into ~/.vimrc.
