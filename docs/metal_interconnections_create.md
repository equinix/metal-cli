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
  -m, --metro string                metro in the interconnection
  -n, --name string                 Name of the interconnection
  -O, --organizationID string       Org ID
  -p, --projectID string            project ID
  -r, --redundancy string           Website URL of the organization.
      --service-token-type string   Type of service token for shared connection. Enum: 'a_side', 'z_side'
  -t, --type string                 type of of interconnection.
      --vlans int32Slice            Array of int vLANs (default [])
      --vrfs strings                Array of strings VRF <uuid>.
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

