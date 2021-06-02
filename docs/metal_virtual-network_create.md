## metal virtual-network create

Creates a virtual network

### Synopsis

Example:

metal virtual-network create --project-id [project_UUID] { --metro [metro_code] --vlan [vlan] | --facility [facility_code] }



```
metal virtual-network create [flags]
```

### Options

```
  -d, --description string   Description of the virtual network
  -f, --facility string      Code of the facility
  -h, --help                 help for create
  -m, --metro string         Code of the metro
  -p, --project-id string    UUID of the project
      --vxlan int            VXLAN id to use (can only be used with --metro)
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

* [metal virtual-network](metal_virtual-network.md)	 - Virtual network operations

