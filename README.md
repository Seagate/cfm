# Composable Fabric Manager ( CFM ) for the Memory Appliance

This is a Go language application that will run as a background service providing functionality to multiple users. The whole software suite contains a few components.  
 __cfm-service__ is the main business logic for memory composition that will provide a north-side OpenAPI interface and use Redfish on the south-side to manage Memory Appliances and CXL Hosts.
 __cfm-clif__ is an interactive program providing functionality to a single users. This will interact with the north-side OpenAPI interface provided by cfm-service.  It is intended to free the users of needing to use curl commands to directly interact with the OpenAPI interface.  
__cfm-webui__ is a single-page application presenting a web UI using Vue.js 3. This application could be auto launched by cfm-service.

## Additional Project Documentation

- [WEBUI-README](webui/README.md) - README file for the __cfm-webui__ component
- [CLI_README](cli/README.md) - README file for the __cfm-cli__ component
- [DOCKER](docs/DOCKER.md) - Information on running __cfm-service__ and components in Docker containers
- [SETUP](docs/SETUP.md) - Information on setting up a development environment
- [LOG](docs/LOG.md) - Information on logging level definations
- [REGRESSION](docs/REGRESSION.md) - TODO - Information on running __cfm-regression__ test suite- [TEMPLATES](docs/TEMPLATES.md) - Information on the template file used in openapi-generaot for generating the go service

## OpenAPI Documentation

The REST API for this service that is used by user interfaces is defined in an OpenAPI specification document. That document was used to generate a Go language service.

- [api/cfm-openapi.yaml](api/cfm-openapi.yaml) The OpenAPI document
- Run `make validate` to verify that any changes made to the OpenAPI document are valid using openapi-generator-cli.
- Run `make generate` to generate a Go language service from the OpenAPI document using openapi-generator-cli.
  - Generating new go files requires manual steps to incorporate the new code.

## Running cfm-service

For development, follow this steps to easily start testing your code changes:

- `cp cfm-service.conf test-cfm-service.conf`
- edit `test-cfm-service.conf` as needed
- run `make test` to build and run your changes using test-cfm-service.conf

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

### Example 1

```bash
./cfm-service -verbosity 4
```

### Example 2

```bash
./cfm-service -config cfm-service.conf -webuiIp 127.0.0.1
```

## Testing

To run a quick test quite of the cfm-service API, run `cfm-service` and then:

- `cd test`
- `./testapi`

Then check the terminal window for cfm-service output.

## Code Standards

Code must pass a gofmt test to be acc epted.

- For VSCode, in settings search for "go fmt" and change the language server from the default of gopls to gofmt.
- For VIM, add "au BufWritePost \*.go !gofmt -w %" into ~/.vimrc.
