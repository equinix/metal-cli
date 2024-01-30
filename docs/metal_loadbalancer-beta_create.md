## metal loadbalancer-beta create

Creates a loadbalancer.

### Synopsis

Creates a loadbalancer with the specified name.

```
metal loadbalancer-beta create -n <loadbalancer_name> -l <location_id_or_metro> [-p <project_UUID>] [--provider <provider_id>] [--port <port_UUID>] [flags]
```

### Examples

```
  # Creates a new loadbalancer named dev-loadbal in the Dallas metro: 
  metal loadbalancer create --name dev-loadbal -l da
  
  # Creates a new loadbalancer named prod-loadbal in the DC metro:
  metal loadbalancer create -n prod-loadbal -l dc
```

### Options

```
  -h, --help                help for create
  -l, --location string     The location's ID. This flag is required.
  -n, --name string         Name of the loadbalancer
      --port strings        The port(s) UUID
  -p, --project-id string   The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.
  -r, --provider string     The provider ID. (default "loadpvd-gOB_-byp5ebFo7A3LHv2B")
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

* [metal loadbalancer-beta](metal_loadbalancer-beta.md)	 - LoadBalancer BETA operations: create, get, update, and delete.

