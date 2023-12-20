## metal ip

IP address, reservations, and assignment operations: assign, unassign, remove, available, request and get.

### Synopsis

IP address and subnet operations, including requesting IPv4 and IPv6 addresses, assigning and removing IPs to servers, and getting information about subnets and their usage. For more information is available on https://metal.equinix.com/developers/docs/networking/ip-addresses/.

### Options

```
  -h, --help   help for ip
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
* [metal ip assign](metal_ip_assign.md)	 - Assigns an IP address to a specified device.
* [metal ip available](metal_ip_available.md)	 - Lists available IP addresses from a reservation.
* [metal ip get](metal_ip_get.md)	 - Retrieves information about IP addresses, IP address reservations, and IP address assignments.
* [metal ip remove](metal_ip_remove.md)	 - Removes an IP address reservation from a project.
* [metal ip request](metal_ip_request.md)	 - Request a block of IP addresses.
* [metal ip unassign](metal_ip_unassign.md)	 - Unassigns an IP address assignment.

