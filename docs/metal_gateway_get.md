## metal gateway get

Lists Metal Gateways.

### Synopsis

Retrieves a list of all VLANs for the specified project.

```
metal gateway get -p <project_UUID> [flags]
```

### Examples

```

  # Lists Metal Gateways for project 3b0795ba-ec9a-4a9e-83a7-043e7e11407c:
  metal virtual-network get -p 3b0795ba-ec9a-4a9e-83a7-043e7e11407c
```

### Options

```
  -h, --help                help for get
  -p, --project-id string   The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.
```

### Options inherited from parent commands

```
      --config string        Path to JSON or YAML configuration file
      --exclude strings      Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray   Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --include strings      Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string        Output format (*table, json, yaml)
      --search string        Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string       Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string      Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string         Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal gateway](metal_gateway.md)	 - Metal Gateways. For more information on Metal Gateways, visit https://metal.equinix.com/developers/docs/networking/metal-gateway/.

