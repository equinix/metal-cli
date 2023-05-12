## metal port get

Retrieves the details of the specified port.

### Synopsis

Retrieves the details of the specified port. Details of an port are only available to its members.

```
metal port get -i <port_UUID> [flags]
```

### Examples

```
  # Retrieves details of a port:
  metal port get -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8
```

### Options

```
  -h, --help             help for get
  -i, --port-id string   The UUID of a port.
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

* [metal port](metal_port.md)	 - Port operations: get, convert, vlans.

