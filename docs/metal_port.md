## metal port

Port operations: get, convert, vlans.

### Synopsis

Information and operations for converting ports between networking modes and managing VLAN assignments to ports. For more information on the different modes, ports, and VLANs, visit https://deploy.equinix.com/developers/docs/metal/layer2-networking/overview/.

### Options

```
  -h, --help   help for port
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
* [metal port convert](metal_port_convert.md)	 - Converts a list of ports or the details of the specified port.
* [metal port get](metal_port_get.md)	 - Retrieves the details of the specified port.
* [metal port vlan](metal_port_vlan.md)	 - Modifies VLAN assignments on a port

