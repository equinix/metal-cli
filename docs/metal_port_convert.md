## metal port convert

Converts a list of ports or the details of the specified port.

### Synopsis

Converts a list of ports or the details of the specified port. Details of an port are only available to its members.

```
metal port convert -i <port_UUID> [--bonded] [--bulk] --layer2 [--force] [--public-ipv4] [--public-ipv6] [flags]
```

### Examples

```
  # Converts list of the current user's ports:
  metal port convert -i <port_UUID> [--bonded] [--bulk] [--layer2] [--force] [--public-ipv4] [--public-ipv6]

  # Converts port to layer-2 unbonded:
  metal port convert -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8 --layer2 --bonded=false

  # Converts port to layer-2 bonded:
  metal port convert -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8 --layer2 --bonded

  # Converts port to layer-3 bonded with public IPv4 and public IPv6:
  metal port convert -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8 -2=false -b -4 -6
```

### Options

```
  -b, --bonded           Convert to layer-2 bonded.
      --bulk             Affect both ports in a bond.
  -f, --force            Force conversion to layer-2 bonded.
  -h, --help             help for convert
  -2, --layer2           Convert to layer-2 unbonded.
  -i, --port-id string   The UUID of a port.
  -4, --public-ipv4      Convert to layer-2 bonded with public IPv4.
  -6, --public-ipv6      Convert to layer-2 bonded with public IPv6.
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

* [metal port](metal_port.md)	 - Port operations: get, convert, vlans.

