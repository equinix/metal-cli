## metal virtual-network create

Creates a virtual network.

### Synopsis

Creates a VLAN in the specified project. If you are creating a VLAN in a metro, you can optionally specify the VXLAN ID otherwise it is auto-assigned. If you are creating a VLAN in a facility, the VXLAN ID is auto-assigned.

```
metal virtual-network create -p <project_UUID>  [-m <metro_code> -vxlan <vlan> | -f <facility_code>] [-d <description>] [flags]
```

### Examples

```
  # Creates a VLAN with vxlan ID 1999 in the Dallas metro:
  metal virtual-network create -p $METAL_PROJECT_ID -m da --vxlan 1999

  # Creates a VLAN in the sjc1 facility
  metal virtual-network create -p $METAL_PROJECT_ID -f sjc1
```

### Options

```
  -d, --description string   A user-friendly description of the virtual network.
  -f, --facility string      Code of the facility.
  -h, --help                 help for create
  -m, --metro string         Code of the metro.
  -p, --project-id string    The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.
      --vxlan int            Optional VXLAN ID. Must be between 2 and 3999 and can only be used with --metro.
```

### Options inherited from parent commands

```
      --config string         Path to JSON or YAML configuration file
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

