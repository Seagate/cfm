# Customize openapi-generator-cli generation

# go-server

This target is used to generate the go-server. The command option used to generate the go-server is `openapi-generator-cli generate -i api/cfm-openapi.yaml -g go-server`.

If you want to customize the files generated, you can retrieve the current template files and modify the ones you want to change.

## Example:

The desire was to remove the `API version: 1.0.5` text that is output to a header for every go file generated.
The main reason for removing this comment is that when the version changes, all files get updated even if the only change is the new version.
We already have source control to track versions, so this comment causes extra work even when the file's code didn't change.

### Step 1: Retrieve the current template files.

```
$ openapi-generator-cli author template -g go-server
[main] INFO  o.o.codegen.cmd.AuthorTemplate - Extracted templates to 'out' directory. Refer to https://openapi-generator.tech/docs/templating for customization details.
```

This command will pull down every mustache template file used to generate the go files.

### Step 2: Find the mustache file you want to change.

In this case, we want to change `out/partial_header.mustache`. I removed the following lines.

```
 {{#version}}
 * API version: {{{.}}}
 {{/version}}

```

### Step 3: Copy all modified files to `templates/go-server`.

The templates folder can hold multiple folders for each type of target. Since ours if `go-sever` only a single file was added to this folder.
You can copy all files, **but** only the files change have to be copied. The files **not** copied will be used from their original source. The **out/** folder is no longer needed
and should not be checked in.

```
mv out/partial_header.mustache templates/go-server/
rm -rf out/
```

### Step 4: Add `-t templates/go-server` to the openapi-generator-cli call

Change, if needed the `make generate` target to use this new template folder. The option is: `-t templates/go-server`.

### Step 5: Generate new server files


# References

- [OpenAPI Generator Templating](https://openapi-generator.tech/docs/templating)
- [{{mustache}}](https://mustache.github.io/)
