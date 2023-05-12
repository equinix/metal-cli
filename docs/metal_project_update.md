## metal project update

Updates a project.

### Synopsis

Updates the specified project with a new name, a new payment method, or both.

```
metal project update -i <project_UUID> [-n <name>] [-m <payment_method_UUID>] [flags]
```

### Examples

```
  # Updates the specified project with a new name:
  metal project update -i $METAL_PROJECT_ID -n new-prod-cluster05
  
  # Updates the specified project with a new payment method:
  metal project update -i $METAL_PROJECT_ID -m e2fcdf91-b6dc-4d6a-97ad-b26a14b66839
```

### Options

```
  -h, --help                       help for update
  -i, --id string                  The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.
  -n, --name string                The new name for the project.
  -m, --payment-method-id string   The UUID of the new payment method.
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

* [metal project](metal_project.md)	 - Project operations: create, get, update, delete, and bgpenable, bgpconfig, bgpsessions.

