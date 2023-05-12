## metal capacity

Capacity operations: get, check

### Synopsis

Capacity operations. For more information on capacity in metros, visit https://metal.equinix.com/developers/docs/locations/metros/. For more information on capacity in facilities, visit https://metal.equinix.com/developers/docs/locations/facilities/.

### Options

```
  -h, --help   help for capacity
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
* [metal capacity check](metal_capacity_check.md)	 - Validates if the number of the specified server plan is available in the specified metro or facility.
* [metal capacity get](metal_capacity_get.md)	 - Returns capacity of metros or facilities, with optional filtering.

