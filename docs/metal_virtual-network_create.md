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
  -p, --project-id string    Project ID (METAL_PROJECT_ID)
      --vxlan int            VXLAN id to use (can only be used with --metro)
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

