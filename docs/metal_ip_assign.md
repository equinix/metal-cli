## metal ip assign

Assigns an IP address to a specified device.

### Synopsis

Assigns an IP address and subnet to a specified device. Returns an assignment ID.

```
metal ip assign -a <IP_address> -d <device_UUID> [flags]
```

### Examples

```
  # Assigns an IP address to a server:
  metal ip assign -d 060d1626-2481-475a-9789-c6f4bb927303  -a 198.51.100.3/31
```

### Options

```
  -a, --address string     IP address and CIDR you would like to assign.
  -d, --device-id string   The UUID of the device.
  -h, --help               help for assign
```

### Options inherited from parent commands

```
      --config string        Path to JSON or YAML configuration file
      --exclude strings      Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray   Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --include strings      Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string        Output format (*table, json, yaml). env output formats are (*sh, terraform, capp).
      --search string        Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string       Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string      Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string         Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal ip](metal_ip.md)	 - IP address, reservations, and assignment operations: assign, unassign, remove, available, request and get.

