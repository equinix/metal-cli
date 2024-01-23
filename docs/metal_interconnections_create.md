## metal interconnections create

Creates an interconnection.

### Synopsis

Creates a new interconnection as per the organization ID or project ID 

```
metal interconnections create -n <name> [-m <metro>] [-r <redundancy> ] [-t <type> ] [-p <project_id> ] | [-O <organization_id> ] [flags]
```

### Examples

```
  # Creates a new interconnection named "it-interconnection":
  metal interconnections create -n <name> [-m <metro>] [-r <redundancy>] [-t "dedicated" ] [-p <project_id>] | [-O <organization_id>]

  metal interconnections create -n <name> [-m <metro>] [-r <redundancy>] [-t "shared" ] [-p <project_id>] | [-O <organization_id>] -T <service_token_type>

  metal interconnections create -n <name> [-m <metro>] [-r <redundancy>] [-t "shared" ] [-p <project_id>] | [-O <organization_id>] -T <service_token_type> -v <vrfs>
```

### Options

```
  -h, --help                        help for create
  -m, --metro string                Metro Id or Metro Code from where the interconnection will be originated.
  -n, --name string                 Name of the interconnection.
      --organization-id string      The Organization's UUID to be used for creating org level interconnection request. Either one of this flag or --project-id is required.
  -p, --project-id string           The project's UUID. Either one of this flag or --organization-id is required.
  -r, --redundancy string           Types of redundancy for the interconnection. Either 'primary' or 'redundant'.
  -T, --service-token-type string   Type of service token for shared connection. Enum: 'a_side', 'z_side'.
  -s, --speed int32                 The maximum speed of the interconnections. (default 1000000000)
  -t, --type string                 Type of of interconnection. Either 'dedicated' or 'shared' when requesting for a Fabric VC.
      --vlan int32Slice             A list of VLANs to attach to the Interconnection. Ex: --vlans 1000, 1001 . (default [])
      --vrf strings                 A list of VRFs to attach to the Interconnection. Ex: --vrfs uuid1, uuid2 .
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

* [metal interconnections](metal_interconnections.md)	 - interconnections operations: create, get, update, delete

