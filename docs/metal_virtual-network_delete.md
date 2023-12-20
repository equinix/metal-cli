## metal virtual-network delete

Deletes a virtual network.

### Synopsis

Deletes the specified VLAN with a confirmation prompt. To skip the confirmation use --force. You are not able to delete a VLAN that is attached to any ports.

```
metal virtual-network delete -i <virtual_network_UUID> [-f] [flags]
```

### Examples

```
  # Deletes a VLAN, with confirmation.
  metal virtual-network delete -i 77e6d57a-d7a4-4816-b451-cf9b043444e2
  >
  ✔ Are you sure you want to delete virtual network 77e6d57a-d7a4-4816-b451-cf9b043444e2 [Y/N]: y
		
  # Deletes a VLAN, skipping confirmation.
  metal virtual-network delete -f -i 77e6d57a-d7a4-4816-b451-cf9b043444e2
```

### Options

```
  -f, --force       Skips confirmation for the removal of the virtual network.
  -h, --help        help for delete
  -i, --id string   UUID of the VLAN.
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

* [metal virtual-network](metal_virtual-network.md)	 - Virtual network (VLAN) operations : create, get, delete.

