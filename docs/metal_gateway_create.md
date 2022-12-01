## metal gateway create

Creates a Metal Gateway.

### Synopsis

Creates a Metal Gateway on the VLAN. Either an IP Reservation ID or a Private Subnet Size must be specified.

```
metal gateway create -p <project_UUID> --virtual-network <virtual_network_UUID> (--ip-reservation-id <ip_reservation_UUID> | --private-subnet-size <size>) [flags]
```

### Examples

```
  # Creates a Metal Gateway on the VLAN with a given IP Reservation ID:
  metal gateway create -p $METAL_PROJECT_ID -v 77e6d57a-d7a4-4816-b451-cf9b043444e2 -r 50052f72-02b7-4b40-ac9d-253713e1e178

  # Creates a Metal Gateway on the VLAN with a Private 10.x.x.x/28 subnet:
  metal virtual-network create -p $METAL_PROJECT_ID -s 16
```

### Options

```
  -h, --help                       help for create
  -r, --ip-reservation-id string   UUID of the Public or VRF IP Reservation to assign.
  -s, --private-subnet-size int    Size of the private subnet to request (8 for /29)
  -p, --project-id string          The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.
  -v, --virtual-network string     UUID of the Virtual Network to assign.
```

### Options inherited from parent commands

```
      --config string        Path to JSON or YAML configuration file
      --exclude strings      Comma separated Href references to collapse in results, may be dotted three levels deep
      --filter stringArray   Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.
      --include strings      Comma separated Href references to expand in results, may be dotted three levels deep
  -o, --output string        Output format (*table, json, yaml)
      --search string        Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string       Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string      Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string         Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal gateway](metal_gateway.md)	 - Metal Gateway operations: create, delete, and retrieve.

