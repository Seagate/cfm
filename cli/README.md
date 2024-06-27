# Composable Fabric Manager ( CFM ) Command Line Tool

## Description

**cfm-cli** is a Go language interactive program providing cfm-service command line functionality to a single user. This will interact with the north-side (frontend) OpenAPI interface provided by cfm-service. It is intended to free the users of needing to use curl commands to directly interact with cfm-service's OpenAPI interface.

## Installation

Use the included Makefile to build a local copy of the executable.

```bash
make build-cli
```

## Usage

`./cfm-cli <args> [flags]`

Example:

```bash
./cfm-cli list appliances [flags]
```

- Use -h for help
  - `./cfm-cli -h`
  - `./cfm-cli arg1 -h`
  - `./cfm-cli arg1 arg2 -h`
- Note: cfm-service **must** be running

## Accessing Client Libraries from cfm-service

This project uses an OpenApi generated client library for interacting with the OpenApi generated API client of the cfm-service.

The library may be imported within each go module like:
`service "cfm/pkg/client"`

## Support - TODO

Tell people where they can go to for help. It can be any combination of an issue tracker, a chat room, an email address, etc.

## Roadmap - TODO

If you have ideas for releases in the future, it is a good idea to list them in the README.

## Contributing - TODO

State if you are open to contributions and what your requirements are for accepting them.

Code must pass a gofmt test to be accepted.

- For VSCode, in settings search for "go fmt" and change the language server from the default of gopls to gofmt.
- For VIM, add "au BufWritePost \*.go !gofmt -w %" into ~/.vimrc.

## Authors and acknowledgment - TODO

Show your appreciation to those who have contributed to the project.

## License

Apache 2.0
Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

## Project status - TODO

If you have run out of energy or time for your project, put a note at the top of the README saying that development has slowed down or stopped completely. Someone may choose to fork your project or volunteer to step in as a maintainer or owner, allowing your project to keep going. You can also make an explicit request for maintainers.
