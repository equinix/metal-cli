## packet virtual-network delete

Deletes a Virtual Network

### Synopsis

Example:

packet virtual-network delete -i [virtual_network_UUID]

	

```
packet virtual-network delete [flags]
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
  -j, --json              JSON output
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
  -y, --yaml              YAML output
```

### SEE ALSO

* [packet virtual-network](packet_virtual-network.md)	 - Virtual network operations

