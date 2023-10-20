## metal virtual-circuit

virtual-circuit operations: create, get, update, delete

### Synopsis

For more information on https://deploy.equinix.com/developers/docs/metal/interconnections.

### Options

```
  -h, --help   help for virtual-circuit
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

* [metal](metal.md)	 - Command line interface for Equinix Metal
* [metal virtual-circuit create](metal_virtual-circuit_create.md)	 - Creates an create-virtual-circuit for specific interconnection.
* [metal virtual-circuit delete](metal_virtual-circuit_delete.md)	 - Deletes a virtual-circuit.
* [metal virtual-circuit get](metal_virtual-circuit_get.md)	 - Retrieves virtual circuit for a specific circuit Id.
* [metal virtual-circuit update](metal_virtual-circuit_update.md)	 - Updates a virtualcircuit.

