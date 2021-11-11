## metal virtual-network delete

Deletes a Virtual Network

### Synopsis

Example:

metal virtual-network delete -i [virtual_network_UUID]

	

```
metal virtual-network delete [flags]
```

### Options

```
  -f, --force       Force removal of the virtual network
  -h, --help        help for delete
  -i, --id string   UUID of the vlan
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -o, --output string     Output format (*table, json, yaml)
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
      --sort-by string    Sort fields for use in 'get' actions. Sort is not supported by all resources.
      --sort-dir string   Sort field direction for use in 'get' actions. Sort is not supported by all resources.
      --token string      Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal virtual-network](metal_virtual-network.md)	 - Virtual network operations

