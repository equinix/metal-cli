## metal ip available

Lists available IP addresses from a reservation.

### Synopsis

Lists available IP addresses in a specified reservation for the desired subnet size.

```
metal ip available -r <reservation_UUID> -c <size_of_subnet> [flags]
```

### Examples

```
  # Lists available IP addresses in a reservation for a /31 subnet:
  metal ip available --reservation-id da1bb048-ea6e-4911-8ab9-b95635ca127a --cidr 31
```

### Options

```
  -c, --cidr int                The size of the desired subnet in bits.
  -h, --help                    help for available
  -r, --reservation-id string   The UUID of the IP address reservation.
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

* [metal ip](metal_ip.md)	 - IP address, reservations, and assignment operations: assign, unassign, remove, available, request and get.

