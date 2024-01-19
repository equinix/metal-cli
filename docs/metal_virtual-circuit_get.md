## metal virtual-circuit get

Retrieves virtual circuit for a specific circuit Id.

### Synopsis

Retrieves virtual circuit for a specific circuit Id.

```
metal virtual-circuit get -i <id> [flags]
```

### Examples

```
  # Retrieve virtual circuit for a specific circuit::

  # Retrieve the details of a specific virtual-circuit:
  metal vc get -i e9a969b3-8911-4667-9d99-57cd3dd4ef6f
```

### Options

```
  -h, --help        help for get
  -i, --id string   Specify UUID of the virtual-circuit
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

* [metal virtual-circuit](metal_virtual-circuit.md)	 - virtual-circuit operations: create, get, update, delete

