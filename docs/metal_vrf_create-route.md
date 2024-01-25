## metal vrf create-route

Create a route in a VRF. Currently only static default routes are supported.

### Synopsis

Create a route in a VRF. Currently only static default routes are supported.

```
metal vrf create-route [-i <VrfId>] [-p <Prefix>] [-n NextHop] [-t <tags> ] [flags]
```

### Examples

```
 # Create a route in a VRF. Currently only static default routes are supported..
	
	metal vrf create-route [-i <VrfID>] [-p <prefix>] [-n nextHop] [-t <tags> ]
```

### Options

```
  -h, --help             help for create-route
  -i, --id string        Specify the VRF UUID activate route configurations
  -n, --nextHop string   The IPv4 address within the VRF of the host that will handle this route
  -p, --prefix string    The IPv4 prefix for the route, in CIDR-style notation. For a static default route, this will always be '0.0.0.0/0'
  -t, --tags strings     Adds the tags for the connection --tags="tag1,tag2".
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

* [metal vrf](metal_vrf.md)	 - VRF operations : create, get, delete

