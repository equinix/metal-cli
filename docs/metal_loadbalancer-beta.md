## metal loadbalancer-beta

LoadBalancer BETA operations: create, get, update, and delete.

### Synopsis

Information and management for LoadBalancers is on https://deploy.equinix.com/developers/docs/metal/networking/load-balancers/.

### Options

```
  -h, --help   help for loadbalancer-beta
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

* [metal](metal.md)	 - Command line interface for Equinix Metal
* [metal loadbalancer-beta create](metal_loadbalancer-beta_create.md)	 - Creates a loadbalancer.
* [metal loadbalancer-beta delete](metal_loadbalancer-beta_delete.md)	 - Deletes a loadbalancer.
* [metal loadbalancer-beta get](metal_loadbalancer-beta_get.md)	 - Retrieves all the project loadbalancers or the details of a specified loadbalancer.

