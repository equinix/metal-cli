## metal gateway delete

Deletes a Metal Gateway.

### Synopsis

Deletes the specified Gateway with a confirmation prompt. To skip the confirmation use --force.

```
metal gateway delete -i <metal_gateway_UUID> [-f] [flags]
```

### Examples

```
  # Deletes a Gateway, with confirmation.
  metal gateway delete -i 77e6d57a-d7a4-4816-b451-cf9b043444e2
  >
  ✔ Are you sure you want to delete Metal Gateway 77e6d57a-d7a4-4816-b451-cf9b043444e2: y

  # Deletes a Gateway, skipping confirmation.
  metal gateway delete -f -i 77e6d57a-d7a4-4816-b451-cf9b043444e2
```

### Options

```
  -f, --force       Skips confirmation for the removal of the Metal Gateway.
  -h, --help        help for delete
  -i, --id string   UUID of the Gateway.
```

### Options inherited from parent commands

```
      --config string        Path to JSON or YAML configuration file
      --exclude strings      Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray   Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --include strings      Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string        Output format (*table, json, yaml). env output formats are (*sh, terraform, capp).
      --search string        Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string       Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string      Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string         Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal gateway](metal_gateway.md)	 - Metal Gateway operations: create, delete, and retrieve.

