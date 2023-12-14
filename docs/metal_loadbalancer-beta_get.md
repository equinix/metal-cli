## metal loadbalancer-beta get

Retrieves all the project loadbalancers or the details of a specified loadbalancer.

### Synopsis

Retrieves all the project loadbalancers or the details of a specified loadbalancer. You can specify which loadbalancer by UUID or name.

```
metal loadbalancer-beta get [-i <loadbalancer_UUID> | -n <loadbalancer_name>] [flags]
```

### Examples

```
  # Retrieve all loadbalancers:
  metal loadbalancer get
  
  # Retrieve a specific loadbalancer by UUID: 
  metal loadbalancer get -i 2008f885-1aac-406b-8d99-e6963fd21333

  # Retrieve a specific loadbalancer by name:
  metal loadbalancer get -n dev-cluster03
```

### Options

```
  -h, --help                  help for get
  -i, --id string             The loadbalancer's UUID, which can be specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.
  -n, --loadbalancer string   The name of the loadbalancer.
  -p, --project-id string     The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.
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

* [metal loadbalancer-beta](metal_loadbalancer-beta.md)	 - LoadBalancer BETA operations: create, get, update, and delete.

