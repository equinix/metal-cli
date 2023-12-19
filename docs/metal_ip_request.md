## metal ip request

Request a block of IP addresses.

### Synopsis

Requests either a block of public IPv4 addresses or global IPv4 addresses for your project in a specific metro or facility.

```
metal ip request -p <project-id> -t <ip_address_type> -q <quantity> (-m <metro> | -f <facility>) [-f <flags>] [-c <comments>] [flags]
```

### Examples

```
  # Requests a block of 4 public IPv4 addresses in Dallas:
  metal ip request -p $METAL_PROJECT_ID -t public_ipv4 -q 4 -m da

  metal ip request -v df18fbd8-2919-4104-a042-5d42a05b8eed -t vrf --cidr 24 -n 172.89.1.0 --tags foo --tags bar --customdata '{"my":"goodness"}' --details "i don't think VRF users need this or will see it after submitting the request"
```

### Options

```
      --cidr int            The size of the desired subnet in bits.
  -c, --comments string     General comments or description.
      --customdata string   customdata is to add to the reservation, in a comma-separated list.
      --details string      VRF IP Reservation's details
  -f, --facility string     Code of the facility where the IP Reservation will be created
  -h, --help                help for request
  -m, --metro string        Code of the metro where the IP Reservation will be created
  -n, --network string      The starting address for this VRF IP Reservation's subnet
  -p, --project-id string   The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.
  -q, --quantity int        Number of IP addresses to reserve.
      --tags strings        Adds the tags for the IP Reservations --tags "tag1,tag2" OR --tags "tag1" --tags "tag2"
  -t, --type string         The type of IP Address, either public_ipv4 or global_ipv4.
  -v, --vrf-id string       Specify the VRF UUID.
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

