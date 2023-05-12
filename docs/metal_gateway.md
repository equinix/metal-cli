## metal gateway

Metal Gateway operations: create, delete, and retrieve.

### Synopsis

A Metal Gateway provides a single IPv4 address as a gateway for a subnet. For more information, visit https://metal.equinix.com/developers/docs/networking/metal-gateway/.

### Options

```
  -h, --help   help for gateway
```

### Options inherited from parent commands

```
      --config string         Path to JSON or YAML configuration file
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

* [metal](metal.md)	 - Command line interface for Equinix Metal
* [metal gateway create](metal_gateway_create.md)	 - Creates a Metal Gateway.
* [metal gateway delete](metal_gateway_delete.md)	 - Deletes a Metal Gateway.
* [metal gateway get](metal_gateway_get.md)	 - Lists Metal Gateways.

