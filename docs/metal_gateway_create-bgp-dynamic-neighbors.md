## metal gateway create-bgp-dynamic-neighbors

Creates a BGP Dynamic Neighbor

### Synopsis

Creates the BGP Dynamic Neighbor for the metal gateway with the specified IP Range and ASN

```
metal gateway create-bgp-dynamic-neighbors [flags]
```

### Examples

```
# Create a BGP Dynamic Neighbor using ip range and asn for the metal gateway id

	metal gateways create-bgp-dynamic-neighbor --id "9c56fa1d-ec05-470b-a938-0e5dd6a1540c" --bgp-neighbor-range "10.70.43.226/29" --asn 65000

```

### Options

```
      --asn int32                   ASN for the BGP Dynamic Neighbor IP range.
      --bgp-neighbor-range string   BGP Dynamic Neighbor IP Range from gateway.
  -h, --help                        help for create-bgp-dynamic-neighbors
  -i, --id string                   Metal Gateway ID for which the BGP Dynamic Neighbor to be created.
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

