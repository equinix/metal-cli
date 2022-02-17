## metal virtual-network

Virtual network (VLAN) operations. For more information on how VLANs work in Equinix Metal, visit https://metal.equinix.com/developers/docs/layer2-networking/vlans/.

### Synopsis

Virtual network (VLAN) operations : create, get, delete.

### Options

```
  -h, --help   help for virtual-network
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

* [metal](metal.md)	 - Command line interface for Equinix Metal
* [metal virtual-network create](metal_virtual-network_create.md)	 - Creates a virtual network.
* [metal virtual-network delete](metal_virtual-network_delete.md)	 - Deletes a virtual network.
* [metal virtual-network get](metal_virtual-network_get.md)	 - Lists virtual networks.

