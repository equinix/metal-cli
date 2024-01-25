## metal vrf create

Creates a Virtual Routing and Forwarding(VRF) for a specified project.

### Synopsis

Creates a Creates a Virtual Routing and Forwarding(VRF) for a specified project.

```
metal vrf create [-p <project_id] [-d <description>] [-m <metro>] [-n <name>] [-a <localASN>] [-r <IPranges>] [-t <tags> ] [flags]
```

### Examples

```
 # Creates an Creates a Virtual Routing and Forwarding(VRF) for a specified project.
	
	metal vrf create [-p <project_id] [-d <description>] [-m <metro>] [-n <name>] [-a <localASN>] [-r <ipranges>] [-t <tags> ]
```

### Options

```
  -d, --description string   Description of the Virtual Routing and Forwarding.
  -h, --help                 help for create
  -r, --ipranges strings     A list of CIDR network addresses. Like [10.0.0.0/16, 2001:d78::/56]. IPv4 blocks must be between /8 and /29 in size. IPv6 blocks must be between /56 and /64.
  -a, --local-asn int32      Local ASN for the VRF
  -m, --metro string         The UUID (or metro code) for the Metro in which to create this Virtual Routing and Forwarding
  -n, --name string          Name of the Virtual Routing and Forwarding
  -p, --project-id string    The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.
  -t, --tags strings         Adds the tags for the virtual-circuit --tags "tag1,tag2" OR --tags "tag1" --tags "tag2"
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

