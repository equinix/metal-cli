## metal loadbalancer-beta delete

Deletes a loadbalancer.

### Synopsis

Deletes the specified loadbalancer with a confirmation prompt. To skip the confirmation use --force.

```
metal loadbalancer-beta delete --id <loadbalancer_UUID> [--force] [flags]
```

### Examples

```
  # Deletes loadbalancer 50693ba9-e4e4-4d8a-9eb2-4840b11e9375:
  metal loadbalancer delete -i 50693ba9-e4e4-4d8a-9eb2-4840b11e9375
  >
  ✔ Are you sure you want to delete loadbalancer 50693ba9-e4e4-4d8a-9eb2-4840b11e9375: y
  
  # Deletes loadbalancer 50693ba9-e4e4-4d8a-9eb2-4840b11e9375, skipping confirmation:
  metal loadbalancer delete -i 50693ba9-e4e4-4d8a-9eb2-4840b11e9375 -f
```

### Options

```
  -f, --force       Force removal of the loadbalancer
  -h, --help        help for delete
  -i, --id string   The loadbalancer's ID. This flag is required.
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

