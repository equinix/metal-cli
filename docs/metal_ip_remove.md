## metal ip remove

Removes an IP address reservation from a project.

### Synopsis

Removes an IP address reservation from a project. Any subnets and IP addresses in the reservation will no longer be able to be used by your devices.

```
metal ip remove -i <reservation_UUID> [flags]
```

### Examples

```
  # Removes an IP address reservation:
  metal ip remove --id a9dfc9d5-ba1a-4d11-8cfc-6e30b9630876
```

### Options

```
  -h, --help        help for remove
  -i, --id string   UUID of the reservation
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

* [metal ip](metal_ip.md)	 - IP address, reservations, and assignment operations: assign, unassign, remove, available, request and get.

