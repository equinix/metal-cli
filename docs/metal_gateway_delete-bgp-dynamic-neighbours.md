## metal gateway delete-bgp-dynamic-neighbours

Deletes a BGP Dynamic Neighbour

### Synopsis

Deletes the BGP Dynamic Neighbour for the metal gateway with the specified ID

```
metal gateway delete-bgp-dynamic-neighbours [flags]
```

### Examples

```
# Deletes a BGP Dynamic Neighbour using the bgp dynamic neighbour ID

	$ metal gateways delete-bgp-dynamic-neighbour --id "9c56fa1d-ec05-470b-a938-0e5dd6a1540c"

	BGP Dynamic Neighbour deletion initiated. Please check 'metal gateway get-bgp-dynamic-neighbour -i 9c56fa1d-ec05-470b-a938-0e5dd6a1540c for status

```

### Options

```
  -h, --help        help for delete-bgp-dynamic-neighbours
  -i, --id string   
```

### Options inherited from parent commands

```
      --config string         Path to JSON or YAML configuration file (METAL_CONFIG)
      --exclude strings       Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray    Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --http-header strings   Headers to add to requests (in format key=value)
      --include strings       Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string         Output format (*table, json, yaml). env output formats are (*sh, terraform, capp).
      --search string         Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string        Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string       Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string          Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal gateway](metal_gateway.md)	 - Metal Gateway operations: create, delete, and retrieve.

