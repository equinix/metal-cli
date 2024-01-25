## metal gateway create-bgp-dynamic-neighbours

Creates a BGP Dynamic Neighbour

### Synopsis

Creates the BGP Dynamic Neighbour for the metal gateway with the specified IP Range and ASN

```
metal gateway create-bgp-dynamic-neighbours [flags]
```

### Examples

```
# Create a BGP Dynamic Neighbour using ip range and asn for the gateway-id

	metal gateways create-bgp-dynamic-neighbour --gateway-id "9c56fa1d-ec05-470b-a938-0e5dd6a1540c" --bgp-neighbour-range "10.70.43.226/29" --asn 65000

```

### Options

```
      --asn int32                    
      --bgp-neighbour-range string   
      --gateway-id string            
  -h, --help                         help for create-bgp-dynamic-neighbours
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

